package stytch

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
)

const apiVersion = "1.1.4"

type Client struct {
	Config     *config
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
func (c *Client) newRequest(method string, path string, queryParams map[string]string,
	body []byte, v interface{}) error {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = string(c.Config.baseURI) + path

	req, err := http.NewRequest(method, path, bytes.NewReader(body))
	if err != nil {
		return newInternalServerError("Oops, something seems to have gone " +
			"wrong creating a new http request")
	}

	// add query params
	q := req.URL.Query()
	for k, v := range queryParams {
		if v != "" {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()

	// append basic auth headers
	if len(c.Config.projectID) > 1 || len(c.Config.secret) > 1 {
		authToken := base64.StdEncoding.EncodeToString(
			[]byte(c.Config.projectID + ":" + c.Config.secret))
		req.Header.Set("Authorization", "Basic "+authToken)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Stytch Go v"+apiVersion)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return newInternalServerError("Oops, something seems to have gone " +
			"wrong sending the http request")
	}
	defer func() {
		res.Body.Close()
	}()

	// Successful response
	if res.StatusCode == 200 || res.StatusCode == 201 {
		if err = json.NewDecoder(res.Body).Decode(v); err != nil {
			return newInternalServerError("Oops, something seems to have gone wrong " +
				"decoding the successful response body")
		}
		return nil
	}

	// Attempt to unmarshal into Stytch error format
	var stytchErr Error
	if err = json.NewDecoder(res.Body).Decode(&stytchErr); err != nil {
		return newInternalServerError("Oops, something seems to have gone wrong " +
			"decoding the unsuccessful response body")
	}
	stytchErr.StatusCode = res.StatusCode
	return stytchErr
}
