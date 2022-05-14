// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type JobRemove struct{}

func NewJobRemove(extra string) (JobRemove, error) {
	var m JobRemove

	if err := m.Deserialize(extra); err != nil {
		return JobRemove{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m JobRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.JobRemove
}

func (m JobRemove) MessageName() string {
	return "JobRemove"
}

func (m JobRemove) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *JobRemove) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
