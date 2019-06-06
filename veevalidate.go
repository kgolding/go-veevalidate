package veevalidate

import (
	"strconv"
	"strings"
)

type VBuilder struct {
	rules []string
}

func New() VBuilder {
	return VBuilder{
		rules: []string{},
	}
}

func (b VBuilder) MarshalJSON() ([]byte, error) {
	return []byte(`"` + b.String() + `"`), nil
}

func (b VBuilder) String() string {
	return strings.Join(b.rules, "|")
}

func (b VBuilder) Required() VBuilder {
	b.rules = append(b.rules, "required")
	return b
}

func (b VBuilder) IP_or_FQDN() VBuilder {
	b.rules = append(b.rules, "ip_or_fqdn")
	return b
}

type UrlOption int

const (
	UrlRequireProtocol = UrlOption(iota)
	UrlDoNotRequireProtocol
	UrlRequireTLD
	UrlDoNotRequireTLD
	UrlRequireHost
	UrlDoNotRequireHost
)

func (b VBuilder) Url(options ...interface{}) VBuilder {
	rule := "url"
	params := []string{}
	for _, option := range options {
		switch option.(type) {
		case []string:
			params = append(params, "protocols:"+strings.Join(option.([]string), ","))
		case UrlOption:
			switch option.(UrlOption) {
			case UrlRequireProtocol:
				params = append(params, "require_protocol:true")
			case UrlDoNotRequireProtocol:
				params = append(params, "require_protocol:false")
			case UrlRequireHost:
				params = append(params, "require_host:true")
			case UrlDoNotRequireHost:
				params = append(params, "require_host:false")
			case UrlRequireTLD:
				params = append(params, "require_tld:true")
			case UrlDoNotRequireTLD:
				params = append(params, "require_tld:false")
			}
		}
	}
	if len(params) > 0 {
		rule += ":" + strings.Join(params, ":")
	}
	b.rules = append(b.rules, rule)
	return b
}

func (b VBuilder) MinValue(value int) VBuilder {
	b.rules = append(b.rules, "min_value:"+strconv.Itoa(value))
	return b
}

func (b VBuilder) MaxValue(value int) VBuilder {
	b.rules = append(b.rules, "max_value:"+strconv.Itoa(value))
	return b
}
