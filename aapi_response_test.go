package emailvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAAValidateEmailResp_IsValid(t *testing.T) {

	tests := []struct {
		name     string
		response AAValidateEmailResp
		isValid  bool
	}{
		{
			name: "Should return true if email is found to be valid in response",
			response: AAValidateEmailResp{
				IsValidFormat: struct {
					Value bool   "json:\"value\""
					Text  string "json:\"text\""
				}{
					Value: true, Text: "true",
				}},
			isValid: true,
		},
		{
			name: "Should return false if email is found to be invalid in response",
			response: AAValidateEmailResp{
				IsValidFormat: struct {
					Value bool   "json:\"value\""
					Text  string "json:\"text\""
				}{
					Value: false, Text: "false",
				}},
			isValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isValid, tt.response.IsValid())
		})
	}
}

func TestAAValidateEmailResp_IsDeliverable(t *testing.T) {

	tests := []struct {
		name     string
		response AAValidateEmailResp
		isValid  bool
	}{
		{
			name: "Should return true if email is found to be deliverable in response",
			response: AAValidateEmailResp{
				Deliverability: "DELIVERABLE",
			},
			isValid: true,
		},
		{
			name: "Should return false if email is found to be undeliverable in response",
			response: AAValidateEmailResp{
				Deliverability: "UNDELIVERABLE",
			},
			isValid: false,
		},
		{
			name: "Should return false if email is found to be risky in response",
			response: AAValidateEmailResp{
				Deliverability: "RISKY",
			},
			isValid: false,
		},
		{
			name: "Should return false if email is found to be unknown in response",
			response: AAValidateEmailResp{
				Deliverability: "UNKNOWN",
			},
			isValid: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isValid, tt.response.IsDeliverable())
		})
	}
}
