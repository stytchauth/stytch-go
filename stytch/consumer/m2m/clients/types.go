package clients

// !!!
// WARNING: This file is autogenerated
// Only modify code within MANUAL() sections
// or your changes may be overwritten later!
// !!!

import (
	"github.com/stytchauth/stytch-go/v17/stytch/consumer/m2m"
)

// CreateParams: Request type for `Clients.Create`.
type CreateParams struct {
	// Scopes: An array of scopes assigned to the client.
	Scopes []string `json:"scopes,omitempty"`
	// ClientID: If provided, the ID of the client to create. If not provided, Stytch will generate this value
	// for you. The `client_id` must be unique within your project.
	ClientID string `json:"client_id,omitempty"`
	// ClientSecret: If provided, the stored secret of the client to create. If not provided, Stytch will
	// generate this value for you. If provided, the `client_secret` must be at least 8 characters long and
	// pass entropy requirements.
	ClientSecret string `json:"client_secret,omitempty"`
	// ClientName: A human-readable name for the client.
	ClientName string `json:"client_name,omitempty"`
	// ClientDescription: A human-readable description for the client.
	ClientDescription string `json:"client_description,omitempty"`
	// TrustedMetadata: The `trusted_metadata` field contains an arbitrary JSON object of application-specific
	// data. See the [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior
	// details.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
}

// DeleteParams: Request type for `Clients.Delete`.
type DeleteParams struct {
	// ClientID: The ID of the client.
	ClientID string `json:"client_id,omitempty"`
}

// GetParams: Request type for `Clients.Get`.
type GetParams struct {
	// ClientID: The ID of the client.
	ClientID string `json:"client_id,omitempty"`
}

// SearchParams: Request type for `Clients.Search`.
type SearchParams struct {
	// Cursor: The `cursor` field allows you to paginate through your results. Each result array is limited to
	// 1000 results. If your query returns more than 1000 results, you will need to paginate the responses
	// using the `cursor`. If you receive a response that includes a non-null `next_cursor` in the
	// `results_metadata` object, repeat the search call with the `next_cursor` value set to the `cursor` field
	// to retrieve the next page of results. Continue to make search calls until the `next_cursor` in the
	// response is null.
	Cursor string `json:"cursor,omitempty"`
	// Limit: The number of search results to return per page. The default limit is 100. A maximum of 1000
	// results can be returned by a single search request. If the total size of your result set is greater than
	// one page size, you must paginate the response. See the `cursor` field.
	Limit uint32 `json:"limit,omitempty"`
	// Query: The optional query object contains the operator, i.e. `AND` or `OR`, and the operands that will
	// filter your results. Only an operator is required. If you include no operands, no filtering will be
	// applied. If you include no query object, it will return all results with no filtering applied.
	Query *m2m.M2MSearchQuery `json:"query,omitempty"`
}

// UpdateParams: Request type for `Clients.Update`.
type UpdateParams struct {
	// ClientID: The ID of the client.
	ClientID string `json:"client_id,omitempty"`
	// ClientName: A human-readable name for the client.
	ClientName string `json:"client_name,omitempty"`
	// ClientDescription: A human-readable description for the client.
	ClientDescription string `json:"client_description,omitempty"`
	// Status: The status of the client - either `active` or `inactive`.
	Status *UpdateRequestStatus `json:"status,omitempty"`
	// Scopes: An array of scopes assigned to the client.
	Scopes []string `json:"scopes,omitempty"`
	// TrustedMetadata: The `trusted_metadata` field contains an arbitrary JSON object of application-specific
	// data. See the [Metadata](https://stytch.com/docs/api/metadata) reference for complete field behavior
	// details.
	TrustedMetadata map[string]any `json:"trusted_metadata,omitempty"`
}

// CreateResponse: Response type for `Clients.Create`.
type CreateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// M2MClient: The M2M Client created by this API call.
	M2MClient m2m.M2MClientWithClientSecret `json:"m2m_client,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// DeleteResponse: Response type for `Clients.Delete`.
type DeleteResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// ClientID: The ID of the client.
	ClientID string `json:"client_id,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// GetResponse: Response type for `Clients.Get`.
type GetResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// M2MClient: The M2M Client affected by this operation.
	M2MClient m2m.M2MClient `json:"m2m_client,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// SearchResponse: Response type for `Clients.Search`.
type SearchResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// M2MClients: An array of M2M Clients that match your search query.
	M2MClients []m2m.M2MClient `json:"m2m_clients,omitempty"`
	// ResultsMetadata: The search `results_metadata` object contains metadata relevant to your specific query
	// like total and `next_cursor`.
	ResultsMetadata m2m.ResultsMetadata `json:"results_metadata,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

// UpdateResponse: Response type for `Clients.Update`.
type UpdateResponse struct {
	// RequestID: Globally unique UUID that is returned with every API call. This value is important to log for
	// debugging purposes; we may ask for this value to help identify a specific API call when helping you
	// debug an issue.
	RequestID string `json:"request_id,omitempty"`
	// M2MClient: The M2M Client affected by this operation.
	M2MClient m2m.M2MClient `json:"m2m_client,omitempty"`
	// StatusCode: The HTTP status code of the response. Stytch follows standard HTTP response status code
	// patterns, e.g. 2XX values equate to success, 3XX values are redirects, 4XX are client errors, and 5XX
	// are server errors.
	StatusCode int32 `json:"status_code,omitempty"`
}

type UpdateRequestStatus string

const (
	UpdateRequestStatusActive   UpdateRequestStatus = "active"
	UpdateRequestStatusInactive UpdateRequestStatus = "inactive"
)
