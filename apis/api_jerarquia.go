package apis

import (
	"fmt"
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CRUD for Jerarquia table

func JerarquiaGET(c *gin.Context) {
	var lis []models.Jerarquia

	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Grado.Nivel").Preload("Sucursal").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las Jerarquias",
		"r":   lis,
	})
}
func JerarquiaGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Jerarquia
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Sucursal").Preload("Grado.Nivel").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func JerarquiaPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	jerarquia := models.Jerarquia{}
	err := c.ShouldBindJSON(&jerarquia)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = jerarquia.Register(&conn)
	if err != nil {
		fmt.Println("Error in user.Register()")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, jerarquia)
}
func JerarquiaUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Jerarquia
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.BindJSON(&d)
	conn.Save(&d)
	c.JSON(http.StatusOK, &d)
}

func JerarquiaDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Jerarquia
	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
	c.JSON(http.StatusOK, gin.H{
		"delete": &d,
	})
}
