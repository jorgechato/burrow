package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

// GetAllFiles from a path
func GetAllFiles(path string) []string {
	var allFiles []string
	filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		filename := f.Name()
		if strings.HasSuffix(filename, ".tmpl") {
			allFiles = append(allFiles, path)
		}
		return nil
	})

	return allFiles
}

// Render the template to an actual go file
func Render(tmpl string, app Generator) error {
	filename := filepath.Base(tmpl)
	t := template.Must(
		template.
			New(filename).
			Funcs(TemplateMap).
			ParseFiles(tmpl),
	)

	subfolder := ""
	if strings.Contains(tmpl, filepath.Join("templates", "project")) {
		os.MkdirAll(
			filepath.Join(app.Root, app.Name),
			0755,
		)

		subfolder = app.Name
	}
	if strings.Contains(tmpl, filepath.Join("templates", "swagger-ui")) {
		os.MkdirAll(
			filepath.Join(app.Root, "swagger-ui"),
			0755,
		)

		subfolder = "swagger-ui"
	}

	goFile := filepath.Join(
		app.Root,
		subfolder,
		strings.Replace(filename, ".tmpl", "", -1),
	)

	if _, err := os.Stat(goFile); os.IsNotExist(err) {
		f, _ := os.Create(goFile)
		defer f.Close()

		err := t.Execute(f, app)
		return err
	}

	return nil
}

// TemplateMap is a map of custom functions to work with tmpl variables
var TemplateMap = template.FuncMap{
	"ToUpper": strings.ToUpper,
	"Title":   strings.Title,
	"Nose": func(s string) string {
		return filepath.Dir(s)
	},
}

// InstallDependencies will install all dependencies in the app/project
func InstallDependencies(path string) {
	c := exec.Command(
		"go",
		"get",
		"-v",
		"-d",
		fmt.Sprintf("./%v/...", path),
	)

	c.Run()
}
