package b2b

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-go/v15/stytch"
	"github.com/stytchauth/stytch-go/v15/stytch/b2b/recoverycodes"
	"github.com/stytchauth/stytch-go/v15/stytch/stytcherror"
)

type RecoveryCodesClient struct {
	C stytch.Client
}

func NewRecoveryCodesClient(c stytch.Client) *RecoveryCodesClient {
	return &RecoveryCodesClient{
		C: c,
	}
}

// Recover: Allows a to complete an MFA flow by consuming a recovery code. This consumes the recovery code
// and returns a session token that can be used to authenticate the Member.
func (c *RecoveryCodesClient) Recover(
	ctx context.Context,
	body *recoverycodes.RecoverParams,
) (*recoverycodes.RecoverResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal recoverycodes.RecoverResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/recovery_codes/recover",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Get: Returns a's full set of active recovery codes.
func (c *RecoveryCodesClient) Get(
	ctx context.Context,
	body *recoverycodes.GetParams,
) (*recoverycodes.GetResponse, error) {
	headers := make(map[string][]string)

	var retVal recoverycodes.GetResponse
	err := c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "GET",
			Path:        fmt.Sprintf("/v1/b2b/recovery_codes/%s/%s", body.OrganizationID, body.MemberID),
			QueryParams: nil,
			Body:        nil,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}

// Rotate a's recovery codes. This invalidates all existing recovery codes and generates a new set of
// recovery codes.
func (c *RecoveryCodesClient) Rotate(
	ctx context.Context,
	body *recoverycodes.RotateParams,
) (*recoverycodes.RotateResponse, error) {
	var jsonBody []byte
	var err error
	if body != nil {
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, stytcherror.NewClientLibraryError("error marshaling request body")
		}
	}

	headers := make(map[string][]string)

	var retVal recoverycodes.RotateResponse
	err = c.C.NewRequest(
		ctx,
		stytch.RequestParams{
			Method:      "POST",
			Path:        "/v1/b2b/recovery_codes/rotate",
			QueryParams: nil,
			Body:        jsonBody,
			V:           &retVal,
			Headers:     headers,
		},
	)
	return &retVal, err
}
