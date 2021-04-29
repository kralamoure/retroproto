// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type EmotesRemove struct{}

func (m EmotesRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.EmotesRemove
}

func (m EmotesRemove) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EmotesRemove) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
