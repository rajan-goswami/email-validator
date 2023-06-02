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
	defaultHunterAPIURL         = "https://api.hunter.io"
	defaultHunterAPIVersion     = "/v2"
	defaultEmailVerifierSubPath = "/email-verifier"
)

// HunterAPIClient defines http client for making Hunter's REST calls
type HunterAPIClient struct {
	apiKey  string
	baseURL *url.URL
	client  *internal.HTTPClient
}

// NewHunterAPIClient creates new Hunter API client for calling Hunter's REST APIs
func NewHunterAPIClient(apiKey string, options ...HunterAPIOptionFunc) (*HunterAPIClient, error) {
	if apiKey == "" {
		return nil, ErrEmptyAPIKey
	}

	opts, err := parseHunterAPIOptions(options...)
	if err != nil {
		return nil, err
	}

	baseURL := defaultHunterAPIURL
	apiVersion := defaultHunterAPIVersion

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

	clientOpts := []func(*internal.Option){
		internal.WithLimit(opts.rate.Limit),
		internal.WithLimitInterval(opts.rate.Interval),
	}
	if opts.blocking {
		clientOpts = append(clientOpts, internal.WithBlocking())
	}

	client := internal.NewClient(clientOpts...)

	return &HunterAPIClient{
		apiKey:  apiKey,
		baseURL: u,
		client:  client,
	}, nil
}

// Validate validates email address and returns hunter api's response
func (hc *HunterAPIClient) Validate(email string) (*HunterValidateEmailResp, error) {
	// Add Query params
	params := url.Values{}
	params.Add("api_key", hc.apiKey)
	params.Add("email", email)
	hc.baseURL.RawQuery = params.Encode()

	// Create HTTP Request
	url := path.Join(hc.baseURL.Path, defaultEmailVerifierSubPath)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Make REST Call
	resp, err := hc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	validateResp := &HunterValidateEmailResp{}
	if err = json.Unmarshal(body, validateResp); err != nil {
		return nil, err
	}

	return validateResp, nil
}
