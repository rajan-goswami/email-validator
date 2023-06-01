package emailvalidator

// HunterValidateEmailResp is returned by hunter's email verify API
type HunterValidateEmailResp struct {
	Data struct {
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
	} `json:"data"`
	Meta struct {
		Params struct {
			Email string `json:"email"`
		} `json:"params"`
	} `json:"meta"`
}

func (resp *HunterValidateEmailResp) IsValid() bool {
	return resp.Data.Status == "valid"
}
