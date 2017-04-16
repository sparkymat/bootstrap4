package bootstrap4

import "testing"

func TestContainer(t *testing.T) {
	expectedString := `<div class="container"></div>`

	node := Container()

	if node.String() != expectedString {
		t.Errorf("Expected: %v, Got: %v\n", expectedString, node.String())
	}
}

func TestNavbar(t *testing.T) {
	expectedString := ` <div class="navbar navbar-toggleable-md navbar-light bg-faded"><button class="navbar-toggler navbar-toggler-right" aria-label="Toggle navigation" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false"><span class="navbar-toggler-icon"></span></button><a class="navbar-brand" href="#">Test</a></div>`

	node := Navbar("Test")

	if node.String() != expectedString {
		t.Errorf("Expected: %v, Got: %v\n", expectedString, node.String())
	}
}
