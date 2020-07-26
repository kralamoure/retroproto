// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type SubwayUse struct{}

func (m SubwayUse) ProtocolId() d1proto.MsgCliId {
	return d1proto.SubwayUse
}

func (m SubwayUse) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *SubwayUse) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
