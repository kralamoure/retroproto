// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type FightsBlockJoinerExceptParty struct{}

func (m FightsBlockJoinerExceptParty) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.FightsBlockJoinerExceptParty
}

func (m FightsBlockJoinerExceptParty) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *FightsBlockJoinerExceptParty) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
