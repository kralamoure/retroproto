// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type SubwayRequestLeave struct{}

func (m SubwayRequestLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.SubwayRequestLeave
}

func (m SubwayRequestLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SubwayRequestLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
