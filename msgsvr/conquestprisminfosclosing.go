// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismInfosClosing struct{}

func NewConquestPrismInfosClosing(extra string) (ConquestPrismInfosClosing, error) {
	var m ConquestPrismInfosClosing

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismInfosClosing{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismInfosClosing) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismInfosClosing
}

func (m ConquestPrismInfosClosing) MessageName() string {
	return "ConquestPrismInfosClosing"
}

func (m ConquestPrismInfosClosing) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismInfosClosing) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
