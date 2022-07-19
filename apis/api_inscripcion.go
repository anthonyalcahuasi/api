package apis

import (
	"fmt"
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CRUD for Jerarquia table

func InscripcionGET(c *gin.Context) {
	var lis []models.Inscripcion

	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("CargaDocente.Docente.User").Preload("Estudiante.User").Preload("CargaDocente.Periodo").Preload("CargaDocente.Grupo").Preload("CargaDocente.Plan.Curso").Preload("CargaDocente.Plan.Jerarquia.Grado.Nivel").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las Inscripcions",
		"r":   lis,
	})
}
func InscripcionGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Inscripcion
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("CargaDocente.Docente.User").Preload("Estudiante.User").Preload("CargaDocente.Periodo").Preload("CargaDocente.Grupo").Preload("CargaDocente.Plan.Curso").Preload("CargaDocente.Plan.Jerarquia.Grado.Nivel").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func InscripcionPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	inscripcion := models.Inscripcion{}
	err := c.ShouldBindJSON(&inscripcion)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = inscripcion.Register(&conn)
	if err != nil {
		fmt.Println("Error in user.Register()")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, inscripcion)
}
func InscripcionUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Inscripcion
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

func InscripcionDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Inscripcion
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
