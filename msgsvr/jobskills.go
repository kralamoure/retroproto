// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type JobSkills struct{}

func (m JobSkills) ProtocolId() retroproto.MsgSvrId {
	return retroproto.JobSkills
}

func (m JobSkills) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobSkills) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
