package generator

import (
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/gobuffalo/envy"
)

// Generator model
type Generator struct {
	App
	DB   bool   `json:"db"`
	Name string `json:"name"`
}

// New app generator
func New(name string, db bool) (Generator, error) {
	g := Generator{
		App:  newApp("."),
		DB:   db,
		Name: name,
	}

	if name == "." {
		g.Name = g.App.Name
	}
	g.Root = filepath.Join(g.Root, name)
	g.Pkg = filepath.Join(g.Pkg, name)

	return g, g.inGoPath()
}

func (g Generator) inGoPath() error {
	gpMultiple := envy.GoPaths()

	larp := strings.ToLower(g.Root)
	for i := 0; i < len(gpMultiple); i++ {
		lgpm := strings.ToLower(filepath.Join(gpMultiple[i], "src"))
		if strings.HasPrefix(larp, lgpm) {
			return nil
		}
	}

	return ErrorNotInGoPath
}

func (g Generator) String() string {
	b, _ := json.Marshal(g)
	return string(b)
}
