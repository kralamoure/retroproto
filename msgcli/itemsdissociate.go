// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ItemsDissociate struct{}

func (m ItemsDissociate) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ItemsDissociate
}

func (m ItemsDissociate) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ItemsDissociate) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
