// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type SubwayUse struct{}

func (m SubwayUse) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.SubwayUse
}

func (m SubwayUse) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *SubwayUse) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
