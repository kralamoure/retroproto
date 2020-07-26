// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type JobRemove struct{}

func (m JobRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.JobRemove
}

func (m JobRemove) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *JobRemove) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
