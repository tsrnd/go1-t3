package main

import (
	"net/http"
	"os"
	"encoding/json"
	route "github.com/goweb3/app/routers"
	"github.com/goweb3/app/shared/jsonconfig"
	"github.com/goweb3/app/shared/view"
	"github.com/goweb3/app/shared/database"
)
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database database.Info	  `json:"Database"`
	Template  view.Template   `json:"Template"`
	View      view.View       `json:"View"`
}
func main() {
	// Load the configuration file
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)
	// database.Connect(config.Database)
	// Setup the views
	view.Configure(config.View)
	view.LoadTemplates(config.Template.Root, config.Template.Children)
	
	r := route.Init()
	
	http.ListenAndServe(":8080", r)
	
}
// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

