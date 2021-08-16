package main

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	HOST     = "https://dev-852842.okta.com"
	USERNAME = "0oa43rs29g4wXhT804x7"
	PASSWORD = "saqxXSxdT8RKlL1YxoMJpzbQbXVYlrUvHyaQedQc"
)

func main() {
	e := echo.New()

	AuthToken := USERNAME + ":" + PASSWORD

	PORT := os.Getenv("PORT")
	if PORT == ""{
		PORT = "8080"
	}


	Base64Token := b64.StdEncoding.EncodeToString([]byte(AuthToken))

	e.POST("/token", func(c echo.Context) error {

		var RequestData map[string]interface{}

		receivedJSON, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}
		defer c.Request().Body.Close()

		err = json.Unmarshal([]byte(receivedJSON), &RequestData)
		if err != nil {
			return err
		}
		Code := RequestData["code"]
		if Code == nil {
			return errors.New("code not found")
		}
		code := Code.(string)

		form := url.Values{}
		form.Add("grant_type", "authorization_code")
		form.Add("code", code)
		form.Add("redirect_uri", "http://localhost:3000")
		req, err := http.NewRequest("POST", HOST+"/oauth2/default/v1/token", strings.NewReader(form.Encode()))
		if err != nil {
			e.Logger.Info("Error: %v", err)
			return err
		}

		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Basic "+Base64Token)

		respNew, _ := http.DefaultClient.Do(req)

		body, err := ioutil.ReadAll(respNew.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var ResponseData map[string]interface{}
		err = json.Unmarshal(body, &ResponseData)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, ResponseData)
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowCredentials: true,
	}))

	e.Logger.Fatal(e.Start(":"+PORT))
}
