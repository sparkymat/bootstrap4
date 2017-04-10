package bootstrap4

import (
	"github.com/sparkymat/bootstrap4/classes"
	"github.com/sparkymat/webdsl/html"
)

func Container(children ...html.ChildNode) *html.Node {
	return html.Div().Class(classes.Container).Add(children...)
}

func ContainerFluid(children ...html.ChildNode) *html.Node {
	return html.Div().Class(classes.ContainerFluid).Add(children...)
}

func Row(children ...html.ChildNode) *html.Node {
	return html.Div().Class(classes.Row).Add(children...)
}

func Col(children ...html.ChildNode) *html.Node {
	return html.Div().Class(classes.Col).Add(children...)
}
