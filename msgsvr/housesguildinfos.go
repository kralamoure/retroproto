// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type HousesGuildInfos struct{}

func (m HousesGuildInfos) ProtocolId() d1proto.MsgSvrId {
	return d1proto.HousesGuildInfos
}

func (m HousesGuildInfos) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *HousesGuildInfos) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
