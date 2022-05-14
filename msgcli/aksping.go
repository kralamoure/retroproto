// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type AksPing struct{}

func (m AksPing) MessageId() retroproto.MsgCliId {
	return retroproto.AksPing
}

func (m AksPing) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *AksPing) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
