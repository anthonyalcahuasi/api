package routers

import (
	"net/http"
	"strings"

	"github.com/202lp2/go2/apis"
	"github.com/202lp2/go2/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	conn, err := connectDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
		//return
	}
	// Migrate the schema
	conn.AutoMigrate(
		&models.User{},

		&models.Rol{},
		&models.Estudiante{},
		&models.Docente{},

		&models.Curso{},
		&models.Nivel{},
		&models.Grado{},
		&models.Jerarquia{},
		&models.Plan{},

		&models.Grupo{},
		&models.Periodo{},
		&models.Sucursal{},
		&models.CargaDocente{},
		&models.Horario{},

		&models.Inscripcion{},
		&models.Unidad{},
		&models.Sesion{},

		&models.TipoRecurso{},
		&models.Recurso{},
		&models.RecursoView{},

		&models.ToolVideo{},
		&models.ToolEnlace{},
		&models.ToolAnuncio{},
		&models.ToolDocumento{},
		&models.ToolVideoConferencia{},

		&models.ToolExamenCategoria{},
		&models.ToolExamen{},
		&models.ToolExamenPublicacion{},
		&models.ToolExamenPregunta{},
		&models.ToolExamenAlternativa{},

		&models.ToolExamenDetalleUsuario{},
		&models.Supervicion{},

		&models.ToolExamen2{},

		/*

			&models.Trabajo{},
			&models.Foro{},
			&models.ToolForo{},
			&models.ToolActividad{},
		*/
	)
	r := gin.Default()
	r.Use(CORSMiddleware())

	r.Use(dbMiddleware(*conn))

	v1 := r.Group("/v1")

	{
		v1.GET("/users", apis.UsersIndex)
		v1.POST("/users", apis.UsersCreate)
		v1.GET("/users/:id", authMiddleWare(), apis.UsersGet)
		v1.PUT("/users/:id", authMiddleWare(), apis.UsersUpdate)
		v1.DELETE("/users/:id", authMiddleWare(), apis.UsersDelete)
		v1.POST("/login", apis.UsersLogin)
		v1.POST("/logout", apis.UsersLogout)

		v1.GET("/rol", authMiddleWare(), apis.RolGet)
		v1.POST("/rol", authMiddleWare(), apis.RolCreate)
		v1.GET("/rol/:id", authMiddleWare(), apis.RolGetID)
		v1.PUT("/rol/:id", authMiddleWare(), apis.RolUpdate)
		v1.DELETE("/rol/:id", authMiddleWare(), apis.RolDelete)

		v1.GET("/rolget/:nombre", authMiddleWare(), apis.RolGETNombre)

		v1.GET("/userrol", authMiddleWare(), apis.UserRolGET)
		v1.POST("/userrol", authMiddleWare(), apis.UserRolPOST)
		v1.GET("/userrol/:id", authMiddleWare(), apis.UserRolGETID)
		v1.PUT("/userrol/:id", authMiddleWare(), apis.UserRolUpdate)
		v1.DELETE("/userrol/:user_id/:rol_id", authMiddleWare(), apis.UserRolDelete)

		v1.GET("/estudiante", authMiddleWare(), apis.EstudianteGET)
		v1.GET("/estudiante/:id", authMiddleWare(), apis.EstudianteGETID)
		v1.PUT("/estudiante/:id", authMiddleWare(), apis.EstudianteUpdate)
		v1.DELETE("/estudiante/:id", authMiddleWare(), apis.EstudianteDelete)

		v1.GET("/docente", authMiddleWare(), apis.DocenteGET)
		v1.GET("/docente/:id", authMiddleWare(), apis.DocenteGETID)
		v1.PUT("/docente/:id", authMiddleWare(), apis.DocenteUpdate)
		v1.DELETE("/docente/:id", authMiddleWare(), apis.DocenteDelete)

		v1.GET("/curso", authMiddleWare(), apis.CursoIndex)
		v1.POST("/curso", authMiddleWare(), authMiddleWare(), apis.CursoCreate)
		v1.GET("/curso/:id", authMiddleWare(), apis.CursoGet)
		v1.PUT("/curso/:id", authMiddleWare(), apis.CursoUpdate)
		v1.DELETE("/curso/:id", authMiddleWare(), apis.CursoDelete)

		v1.GET("/sucursal", authMiddleWare(), apis.SucursalGET)
		v1.POST("/sucursal", authMiddleWare(), apis.SucursalPOST)
		v1.GET("/sucursal/:id", authMiddleWare(), apis.SucursalGETID)
		v1.PUT("/sucursal/:id", authMiddleWare(), apis.SucursalUpdate)
		v1.DELETE("/sucursal/:id", authMiddleWare(), apis.SucursalDelete)

		v1.GET("/nivel", authMiddleWare(), apis.NivelGET)
		v1.POST("/nivel", authMiddleWare(), apis.NivelPOST)
		v1.GET("/nivel/:id", authMiddleWare(), apis.NivelGETID)
		v1.PUT("/nivel/:id", authMiddleWare(), apis.NivelUpdate)
		v1.DELETE("/nivel/:id", authMiddleWare(), apis.NivelDelete)

		v1.GET("/grado", authMiddleWare(), apis.GradoGET)
		v1.POST("/grado", authMiddleWare(), apis.GradoPOST)
		v1.GET("/grado/:id", authMiddleWare(), apis.GradoGETID)
		v1.PUT("/grado/:id", authMiddleWare(), apis.GradoUpdate)
		v1.DELETE("/grado/:id", authMiddleWare(), apis.GradoDelete)

		v1.GET("/jerarquia", authMiddleWare(), apis.JerarquiaGET)
		v1.POST("/jerarquia", authMiddleWare(), apis.JerarquiaPOST)
		v1.GET("/jerarquia/:id", authMiddleWare(), apis.JerarquiaGETID)
		v1.PUT("/jerarquia/:id", authMiddleWare(), apis.JerarquiaUpdate)
		v1.DELETE("/jerarquia/:id", authMiddleWare(), apis.JerarquiaDelete)

		v1.GET("/plan", authMiddleWare(), apis.PlanGET)
		v1.POST("/plan", authMiddleWare(), apis.PlanPOST)
		v1.GET("/plan/:id", authMiddleWare(), apis.PlanGETID)
		v1.PUT("/plan/:id", authMiddleWare(), apis.PlanUpdate)
		v1.DELETE("/plan/:id", authMiddleWare(), apis.PlanDelete)

		v1.GET("/cargadocente", authMiddleWare(), apis.CargaDocenteGET)
		v1.POST("/cargadocente", authMiddleWare(), apis.CargaDocentePOST)
		v1.GET("/cargadocente/:id", authMiddleWare(), apis.CargaDocenteGETID)
		v1.PUT("/cargadocente/:id", authMiddleWare(), apis.CargaDocenteUpdate)
		v1.DELETE("/cargadocente/:id", authMiddleWare(), apis.CargaDocenteDelete)

		v1.GET("/cargadocentecodigo/:codigo", authMiddleWare(), apis.CargaDocenteGETCODIGO)

		v1.GET("/horario", authMiddleWare(), apis.HorarioGET)
		v1.POST("/horario", authMiddleWare(), apis.HorarioPOST)
		v1.GET("/horario/:id", authMiddleWare(), apis.HorarioGETID)
		v1.PUT("/horario/:id", authMiddleWare(), apis.HorarioUpdate)
		v1.DELETE("/horario/:id", authMiddleWare(), apis.HorarioDelete)

		v1.GET("/inscripcion", authMiddleWare(), apis.InscripcionGET)
		v1.POST("/inscripcion", authMiddleWare(), apis.InscripcionPOST)
		v1.GET("/inscripcion/:id", authMiddleWare(), apis.InscripcionGETID)
		v1.PUT("/inscripcion/:id", authMiddleWare(), apis.InscripcionUpdate)
		v1.DELETE("/inscripcion/:id", authMiddleWare(), apis.InscripcionDelete)

		v1.GET("/grupo", authMiddleWare(), apis.GrupoGET)
		v1.POST("/grupo", authMiddleWare(), apis.GrupoPOST)
		v1.GET("/grupo/:id", authMiddleWare(), apis.GrupoGETID)
		v1.PUT("/grupo/:id", authMiddleWare(), apis.GrupoUpdate)
		v1.DELETE("/grupo/:id", authMiddleWare(), apis.GrupoDelete)

		v1.GET("/periodo", authMiddleWare(), apis.PeriodoGET)
		v1.POST("/periodo", authMiddleWare(), authMiddleWare(), apis.PeriodoPOST)
		v1.GET("/periodo/:id", authMiddleWare(), apis.PeriodoGETID)
		v1.PUT("/periodo/:id", authMiddleWare(), apis.PeriodoUpdate)
		v1.DELETE("/periodo/:id", authMiddleWare(), apis.PeriodoDelete)

		v1.GET("/unidad", authMiddleWare(), apis.UnidadGET)
		v1.POST("/unidad", authMiddleWare(), apis.UnidadPOST)
		v1.GET("/unidad/:id", authMiddleWare(), apis.UnidadGETID)
		v1.PUT("/unidad/:id", authMiddleWare(), apis.UnidadUpdate)
		v1.DELETE("/unidad/:id", authMiddleWare(), apis.UnidadDelete)

		v1.GET("/sesion", authMiddleWare(), apis.SesionGET)
		v1.POST("/sesion", authMiddleWare(), apis.SesionPOST)
		v1.GET("/sesion/:id", authMiddleWare(), apis.SesionGETID)
		v1.PUT("/sesion/:id", authMiddleWare(), apis.SesionUpdate)
		v1.DELETE("/sesion/:id", authMiddleWare(), apis.SesionDelete)

		v1.GET("/tiporecurso", authMiddleWare(), apis.TipoRecursoGET)
		v1.POST("/tiporecurso", authMiddleWare(), apis.TipoRecursoPOST)
		v1.GET("/tiporecurso/:id", authMiddleWare(), apis.TipoRecursoGETID)
		v1.PUT("/tiporecurso/:id", authMiddleWare(), apis.TipoRecursoUpdate)
		v1.DELETE("/tiporecurso/:id", authMiddleWare(), apis.TipoRecursoDelete)

		v1.GET("/tiporecursonombre/:nombre", authMiddleWare(), apis.TipoRecursoGETNOMBRE)

		v1.GET("/recurso", authMiddleWare(), apis.RecursoGET)
		v1.POST("/recurso", authMiddleWare(), apis.RecursoPOST)
		v1.GET("/recurso/:id", authMiddleWare(), apis.RecursoGETID)
		v1.PUT("/recurso/:id", authMiddleWare(), apis.RecursoUpdate)
		v1.DELETE("/recurso/:id", authMiddleWare(), apis.RecursoDelete)

		v1.GET("/recursoview", authMiddleWare(), apis.RecursoViewGET)
		v1.POST("/recursoview", authMiddleWare(), apis.RecursoViewPOST)
		v1.GET("/recursoview/:id", authMiddleWare(), apis.RecursoViewGETID)
		v1.PUT("/recursoview/:id", authMiddleWare(), apis.RecursoViewUpdate)
		v1.DELETE("/recursoview/:id", authMiddleWare(), apis.RecursoViewDelete)

		v1.GET("/toolvideo", authMiddleWare(), apis.ToolVideoGET)
		v1.POST("/toolvideo", authMiddleWare(), apis.ToolVideoPOST)
		v1.GET("/toolvideo/:id", authMiddleWare(), apis.ToolVideoGETID)
		v1.PUT("/toolvideo/:id", authMiddleWare(), apis.ToolVideoUpdate)
		v1.DELETE("/toolvideo/:id", authMiddleWare(), apis.ToolVideoDelete)

		v1.GET("/toolenlace", authMiddleWare(), apis.ToolEnlaceGET)
		v1.POST("/toolenlace", authMiddleWare(), apis.ToolEnlacePOST)
		v1.GET("/toolenlace/:id", authMiddleWare(), apis.ToolEnlaceGETID)
		v1.PUT("/toolenlace/:id", authMiddleWare(), apis.ToolEnlaceUpdate)
		v1.DELETE("/toolenlace/:id", authMiddleWare(), apis.ToolEnlaceDelete)

		v1.GET("/toolanuncio", authMiddleWare(), apis.ToolAnuncioGET)
		v1.POST("/toolanuncio", authMiddleWare(), apis.ToolAnuncioPOST)
		v1.GET("/toolanuncio/:id", authMiddleWare(), apis.ToolAnuncioGETID)
		v1.PUT("/toolanuncio/:id", authMiddleWare(), apis.ToolAnuncioUpdate)
		v1.DELETE("/toolanuncio/:id", authMiddleWare(), apis.ToolAnuncioDelete)

		v1.GET("/tooldocumento", authMiddleWare(), apis.ToolDocumentoGET)
		v1.POST("/tooldocumento", authMiddleWare(), apis.ToolDocumentoPOST)
		v1.GET("/tooldocumento/:id", authMiddleWare(), apis.ToolDocumentoGETID)
		v1.PUT("/tooldocumento/:id", authMiddleWare(), apis.ToolDocumentoUpdate)
		v1.DELETE("/tooldocumento/:id", authMiddleWare(), apis.ToolDocumentoDelete)

		v1.GET("/toolvideoconferencia", authMiddleWare(), apis.ToolVideoConferenciaGET)
		v1.POST("/toolvideoconferencia", authMiddleWare(), apis.ToolVideoConferenciaPOST)
		v1.GET("/toolvideoconferencia/:id", authMiddleWare(), apis.ToolVideoConferenciaGETID)
		v1.PUT("/toolvideoconferencia/:id", authMiddleWare(), apis.ToolVideoConferenciaUpdate)
		v1.DELETE("/toolvideoconferencia/:id", authMiddleWare(), apis.ToolVideoConferenciaDelete)

		v1.GET("/toolexamen", authMiddleWare(), apis.ToolExamenGET)
		v1.POST("/toolexamen", authMiddleWare(), apis.ToolExamenPOST)
		v1.GET("/toolexamen/:id", authMiddleWare(), apis.ToolExamenGETID)
		v1.PUT("/toolexamen/:id", authMiddleWare(), apis.ToolExamenUpdate)
		v1.DELETE("/toolexamen/:id", authMiddleWare(), apis.ToolExamenDelete)

		v1.GET("/toolexamencategoria", authMiddleWare(), apis.ToolExamenCategoriaGET)
		v1.POST("/toolexamencategoria", authMiddleWare(), apis.ToolExamenCategoriaPOST)
		v1.GET("/toolexamencategoria/:id", authMiddleWare(), apis.ToolExamenCategoriaGETID)
		v1.PUT("/toolexamencategoria/:id", authMiddleWare(), apis.ToolExamenCategoriaUpdate)
		v1.DELETE("/toolexamencategoria/:id", authMiddleWare(), apis.ToolExamenCategoriaDelete)

		v1.GET("/toolexamenpregunta", authMiddleWare(), apis.ToolExamenPreguntaGET)
		v1.POST("/toolexamenpregunta", authMiddleWare(), apis.ToolExamenPreguntaPOST)
		v1.GET("/toolexamenpregunta/:id", authMiddleWare(), apis.ToolExamenPreguntaGETID)
		v1.PUT("/toolexamenpregunta/:id", authMiddleWare(), apis.ToolExamenPreguntaUpdate)
		v1.DELETE("/toolexamenpregunta/:id", authMiddleWare(), apis.ToolExamenPreguntaDelete)

		v1.GET("/toolexamenalternativa", authMiddleWare(), apis.ToolExamenAlternativaGET)
		v1.POST("/toolexamenalternativa", authMiddleWare(), apis.ToolExamenAlternativaPOST)
		v1.GET("/toolexamenalternativa/:id", authMiddleWare(), apis.ToolExamenAlternativaGETID)
		v1.PUT("/toolexamenalternativa/:id", authMiddleWare(), apis.ToolExamenAlternativaUpdate)
		v1.DELETE("/toolexamenalternativa/:id", authMiddleWare(), apis.ToolExamenAlternativaDelete)

		v1.GET("/toolexamenpublicacion", authMiddleWare(), apis.ToolExamenPublicacionGET)
		v1.POST("/toolexamenpublicacion", authMiddleWare(), apis.ToolExamenPublicacionPOST)
		v1.GET("/toolexamenpublicacion/:id", authMiddleWare(), apis.ToolExamenPublicacionGETID)
		v1.PUT("/toolexamenpublicacion/:id", authMiddleWare(), apis.ToolExamenPublicacionUpdate)
		v1.DELETE("/toolexamenpublicacion/:id", authMiddleWare(), apis.ToolExamenPublicacionDelete)

		v1.GET("/toolexamendetalles", authMiddleWare(), apis.ToolExamenDetalleUsuarioGET)
		v1.POST("/toolexamendetalles", authMiddleWare(), apis.ToolExamenDetalleUsuarioPOST)
		v1.GET("/toolexamendetalles/:id", authMiddleWare(), apis.ToolExamenDetalleUsuarioGETID)
		v1.PUT("/toolexamendetalles/:id", authMiddleWare(), apis.ToolExamenDetalleUsuarioUpdate)
		v1.DELETE("/toolexamendetalles/:id", authMiddleWare(), apis.ToolExamenDetalleUsuarioDelete)

		v1.GET("/toolexamensupervicion", authMiddleWare(), apis.SupervicionGET)
		v1.POST("/toolexamensupervicion", authMiddleWare(), apis.SupervicionPOST)
		v1.GET("/toolexamensupervicion/:id", authMiddleWare(), apis.SupervicionGETID)
		v1.PUT("/toolexamensupervicion/:id", authMiddleWare(), apis.SupervicionUpdate)
		v1.DELETE("/toolexamensupervicion/:id", authMiddleWare(), apis.SupervicionDelete)

		v1.GET("/toolexamen2", authMiddleWare(), apis.ToolExamen2GET)
		v1.POST("/toolexamen2", authMiddleWare(), apis.ToolExamen2POST)
		v1.GET("/toolexamen2/:id", authMiddleWare(), apis.ToolExamen2GETID)
		v1.PUT("/toolexamen2/:id", authMiddleWare(), apis.ToolExamen2Update)
		v1.DELETE("/toolexamen2/:id", authMiddleWare(), apis.ToolExamen2Delete)

		/*

			v1.GET("/trabajo", apis.TrabajoGET)
			v1.POST("/trabajo", authMiddleWare(), apis.TrabajoPOST)
			v1.GET("/trabajo/:id", apis.TrabajoGETID)
			v1.PUT("/trabajo/:id", apis.TrabajoUpdate)
			v1.DELETE("/trabajo/:id", apis.TrabajoDelete)

			v1.GET("/foro", apis.ForoGET)
			v1.POST("/foro", authMiddleWare(), apis.ForoPOST)
			v1.GET("/foro/:id", apis.ForoGETID)
			v1.PUT("/foro/:id", apis.ForoUpdate)
			v1.DELETE("/foro/:id", apis.ForoDelete)

			v1.GET("/toolforo", apis.ToolForoGET)
			v1.POST("/toolforo", authMiddleWare(), apis.ToolForoPOST)
			v1.GET("/toolforo/:id", apis.ToolForoGETID)
			v1.PUT("/toolforo/:id", apis.ToolForoUpdate)
			v1.DELETE("/toolforo/:id", apis.ToolForoDelete)

			v1.GET("/toolactividad", apis.ToolActividadGET)
			v1.POST("/toolactividad", authMiddleWare(), apis.ToolActividadPOST)
			v1.GET("/toolactividad/:id", apis.ToolActividadGETID)
			v1.PUT("/toolactividad/:id", apis.ToolActividadUpdate)
			v1.DELETE("/toolactividad/:id", apis.ToolActividadDelete)


		*/
	}
	return r
}

