// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type EmotesUseError struct{}

func (m EmotesUseError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.EmotesUseError
}

func (m EmotesUseError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EmotesUseError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
