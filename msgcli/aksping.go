// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type AksPing struct{}

func (m AksPing) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.AksPing
}

func (m AksPing) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *AksPing) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
