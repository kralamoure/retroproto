package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ItemsTool struct {
	JobId int
}

func NewItemsTool(extra string) (ItemsTool, error) {
	var m ItemsTool

	if err := m.Deserialize(extra); err != nil {
		return ItemsTool{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ItemsTool) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsTool
}

func (m ItemsTool) MessageName() string {
	return "ItemsTool"
}

func (m ItemsTool) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.JobId), nil
}

func (m *ItemsTool) Deserialize(extra string) error {
	if len(extra) < 1 {
		return retroproto.ErrInvalidMsg
	}

	jobId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.JobId = int(jobId)

	return nil
}
