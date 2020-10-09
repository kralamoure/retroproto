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
	return "", d1proto.ErrNotImplemented
}

func (m *ItemsDissociate) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
