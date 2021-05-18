// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type HousesGuildInfos struct{}

func (m HousesGuildInfos) ProtocolId() retroproto.MsgSvrId {
	return retroproto.HousesGuildInfos
}

func (m HousesGuildInfos) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesGuildInfos) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
