package bootstrap4

import (
	"github.com/sparkymat/bootstrap4/classes"
	"github.com/sparkymat/webdsl/html"
)

func Navbar() *html.Node {
	return html.Div().Class(classes.Navbar, classes.NavbarToggleableMd, classes.NavbarLight, classes.BgFaded).Add(
		html.Button().Class(classes.NavbarToggler, classes.NavbarTogglerRight).Attr("type", "button").Data("toggle", "collapse").Data("target", "#navbarSupportedContent").Attr("aria-controls", "navbarSupportedContent").Attr("aria-expanded", "false").Attr("aria-label", "Toggle navigation"),
	)
}
