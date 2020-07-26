// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismFightAddPlayerAdd struct{}

func (m ConquestPrismFightAddPlayerAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestPrismFightAddPlayerAdd
}

func (m ConquestPrismFightAddPlayerAdd) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestPrismFightAddPlayerAdd) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
