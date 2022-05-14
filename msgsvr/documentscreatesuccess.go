// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type DocumentsCreateSuccess struct{}

func NewDocumentsCreateSuccess(extra string) (DocumentsCreateSuccess, error) {
	var m DocumentsCreateSuccess

	if err := m.Deserialize(extra); err != nil {
		return DocumentsCreateSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m DocumentsCreateSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.DocumentsCreateSuccess
}

func (m DocumentsCreateSuccess) MessageName() string {
	return "DocumentsCreateSuccess"
}

func (m DocumentsCreateSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *DocumentsCreateSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
