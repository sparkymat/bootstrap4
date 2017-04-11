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
	expectedString := `<div class="navbar navbar-toggleable-md navbar-light bg-faded"><button class="navbar-toggler navbar-toggler-right" aria-expanded="false" aria-label="Toggle navigation" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent"></button></div>`

	node := Navbar()

	if node.String() != expectedString {
		t.Errorf("Expected: %v, Got: %v\n", expectedString, node.String())
	}
}
