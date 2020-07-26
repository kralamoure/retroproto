// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type EmotesUseError struct{}

func (m EmotesUseError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.EmotesUseError
}

func (m EmotesUseError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *EmotesUseError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
