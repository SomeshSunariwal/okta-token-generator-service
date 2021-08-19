package api

import (
	b64 "encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/SomeshSunariwal/okta-token-generator-service/config"
	"github.com/labstack/echo/v4"
)

type Handler struct {
}

func (handler Handler) Token(context echo.Context) error {

	AuthToken := config.CLIENT_ID + ":" + config.CLIENT_SECRET

	Base64Token := b64.StdEncoding.EncodeToString([]byte(AuthToken))

	var RequestData map[string]interface{}
		receivedJSON, err := ioutil.ReadAll(context.Request().Body)
		defer context.Request().Body.Close()

		if err != nil {
			context.Echo().Logger.Info("Error : %v", err)
			return context.JSON(http.StatusBadRequest, map[string]string{"message" : "Unable to Read Request body"})
		}

		err = json.Unmarshal([]byte(receivedJSON), &RequestData)
		if err != nil {
			context.Echo().Logger.Info("Error : %v", err)
			return context.JSON(http.StatusBadRequest, map[string]string{"message" : "Bad Request Body"})
		}

		Code := RequestData["code"]
		if Code == nil {
			context.Echo().Logger.Info("Error : %v", err)
			return context.JSON(http.StatusBadRequest, map[string]string{"message" : "Authorization Code Not Found"})
		}

		code := Code.(string)

		/// Creating application/x-www-form-urlencoded type request parameter
		form := url.Values{}
		form.Add("grant_type", "authorization_code")
		form.Add("code", code)
		form.Add("redirect_uri", "http://localhost:3000")
		
		//Creating Request 
		req, err := http.NewRequest("POST", config.HOST+ "/oauth2/default/v1/token", strings.NewReader(form.Encode()))
		if err != nil {
			context.Echo().Logger.Info("Error : %v", err)
			return context.JSON(http.StatusBadRequest, map[string]string{"message" : "Request Errorr"})
		}

		// Attaching Necessary Headers to make request
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Authorization", "Basic "+Base64Token)

		// Hitting the web server
		respNew, _ := http.DefaultClient.Do(req)

		// Reading Response
		body, err := ioutil.ReadAll(respNew.Body)
		if err != nil {
			context.Echo().Logger.Info("Error : %v", err)
			return context.JSON(http.StatusInternalServerError, map[string]string{"message" : "No Response Error"})
		}

		
		// Making Response Json.
		var ResponseData map[string]interface{}
		err = json.Unmarshal(body, &ResponseData)
		if err != nil {
			context.Echo().Logger.Info("Error : %v", err)
			return context.JSON(http.StatusInternalServerError, map[string]string{"message" : "Unmarshal Error"})
		}
		
		if respNew.StatusCode != http.StatusOK {
			context.Echo().Logger.Info("Error : %v", ResponseData)
			return context.JSON(respNew.StatusCode, ResponseData)
		}

		return context.JSON(http.StatusOK, ResponseData)
}