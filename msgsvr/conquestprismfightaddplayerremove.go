// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismFightAddPlayerRemove struct{}

func (m ConquestPrismFightAddPlayerRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestPrismFightAddPlayerRemove
}

func (m ConquestPrismFightAddPlayerRemove) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismFightAddPlayerRemove) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
