// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismFightAddEnemyRemove struct{}

func (m ConquestPrismFightAddEnemyRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestPrismFightAddEnemyRemove
}

func (m ConquestPrismFightAddEnemyRemove) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismFightAddEnemyRemove) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
