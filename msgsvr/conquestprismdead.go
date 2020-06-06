// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismDead struct{}

func (m ConquestPrismDead) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestPrismDead
}

func (m ConquestPrismDead) Serialized() (string, error) {
	return "", nil
}

func (m *ConquestPrismDead) Deserialize(extra string) error {
	return nil
}
