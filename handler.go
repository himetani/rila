package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"

	"github.com/himetani/rila/pocket"
	"github.com/labstack/echo"
)

func top(c echo.Context) error {
	return c.String(http.StatusOK, c.QueryParam("msg"))
}

func auth(c echo.Context) error {
	consumerKey := c.QueryParam("consumer_key")
	if consumerKey == "" {
		return c.Redirect(http.StatusFound, "/?msg=consumerKey is not specified")
	}

	pocket := pocket.NewPocket()

	host := "http://" + c.Request().Host

	base, err := url.Parse(host)
	if err != nil {
		return err
	}
	base.Path = path.Join(base.Path, "pocket/redirect")
	redirect := base.String() + "?consumer_key=" + consumerKey

	code, err := pocket.GetRequestCode(consumerKey, redirect)
	if err != nil {
		return err
	}

	authURL := "https://getpocket.com/auth/authorize?request_token=" + code + "&redirect_uri=" + redirect

	requestCodes.Set(consumerKey, code)

	return c.Redirect(http.StatusFound, authURL)
}

func redirect(c echo.Context) error {
	consumerKey := c.QueryParam("consumer_key")
	if consumerKey == "" {
		return errors.New("consumer_key is empty")
	}

	pocket := pocket.NewPocket()
	username, token, err := pocket.GetAccessToken(consumerKey, requestCodes.Get(consumerKey))
	if err != nil {
		return err
	}

	fmt.Printf("Username: %s, Token: %s", username, token)
	return c.Redirect(http.StatusFound, "/pocket/token")

}

func token(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/?msg=pocket authorization was success")
}

func list(c echo.Context) error {
	return nil
}

func redirectToAuth(c echo.Context) error {
	return c.Redirect(http.StatusFound, "/pocket/auth")
}
