// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FightsNeedHelp struct{}

func (m FightsNeedHelp) ProtocolId() d1proto.MsgCliId {
	return d1proto.FightsNeedHelp
}

func (m FightsNeedHelp) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *FightsNeedHelp) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}