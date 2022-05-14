// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type DocumentsLeave struct{}

func NewDocumentsLeave(extra string) (DocumentsLeave, error) {
	var m DocumentsLeave

	if err := m.Deserialize(extra); err != nil {
		return DocumentsLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m DocumentsLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.DocumentsLeave
}

func (m DocumentsLeave) MessageName() string {
	return "DocumentsLeave"
}

func (m DocumentsLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *DocumentsLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
