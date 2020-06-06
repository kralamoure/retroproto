// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type SubwayRequestPrismLeave struct{}

func (m SubwayRequestPrismLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.SubwayRequestPrismLeave
}

func (m SubwayRequestPrismLeave) Serialized() (string, error) {
	return "", nil
}

func (m *SubwayRequestPrismLeave) Deserialize(extra string) error {
	return nil
}
