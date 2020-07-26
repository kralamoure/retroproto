// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type JobOptions struct{}

func (m JobOptions) ProtocolId() d1proto.MsgSvrId {
	return d1proto.JobOptions
}

func (m JobOptions) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *JobOptions) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
