// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type FightsBlockJoinerExceptParty struct{}

func (m FightsBlockJoinerExceptParty) MessageId() retroproto.MsgCliId {
	return retroproto.FightsBlockJoinerExceptParty
}

func (m FightsBlockJoinerExceptParty) MessageName() string {
	return "FightsBlockJoinerExceptParty"
}

func (m FightsBlockJoinerExceptParty) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *FightsBlockJoinerExceptParty) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
