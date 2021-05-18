// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type DocumentsRequestLeave struct{}

func (m DocumentsRequestLeave) ProtocolId() retroproto.MsgCliId {
	return retroproto.DocumentsRequestLeave
}

func (m DocumentsRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *DocumentsRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
