// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type InfosInfoMaps struct{}

func (m InfosInfoMaps) MessageId() retroproto.MsgSvrId {
	return retroproto.InfosInfoMaps
}

func (m InfosInfoMaps) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *InfosInfoMaps) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
