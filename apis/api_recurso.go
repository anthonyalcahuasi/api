package apis

import (
	"net/http"

	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RecursoGET(c *gin.Context) {
	var lis []models.Recurso

	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Sesion").Preload("TipoRecurso").Preload("RecursoView").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las Recursos",
		"r":   lis,
	})
}
func RecursoGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Recurso
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Sesion").Preload("TipoRecurso").Preload("RecursoView").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func RecursoPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.Recurso
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func RecursoUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Recurso
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

func RecursoDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Recurso
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

func TipoRecursoGET(c *gin.Context) {
	var lis []models.TipoRecurso

	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Recurso").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las TipoRecursos",
		"r":   lis,
	})
}
func TipoRecursoGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.TipoRecurso
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Recurso").Find(&d)
	c.JSON(http.StatusOK, &d)
}

func TipoRecursoGETNOMBRE(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	nombre := c.Param("nombre")
	var d models.TipoRecurso
	if err := conn.First(&d, "nombre = ?", nombre).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":     &d.ID,
		"nombre": &d.NOMBRE,
	})
}

func TipoRecursoPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.TipoRecurso
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func TipoRecursoUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.TipoRecurso
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

func TipoRecursoDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.TipoRecurso
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

func ToolVideoConferenciaGET(c *gin.Context) {
	var lis []models.ToolVideoConferencia
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Recurso").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las ToolVideoConferencias",
		"r":   lis,
	})
}
func ToolVideoConferenciaGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolVideoConferencia
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Recurso.TipoRecurso").Preload("Recurso.RecursoView").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolVideoConferenciaPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.ToolVideoConferencia
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolVideoConferenciaUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolVideoConferencia
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

func ToolVideoConferenciaDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolVideoConferencia
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

func ToolVideoGET(c *gin.Context) {
	var lis []models.ToolVideo
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Recurso").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las ToolVideos",
		"r":   lis,
	})
}
func ToolVideoGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolVideo
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Recurso.TipoRecurso").Preload("Recurso.RecursoView").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolVideoPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.ToolVideo
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolVideoUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolVideo
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

func ToolVideoDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolVideo
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

func ToolEnlaceGET(c *gin.Context) {
	var lis []models.ToolEnlace
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Recurso").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las ToolEnlaces",
		"r":   lis,
	})
}
func ToolEnlaceGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolEnlace
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Recurso.TipoRecurso").Preload("Recurso.RecursoView").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolEnlacePOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.ToolEnlace
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolEnlaceUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolEnlace
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

func ToolEnlaceDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolEnlace
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

func ToolDocumentoGET(c *gin.Context) {
	var lis []models.ToolDocumento
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Recurso").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las ToolDocumentos",
		"r":   lis,
	})
}
func ToolDocumentoGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolDocumento
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Recurso.TipoRecurso").Preload("Recurso.RecursoView").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolDocumentoPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.ToolDocumento
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolDocumentoUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolDocumento
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

func ToolDocumentoDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolDocumento
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

func ToolActividadGET(c *gin.Context) {
	var lis []models.ToolActividad
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Recurso").Preload("Trabajo").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las ToolActividads",
		"r":   lis,
	})
}
func ToolActividadGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolActividad
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Recurso").Preload("Trabajo").Preload("Recurso.RecursoView").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolActividadPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.ToolActividad
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolActividadUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolActividad
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

func ToolActividadDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolActividad
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

func ToolForoGET(c *gin.Context) {
	var lis []models.ToolForo
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Recurso").Preload("Foro").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las ToolForos",
		"r":   lis,
	})
}
func ToolForoGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolForo
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Recurso").Preload("Foro").Preload("Recurso.RecursoView").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolForoPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.ToolForo
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolForoUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolForo
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

func ToolForoDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolForo
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

func ToolAnuncioGET(c *gin.Context) {
	var lis []models.ToolAnuncio
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Recurso").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las ToolAnuncios",
		"r":   lis,
	})
}
func ToolAnuncioGETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolAnuncio
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Recurso").Preload("Recurso.RecursoView").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolAnuncioPOST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.ToolAnuncio
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolAnuncioUpdate(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolAnuncio
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

func ToolAnuncioDelete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolAnuncio
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


//ToolExamen
func ToolExamen2GET(c *gin.Context) {
	var lis []models.ToolExamen2
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Find(&lis)
	conn.Preload("Recurso").Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "Lista de las ToolExamen2s",
		"r":   lis,
	})
}
func ToolExamen2GETID(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolExamen2
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Preload("Recurso.TipoRecurso").Preload("Recurso.RecursoView").Find(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolExamen2POST(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	var d models.ToolExamen2
	if err := c.BindJSON(&d); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func ToolExamen2Update(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolExamen2
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

func ToolExamen2Delete(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.ToolExamen2
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
