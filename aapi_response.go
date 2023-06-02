package emailvalidator

// AAValidateEmailResp is returned by Abstract API's email verify API
type AAValidateEmailResp struct {
	Email          string `json:"email"`
	Autocorrect    string `json:"autocorrect"`
	Deliverability string `json:"deliverability"`
	QualityScore   string `json:"quality_score"`
	IsValidFormat  struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_valid_format"`
	IsFreeEmail struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_free_email"`
	IsDisposableEmail struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_disposable_email"`
	IsRoleEmail struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_role_email"`
	IsCatchallEmail struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_catchall_email"`
	IsMxFound struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_mx_found"`
	IsSMTPValid struct {
		Value bool   `json:"value"`
		Text  string `json:"text"`
	} `json:"is_smtp_valid"`
}

// IsValid checks if email is valid
func (resp *AAValidateEmailResp) IsValid() bool {
	return resp.IsValidFormat.Value
}
