package emailvalidator

// HunterRespData defines data in email verify API response
type HunterRespData struct {
	Status     string `json:"status"`
	Result     string `json:"result"`
	Score      int    `json:"score"`
	Email      string `json:"email"`
	Regexp     bool   `json:"regexp"`
	Gibberish  bool   `json:"gibberish"`
	Disposable bool   `json:"disposable"`
	Webmail    bool   `json:"webmail"`
	MxRecords  bool   `json:"mx_records"`
	SMTPServer bool   `json:"smtp_server"`
	SMTPCheck  bool   `json:"smtp_check"`
	AcceptAll  bool   `json:"accept_all"`
	Block      bool   `json:"block"`
	Sources    []struct {
		Domain      string `json:"domain"`
		URI         string `json:"uri"`
		ExtractedOn string `json:"extracted_on"`
		LastSeenOn  string `json:"last_seen_on"`
		StillOnPage bool   `json:"still_on_page"`
	} `json:"sources"`
}

// HunterRespMeta defines meta data in email verify API response
type HunterRespMeta struct {
	Params struct {
		Email string `json:"email"`
	} `json:"params"`
}

// HunterValidateEmailResp is returned by hunter's email verify API
type HunterValidateEmailResp struct {
	Data HunterRespData `json:"data"`
	Meta HunterRespMeta `json:"meta"`
}

// IsValid checks if email is valid
func (resp *HunterValidateEmailResp) IsValid() bool {
	return resp.Data.Status == "valid"
}

// IsDeliverable checks if email is deliverable
func (resp *HunterValidateEmailResp) IsDeliverable() bool {
	return resp.Data.Result == "deliverable"
}
