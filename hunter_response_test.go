package emailvalidator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHunterValidateEmailResp_IsValid(t *testing.T) {

	tests := []struct {
		name     string
		response HunterValidateEmailResp
		isValid  bool
	}{
		{
			name:     "Should return true if email is found to be valid in response",
			response: HunterValidateEmailResp{Data: HunterRespData{Status: "valid"}},
			isValid:  true,
		},
		{
			name:     "Should return false if email is found to be invalid in response",
			response: HunterValidateEmailResp{Data: HunterRespData{Status: "invalid"}},
			isValid:  false,
		},
		{
			name:     "Should return false if email is found to be unknown in response",
			response: HunterValidateEmailResp{Data: HunterRespData{Status: "unknown"}},
			isValid:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isValid, tt.response.IsValid())
		})
	}
}

func TestHunterValidateEmailResp_IsDeliverable(t *testing.T) {

	tests := []struct {
		name     string
		response HunterValidateEmailResp
		isValid  bool
	}{
		{
			name:     "Should return true if email is found to be deliverable in response",
			response: HunterValidateEmailResp{Data: HunterRespData{Result: "deliverable"}},
			isValid:  true,
		},
		{
			name:     "Should return false if email is found to be undeliverable in response",
			response: HunterValidateEmailResp{Data: HunterRespData{Result: "undeliverable"}},
			isValid:  false,
		},
		{
			name:     "Should return false if email is found to be risky in response",
			response: HunterValidateEmailResp{Data: HunterRespData{Result: "risky"}},
			isValid:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.isValid, tt.response.IsDeliverable())
		})
	}
}
