// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type FightsBlockJoinerExceptParty struct{}

func (m FightsBlockJoinerExceptParty) ProtocolId() d1proto.MsgCliId {
	return d1proto.FightsBlockJoinerExceptParty
}

func (m FightsBlockJoinerExceptParty) Serialized() (string, error) {
	return "", nil
}

func (m *FightsBlockJoinerExceptParty) Deserialize(extra string) error {
	return nil
}
