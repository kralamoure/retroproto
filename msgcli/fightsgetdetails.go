// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FightsGetDetails struct{}

func (m FightsGetDetails) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FightsGetDetails
}

func (m FightsGetDetails) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FightsGetDetails) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
