// main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/vault/api"
)

var (
	router          *gin.Engine
	vaultServer     = flag.String("vaultip", "127.0.0.1", "(optional) IP of vault server")
	vaultPort       = flag.Int("vaultport", 8200, "(optional) TCP port of vault server")
	vaultMountPoint = flag.String("vaultmountpoint", "userpass", "(optional) userpass auth mount point")
	vaultClient     *api.Client
)

func main() {
	flag.Parse()

	// Setup initial vault config
	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("https://%s:%d", *vaultServer, *vaultPort)

	// Setup TLS config for vault
	t := &api.TLSConfig{
		Insecure: true,
	}
	config.ConfigureTLS(t)

	// Build the vault client
	var err error
	vaultClient, err = api.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("using vault at [%s]", config.Address)
	log.Printf("using vault userpass auth at mount [%s]", *vaultMountPoint)

	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Initialize the routes
	initializeRoutes()

	// Start serving the application
	router.Run()
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}
}
