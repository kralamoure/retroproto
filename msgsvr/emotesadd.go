// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type EmotesAdd struct{}

func (m EmotesAdd) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.EmotesAdd
}

func (m EmotesAdd) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *EmotesAdd) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
