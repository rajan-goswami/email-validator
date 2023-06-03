[![CI](https://github.com/r-goswami/email-validator/actions/workflows/CI.yaml/badge.svg)](https://github.com/r-goswami/email-validator/actions/workflows/CI.yaml)
[![codecov](https://codecov.io/gh/r-goswami/email-validator/branch/main/graph/badge.svg?token=VAB3KVPV1P)](https://codecov.io/gh/r-goswami/email-validator)
[![Go Report Card](https://goreportcard.com/badge/github.com/r-goswami/email-validator)](https://goreportcard.com/report/github.com/r-goswami/email-validator)
[![Go Reference](https://pkg.go.dev/badge/github.com/r-goswami/email-validator.svg)](https://pkg.go.dev/github.com/r-goswami/email-validator)

# email-validator
Go library to verify email addresses through Email verification API provided by 3rd party services - abstractapi.com and hunter.io.


## Features

- Completely native (no 3rd party module dependencies)
- Client side rate limiting
- Ability to change rate limits

---


# How to use http clients

## abstract.com client
```go

import (
    "fmt"
	"log"
    "os"

    emailvalidator "github.com/r-goswami/email-validator"
)

func main() {
    apiKey :=  os.GetEnv("API_KEY")
    validator, err := emailvalidator.NewAbstractAPIClient(apiKey)
	if err != nil {
		log.Panic(err)
	}

    email := "abc@xyz.com"
	resp, err := validator.Validate(email)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.IsValid() {
		fmt.Printf("\n%s is a valid email\n", email)
	}
}

```

## hunter.io client
```go

import (
    "fmt"
	"log"
    "os"

    emailvalidator "github.com/r-goswami/email-validator"
)

func main() {
    apiKey :=  os.GetEnv("API_KEY")
    validator, err := emailvalidator.NewHunterAPIClient(apiKey)
	if err != nil {
		log.Panic(err)
	}

    email := "abc@xyz.com"
	resp, err := validator.Validate(email)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.IsValid() {
		fmt.Printf("\n%s is a valid email\n", email)
	}
}

```

---


# How to check clients using commands

## abstract.com client
### Build
```bash
cd cmd/aapi_validate
go build
```

### Run
```bash
cd cmd/aapi_validate
./aapi_validate validate -apiKey API_KEY -email EMAIL
```

## hunter.io client
### Build
```bash
cd cmd/hunter_validate
go build
```

### Run
```bash
cd cmd/hunter_validate
./hunter_validate validate -apiKey API_KEY -email EMAIL
```

