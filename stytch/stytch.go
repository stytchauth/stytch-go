package stytch

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/stytchauth/stytch-go/v12/stytch/config"
	"github.com/stytchauth/stytch-go/v12/stytch/stytcherror"
)

const (
	EnvTest = config.EnvTest
	EnvLive = config.EnvLive
)

type Client interface {
	NewRequest(ctx context.Context, method string, path string, queryParams map[string]string, body []byte, v interface{}, headers map[string][]string) error
	RawRequest(ctx context.Context, method string, path string, queryParams map[string]string, body []byte, headers map[string][]string) ([]byte, error)
	GetConfig() *config.Config
	GetHTTPClient() *http.Client
}

type DefaultClient struct {
	Config     *config.Config
	HTTPClient *http.Client
}

func New(projectID string, secret string) *DefaultClient {
	var detectedEnv config.Env
	if strings.HasPrefix(projectID, "project-live-") {
		detectedEnv = config.EnvLive
	} else {
		detectedEnv = config.EnvTest
	}

	stytchClient := new(DefaultClient)
	stytchClient.Config = config.New()

	stytchClient.Config.SetBasicAuthProjectID(projectID)
	stytchClient.Config.SetBasicAuthSecret(secret)
	stytchClient.Config.SetEnv(detectedEnv)

	stytchClient.HTTPClient = &http.Client{}

	return stytchClient
}

// newRequest is used by Call to generate and Do a http.Request
func (c *DefaultClient) NewRequest(
	ctx context.Context,
	method string,
	path string,
	queryParams map[string]string,
	body []byte,
	v interface{},
	headers map[string][]string,
) error {
	b, err := c.RawRequest(ctx, method, path, queryParams, body, headers)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, v); err != nil {
		return fmt.Errorf("error decoding http request: %w", err)
	}
	return nil
}

// RawRequest sends the request and returns the successful response body as bytes. If the response
// is an error, the response body will be parsed and returned as (nil, stytcherror.Error).
//
// Prefer using NewRequest (which unmarshals the response JSON) unless you need the actual bytes.
func (c *DefaultClient) RawRequest(
	ctx context.Context,
	method string,
	path string,
	queryParams map[string]string,
	body []byte,
	headers map[string][]string,
) ([]byte, error) {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = string(c.Config.BaseURI) + path

	req, err := http.NewRequestWithContext(ctx, method, path, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
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
	if len(c.Config.ProjectID) > 1 || len(c.Config.Secret) > 1 {
		authToken := base64.StdEncoding.EncodeToString(
			[]byte(c.Config.ProjectID + ":" + c.Config.Secret))
		req.Header.Set("Authorization", "Basic "+authToken)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "Stytch Go v"+config.APIVersion)

	for k, vSlice := range headers {
		for _, v := range vSlice {
			req.Header.Add(k, v)
		}
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending http request: %w", err)
	}
	defer func() {
		res.Body.Close()
	}()

	// Successful response
	if res.StatusCode == 200 || res.StatusCode == 201 {
		return io.ReadAll(res.Body)
	}

	// Attempt to unmarshal into Stytch error format
	var stytchErr stytcherror.Error
	if err = json.NewDecoder(res.Body).Decode(&stytchErr); err != nil {
		return nil, fmt.Errorf("error decoding http request: %w", err)
	}
	stytchErr.StatusCode = res.StatusCode
	return nil, stytchErr
}

func (c *DefaultClient) GetConfig() *config.Config {
	return c.Config
}

func (c *DefaultClient) GetHTTPClient() *http.Client {
	return c.HTTPClient
}
