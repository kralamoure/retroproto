// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameRequestReady struct{}

func (m GameRequestReady) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameRequestReady
}

func (m GameRequestReady) Serialized() (string, error) {
	return "", nil
}

func (m *GameRequestReady) Deserialize(extra string) error {
	return nil
}
