// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type SubwayRequestPrismLeave struct{}

func (m SubwayRequestPrismLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.SubwayRequestPrismLeave
}

func (m SubwayRequestPrismLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SubwayRequestPrismLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
