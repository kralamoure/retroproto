// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type HousesGuildInfos struct{}

func NewHousesGuildInfos(extra string) (HousesGuildInfos, error) {
	var m HousesGuildInfos

	if err := m.Deserialize(extra); err != nil {
		return HousesGuildInfos{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m HousesGuildInfos) MessageId() retroproto.MsgSvrId {
	return retroproto.HousesGuildInfos
}

func (m HousesGuildInfos) MessageName() string {
	return "HousesGuildInfos"
}

func (m HousesGuildInfos) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *HousesGuildInfos) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
