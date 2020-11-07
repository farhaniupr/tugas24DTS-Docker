package controller

import (
	"log"
	"tugas/model/mAccounts"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MysqlDsn = `root:root@tcp(172.19.0.10:3306)/mahasiswa?parseTime=True&charset=utf8`
)

type Server struct {
	Router    *gin.Engine
	App       *firebase.App
	Firestore *firestore.Client
	Db        *gorm.DB
}

func InitServer() *Server {

	db, err := gorm.Open(mysql.Open(MysqlDsn), &gorm.Config{})
	if err != nil {
		log.Fatalf(`failed to connect to mysql: ` + err.Error())
	}
	db.AutoMigrate(&mAccounts.Account{})
	s := &Server{}
	s.Router = gin.Default()
	s.Db = db

	cfg := cors.DefaultConfig()
	cfg.AllowAllOrigins = true
	cfg.AllowCredentials = true
	cfg.AllowMethods = []string{"GET", "POST"}
	cfg.AllowHeaders = []string{"Authorization", "Origin", "Accept", "X-Requested-With", " Content-Type", "Access-Control-Request-Method", "Access-Control-Request-Headers"}
	s.Router.Use(cors.New(cfg))

	s.Router.LoadHTMLGlob("view/*")
	return s
}

func (s *Server) Listen(port string) {
	s.Router.Run(port)
}

func (s *Server) AssignHandler(route string, handler Handler) {
	s.Router.Any(route, func(context *gin.Context) {
		handler(&Ctx{
			Server:  s,
			Context: context,
		})
	})
}
