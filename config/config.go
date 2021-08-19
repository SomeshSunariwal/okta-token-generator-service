package config

import "os"

var (
	HOST          = "Your_HOST"
	CLIENT_ID     = "Your_App_Client_Id"
	CLIENT_SECRET = "Your_App_Client_Secret"
	SSWS_KEY 	  = "YOUR_API_SECRET_Key"
)

type Config struct {
	HOST          string
	CLIENT_ID     string
	CLIENT_SECRET string
	SSWS_KEY      string
}

func ConfigGenerator() Config {
	config := Config{}
	host := os.Getenv("HOST")
	if host == "" {
		config.HOST = HOST
	}
	config.HOST = host

	clientId := os.Getenv("CLIENT_ID")
	if clientId == "" {
		config.HOST = CLIENT_ID
	}
	config.CLIENT_ID = clientId

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		config.HOST = CLIENT_SECRET
	}
	config.CLIENT_SECRET = clientSecret
	
	sswsKey := os.Getenv("SSWS_KEY")
	if sswsKey == "" {
		config.HOST = SSWS_KEY
	}
	config.CLIENT_SECRET = sswsKey

    return config
}