package main

import (
	"fmt"
	"os"

	emailvalidator "github.com/r-goswami/email-validator"
)

func main() {

	apiKey := os.Getenv("API_KEY")
	email := os.Getenv("EMAIL")
	validator, err := emailvalidator.NewAbstractAPIClient(apiKey)
	if err != nil {
		panic(err)
	}

	resp, err := validator.Validate(email)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", resp)
	if resp.IsValid() {
		fmt.Printf("\n%s is a valid email\n", email)
	}
}
