package generator

import (
	"path/filepath"

	"github.com/gobuffalo/envy"
	"github.com/pkg/errors"
)

var (
	// ErrorNotInGoPath GOPATH error
	ErrorNotInGoPath = errors.New("Currently not in a $GOPATH")
)

// TemplateFolder get where the templates lives
func TemplateFolder(spec string) string {
	return filepath.Join(
		envy.GoPath(),
		"src",
		"tools.adidas-group.com",
		"bitbucket",
		"chatojor",
		"burrow",
		"generator",
		spec,
		"templates",
	)
}
