// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type InfosInfoCoordinatesPHighlight struct{}

func NewInfosInfoCoordinatesPHighlight(extra string) (InfosInfoCoordinatesPHighlight, error) {
	var m InfosInfoCoordinatesPHighlight

	if err := m.Deserialize(extra); err != nil {
		return InfosInfoCoordinatesPHighlight{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m InfosInfoCoordinatesPHighlight) MessageId() retroproto.MsgSvrId {
	return retroproto.InfosInfoCoordinatesPHighlight
}

func (m InfosInfoCoordinatesPHighlight) MessageName() string {
	return "InfosInfoCoordinatesPHighlight"
}

func (m InfosInfoCoordinatesPHighlight) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *InfosInfoCoordinatesPHighlight) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
