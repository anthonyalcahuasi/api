package apis

import (
	"fmt"
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CRUD for Jerarquia table

func RecursoViewGET(c *gin.Context) {
	var lis []models.RecursoView
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	if err := conn.Find(&lis).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de los Recursos vistos",
		"r":   lis,
	})
}
func RecursoViewGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.RecursoView
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Find(&d)
	c.JSON(http.StatusOK, &d)
}
func RecursoViewPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	recurso := models.RecursoView{}
	err := c.ShouldBindJSON(&recurso)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = recurso.Register(&conn)
	if err != nil {
		fmt.Println("Error in user.Register()")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn.Find(&recurso)
	c.JSON(http.StatusOK, recurso)
}
func RecursoViewUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.RecursoView
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

func RecursoViewDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.RecursoView
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
