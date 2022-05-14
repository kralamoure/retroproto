// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightAddEnemyAdd struct{}

func (m ConquestPrismFightAddEnemyAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismFightAddEnemyAdd
}

func (m ConquestPrismFightAddEnemyAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightAddEnemyAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
