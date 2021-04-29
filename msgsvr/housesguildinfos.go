// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type HousesGuildInfos struct{}

func (m HousesGuildInfos) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.HousesGuildInfos
}

func (m HousesGuildInfos) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *HousesGuildInfos) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
