package magic_links

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-go/v7/stytch"
	"github.com/stytchauth/stytch-go/v7/stytch/stytcherror"
)

type Client struct {
	C *stytch.Client
}

const subURL = "magic_links"

func (c *Client) Magiclinkredirect(
	ctx context.Context,
	body *stytch.MagiclinkredirectParams,
) (*stytch.MagiclinkredirectResponse, error) {
	path := subURL + "//v1/magic_links/redirect"

	var retVal stytch.MagiclinkredirectResponse
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong marshalling the request body")
		}
	}

	err = c.C.NewRequest(ctx, "GET", path, nil, jsonBody, &retVal)

	return &retVal, err
}

func (c *Client) Magiclinksredirectcaptcha(
	ctx context.Context,
	body *stytch.MagiclinksredirectcaptchaParams,
) (*stytch.MagiclinksredirectcaptchaResponse, error) {
	path := subURL + "//v1/magic_links/redirect/captcha"

	var retVal stytch.MagiclinksredirectcaptchaResponse
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("Oops, something seems to have gone wrong marshalling the request body")
		}
	}

	err = c.C.NewRequest(ctx, "POST", path, nil, jsonBody, &retVal)

	return &retVal, err
}
