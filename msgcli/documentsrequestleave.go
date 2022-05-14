// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type DocumentsRequestLeave struct{}

func NewDocumentsRequestLeave(extra string) (DocumentsRequestLeave, error) {
	var m DocumentsRequestLeave

	if err := m.Deserialize(extra); err != nil {
		return DocumentsRequestLeave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m DocumentsRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.DocumentsRequestLeave
}

func (m DocumentsRequestLeave) MessageName() string {
	return "DocumentsRequestLeave"
}

func (m DocumentsRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *DocumentsRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
