package stytch

import "encoding/json"

func (c *Client) DeleteEmail (emailID string, userID string) (*DeleteEmailResponse, error) {
	path := "emails/" + emailID + "/users/" + userID

	var retVal *DeleteEmailResponse
	err := c.newRequest("DELETE", path, nil, &retVal)
	return retVal, err
}

func (c *Client) SendEmailVerification (emailID string, body *SendEmailVerification) (*SendEmailVerificationResponse, error) {
	path := "/emails/" + emailID + "/send_verification"

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var retVal *SendEmailVerificationResponse
	err = c.newRequest("POST", path, jsonBody, &retVal)
	return retVal, err
}

/**
* Verify that a user supplied the correct email during signup.
* @param    string        token     parameter: Required
* @return	Returns the *models_pkg.VerifyEmailResponse response from the API call
*/
func (c *Client) VerifyEmail (token string) (*VerifyEmailResponse, error) {
	path := "/emails/" + token + "/verify"

	var retVal *VerifyEmailResponse
	err := c.newRequest("POST", path, nil, &retVal)
	return retVal, err
}