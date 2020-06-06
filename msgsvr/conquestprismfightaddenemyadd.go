// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismFightAddEnemyAdd struct{}

func (m ConquestPrismFightAddEnemyAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestPrismFightAddEnemyAdd
}

func (m ConquestPrismFightAddEnemyAdd) Serialized() (string, error) {
	return "", nil
}

func (m *ConquestPrismFightAddEnemyAdd) Deserialize(extra string) error {
	return nil
}
