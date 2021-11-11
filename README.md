# Stytch Go Library

The Stytch Go library makes it easy to use the Stytch user infrastructure API in Go applications.

It pairs well with the Stytch [Web SDK](https://www.npmjs.com/package/@stytch/stytch-js) or your own custom authentication flow.

## Install

```console
$ go get github.com/stytchauth/stytch-go/v3
```

## Usage

You can find your API credentials in the [Stytch Dashboard](https://stytch.com/dashboard/api-keys).

This client library supports all of Stytch's live products:
  - [x] [Email Magic Links](https://stytch.com/docs/api/send-by-email)
  - [x] [Embeddable Magic Links](https://stytch.com/docs/api/create-magic-link-overview)
  - [x] [OAuth logins](https://stytch.com/docs/api/oauth-overview)
  - [x] [SMS passcodes](https://stytch.com/docs/api/send-otp-by-sms)
  - [x] [WhatsApp passcodes](https://stytch.com/docs/api/whatsapp-send)
  - [x] [Email passcodes](https://stytch.com/docs/api/send-otp-by-email)
  - [x] [Session Management (Beta)](https://stytch.com/docs/api/sessions-overview
  - [ ] [WebAuthn (Coming soon)](https://stytch.com/docs/api/webauthn-overview)

### Example usage
Create an API client:
```go
import (
	"os"

	"github.com/stytchauth/stytch-go/v3/stytch"
	"github.com/stytchauth/stytch-go/v3/stytch/stytchapi"
)

stytchAPIClient := stytchapi.NewAPIClient(
	stytch.EnvTest, // available environments are EnvTest and EnvLive
	"project-live-c60c0abe-c25a-4472-a9ed-320c6667d317",
	"secret-live-80JASucyk7z_G8Z-7dVwZVGXL5NT_qGAQ2I=",
)
```

Send a magic link by email:
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

Authenticate the token from the magic link:
```go
	res, err := stytchAPIClient.MagicLinks.Authenticate(
		&stytch.MagicLinksAuthenticateParams{
			Token:      "DOYoip3rvIMMW5lgItikFK-Ak1CfMsgjuiCyI7uuU94=",
			Options:    stytch.Options{IPMatchRequired: true},
			Attributes: stytch.Attributes{IPAddress: "10.0.0.0"},
		})
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

If you have questions or want help troubleshooting, join us in [Slack](https://join.slack.com/t/stytch/shared_invite/zt-nil4wo92-jApJ9Cl32cJbEd9esKkvyg) or email support@stytch.com.

If you've found a security vulnerability, please follow our [responsible disclosure instructions](https://stytch.com/docs/security).

## Development

See [DEVELOPMENT.md](DEVELOPMENT.md)

## Code of Conduct

Everyone interacting in the Stytch project's codebases, issue trackers, chat rooms and mailing lists is expected to follow the [code of conduct](CODE_OF_CONDUCT.md).
