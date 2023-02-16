package main

import (
	"github.com/aziemp66/freya-be/common/env"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cfg := env.LoadConfig()

	router.Run(":" + cfg.Port)
}
