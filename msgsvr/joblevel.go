// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type JobLevel struct{}

func NewJobLevel(extra string) (JobLevel, error) {
	var m JobLevel

	if err := m.Deserialize(extra); err != nil {
		return JobLevel{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m JobLevel) MessageId() retroproto.MsgSvrId {
	return retroproto.JobLevel
}

func (m JobLevel) MessageName() string {
	return "JobLevel"
}

func (m JobLevel) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobLevel) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
