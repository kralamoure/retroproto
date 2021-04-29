// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type InfosInfoCoordinatesPHighlight struct{}

func (m InfosInfoCoordinatesPHighlight) ProtocolId() d1proto.MsgSvrId {
	return d1proto.InfosInfoCoordinatesPHighlight
}

func (m InfosInfoCoordinatesPHighlight) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *InfosInfoCoordinatesPHighlight) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
