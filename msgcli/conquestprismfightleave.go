// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismFightLeave struct{}

func (m ConquestPrismFightLeave) ProtocolId() retroproto.MsgCliId {
	return retroproto.ConquestPrismFightLeave
}

func (m ConquestPrismFightLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismFightLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
