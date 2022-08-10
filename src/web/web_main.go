package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type User struct {
	Id        string
	AddressID string
}

func main() {

	server := NewServer()

	server.HandleFunc("GET", "/public", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, "public!")
	})
	server.HandleFunc("GET", "/", func(c *Context) {
		c.RenderTemplate("/public/index.html",
			map[string]interface{}{"time": time.Now()})
	})
	server.HandleFunc("GET", "/users/:id", func(c *Context) {
		u := User{Id: c.Params["id"].(string)}
		c.RenderJson(u)
	})

	server.HandleFunc("POST", "/users", func(c *Context) {
		fmt.Fprintln(c.ResponseWriter, c.Params)
	})

	server.HandleFunc("GET", "/login", func(c *Context) {
		c.RenderTemplate("/public/login.html",
			map[string]interface{}{"message": "Login is needed"})
	})

	server.HandleFunc("POST", "/login", func(c *Context) {
		if CheckLogin(c.Params["username"].(string), c.Params["password"].(string)) {
			http.SetCookie(c.ResponseWriter, &http.Cookie{
				Name:  "X_AUTH",
				Value: Sign(VerifyMessage),
				Path:  "/",
			})
			c.Redirect("/")
		}

		c.RenderTemplate("/public/login.html", map[string]interface{}{"message": "id or password is invalid."})
	})

	server.Use(AuthHandler)

	server.Run(":8080")
}

const VerifyMessage = "verifed"

func AuthHandler(next HandlerFunc) HandlerFunc {
	ignore := []string{"/login", "public/index.html"}
	return func(c *Context) {
		for _, s := range ignore {
			if strings.HasPrefix(c.Request.URL.Path, s) {
				next(c)
				return
			}
		}

		if v, err := c.Request.Cookie("X_AUTH"); err == http.ErrNoCookie {
			c.Redirect("/login")
			return
		} else if err != nil {
			c.RenderErr(http.StatusInternalServerError, err)
			return
		} else if Verify(VerifyMessage, v.Value) {
			next(c)
			return
		}

		c.Redirect("/login")
	}
}

func Verify(message, sig string) bool {
	return hmac.Equal([]byte(sig), []byte(Sign(message)))
}

func CheckLogin(username, password string) bool {
	const (
		USERNAME = "tester"
		PASSWORD = "1234"
	)

	return username == USERNAME && password == PASSWORD
}

func Sign(message string) string {
	secretKey := []byte("secretkey")
	if len(secretKey) == 0 {
		return " "
	}

	mac := hmac.New(sha1.New, secretKey)
	io.WriteString(mac, message)
	return hex.EncodeToString(mac.Sum(nil))
}
