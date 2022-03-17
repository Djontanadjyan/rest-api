package app

import (
	"api/configs"
	"api/internal/controller/web"
	"api/internal/controller/web/middleware"
	"api/internal/infrastructure/repository"
	"api/internal/usecase"
	"api/pkg/httpserver"
	"api/pkg/mongo"

	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const db = "userdb"
const collection = "users"

func Run() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := configs.GetConfig()

	fmt.Println(config.Env)
	fmt.Println(config.Port)
	fmt.Println(config.MongoDB)

	conn := mongo.New(config.MongoDB.URI, db, collection)

	rp := repository.New(conn)

	us := usecase.New(rp)

	c := gin.Default()

	middle := middleware.InitMiddleware()

	c.Use(middle.CORS())

	web.NewRouter(c, us)

	httpServer := httpserver.New(c, httpserver.Port(config.Port.PORT))
	log.Printf("Start Server ...  ")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Println("app - Run - signal: " + s.String())
		log.Println("Shutdown Server ... ")
	case err := <-httpServer.Notify():
		log.Printf("app - Run - httpServer.Notify: %w", err)
		log.Println("Shutdown Server ... ")
	}

	err := httpServer.Shutdown()
	if err != nil {
		log.Printf("app - Run - httpServer.Shutdown: %w", err)
		log.Println("Shutdown Server ... ")
	}
	log.Println("Server exiting")

}
