// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismFightAddPlayerRemove struct{}

func (m ConquestPrismFightAddPlayerRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestPrismFightAddPlayerRemove
}

func (m ConquestPrismFightAddPlayerRemove) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestPrismFightAddPlayerRemove) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
