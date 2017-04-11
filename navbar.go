package bootstrap4

import (
	"github.com/sparkymat/bootstrap4/classes"
	"github.com/sparkymat/webdsl/html"
)

func Navbar(title string) *html.Node {
	return html.Div().Class(classes.Navbar, classes.NavbarToggleableMd, classes.NavbarLight, classes.BgFaded).Add(
		html.Button().Class(classes.NavbarToggler, classes.NavbarTogglerRight).Attr("type", "button").Data("toggle", "collapse").Data("target", "#navbarSupportedContent").Attr("aria-controls", "navbarSupportedContent").Attr("aria-expanded", "false").Attr("aria-label", "Toggle navigation").Add(
			html.Span().Class(classes.NavbarTogglerIcon),
		),
		html.LinkTo(title, "#").Class(classes.NavbarBrand),
	)
}
