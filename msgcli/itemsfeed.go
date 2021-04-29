// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type ItemsFeed struct{}

func (m ItemsFeed) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.ItemsFeed
}

func (m ItemsFeed) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ItemsFeed) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
