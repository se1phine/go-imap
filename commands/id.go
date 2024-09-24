package commands

import (
	"errors"
	"github.com/se1phine/go-imap"
)

type ID struct {
	IDString string
}

func (cmd *ID) Command() *imap.Command {
	return &imap.Command{
		Name:      "ID",
		Arguments: []interface{}{cmd.IDString},
	}
}

func (cmd *ID) Parse(fields []interface{}) error {
	if len(fields) < 1 {
		return errors.New("No enough arguments")
	}

	return nil
}
