package msgsvr

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type BasicsSubscriberRestrictionAdd struct {
	DialogId int
}

func (m BasicsSubscriberRestrictionAdd) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsSubscriberRestrictionAdd
}

func (m BasicsSubscriberRestrictionAdd) MessageName() string {
	return "BasicsSubscriberRestrictionAdd"
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
