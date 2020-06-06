// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ItemsDissociate struct{}

func (m ItemsDissociate) ProtocolId() d1proto.MsgCliId {
	return d1proto.ItemsDissociate
}

func (m ItemsDissociate) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsDissociate) Deserialize(extra string) error {
	return nil
}
