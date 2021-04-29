// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsRequestAveragePing struct{}

func (m BasicsRequestAveragePing) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.BasicsRequestAveragePing
}

func (m BasicsRequestAveragePing) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsRequestAveragePing) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
