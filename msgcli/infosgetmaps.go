// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type InfosGetMaps struct{}

func (m InfosGetMaps) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.InfosGetMaps
}

func (m InfosGetMaps) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *InfosGetMaps) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
