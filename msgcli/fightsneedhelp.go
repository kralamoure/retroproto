// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FightsNeedHelp struct{}

func (m FightsNeedHelp) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FightsNeedHelp
}

func (m FightsNeedHelp) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FightsNeedHelp) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
