package view

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	service "github.com/goweb3/app/services"
	"github.com/goweb3/app/shared/flash"
	"github.com/jianfengye/web-golang/web/session"
)

var (
	viewInfo           View
	childTemplates     []string
	rootTemplate       string
	pluginCollection   = make(template.FuncMap)
	templateCollection = make(map[string]*template.Template)
	mutex              sync.RWMutex
	mutexPlugins       sync.RWMutex
)

type Template struct {
	Root     string   `json:"Root"`
	Children []string `json:"Children"`
}

type View struct {
	BaseURI   string
	Extension string
	Folder    string
	Name      string
	Caching   bool
	Vars      map[string]interface{}
	request   *http.Request
}

// Configure sets the view information
func Configure(vi View) {
	viewInfo = vi
}

// ReadConfig returns the configuration
func ReadConfig() View {
	return viewInfo
}

// LoadTemplates will set the root and child templates
func LoadTemplates(rootTemp string, childTemps []string) {
	rootTemplate = rootTemp
	childTemplates = childTemps
}

/**
*
* Constructor View
*
**/
func New(r *http.Request) *View {
	v := &View{}
	v.Vars = make(map[string]interface{})
	v.Vars["AuthLevel"] = "anon"

	v.BaseURI = viewInfo.BaseURI
	v.Extension = viewInfo.Extension
	v.Folder = viewInfo.Folder
	v.Name = viewInfo.Name

	// Make sure BaseURI is available in the templates
	// v.Vars["BaseURI"] = v.BaseURI
	v.Vars["BaseURI"] = "/"

	// This is required for the view to access the request
	v.request = r
	return v
}

/**
*
* Render view from controller
*
**/
func (v *View) Render(w http.ResponseWriter) {

	var templateList []string
	templateList = append(templateList, rootTemplate)
	templateList = append(templateList, childTemplates...)
	templateList = append(templateList, v.Name)

	// Loop through each template and test the full path
	for i, name := range templateList {
		// Get the absolute path of the root template
		path, err := filepath.Abs(v.Folder + string(os.PathSeparator) + name + "." + v.Extension)
		if err != nil {
			http.Error(w, "Template Path Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		templateList[i] = path
	}
	// Determine if there is an error in the template syntax
	templates, err := template.New(v.Name).ParseFiles(templateList...)
	if err != nil {
		http.Error(w, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Page url
	v.Vars["url"] = GetUrl(v.request)

	// User name
	sess, _ := session.SessionStart(v.request, w)
	userName := sess.Get("name")
	v.Vars["name"] = userName

	// Get flash message
	fm, err := flash.GetFlash(w, v.request)
	if err == nil && (flash.Flash{}) != fm {
		var flashes = make([]flash.Flash, 0)
		flashes = append(flashes, fm)
		v.Vars["flashes"] = flashes
	}

	// Get count cart product
	userID, _ := strconv.Atoi(sess.Get("id"))
	count := service.ProcessGetCountCartProduct(uint(userID))
	v.Vars["count"] = count

	err = templates.ExecuteTemplate(w, "layout."+v.Extension, v.Vars)
	if err != nil {
		http.Error(w, "Template File Error: "+err.Error(), http.StatusInternalServerError)
	}
}

/**
*
* Get page url
*
**/
func GetUrl(r *http.Request) string {
	return r.URL.Path
}
