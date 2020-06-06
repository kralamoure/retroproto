// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SubwayPrismLeave struct{}

func (m SubwayPrismLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SubwayPrismLeave
}

func (m SubwayPrismLeave) Serialized() (string, error) {
	return "", nil
}

func (m *SubwayPrismLeave) Deserialize(extra string) error {
	return nil
}
