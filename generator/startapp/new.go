package startapp

import (
	"os"
	"strings"
	"sync"

	"github.com/jorgechato/burrow/generator"
)

// Run create the file system for the app
func Run(app generator.Generator) {
	os.MkdirAll(app.Root, 0755)

	templates := generator.GetAllFiles(
		generator.TemplateFolder("startapp"),
	)

	wg := new(sync.WaitGroup)

	for _, tmpl := range templates {
		wg.Add(1)
		go func(tmpl string, app generator.Generator) {
			if !app.DB && strings.Contains(tmpl, "database.go") {
			} else {
				err := generator.Render(tmpl, app)
				if err != nil {
					panic(err)
				}
			}

			wg.Done()
		}(tmpl, app)
	}

	wg.Wait()

	generator.InstallDependencies(app.Name)
}
