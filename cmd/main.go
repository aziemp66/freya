package main

import (
	"embed"

	"github.com/aziemp66/freya-be/common/env"
	"github.com/gin-gonic/gin"
)

//go:embed common/mail/templates/*.html
var mailTemplates embed.FS

func main() {
	router := gin.Default()
	cfg := env.LoadConfig()

	router.Run(":" + cfg.Port)
}
