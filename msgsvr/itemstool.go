package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type ItemsTool struct {
	JobId int
}

func (m ItemsTool) ProtocolId() retroproto.MsgSvrId {
	return retroproto.ItemsTool
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
