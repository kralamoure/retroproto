// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ConquestRequestBalance struct{}

func (m ConquestRequestBalance) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestRequestBalance
}

func (m ConquestRequestBalance) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestRequestBalance) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
