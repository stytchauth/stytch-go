# Stytch Go Library

The Stytch Go library makes it easy to use the Stytch user infrastructure API in Go applications.

It pairs well with the Stytch [Web SDK](https://www.npmjs.com/package/@stytch/vanilla-js) or your own custom authentication flow.

## Install

```console
$ go get github.com/stytchauth/stytch-go/v12
```

## Usage

You can find your API credentials in the [Stytch Dashboard](https://stytch.com/dashboard/api-keys).

This client library supports all Stytch's live products:
  - [x] [Email Magic Links](https://stytch.com/docs/api/send-by-email)
  - [x] [Embeddable Magic Links](https://stytch.com/docs/guides/magic-links/embeddable-magic-links/api)
  - [x] [OAuth logins](https://stytch.com/docs/guides/oauth/idp-overview)
  - [x] [SMS passcodes](https://stytch.com/docs/api/send-otp-by-sms)
  - [x] [WhatsApp passcodes](https://stytch.com/docs/api/whatsapp-send)
  - [x] [Email passcodes](https://stytch.com/docs/api/send-otp-by-email)
  - [x] [Session Management](https://stytch.com/docs/guides/sessions/using-sessions)
  - [x] [WebAuthn](https://stytch.com/docs/guides/webauthn/api)
  - [x] [Time-based one-time passcodes (TOTPs)](https://stytch.com/docs/guides/totp/api)
  - [x] [Crypto wallets](https://stytch.com/docs/guides/web3/api)
  - [x] [Passwords](https://stytch.com/docs/guides/passwords/api)
  - [x] [M2M](https://stytch.com/docs/api/m2m-client)


### Example usage
Create an API client:
```go
import (
	"context"

	"github.com/stytchauth/stytch-go/v12/stytch"
	"github.com/stytchauth/stytch-go/v12/stytch/consumer/stytchapi"
)

stytchAPIClient, err := stytchapi.NewClient(
	"project-live-c60c0abe-c25a-4472-a9ed-320c6667d317",
	"secret-live-80JASucyk7z_G8Z-7dVwZVGXL5NT_qGAQ2I=",
)
if err != nil {
	panic(err)
}
```

Send a magic link by email:
```go
	res, err := stytchAPIClient.MagicLinks.Email.Send(
		context.Background(),
		&stytch.MagicLinksEmailSendParams{
			Email: "sandbox@stytch.com",
			Attributes: stytch.Attributes{
				IPAddress: "10.0.0.0",
			},
		},
	)
```

Authenticate the token from the magic link:
```go
	res, err := stytchAPIClient.MagicLinks.Authenticate(
		context.Background(),
		&stytch.MagicLinksAuthenticateParams{
			Token:      "DOYoip3rvIMMW5lgItikFK-Ak1CfMsgjuiCyI7uuU94=",
			Options:    stytch.Options{IPMatchRequired: true},
			Attributes: stytch.Attributes{IPAddress: "10.0.0.0"},
		})
```

Get all users
```go
    res, err := stytchAPIClient.Users.Search(
		context.Background(),
		&stytch.UsersSearchParams{
			Limit: 1000
		})
```

Search users
```go
	res, err := stytchAPIClient.Users.Search(
		context.Background(),
		&stytch.UsersSearchParams{
			Limit: 1000,
			Query: stytch.UsersSearchQuery{
				Operator: stytch.UserSearchOperatorOR,
				Operands: []json.Marshaler{
					stytch.UsersSearchQueryPhoneVerifiedFilter{true},
					stytch.UsersSearchQueryEmailVerifiedFilter{true},
					stytch.UsersSearchQueryWebAuthnRegistrationVerifiedFilter{true},
				}
			}
		})
```

Iterate over all pages of users for a search query
```go
	var users []stytch.User
	iter := stytchAPIClient.Users.SearchAll(&stytch.UsersSearchParams{})
	for iter.HasNext() {
		res, err := iter.Next(context.Background())
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, res...)
	}
```

## Handling Errors

When possible Stytch returns an error prepended with `Stytch Error`.
Additionally, the error should include a type that can be used to distinguish errors.

Learn more about errors in the [docs](https://stytch.com/docs/api/errors).

## Documentation

See example requests and responses for all the endpoints in the [Stytch API Reference](https://stytch.com/docs/api).

Follow one of the [integration guides](https://stytch.com/docs/guides) or start with one of our [example apps](https://stytch.com/docs/example-apps).

## Support

If you've found a bug, [open an issue](https://github.com/stytchauth/stytch-go/issues/new)!

If you have questions or want help troubleshooting, join us in [Slack](https://stytch.slack.com/join/shared_invite/zt-2f0fi1ruu-ub~HGouWRmPARM1MTwPESA) or email support@stytch.com.

If you've found a security vulnerability, please follow our [responsible disclosure instructions](https://stytch.com/docs/resources/security-and-trust/security#:~:text=Responsible%20disclosure%20program).

## Development

See [DEVELOPMENT.md](DEVELOPMENT.md)

## Code of Conduct

Everyone interacting in the Stytch project's codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](CODE_OF_CONDUCT.md).
