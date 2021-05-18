// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestConquestBalance struct{}

func (m ConquestConquestBalance) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ConquestConquestBalance
}

func (m ConquestConquestBalance) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestConquestBalance) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
