package testtoml

import "log"

// ExampleGetTOML gets the test.toml file for coins.asia
func ExampleClient_GetTestToml() {
	_, err := DefaultClient.GetTestToml("coins.asia")
	if err != nil {
		log.Fatal(err)
	}
}
