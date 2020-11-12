package stytch

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

type Client struct {
	Config  *config
	HTTPClient *http.Client
}

func NewClient(env Env, projectID string, secret string) *Client {
	stytchClient := new(Client)
	stytchClient.Config = newConfig()

	stytchClient.Config.SetBasicAuthProjectID(projectID)
	stytchClient.Config.SetBasicAuthSecret(secret)
	stytchClient.Config.SetEnv(env)

	stytchClient.HTTPClient = &http.Client{}

	return stytchClient
}

// newRequest is used by Call to generate and Do a http.Request
func (c *Client) newRequest(method string, path string, body []byte, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = string(c.Config.baseURI) + path

	req, err := http.NewRequest(method, path, bytes.NewReader(body))
	if err != nil {
		return err
	}

	//append basic auth headers
	if len(c.Config.projectID) > 1 || len(c.Config.secret) > 1 {
		authToken := base64.StdEncoding.EncodeToString([]byte(c.Config.projectID + ":" + c.Config.secret))
		req.Header.Set("Authorization", "Basic " + authToken)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Stytch Go v1.0.0")

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	// Successful response
	if res.StatusCode == 200 {
		return json.NewDecoder(res.Body).Decode(v)
	}

	// Attempt to unmarshal into Stytch error format
	var stytchErr stytchError
	if err = json.NewDecoder(res.Body).Decode(&stytchErr); err != nil {
		return err
	}
	stytchErr.Status = res.StatusCode
	return stytchErr
}
