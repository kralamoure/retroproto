// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestConquestBalance struct{}

func (m ConquestConquestBalance) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestConquestBalance
}

func (m ConquestConquestBalance) Serialized() (string, error) {
	return "", nil
}

func (m *ConquestConquestBalance) Deserialize(extra string) error {
	return nil
}
