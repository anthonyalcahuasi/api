package apis

import (
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CRUD for items table
func EstudianteGET(c *gin.Context) {
	var lis []models.Estudiante
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("User.Rol").Preload("Inscripcion").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de los estudiantes",
		"r":   lis,
	})

}
func EstudianteGETID(c *gin.Context) {

	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Estudiante
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("User.Rol").Preload("Inscripcion.CargaDocente.Docente.User").Preload("Inscripcion.CargaDocente.Periodo").Preload("Inscripcion.CargaDocente.Grupo").Preload("Inscripcion.CargaDocente.Plan.Curso").Preload("Inscripcion.CargaDocente.Plan.Jerarquia.Grado.Nivel").Preload("Inscripcion.CargaDocente.Plan.Jerarquia.Sucursal").Preload("Inscripcion.CargaDocente.Unidad.Sesion.Recurso.TipoRecurso").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func EstudianteUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Estudiante
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

func EstudianteDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	id := c.Param("id")
	var d models.Estudiante

	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
