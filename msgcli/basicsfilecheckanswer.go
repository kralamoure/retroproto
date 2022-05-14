// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsFileCheckAnswer struct{}

func NewBasicsFileCheckAnswer(extra string) (BasicsFileCheckAnswer, error) {
	var m BasicsFileCheckAnswer

	if err := m.Deserialize(extra); err != nil {
		return BasicsFileCheckAnswer{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsFileCheckAnswer) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsFileCheckAnswer
}

func (m BasicsFileCheckAnswer) MessageName() string {
	return "BasicsFileCheckAnswer"
}

func (m BasicsFileCheckAnswer) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsFileCheckAnswer) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
