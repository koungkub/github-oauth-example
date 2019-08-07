package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/random"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("env")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.AutomaticEnv()
}

func main() {

	e := echo.New()

	githubOAuth := oauth2.Config{
		ClientID:     viper.GetString("ID"),
		ClientSecret: viper.GetString("SECRET"),
		RedirectURL:  viper.GetString("REDIRECTURL"),
		Scopes: []string{
			"repo",
			"repo:status",
		},
		Endpoint: github.Endpoint,
	}

	state := random.String(8)

	e.Use(
		middleware.BodyDump(func(c echo.Context, req, res []byte) {
			fmt.Print("req")
			fmt.Println(string(req))
			fmt.Print("res")
			fmt.Println(string(res))
		}),
	)

	e.GET("/", func(c echo.Context) error {
		html := `
			<a href="/oauth/github">login github here !!</a>	
		`
		return c.HTML(http.StatusOK, html)
	})

	e.GET("/oauth/github", func(c echo.Context) error {
		// allowSignup := oauth2.SetAuthURLParam("allow_signup", "true")
		// url := githubOAuth.AuthCodeURL(code, allowSignup)
		url := githubOAuth.AuthCodeURL(state)
		return c.Redirect(http.StatusTemporaryRedirect, url)
	})

	e.GET("/oauth/callback", func(c echo.Context) error {
		code := c.QueryParam("code")
		newState := c.QueryParam("state")

		fmt.Println(code, newState)

		html := `
			<h1>koung</h1>
		`
		return c.HTML(http.StatusOK, html)
	})

	log.Fatal(e.Start(":1323"))
}
