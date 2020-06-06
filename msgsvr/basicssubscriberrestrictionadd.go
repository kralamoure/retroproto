package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type BasicsSubscriberRestrictionAdd struct {
	DialogId int
}

func (m BasicsSubscriberRestrictionAdd) ProtocolId() d1proto.MsgSvrId {
	return d1proto.BasicsSubscriberRestrictionAdd
}

func (m BasicsSubscriberRestrictionAdd) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.DialogId), nil
}

func (m *BasicsSubscriberRestrictionAdd) Deserialize(extra string) error {
	dialogId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.DialogId = int(dialogId)

	return nil
}
