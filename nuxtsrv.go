package nuxtsrv

import (
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func RegisterNuxtHandlers(e *echo.Echo, frontendDir string) {
	e.Static("/_nuxt", filepath.Join(frontendDir, "_nuxt"))

	var routes = map[string]string{
		"/*":                      "200.html",
		"/favicon.ico":            "favicon.ico",
		"/favicon.svg":            "favicon.svg",
		"/icon-logo-not-name.svg": "icon-logo-not-name.svg",
		"/icon-logo-text.svg":     "icon-logo-text.svg",
		"/sw.js":                  "sw.js",
		"/custom-sw.js":           "custom-sw.js",
	}

	for route, fname := range routes {
		e.File(route, filepath.Join(frontendDir, fname))
	}
}

func RegisterStaticHandlers(e *echo.Echo, frontendDir string) {
	e.Static("/", frontendDir)
	e.File("/*", filepath.Join(frontendDir, "index.html"))
}
