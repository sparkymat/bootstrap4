package bootstrap4

import (
	"github.com/sparkymat/bootstrap4/classes"
	"github.com/sparkymat/webdsl/html"
)

func AlertSuccess(message string) *html.Node {
	return html.Div().Class(classes.Alert, classes.AlertSuccess).Attr("role", "alert").Add(
		html.T(message),
	)
}

func AlertInfo(message string) *html.Node {
	return html.Div().Class(classes.Alert, classes.AlertInfo).Attr("role", "alert").Add(
		html.T(message),
	)
}

func AlertWarning(message string) *html.Node {
	return html.Div().Class(classes.Alert, classes.AlertWarning).Attr("role", "alert").Add(
		html.T(message),
	)
}

func AlertDanger(message string) *html.Node {
	return html.Div().Class(classes.Alert, classes.AlertDanger).Attr("role", "alert").Add(
		html.T(message),
	)
}
