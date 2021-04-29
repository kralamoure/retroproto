// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FightsGetList struct{}

func (m FightsGetList) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FightsGetList
}

func (m FightsGetList) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FightsGetList) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
