package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	HOST     = "https://dev-852842.okta.com"
	CLIENT_ID = "0oa43rs29g4wXhT804x7"
	CLIENT_SECRET = "saqxXSxdT8RKlL1YxoMJpzbQbXVYlrUvHyaQedQc"
)

func main() {
	e := echo.New()
	
	AuthToken := CLIENT_ID + ":" + CLIENT_SECRET

	// Require Parameter For Heroku to start service on some port
	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}


	Base64Token := b64.StdEncoding.EncodeToString([]byte(AuthToken))

	e.POST("/token", func(c echo.Context) error {

		var RequestData map[string]interface{}

		receivedJSON, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			e.Logger.Info("Error: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"message" : "Unable to Read Request body"})
		}
		defer c.Request().Body.Close()

		err = json.Unmarshal([]byte(receivedJSON), &RequestData)
		if err != nil {
			e.Logger.Info("Error: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"message" : "Bad Request Body"})
		}
		Code := RequestData["code"]
		if Code == nil {
			e.Logger.Info("Error: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"message" : "Authorization Code Not Found"})
		}

		code := Code.(string)

		/// Creating application/x-www-form-urlencoded type request parameter
		form := url.Values{}
		form.Add("grant_type", "authorization_code")
		form.Add("code", code)
		form.Add("redirect_uri", "http://localhost:3000")
		
		//Creating Request 
		req, err := http.NewRequest("POST", HOST+"/oauth2/default/v1/token", strings.NewReader(form.Encode()))
		if err != nil {
			e.Logger.Info("Error: %v", err)
			return c.JSON(http.StatusBadRequest, map[string]string{"message" : "Request Errorr"})
		}

		// Attaching Necessary Headers to make request
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Basic "+Base64Token)

		// Hitting the web server
		respNew, _ := http.DefaultClient.Do(req)

		body, err := ioutil.ReadAll(respNew.Body)
		if err != nil {
			e.Logger.Info("Error: %v", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message" : "No Response Error"})
		}

		var ResponseData map[string]interface{}
		err = json.Unmarshal(body, &ResponseData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message" : "Unmarshal Error"})
		}
		return c.JSON(http.StatusOK, ResponseData)
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// Add your host URL Here from which you want to hit this service API.
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
	}))

	e.Logger.Fatal(e.Start(":"+PORT))
}
