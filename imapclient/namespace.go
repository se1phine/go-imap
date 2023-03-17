package imapclient

import (
	"fmt"

	"github.com/emersion/go-imap/v2/internal/imapwire"
)

// Namespace sends a NAMESPACE command.
//
// This command requires support for IMAP4rev2 or the NAMESPACE extension.
func (c *Client) Namespace() *NamespaceCommand {
	cmd := &NamespaceCommand{}
	c.beginCommand("NAMESPACE", cmd).end()
	return cmd
}

// NamespaceCommand is a NAMESPACE command.
type NamespaceCommand struct {
	cmd
	data NamespaceData
}

func (cmd *NamespaceCommand) Wait() (*NamespaceData, error) {
	return &cmd.data, cmd.cmd.Wait()
}

// NamespaceData is the data returned by the NAMESPACE command.
type NamespaceData struct {
	Personal []NamespaceDescriptor
	Other    []NamespaceDescriptor
	Shared   []NamespaceDescriptor
}

// NamespaceDescriptor describes a namespace.
type NamespaceDescriptor struct {
	Prefix string
	Delim  rune
}

func readNamespaceResponse(dec *imapwire.Decoder) (*NamespaceData, error) {
	var (
		data NamespaceData
		err  error
	)

	data.Personal, err = readNamespace(dec)
	if err != nil {
		return nil, err
	}

	if !dec.ExpectSP() {
		return nil, dec.Err()
	}

	data.Other, err = readNamespace(dec)
	if err != nil {
		return nil, err
	}

	if !dec.ExpectSP() {
		return nil, dec.Err()
	}

	data.Shared, err = readNamespace(dec)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func readNamespace(dec *imapwire.Decoder) ([]NamespaceDescriptor, error) {
	var l []NamespaceDescriptor
	err := dec.ExpectNList(func() error {
		descr, err := readNamespaceDescr(dec)
		if err != nil {
			return fmt.Errorf("in namespace-descr: %v", err)
		}
		l = append(l, *descr)
		return nil
	})
	return l, err
}

func readNamespaceDescr(dec *imapwire.Decoder) (*NamespaceDescriptor, error) {
	var descr NamespaceDescriptor

	if !dec.ExpectSpecial('(') || !dec.ExpectString(&descr.Prefix) || !dec.ExpectSP() {
		return nil, dec.Err()
	}

	var err error
	descr.Delim, err = readDelim(dec)
	if err != nil {
		return nil, err
	}

	// Skip namespace-response-extensions
	for dec.SP() {
		if !dec.DiscardValue() {
			return nil, dec.Err()
		}
	}

	if !dec.ExpectSpecial(')') {
		return nil, dec.Err()
	}

	return &descr, nil
}