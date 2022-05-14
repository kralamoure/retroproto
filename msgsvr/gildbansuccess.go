// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GildBanSuccess struct{}

func NewGildBanSuccess(extra string) (GildBanSuccess, error) {
	var m GildBanSuccess

	if err := m.Deserialize(extra); err != nil {
		return GildBanSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GildBanSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.GildBanSuccess
}

func (m GildBanSuccess) MessageName() string {
	return "GildBanSuccess"
}

func (m GildBanSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildBanSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
