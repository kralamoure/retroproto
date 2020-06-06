// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SubwayLeave struct{}

func (m SubwayLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SubwayLeave
}

func (m SubwayLeave) Serialized() (string, error) {
	return "", nil
}

func (m *SubwayLeave) Deserialize(extra string) error {
	return nil
}
