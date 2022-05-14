// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type JobSkills struct{}

func NewJobSkills(extra string) (JobSkills, error) {
	var m JobSkills

	if err := m.Deserialize(extra); err != nil {
		return JobSkills{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m JobSkills) MessageId() retroproto.MsgSvrId {
	return retroproto.JobSkills
}

func (m JobSkills) MessageName() string {
	return "JobSkills"
}

func (m JobSkills) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobSkills) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
