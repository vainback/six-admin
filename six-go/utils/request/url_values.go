package request

import (
	"fmt"
	"net/url"
)

type URLParam struct {
	values url.Values
}

func UrlParam(k string, v ...any) *URLParam {
	var p URLParam
	p.values = url.Values{
		k: p.strings(v),
	}
	return &p
}

func (p *URLParam) Add(k string, v ...any) *URLParam {
	p.values[k] = p.strings(v)
	return p
}

func (p *URLParam) Builder() string {
	return p.values.Encode()
}

func (p *URLParam) strings(v []any) []string {
	if len(v) == 0 {
		return nil
	}
	vs := make([]string, len(v))
	if len(v) > 0 {
		for i, vv := range v {
			vs[i] = fmt.Sprint(vv)
		}
	}
	return vs
}
