package controller

import (
	"log"
	"tugas/model/mAccounts"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	MysqlDsn = `root:@/mahasiswa?parseTime=True&charset=utf8`
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
