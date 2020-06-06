// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismFightAddEnemyRemove struct{}

func (m ConquestPrismFightAddEnemyRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestPrismFightAddEnemyRemove
}

func (m ConquestPrismFightAddEnemyRemove) Serialized() (string, error) {
	return "", nil
}

func (m *ConquestPrismFightAddEnemyRemove) Deserialize(extra string) error {
	return nil
}
