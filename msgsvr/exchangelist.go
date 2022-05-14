// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ExchangeList struct{}

func NewExchangeList(extra string) (ExchangeList, error) {
	var m ExchangeList

	if err := m.Deserialize(extra); err != nil {
		return ExchangeList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ExchangeList) MessageId() retroproto.MsgSvrId {
	return retroproto.ExchangeList
}

func (m ExchangeList) MessageName() string {
	return "ExchangeList"
}

func (m ExchangeList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ExchangeList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
