package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type ItemsItemSetRemove struct {
	Id int
}

func (m ItemsItemSetRemove) MessageId() retroproto.MsgSvrId {
	return retroproto.ItemsItemSetRemove
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
