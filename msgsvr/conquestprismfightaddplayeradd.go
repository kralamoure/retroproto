// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismFightAddPlayerAdd struct{}

func (m ConquestPrismFightAddPlayerAdd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestPrismFightAddPlayerAdd
}

func (m ConquestPrismFightAddPlayerAdd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismFightAddPlayerAdd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
