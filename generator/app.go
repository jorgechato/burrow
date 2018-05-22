package generator

import (
	"encoding/json"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/gobuffalo/envy"
)

// App represents meta data for a Buffalo application on disk
type App struct {
	Pwd    string `json:"pwd"`
	Root   string `json:"root"`
	GoPath string `json:"go_path"`
	Name   string `json:"name"`
	Bin    string `json:"bin"`
	Pkg    string `json:"package_path"`
}

func newApp(root string) App {
	pwd, _ := os.Getwd()
	if root == "." {
		root = pwd
	}
	name := filepath.Base(root)
	pkg := envy.CurrentPackage()
	if filepath.Base(pkg) != string(name) {
		pkg = path.Join(pkg, string(name))
	}

	app := App{
		Pwd:    pwd,
		Root:   root,
		GoPath: envy.GoPath(),
		Name:   name,
		Pkg:    pkg,
	}

	app.Bin = filepath.Join("bin", filepath.Base(root))

	if runtime.GOOS == "windows" {
		app.Bin += ".exe"
	}

	return app
}

func (a App) String() string {
	b, _ := json.Marshal(a)
	return string(b)
}
