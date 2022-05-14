// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FightsNeedHelp struct{}

func (m FightsNeedHelp) MessageId() retroproto.MsgCliId {
	return retroproto.FightsNeedHelp
}

func (m FightsNeedHelp) MessageName() string {
	return "FightsNeedHelp"
}

func (m FightsNeedHelp) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsNeedHelp) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
