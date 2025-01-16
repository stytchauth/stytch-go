package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v17/stytch"
	"github.com/stytchauth/stytch-go/v17/stytch/b2b/magiclinks/email"
	"github.com/stytchauth/stytch-go/v17/stytch/stytcherror"
)

type MagicLinksEmailClient struct {
	C         stytch.Client
	Discovery *MagicLinksEmailDiscoveryClient
}

func NewMagicLinksEmailClient(c stytch.Client) *MagicLinksEmailClient {
	return &MagicLinksEmailClient{
		C: c,

		Discovery: NewMagicLinksEmailDiscoveryClient(c),
	}
}

// LoginOrSignup: Send either a login or signup magic link to a Member. A new, pending, or invited Member
// will receive a signup Email Magic Link. Members will have a `pending` status until they successfully
// authenticate. An active Member will receive a login Email Magic Link.
//
// The magic link is valid for 60 minutes.
func (c *MagicLinksEmailClient) LoginOrSignup(
	ctx context.Context,
	body *email.LoginOrSignupParams,
) (*email.LoginOrSignupResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal email.LoginOrSignupResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/magic_links/email/login_or_signup",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Invite: Send an invite email to a new to join an. The Member will be created with an `invited` status
// until they successfully authenticate. Sending invites to `pending` Members will update their status to
// `invited`. Sending invites to already `active` Members will return an error.
//
// The magic link invite will be valid for 1 week.
func (c *MagicLinksEmailClient) Invite(
	ctx context.Context,
	body *email.InviteParams,
	methodOptions ...*email.InviteRequestOptions,
) (*email.InviteResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)
	for _, methodOption := range methodOptions {
		headers = methodOption.AddHeaders(headers)
	}

	var retVal email.InviteResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/magic_links/email/invite",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}
