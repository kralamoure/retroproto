// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FightsGetList struct{}

func (m FightsGetList) MessageId() retroproto.MsgCliId {
	return retroproto.FightsGetList
}

func (m FightsGetList) MessageName() string {
	return "FightsGetList"
}

func (m FightsGetList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsGetList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
