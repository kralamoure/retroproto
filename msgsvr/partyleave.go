// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type PartyLeave struct{}

func (m PartyLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.PartyLeave
}

func (m PartyLeave) Serialized() (string, error) {
	return "", nil
}

func (m *PartyLeave) Deserialize(extra string) error {
	return nil
}
