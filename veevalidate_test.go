package veevalidate

import (
	"encoding/json"
	"testing"
)

func TestAll(t *testing.T) {
	tests := map[string]string{
		"":                          New().String(),
		"required":                  New().Required().String(),
		"required|min_value:0":      New().Required().MinValue(0).String(),
		"required|max_value:999":    New().Required().MaxValue(999).String(),
		"url":                       New().Url(nil).String(),
		"url:protocols:ftp":         New().Url([]string{"ftp"}).String(),
		"url:require_protocol:true": New().Url(UrlRequireProtocol).String(),
		"url:require_protocol:false:require_host:true": New().Url(UrlDoNotRequireProtocol, UrlRequireHost).String(),
	}

	for expect, got := range tests {
		if expect != got {
			t.Errorf("Expected '%s' got '%s'", expect, got)
		}
	}
}

func TestJson(t *testing.T) {
	v := New().Required().IP_or_FQDN()

	j, err := json.Marshal(v)
	if err != nil {
		t.Errorf(err.Error())
	}
	if string(j) != `"required|ip_or_fqdn"` {
		t.Errorf("Bad json: %s", string(j))
	}
}
