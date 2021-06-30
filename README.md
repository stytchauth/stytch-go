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
$ go get github.com/stytchauth/stytch-go/v3
```

## Documentation

The module supports all Stytch API endpoints. Full documentation can be found [here](https://stytch.com/docs).

## Getting Started

### Calling Endpoints

To call an endpoint you must first create a Stytch `API` object.

```go
import (
	"os"

	"github.com/stytchauth/stytch-go/v3/stytch"
	"github.com/stytchauth/stytch-go/v3/stytch/stytchapi"
)

stytchAPIClient := stytchapi.NewAPIClient(
	stytch.EnvTest, // available environments are EnvTest and EnvLive
	os.Getenv("STYTCH_PROJECT_ID"),
	os.Getenv("STYTCH_SECRET"), 
)
```

Each endpoint returns an object which contains the parsed JSON from the HTTP response.

#### Users - Create
```go
	res, err := stytchAPIClient.Users.Create(&stytch.UsersCreateParams{
		Email:      "sandbox@stytch.com",
		Name: stytch.Name{
			FirstName:  "Clark",
			MiddleName: "Joseph",
			LastName:   "Kent",
		},
	})
```

#### Users - Get
```go
	res, err := stytchAPIClient.Users.Get("user-test-e3ca2fde-0cbe-4248-a8b8-b1dd68a4514d")
```

#### Magic Links - Email - Send
```go
	res, err := stytchAPIClient.MagicLinks.Email.Send(&stytch.MagicLinksEmailSendParams{
		Email:              "sandbox@stytch.com",
		LoginMagicLinkURL:  "https://example.com/login",
		SignupMagicLinkURL: "https://example.com/signup",
		Attributes:         stytch.Attributes{
			IPAddress: "10.0.0.0",
		},
    })
```

#### Magic Links - Authenticate
```go
	res, err := stytchAPIClient.MagicLinks.Authenticate(
		&stytch.MagicLinksAuthenticateParams{
			Token:      "GCRzBlufdaQ3mJh2QcygLsbuG__gqGwwvRuIuetv6ZM=",
			Options:    stytch.Options{IPMatchRequired: true},
			Attributes: stytch.Attributes{IPAddress: "10.0.0.0"},
		})
```

#### Magic Links - Email - Login or Create
```go
	res, err := stytchAPIClient.MagicLinks.Email.LoginOrCreate(&stytch.MagicLinksEmailLoginOrCreateParams{
		Email:                  "sandbox@stytch.com",
		LoginMagicLinkURL:      "https://example.com/login",
		SignupMagicLinkURL:     "https://example.com/signup",
		Attributes:             stytch.Attributes{
			IPAddress: "10.0.0.0",
		},
	})
```

#### Magic Links - Email - Invite
```go
	res, err := stytchAPIClient.MagicLinks.Email.Invite(&stytch.MagicLinksEmailInviteParams{
		Email:                   "sandbox@stytch.com",
		InviteMagicLinkURL:      "https://example.com/invite",
		Attributes:              stytch.Attributes{
			IPAddress: "10.0.0.0",
		},
	})
```

#### Magic Links - Email - Revoke Invite
```go
	res, err := stytchAPIClient.MagicLinks.Email.RevokeInvite(&stytch.MagicLinksEmailRevokeInviteParams{
		Email: "sandbox@stytch.com"
	})
```

#### Users - Get Pending
```go
	res, err := stytchAPIClient.Users.GetPending()
```

### Errors

All non-200 responses will return a stytch.Error instance.

For more information on Stytch response codes, head to the [docs](https://stytch.com/docs/api/errors).

## Developing

1. Download this repo into your Go source directory
2. Run `make setup` pull down all dependencies etc

## Support

Open an [issue](https://github.com/stytchauth/stytch-go/issues/new)!

## License

[MIT](LICENSE)
