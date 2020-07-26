// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type InfosInfoMaps struct{}

func (m InfosInfoMaps) ProtocolId() d1proto.MsgSvrId {
	return d1proto.InfosInfoMaps
}

func (m InfosInfoMaps) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *InfosInfoMaps) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
