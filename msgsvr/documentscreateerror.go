// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type DocumentsCreateError struct{}

func NewDocumentsCreateError(extra string) (DocumentsCreateError, error) {
	var m DocumentsCreateError

	if err := m.Deserialize(extra); err != nil {
		return DocumentsCreateError{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m DocumentsCreateError) MessageId() retroproto.MsgSvrId {
	return retroproto.DocumentsCreateError
}

func (m DocumentsCreateError) MessageName() string {
	return "DocumentsCreateError"
}

func (m DocumentsCreateError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *DocumentsCreateError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
