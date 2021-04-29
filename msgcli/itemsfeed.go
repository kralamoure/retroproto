// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ItemsFeed struct{}

func (m ItemsFeed) ProtocolId() d1proto.MsgCliId {
	return d1proto.ItemsFeed
}

func (m ItemsFeed) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ItemsFeed) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
