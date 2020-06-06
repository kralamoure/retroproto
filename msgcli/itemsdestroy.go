// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ItemsDestroy struct{}

func (m ItemsDestroy) ProtocolId() d1proto.MsgCliId {
	return d1proto.ItemsDestroy
}

func (m ItemsDestroy) Serialized() (string, error) {
	return "", nil
}

func (m *ItemsDestroy) Deserialize(extra string) error {
	return nil
}
