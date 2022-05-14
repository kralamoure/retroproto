package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsNothing struct{}

func NewBasicsNothing(extra string) (BasicsNothing, error) {
	var m BasicsNothing

	if err := m.Deserialize(extra); err != nil {
		return BasicsNothing{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsNothing) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsNothing
}

func (m BasicsNothing) MessageName() string {
	return "BasicsNothing"
}

func (m BasicsNothing) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsNothing) Deserialize(extra string) error {
	return nil
}
