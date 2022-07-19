package apis

import (
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CargaDocenteGET(c *gin.Context) {
	var lis []models.CargaDocente
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Docente.User").Preload("Plan.Curso").Preload("Periodo").Preload("Grupo").Preload("Periodo").Preload("Plan.Jerarquia.Grado.Nivel").Preload("Plan.Jerarquia.Sucursal").Preload("Unidad").Preload("Inscripcion").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las Cargas a los Docentes",
		"r":   lis,
	})
}
func CargaDocenteGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.CargaDocente
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Docente.User").Preload("Plan.Curso").Preload("Periodo").Preload("Grupo").Preload("Periodo").Preload("Unidad.Sesion").Preload("Unidad.Sesion.Recurso.RecursoView").Preload("Plan.Jerarquia.Grado.Nivel").Preload("Plan.Jerarquia.Sucursal").Preload("Inscripcion.Estudiante.User").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func CargaDocenteGETCODIGO(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	codigo := c.Param("codigo")
	var d models.CargaDocente
	if err := conn.First(&d, "codigo = ?", codigo).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Find(&d)
	c.JSON(http.StatusOK, gin.H{
		"id": d.ID,
	})
}
func CargaDocentePOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.CargaDocente
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func CargaDocenteUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.CargaDocente
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

func CargaDocenteDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.CargaDocente
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
