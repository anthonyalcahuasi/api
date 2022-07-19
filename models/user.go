package models

import (
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go" //jwt-best-practices
	"golang.org/x/crypto/bcrypt"

	"time"

	"gorm.io/gorm"
)

var (
	tokenSecret = []byte(os.Getenv("TOKEN_SECRET"))
)

type User struct {
	ID              uint      `gorm:"primary_key;column:id" json:"id"`
	NOMBRES         string    `json:"nombres"`
	APPATERNO       string    `json:"appaterno"`
	APMATERNO       string    `json:"apmaterno"`
	CELULAR         string    `json:"celular" gorm:"uniqueIndex"`
	FECHANAC        time.Time `json:"fechanac"`
	SEXO            string    `json:"sexo"`
	FOTO            string    `json:"foto"`
	Email           string    `gorm:"column:email" json:"email"`
	PasswordHash    string    `json:"-"`
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
	Rol             []Rol     `gorm:"many2many:user_rol;"`
}

func (u *User) Register(conn *gorm.DB) error {

	if len(u.Password) < 4 || len(u.PasswordConfirm) < 4 {
		return fmt.Errorf("Password must be at least 4 characters long.")
	}
	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Passwords do not match.")
	}
	if len(u.Email) < 4 {
		return fmt.Errorf("Email must be at least 4 characters long.")
	}
	u.Email = strings.ToLower(u.Email)
	u.CELULAR = strings.ToLower(u.CELULAR)
	var userLookup User
	var err error
	err = conn.First(&userLookup, "celular = ?", u.CELULAR).Error
	if u.CELULAR == strings.ToLower(userLookup.CELULAR) {
		fmt.Println("found user")
		fmt.Println(userLookup.CELULAR)
		return fmt.Errorf("Un usuario está utilizando este número de celular")
	}
	if err = conn.First(&userLookup, "email = ?", u.Email).Error; err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		//return fmt.Errorf("Error p" + err.Error())
	}

	if u.Email == strings.ToLower(userLookup.Email) {
		fmt.Println("found user")
		fmt.Println(userLookup.Email)
		return fmt.Errorf("A user with that email already exists")
	}

	pwdHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("There was an error creating your account.")
	}
	u.PasswordHash = string(pwdHash)
	u.Password = ""
	u.PasswordConfirm = ""
	conn.Create(&u)
	return err // ya te asigna el ID del user
}

func (u *User) GetAuthToken() (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Minute * 120).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	authToken, err := token.SignedString(tokenSecret)
	return authToken, err
}

func DelTokenValid(tokenString string) bool {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok == false {
			return nil, fmt.Errorf("Token signing method is not valid: %v", token.Header["alg"])
		}

		return tokenSecret, nil
	})

	if err != nil {
		fmt.Printf("Err %v \n", err)
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims["authorized"] = false
		claims["user_id"] = nil
		claims["exp"] = time.Now().Add(time.Hour * 0).Unix()
		fmt.Println(claims)
		return true
	} else {
		fmt.Printf("The alg header %v \n", claims["alg"])
		fmt.Println(err)
		return false
	}
}

func IsTokenValid(tokenString string) (bool, string) { //VerifyToken and TokenValid
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok == false {
			return nil, fmt.Errorf("Token signing method is not valid: %v", token.Header["alg"])
		}

		return tokenSecret, nil
	})

	if err != nil {
		fmt.Printf("Err %v \n", err)
		return false, ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
		return true, ""
	} else {
		fmt.Printf("The alg header %v \n", claims["alg"])
		fmt.Println(err)
		return false, "uuid.UUID{}"
	}
}

func (u *User) IsAuthenticated(conn *gorm.DB) error {
	u.Email = strings.ToLower(u.Email)
	var userLookup User
	var err error
	if err = conn.First(&userLookup, "email = ?", u.Email).Error; err != nil {
		//return fmt.Errorf("Error p" + err.Error())
		fmt.Println("User with email not found")
		return fmt.Errorf("Invalid login credentials email")
	}
	u.PasswordHash = string(userLookup.PasswordHash)
	u.ID = userLookup.ID

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(u.Password))
	if err != nil {
		//fmt.Println("pss: " + err.Error())
		return fmt.Errorf("Invalid login credentials pass")
	}

	return nil
}
