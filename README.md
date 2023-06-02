[![CI](https://github.com/r-goswami/email-validator/actions/workflows/CI.yaml/badge.svg)](https://github.com/r-goswami/email-validator/actions/workflows/CI.yaml)

# email-validator
Go library to verify email addresses through Email verification API provided by 3rd party services as abstractApi, hunter.io etc.

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

