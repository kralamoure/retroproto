// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FightsGetDetails struct{}

func (m FightsGetDetails) ProtocolId() retroproto.MsgCliId {
	return retroproto.FightsGetDetails
}

func (m FightsGetDetails) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsGetDetails) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
