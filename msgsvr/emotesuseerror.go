// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type EmotesUseError struct{}

func (m EmotesUseError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.EmotesUseError
}

func (m EmotesUseError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EmotesUseError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
