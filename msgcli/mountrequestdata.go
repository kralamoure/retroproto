package msgcli

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/kralamoure/retroproto"
)

type MountRequestData struct {
	Id       int
	Validity time.Time
}

func (m MountRequestData) MessageId() retroproto.MsgCliId {
	return retroproto.MountRequestData
}

func (m MountRequestData) MessageName() string {
	return "MountRequestData"
}

func (m MountRequestData) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.Id, m.Validity.Unix()*1000), nil
}

func (m *MountRequestData) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")
	if len(sli) != 2 {
		return retroproto.ErrInvalidMsg
	}

	if sli[0] != "undefined" {
		id, err := strconv.ParseInt(sli[0], 10, 64)
		if err != nil {
			return err
		}
		m.Id = int(id)
	}

	if sli[1] != "undefined" {
		millis, err := strconv.ParseInt(sli[1], 10, 64)
		if err != nil {
			return err
		}
		m.Validity = time.Unix(millis/1000, 0)
	}

	return nil
}
