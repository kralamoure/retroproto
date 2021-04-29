// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type JobRemove struct{}

func (m JobRemove) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.JobRemove
}

func (m JobRemove) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *JobRemove) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
