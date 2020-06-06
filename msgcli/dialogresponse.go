// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type DialogResponse struct{}

func (m DialogResponse) ProtocolId() d1proto.MsgCliId {
	return d1proto.DialogResponse
}

func (m DialogResponse) Serialized() (string, error) {
	return "", nil
}

func (m *DialogResponse) Deserialize(extra string) error {
	return nil
}
