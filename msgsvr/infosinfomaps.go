// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type InfosInfoMaps struct{}

func NewInfosInfoMaps(extra string) (InfosInfoMaps, error) {
	var m InfosInfoMaps

	if err := m.Deserialize(extra); err != nil {
		return InfosInfoMaps{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m InfosInfoMaps) MessageId() retroproto.MsgSvrId {
	return retroproto.InfosInfoMaps
}

func (m InfosInfoMaps) MessageName() string {
	return "InfosInfoMaps"
}

func (m InfosInfoMaps) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *InfosInfoMaps) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
