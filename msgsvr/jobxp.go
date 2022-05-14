// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type JobXP struct{}

func NewJobXP(extra string) (JobXP, error) {
	var m JobXP

	if err := m.Deserialize(extra); err != nil {
		return JobXP{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m JobXP) MessageId() retroproto.MsgSvrId {
	return retroproto.JobXP
}

func (m JobXP) MessageName() string {
	return "JobXP"
}

func (m JobXP) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobXP) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
