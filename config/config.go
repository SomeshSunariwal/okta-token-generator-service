package config

import "os"

var (
	HOST          = "https://dev-852842.okta.com"
	CLIENT_ID     = "0oa43rs29g4123wXhT804x7"
	CLIENT_SECRET = "saqxXSxdT8RK123lL1YxoMJpzbQbXVYlrUvHyaQedQc"
)

type Config struct {
	HOST          string
	CLIENT_ID     string
	CLIENT_SECRET string
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

    return config
}