// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type JobOptions struct{}

func NewJobOptions(extra string) (JobOptions, error) {
	var m JobOptions

	if err := m.Deserialize(extra); err != nil {
		return JobOptions{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m JobOptions) MessageId() retroproto.MsgSvrId {
	return retroproto.JobOptions
}

func (m JobOptions) MessageName() string {
	return "JobOptions"
}

func (m JobOptions) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobOptions) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
