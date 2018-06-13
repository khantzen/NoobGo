package views

import (
	"net/http"
	"html/template"
	"bytes"
	"path/filepath"
	"os"
	vm "../model/view"
)

func Render(view string, model interface{}, w http.ResponseWriter) {
	renderedView := RenderNoTemplate(view, model)

	templateViewModel := vm.MainViewTemplate{}

	templateViewModel.Body = template.HTML(renderedView)

	// Recover Css files
	cssFiles, _ := getFilesList("media/css/main")
	cssFiles = append(cssFiles, "media/css/"+view+".css")

	templateViewModel.Css = cssFiles

	// Recover Js files
	jsFiles, _ := getFilesList("media/js/main")
	jsFiles = append(jsFiles, "media/js/"+view+".js")

	templateViewModel.JavaScript = jsFiles;

	renderView("main", templateViewModel, w)

}

func RenderNoTemplate(view string, model interface{}) []byte {
	viewTemplate := getViewTemplate(view)
	tcl := bytes.Buffer{}
	viewTemplate.Execute(&tcl, model)

	return tcl.Bytes()
}

func renderView(view string, model interface{}, w http.ResponseWriter) {
	viewTemplate := getViewTemplate(view)
	viewTemplate.Execute(w, model)
}

func getViewTemplate(view string) *template.Template {
	fileName := "views/" + view + ".html"
	viewTemplate, _ := template.ParseFiles(fileName)
	return viewTemplate
}

func getFilesList(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
