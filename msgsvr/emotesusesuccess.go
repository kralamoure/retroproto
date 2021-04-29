// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type EmotesUseSuccess struct{}

func (m EmotesUseSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.EmotesUseSuccess
}

func (m EmotesUseSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EmotesUseSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
