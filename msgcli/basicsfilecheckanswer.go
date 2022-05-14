// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsFileCheckAnswer struct{}

func (m BasicsFileCheckAnswer) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsFileCheckAnswer
}

func (m BasicsFileCheckAnswer) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsFileCheckAnswer) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
