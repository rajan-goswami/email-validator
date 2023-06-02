package emailvalidator

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/r-goswami/email-validator/internal"
)

const (
	defaultAbstractAPIURL     = "https://emailvalidation.abstractapi.com"
	defaultAbstractAPIVersion = "/v1"
)

// AbstractAPIClient defines http client for making AbstractAPI's REST calls
type AbstractAPIClient struct {
	apiKey  string
	baseURL *url.URL
	client  *internal.HTTPClient
}

// NewAbstractAPIClient creates new AbstractAPI client for calling email validation REST API
func NewAbstractAPIClient(apiKey string, options ...AbstractAPIOptionFunc) (*AbstractAPIClient, error) {
	if apiKey == "" {
		return nil, ErrEmptyAPIKey
	}

	opts, err := ParseAbstractAPIOptions(options...)
	if err != nil {
		return nil, err
	}

	baseURL := defaultAbstractAPIURL
	apiVersion := defaultAbstractAPIVersion

	if opts.baseURL != nil {
		baseURL = opts.baseURL.String()
	}
	if opts.apiVersion != "" {
		apiVersion = opts.apiVersion
	}

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, apiVersion)

	clientOpts := []internal.OptionFunc{
		internal.WithLimit(opts.rate.Limit),
		internal.WithLimitInterval(opts.rate.Interval),
	}
	if opts.blocking {
		clientOpts = append(clientOpts, internal.WithBlocking())
	}

	client := internal.NewClient(clientOpts...)

	return &AbstractAPIClient{
		apiKey:  apiKey,
		baseURL: u,
		client:  client,
	}, nil
}

func (ac *AbstractAPIClient) Validate(email string) (*AAValidateEmailResp, error) {
	// Add Query params
	params := url.Values{}
	params.Add("api_key", ac.apiKey)
	params.Add("email", email)
	params.Add("auto_correct", "false")
	ac.baseURL.RawQuery = params.Encode()

	// Create HTTP Request
	req, err := http.NewRequest(http.MethodGet, ac.baseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	// Make REST Call
	resp, err := ac.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	validateResp := &AAValidateEmailResp{}
	if err = json.Unmarshal(body, validateResp); err != nil {
		return nil, err
	}

	return validateResp, nil
}
