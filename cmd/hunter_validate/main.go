package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	emailvalidator "github.com/r-goswami/email-validator"
)

type cli struct{}

func (cli *cli) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *cli) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  validate -apiKey API_KEY -email EMAIL - Validates if a given email address is valid")
}

func (cli *cli) validate(apiKey string, email string) {
	validator, err := emailvalidator.NewHunterAPIClient(apiKey)
	if err != nil {
		log.Panic(err)
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

func (cli *cli) Run() {
	cli.validateArgs()

	validateCmd := flag.NewFlagSet("validate", flag.ExitOnError)

	apiKey := validateCmd.String("apiKey", "", "client API Key")
	email := validateCmd.String("email", "", "email to validate")

	switch os.Args[1] {
	case "validate":
		err := validateCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if validateCmd.Parsed() {
		if *apiKey == "" || *email == "" {
			validateCmd.Usage()
			os.Exit(1)
		}

		cli.validate(*apiKey, *email)
	}
}

func main() {
	cli := cli{}
	cli.Run()
}
