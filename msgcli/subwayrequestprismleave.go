// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type SubwayRequestPrismLeave struct{}

func (m SubwayRequestPrismLeave) ProtocolId() retroproto.MsgCliId {
	return retroproto.SubwayRequestPrismLeave
}

func (m SubwayRequestPrismLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *SubwayRequestPrismLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
