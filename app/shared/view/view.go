package view

import "net/http"
import "html/template"
import "os"
import "path/filepath"
import "sync"
import "github.com/goweb3/app/shared/session"

var (
	// FlashError is a bootstrap class
	FlashError = "alert-danger"
	// FlashSuccess is a bootstrap class
	FlashSuccess = "alert-success"
	// FlashNotice is a bootstrap class
	FlashNotice = "alert-info"
	// FlashWarning is a bootstrap class
	FlashWarning = "alert-warning"

	viewInfo View
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
// Flash Message
type Flash struct {
	Message string
	Class   string
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

func New(req *http.Request) *View {
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
	v.request = req
	return v
}

func (v *View) Render(res http.ResponseWriter) {

	var templateList []string
	templateList = append(templateList, rootTemplate)
	templateList = append(templateList, childTemplates...)
	templateList = append(templateList, v.Name)
	


	// Loop through each template and test the full path
	for i, name := range templateList {
		// Get the absolute path of the root template
		path, err := filepath.Abs(v.Folder + string(os.PathSeparator) + name + "." + v.Extension)
		if err != nil {
			http.Error(res, "Template Path Error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		templateList[i] = path
	}
	// Determine if there is an error in the template syntax
	templates, err := template.New(v.Name).ParseFiles(templateList...)
	if err != nil {
		http.Error(res, "Template Parse Error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Get session
	sess := session.Instance(v.request)
	// Get the flashes for the template
	if flashes := sess.Flashes(); len(flashes) > 0 {
		v.Vars["flashes"] = make([]Flash, len(flashes))
		for i, f := range flashes {
			switch f.(type) {
				case Flash:
					v.Vars["flashes"].([]Flash)[i] = f.(Flash)
				default:
					v.Vars["flashes"].([]Flash)[i] = Flash{f.(string), "alert-box"}
				}
		}
		sess.Save(v.request, res)
	}
	err = templates.ExecuteTemplate(res, "layout."+v.Extension, v.Vars)
	if err != nil {
		http.Error(res, "Template File Error: "+err.Error(), http.StatusInternalServerError)
	}
}
