// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ConquestWorldInfosLave struct{}

func (m ConquestWorldInfosLave) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestWorldInfosLave
}

func (m ConquestWorldInfosLave) MessageName() string {
	return "ConquestWorldInfosLave"
}

func (m ConquestWorldInfosLave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestWorldInfosLave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
