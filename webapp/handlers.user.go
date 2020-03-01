// handlers.user.go

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/vault/api"
)

var reNewline = regexp.MustCompile(`\n+`)

func vaultUserpassLogin(c *api.Client, mount string, username string, password string) bool {
	// Create payload
	options := map[string]interface{}{
		"password": password,
	}
	path := fmt.Sprintf("auth/%s/login/%s", mount, username)

	// PUT call to get a token
	secret, err := c.Logical().Write(path, options)
	if err != nil {
		//log.Printf("ERROR: %v", err.Error())
		errMsg := reNewline.ReplaceAllString(err.Error(), " ")
		log.Printf("error from vault: [%v]", errMsg)
		return false
	}

	token := secret.Auth.ClientToken
	c.SetToken(token)
	c.Auth().Token().RevokeSelf(token)

	return true
}

func showLoginPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{"title": "Login"}, "login.html")
}

func performLogin(c *gin.Context) {
	// Obtain the POSTed username and password values
	username := c.PostForm("username")
	password := c.PostForm("password")

	var sameSiteCookie http.SameSite

	// Perform vault userpass auth
	successLogin := vaultUserpassLogin(vaultClient, *vaultMountPoint, username, password)

	// Check if vault login was successful
	if successLogin == true {
		// If the username/password is valid generate and set a token in a cookie
		token := generateSessionToken()
		c.SetCookie("token", token, 3600, "", "", sameSiteCookie, false, true)
		c.Set("is_logged_in", true)

		render(c, gin.H{"title": "Successful Login"}, "login-successful.html")

	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func generateSessionToken() string {
	// We're using a random 16 character string as the session token
	// This is NOT a secure way of generating session tokens
	// DO NOT USE THIS IN PRODUCTION
	return strconv.FormatInt(rand.Int63(), 16)
}

func logout(c *gin.Context) {
	var sameSiteCookie http.SameSite

	// Clear the cookie
	c.SetCookie("token", "", -1, "", "", sameSiteCookie, false, true)

	// Redirect to the home page
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
