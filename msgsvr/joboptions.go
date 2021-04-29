// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type JobOptions struct{}

func (m JobOptions) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.JobOptions
}

func (m JobOptions) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *JobOptions) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
