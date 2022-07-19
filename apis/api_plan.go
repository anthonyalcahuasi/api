package apis

import (
	"fmt"
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CRUD for Jerarquia table

func PlanGET(c *gin.Context) {
	var lis []models.Plan
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	if err := conn.Preload("Jerarquia.Grado.Nivel").Preload("Jerarquia.Sucursal").Preload("Curso").Preload("CargaDocente").Find(&lis).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de los Planes",
		"r":   lis,
	})
}
func PlanGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Plan
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Jerarquia.Grado.Nivel").Preload("Jerarquia.Sucursal").Preload("Curso").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func PlanPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	plan := models.Plan{}
	err := c.ShouldBindJSON(&plan)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = plan.Register(&conn)
	if err != nil {
		fmt.Println("Error in user.Register()")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	conn.Preload("Jerarquia.Grado.Nivel").Preload("Jerarquia.Sucursal").Preload("Curso").Find(&plan)
	c.JSON(http.StatusOK, plan)
}
func PlanUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Plan
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

func PlanDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Plan
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
