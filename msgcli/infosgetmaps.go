// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type InfosGetMaps struct{}

func (m InfosGetMaps) ProtocolId() d1proto.MsgCliId {
	return d1proto.InfosGetMaps
}

func (m InfosGetMaps) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *InfosGetMaps) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}