// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestPrismInfosJoined struct{}

func NewConquestPrismInfosJoined(extra string) (ConquestPrismInfosJoined, error) {
	var m ConquestPrismInfosJoined

	if err := m.Deserialize(extra); err != nil {
		return ConquestPrismInfosJoined{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestPrismInfosJoined) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismInfosJoined
}

func (m ConquestPrismInfosJoined) MessageName() string {
	return "ConquestPrismInfosJoined"
}

func (m ConquestPrismInfosJoined) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismInfosJoined) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
