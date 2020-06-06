// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ItemsDrop struct{}

func (m ItemsDrop) ProtocolId() d1proto.MsgCliId {
	return d1proto.ItemsDrop
}

func (m ItemsDrop) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsDrop) Deserialize(extra string) error {
	return nil
}
