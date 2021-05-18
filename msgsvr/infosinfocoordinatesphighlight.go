// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type InfosInfoCoordinatesPHighlight struct{}

func (m InfosInfoCoordinatesPHighlight) ProtocolId() retroproto.MsgSvrId {
	return retroproto.InfosInfoCoordinatesPHighlight
}

func (m InfosInfoCoordinatesPHighlight) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *InfosInfoCoordinatesPHighlight) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
