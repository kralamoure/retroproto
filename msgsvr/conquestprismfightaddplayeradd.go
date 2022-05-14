// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightAddPlayerAdd struct{}

func (m ConquestPrismFightAddPlayerAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismFightAddPlayerAdd
}

func (m ConquestPrismFightAddPlayerAdd) MessageName() string {
	return "ConquestPrismFightAddPlayerAdd"
}

func (m ConquestPrismFightAddPlayerAdd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightAddPlayerAdd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
