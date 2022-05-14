// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type InfosGetMaps struct{}

func NewInfosGetMaps(extra string) (InfosGetMaps, error) {
	var m InfosGetMaps

	if err := m.Deserialize(extra); err != nil {
		return InfosGetMaps{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m InfosGetMaps) MessageId() retroproto.MsgCliId {
	return retroproto.InfosGetMaps
}

func (m InfosGetMaps) MessageName() string {
	return "InfosGetMaps"
}

func (m InfosGetMaps) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *InfosGetMaps) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
