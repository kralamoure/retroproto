// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type BasicsRequestAveragePing struct{}

func (m BasicsRequestAveragePing) ProtocolId() d1proto.MsgCliId {
	return d1proto.BasicsRequestAveragePing
}

func (m BasicsRequestAveragePing) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsRequestAveragePing) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}