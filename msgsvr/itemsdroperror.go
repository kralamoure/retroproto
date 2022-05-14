// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ItemsDropError struct{}

func (m ItemsDropError) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsDropError
}

func (m ItemsDropError) MessageName() string {
	return "ItemsDropError"
}

func (m ItemsDropError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsDropError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
