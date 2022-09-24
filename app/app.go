package app

import (
	"log"
	"os"
	"test/domain"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	r = gin.Default()
)

func StartApp() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT := os.Getenv("PORT")

	Routers()
	domain.ConnDB()
	r.Run(PORT)
}
