package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/zhuharev/nuxtsrv"
)

func main() {
	e := echo.New()

	port := "8008"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	frontendDir := "frontend"
	if fd := os.Getenv("FRONTEND_DIR"); fd != "" {
		frontendDir = fd
	}

	log.Printf("start listen on port=%s dir=%s", port, frontendDir)

	nuxtsrv.RegistreHandlers(e, frontendDir)

	e.Start(":" + port)
}
