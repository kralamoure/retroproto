// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ItemsDissociate struct{}

func (m ItemsDissociate) MessageId() retroproto.MsgCliId {
	return retroproto.ItemsDissociate
}

func (m ItemsDissociate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsDissociate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
