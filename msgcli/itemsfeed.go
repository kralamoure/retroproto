// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type ItemsFeed struct{}

func (m ItemsFeed) MessageId() retroproto.MsgCliId {
	return retroproto.ItemsFeed
}

func (m ItemsFeed) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ItemsFeed) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
