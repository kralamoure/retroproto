// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type SubwayRequestLeave struct{}

func (m SubwayRequestLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.SubwayRequestLeave
}

func (m SubwayRequestLeave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *SubwayRequestLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
