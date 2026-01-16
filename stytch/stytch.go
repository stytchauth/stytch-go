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

	"github.com/stytchauth/stytch-go/v17/stytch/config"
	"github.com/stytchauth/stytch-go/v17/stytch/stytcherror"
)

type BaseURLType string

const (
	baseURLFraud BaseURLType = "FRAUD"
)

const (
	EnvTest = config.EnvTest
	EnvLive = config.EnvLive
)

type Client interface {
	NewRequest(ctx context.Context, params RequestParams) error
	RawRequest(ctx context.Context, params RequestParams) ([]byte, error)
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
	stytchClient.Config.FraudBaseURI = config.BaseURIFraud

	stytchClient.HTTPClient = &http.Client{}

	return stytchClient
}

type RequestParams struct {
	Method      string
	Path        string
	QueryParams map[string]string
	Body        []byte
	V           interface{}
	Headers     map[string][]string
	BaseURLType BaseURLType
}

// newRequest is used by Call to generate and Do a http.Request
func (c *DefaultClient) NewRequest(
	ctx context.Context,
	params RequestParams,
) error {
	b, err := c.RawRequest(ctx, params)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, params.V); err != nil {
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
	params RequestParams,
) ([]byte, error) {
	if !strings.HasPrefix(params.Path, "/") {
		params.Path = "/" + params.Path
	}

	baseURI := c.Config.BaseURI
	if params.BaseURLType == baseURLFraud {
		baseURI = c.Config.FraudBaseURI
	}

	params.Path = string(baseURI) + params.Path

	req, err := http.NewRequestWithContext(ctx, params.Method, params.Path, bytes.NewReader(params.Body))
	if err != nil {
		return nil, fmt.Errorf("error creating http request: %w", err)
	}

	// add query params
	q := req.URL.Query()
	for k, v := range params.QueryParams {
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

	for k, vSlice := range params.Headers {
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
