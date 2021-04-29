// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type InfosInfoMaps struct{}

func (m InfosInfoMaps) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.InfosInfoMaps
}

func (m InfosInfoMaps) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *InfosInfoMaps) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
