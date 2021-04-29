// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsFileCheckAnswer struct{}

func (m BasicsFileCheckAnswer) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.BasicsFileCheckAnswer
}

func (m BasicsFileCheckAnswer) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsFileCheckAnswer) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
