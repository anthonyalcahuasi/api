package apis

import (
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//CRUD for Unidad table

func UnidadGET(c *gin.Context) {
	var lis []models.Unidad

	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("CargaDocente.Docente").Preload("Sesion.Recurso.TipoRecurso").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las Unidads",
		"r":   lis,
	})
}
func UnidadGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Unidad
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload(clause.Associations).Preload("Sesion.Recurso.TipoRecurso").Preload("Sesion.Recurso.Sesion.Unidad").Preload("Sesion.Recurso.RecursoView").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func UnidadPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.Unidad
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func UnidadUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Unidad
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

func UnidadDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Unidad
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
