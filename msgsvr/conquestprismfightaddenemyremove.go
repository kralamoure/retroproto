// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightAddEnemyRemove struct{}

func (m ConquestPrismFightAddEnemyRemove) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismFightAddEnemyRemove
}

func (m ConquestPrismFightAddEnemyRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightAddEnemyRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
