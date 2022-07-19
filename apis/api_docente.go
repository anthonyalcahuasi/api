package apis

import (
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CRUD for items table
func DocenteGET(c *gin.Context) {
	var lis []models.Docente
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("User").Preload("CargaDocente").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de Docentes",
		"r":   lis,
	})

}
func DocenteGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Docente
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("User.Rol").Preload("CargaDocente.Periodo").Preload("CargaDocente.Grupo").Preload("CargaDocente.Docente.User").Preload("CargaDocente.Plan.Curso").Preload("CargaDocente.Plan.Jerarquia.Grado.Nivel").Preload("CargaDocente.Plan.Jerarquia.Sucursal").Find(&d)
	c.JSON(http.StatusOK, &d)
}

func DocenteUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Docente
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

func DocenteDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Docente

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
