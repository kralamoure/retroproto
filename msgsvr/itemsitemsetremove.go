package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
)

type ItemsItemSetRemove struct {
	Id int
}

func (m ItemsItemSetRemove) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ItemsItemSetRemove
}

func (m ItemsItemSetRemove) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *ItemsItemSetRemove) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	if sli[0] != "" {
		id, err := strconv.ParseInt(sli[0], 10, 32)
		if err != nil {
			return err
		}
		m.Id = int(id)
	}

	return nil
}
