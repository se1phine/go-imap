package responses

import (
	"github.com/se1phine/go-imap"
)

type ID struct {
	ResStr []string
}

func (r *ID) Handle(resp imap.Resp) error {
	name, fields, ok := imap.ParseNamedResp(resp)
	if !ok || name != "ID" {
		return ErrUnhandled
	} else if len(fields) < 1 {
		return errNotEnoughFields
	}

	if fields[0] == nil {
		return nil
	}
	res := fields[0].([]interface{})
	var resStrSlices []string
	for i := 0; i < len(res); i++ {
		resStrSlices = append(resStrSlices, res[i].(string))
	}
	r.ResStr = resStrSlices

	return nil
}
