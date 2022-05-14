// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightAddPlayerRemove struct{}

func (m ConquestPrismFightAddPlayerRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismFightAddPlayerRemove
}

func (m ConquestPrismFightAddPlayerRemove) MessageName() string {
	return "ConquestPrismFightAddPlayerRemove"
}

func (m ConquestPrismFightAddPlayerRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightAddPlayerRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
