package tests

import (
	"encoding/json"
	"os"

	"github.com/goweb3/app/shared/database"
	"github.com/goweb3/app/shared/jsonconfig"
	"github.com/goweb3/app/shared/server"
	"github.com/goweb3/app/shared/view"
)

var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	Database database.Info `json:"Database"`
	Template view.Template `json:"Template"`
	View     view.View     `json:"View"`
	Server   server.Server `json:"Server"`
}

func Init() {
	// Load the configuration file
	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)
	database.Connect(config.Database)
	// Setup the views
	view.Configure(config.View)
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}
