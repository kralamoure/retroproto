package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsGetDate struct{}

func NewBasicsGetDate(extra string) (BasicsGetDate, error) {
	var m BasicsGetDate

	if err := m.Deserialize(extra); err != nil {
		return BasicsGetDate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

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
