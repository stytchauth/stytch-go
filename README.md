# stytch-go

A Go client library for the [Stytch API](https://stytch.com/).

## Table of Contents

- [stytch-go](#stytch-go)
    * [Install](#install)
    * [Documentation](#documentation)
    * [Getting Started](#getting-started)
    * [Developing](#developing)
    * [License](#license)

## Install

```console
$ go get github.com/stytchauth/stytch-go
```

## Documentation

The module supports all Stytch API endpoints. Full documentation can be found [here](https://docs.stytch.com/reference).

## Getting Started

### Calling Endpoints

To call an endpoint you must create a `Client` object.

```go
import (
    "os"

    "github.com/stytchauth/stytch-go/stytch"
)

stytchClient := stytch.NewClient(
    stytch.EnvTest, // available environments are EnvTest and EnvLive
    os.Getenv("STYTCH_PROJECT_ID"), 
    os.Getenv("STYTCH_SECRET"), 
)
```

Each endpoint returns an object which contains the parsed JSON from the HTTP response.

#### Create User
```go
    res, err := stytchClient.CreateUser(&stytch.CreateUser{
		Email:      "clark@stytch.com",
		Name: stytch.Name{
			FirstName:  "Clark",
			MiddleName: "Joseph",
			LastName:   "Kent",
		},
    })
```

#### Get User
```go
    res, err := stytchClient.GetUser("user-test-e3ca2fde-0cbe-4248-a8b8-b1dd68a4514d")
```

#### Send Magic Link
```go
    res, err := sc.SendMagicLinkByEmail(&stytch.SendMagicLinkByEmail{
		Email:             "clarck@stytch.com",
		MagicLinkURL:      "https://yoururl.com",
		ExpirationMinutes: 5,
		Attributes:        stytch.Attributes{
			IPAddress: "10.0.0.0",
		},
    })
```

#### Authenticate Magic Link
```go
    res, err := sc.AuthenticateMagicLink(
		"GCRzBlufdaQ3mJh2QcygLsbuG__gqGwwvRuIuetv6ZM=",
		&stytch.AuthenticateMagicLink{
			Options:    stytch.Options{IPMatchRequired: true},
			Attributes: stytch.Attributes{IPAddress: "10.0.0.0"},
		})
```

### Errors

All non-200 responses will return a stytch.Error instance.

For more information on Stytch response codes, head to the [docs](https://docs.stytch.com/reference#errors).

## Developing

1. Download this repo into your Go source directory
2. Run `make setup` pull down all dependencies etc

## Support

Open an [issue](https://github.com/stytchauth/stytch-go/issues/new)!

## License

[MIT](LICENSE)
