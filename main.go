package main

import (
	"net/http"
	"os"

	"github.com/SomeshSunariwal/okta-token-generator-service/api"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MainHandler struct {
	apiHandler api.Handler
}

func main() {
	e := echo.New()

	mainHandler := MainHandler{}

	// Require Parameter For Heroku to start service on some port
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	// POST Request to get okta token using auth code
	e.POST("/token", mainHandler.apiHandler.Token)

	// Revoke All Consent of a user
	e.POST("/revoke-all", mainHandler.apiHandler.RevokeAllGrant)

	e.POST("/user", mainHandler.apiHandler.GetUserId)

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK,
			`
			<div>
				<h1>Hello There !</h1>
			</div>
		`)
	})

	// Adding CORS to allow web browser to hit this service endpoint
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// Add your host URL Here from which you want to hit this service API.
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "{'time':'${time_rfc3339}', 'method'='${method}', 'host'='${host}', 'remote_ip'='${remote_ip}', 'uri'='${uri}', 'latency'='${latency_human}', 'status'='${status}'}\n"}))

	e.Start(":" + PORT)
}
