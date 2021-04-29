// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type InfosInfoCoordinatesPHighlight struct{}

func (m InfosInfoCoordinatesPHighlight) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.InfosInfoCoordinatesPHighlight
}

func (m InfosInfoCoordinatesPHighlight) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *InfosInfoCoordinatesPHighlight) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
