// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestConquestBalance struct{}

func (m ConquestConquestBalance) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestConquestBalance
}

func (m ConquestConquestBalance) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestConquestBalance) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
