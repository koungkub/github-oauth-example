package main

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"

	"github.com/spf13/viper"
)

var (
	githubOAuth = oauth2.Config{
		ClientID:     viper.GetString("ID"),
		ClientSecret: viper.GetString("SECRET"),
		RedirectURL:  viper.GetString("REDIRECTURL"),
		Scopes:       []string{},
		Endpoint:     github.Endpoint,
	}
)

func init() {
	viper.SetConfigName("env")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

func main() {
	fmt.Println(viper.GetString("ID"))
}
