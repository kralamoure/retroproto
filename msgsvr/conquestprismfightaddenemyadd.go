// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismFightAddEnemyAdd struct{}

func (m ConquestPrismFightAddEnemyAdd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestPrismFightAddEnemyAdd
}

func (m ConquestPrismFightAddEnemyAdd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismFightAddEnemyAdd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
