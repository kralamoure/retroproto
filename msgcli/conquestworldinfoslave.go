// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestWorldInfosLave struct{}

func (m ConquestWorldInfosLave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ConquestWorldInfosLave
}

func (m ConquestWorldInfosLave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestWorldInfosLave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