func connectDBmysql() (c *gorm.DB, err error) {
	dsn := "docker:docker@tcp(mysql-db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "docker:docker@tcp(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la Basde de Datos de Mysql: " + err.Error())
	}
	return conn, err
}

func connectDB() (c *gorm.DB, err error) {

	//dsn := "user=kcgewoajynifrc password=5476e5c97df90634425f6581d982b3c286359fe3906113d995837dde89758378 host=ec2-3-234-131-8.compute-1.amazonaws.com dbname=d7il2hpv5f4bsm port=5432 sslmode=require TimeZone=Asia/Shanghai"
	dsn := "user=postgres password=123456789 host=gek-api.crf3g3ly9ukt.us-east-1.rds.amazonaws.com dbname=gekapi port=5432 sslmode=require TimeZone=Asia/Shanghai"
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la Basde de Datos de Postgress: " + err.Error())
	}
	return conn, err
}

func dbMiddleware(conn gorm.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE ")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func authMiddleWare() gin.HandlerFunc { //ExtractToken
	return func(c *gin.Context) {
		bearer := c.Request.Header.Get("Authorization")
		split := strings.Split(bearer, "Bearer ")
		if len(split) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated."})
			c.Abort()
			return
		}
		token := split[1]
		isValid, userID := models.IsTokenValid(token)
		if isValid == false {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated (IsTokenValid)."})
			c.Abort()
		} else {
			c.Set("user_id", userID)
			c.Next()
		}
	}
}
