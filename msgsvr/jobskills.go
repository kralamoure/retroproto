// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type JobSkills struct{}

func (m JobSkills) ProtocolId() d1proto.MsgSvrId {
	return d1proto.JobSkills
}

func (m JobSkills) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *JobSkills) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
