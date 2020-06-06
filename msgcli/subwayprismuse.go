// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type SubwayPrismUse struct{}

func (m SubwayPrismUse) ProtocolId() d1proto.MsgCliId {
	return d1proto.SubwayPrismUse
}

func (m SubwayPrismUse) Serialized() (string, error) {
	return "", nil
}

func (m *SubwayPrismUse) Deserialize(extra string) error {
	return nil
}
