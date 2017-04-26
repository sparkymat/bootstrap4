package bootstrap4

import (
	"github.com/sparkymat/bootstrap4/classes"
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/html"
)

func Navbar(title string, links [][2]string, currentPath string) *html.Node {
	linkNodes := []*html.Node{}

	for _, linkDetails := range links {
		title := linkDetails[0]
		url := linkDetails[1]
		liClasses := []css.Class{classes.NavItem}
		if url == currentPath {
			liClasses = append(liClasses, classes.Active)
		}
		linkNodes = append(linkNodes, html.Li().Class(liClasses...).Add(
			html.LinkTo(title, url).Class(classes.NavLink),
		))
	}

	return html.Div().Class(classes.Navbar, classes.NavbarToggleableMd, classes.NavbarLight, classes.BgFaded).Add(
		html.Button().Class(classes.NavbarToggler, classes.NavbarTogglerRight).Attr("type", "button").Data("toggle", "collapse").Data("target", "#navbarSupportedContent").Attr("aria-controls", "navbarSupportedContent").Attr("aria-expanded", "false").Attr("aria-label", "Toggle navigation").Add(
			html.Span().Class(classes.NavbarTogglerIcon),
		),
		html.LinkTo(title, "#").Class(classes.NavbarBrand),
		html.Ul().Class(classes.NavbarNav).Add(linkNodes...),
	)
}
