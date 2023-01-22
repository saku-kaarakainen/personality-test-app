package server

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/saku-kaarakainen/personality-test-app/api/config"
	"github.com/saku-kaarakainen/personality-test-app/api/db"
	"github.com/saku-kaarakainen/personality-test-app/api/services"
)

type IServer interface {
	SetupCors()
	SetRoutes()
	Run()
}

type Server struct {
	db              db.IDb
	router          *gin.Engine
	cfg             *config.Config
	questionService *services.Question
}

func NewServer(
	db db.IDb,
	router *gin.Engine,
	cfg *config.Config,
	questionService *services.Question,
) *Server {
	return &Server{
		db:              db,
		router:          router,
		cfg:             cfg,
		questionService: questionService,
	}
}

// Utils
func abortMsg(code int, err error, c *gin.Context) {
	c.String(code, "An error occurred")
	// A custom error page with HTML templates can be shown by c.HTML()
	c.Error(err)
	c.Abort()
}

// Interface methods
func (s *Server) SetupCors() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = s.cfg.Server.AllowOrigins
	s.router.Use(cors.New(corsConfig))
}

func (s *Server) SetRoutes() {
	s.router.GET("/ping", func(ctx *gin.Context) {
		// It is more approariate to put the func into it's own file, 'routes/ping.go'.
		// However this goes easily into very big rabbit hole with better framework, or better use of it.
		// Right now it's better to keep the code simpler and leave this as-is.
		ctx.String(200, "pong")
	})
	s.router.GET("/questions", func(ctx *gin.Context) {
		data, err := s.questionService.GetQuestions()
		if err != nil {
			log.Println("Error getting value:", err)
			abortMsg(500, err, ctx)
			return
		}

		log.Println("Return data:")
		log.Println(data)
		ctx.JSON(200, data)
	})

	s.router.GET("/result/calculate", func(ctx *gin.Context) {

	})
}

func (s *Server) Run() error {
	return s.router.Run(s.cfg.Server.Addr)
}
