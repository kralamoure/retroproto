// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameFreeMySoul struct{}

func (m GameFreeMySoul) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameFreeMySoul
}

func (m GameFreeMySoul) Serialized() (string, error) {
	return "", nil
}

func (m *GameFreeMySoul) Deserialize(extra string) error {
	return nil
}
