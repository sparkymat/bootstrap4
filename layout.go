package bootstrap4

import (
	"github.com/sparkymat/webdsl/css"
	"github.com/sparkymat/webdsl/html"
)

func Container(children ...html.ChildNode) *html.Node {
	return html.Div().Class(css.Class("container")).Add(children...)
}
