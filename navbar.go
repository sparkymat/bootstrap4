package bootstrap4

import (
	"github.com/sparkymat/bootstrap4/classes"
	"github.com/sparkymat/webdsl/html"
)

func Navbar() *html.Node {
	return html.Div().Class(classes.Navbar, classes.NavbarToggleableMd, classes.NavbarLight, classes.BgFaded).Add(
		html.Button(),
	)
}
