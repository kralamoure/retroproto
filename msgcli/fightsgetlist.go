// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FightsGetList struct{}

func (m FightsGetList) ProtocolId() retroproto.MsgCliId {
	return retroproto.FightsGetList
}

func (m FightsGetList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsGetList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
