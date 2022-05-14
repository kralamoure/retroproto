// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type EmotesRemove struct{}

func (m EmotesRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.EmotesRemove
}

func (m EmotesRemove) MessageName() string {
	return "EmotesRemove"
}

func (m EmotesRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *EmotesRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
