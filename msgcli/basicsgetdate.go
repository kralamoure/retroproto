package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsGetDate struct{}

func (m BasicsGetDate) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsGetDate
}

func (m BasicsGetDate) MessageName() string {
	return "BasicsGetDate"
}

func (m BasicsGetDate) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsGetDate) Deserialize(extra string) error {
	return nil
}
