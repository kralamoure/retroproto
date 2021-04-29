// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type BasicsFileCheckAnswer struct{}

func (m BasicsFileCheckAnswer) ProtocolId() d1proto.MsgCliId {
	return d1proto.BasicsFileCheckAnswer
}

func (m BasicsFileCheckAnswer) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsFileCheckAnswer) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
