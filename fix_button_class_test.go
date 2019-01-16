package main

/*
func TestFixEmptyButton(t *testing.T) {
	expected := []byte("<button class=\"old-style\"> ")
	result := fixTagClasses([]byte("<button> "), "button", "old-style")
	if !bytes.Equal(result, expected) {
		t.Errorf("expected `%s` but got `%s`", expected, result)
	}
}

func TestFixButtonWithClass(t *testing.T) {
	expected := []byte("<button class=\"old-style some-class good\"")
	result := fixTagClasses([]byte("<button class=\"some-class good\""), "button", "old-style")
	if !bytes.Equal(result, expected) {
		t.Errorf("expected %s but got %s", expected, result)
	}
}

func TestFixButtonWithClassAndNewline(t *testing.T) {
	expected := []byte("<button\n\n\t\tclass=\"old-style some-class good\"")

	result := fixTagClasses([]byte("<button\n\n\t\tclass=\"some-class good\""), "button", "old-style")
	if !bytes.Equal(result, expected) {
		t.Errorf("expected %s but got %s", expected, result)
	}
}
*/
