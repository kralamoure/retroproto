// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type InfosGetMaps struct{}

func (m InfosGetMaps) MessageId() retroproto.MsgCliId {
	return retroproto.InfosGetMaps
}

func (m InfosGetMaps) MessageName() string {
	return "InfosGetMaps"
}

func (m InfosGetMaps) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *InfosGetMaps) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
