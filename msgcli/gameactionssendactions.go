// TODO
package msgcli

import (
	"strconv"
	"strings"

	"github.com/kralamoure/d1proto"
	"github.com/kralamoure/d1proto/enum"
	"github.com/kralamoure/d1proto/typ"
)

type GameActionsSendActions struct {
	ActionType      int
	ActionMovement  GameActionsSendActionsActionMovement
	ActionChallenge GameActionsSendActionsActionChallenge
}

type GameActionsSendActionsActionMovement struct {
	DirAndCells []typ.CommonDirAndCell
}

type GameActionsSendActionsActionChallenge struct {
	Id int
}

func (m GameActionsSendActions) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameActionsSendActions
}

func (m GameActionsSendActions) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GameActionsSendActions) Deserialize(extra string) error {
	if len(extra) < 3 {
		return d1proto.ErrInvalidMsg
	}

	actionType := extra[:3]
	extra = strings.TrimPrefix(extra, actionType)

	actionTypeN, err := strconv.ParseInt(actionType, 10, 32)
	if err != nil {
		return err
	}
	m.ActionType = int(actionTypeN)

	switch m.ActionType {
	case enum.GameActionType.Movement:
		if extra == "" || len(extra)%3 != 0 {
			return d1proto.ErrInvalidMsg
		}

		for i := 0; i < len(extra); i += 3 {
			dirAndCell := &typ.CommonDirAndCell{}

			err := dirAndCell.Deserialize(extra[i : i+3])
			if err != nil {
				return err
			}

			m.ActionMovement.DirAndCells = append(m.ActionMovement.DirAndCells, *dirAndCell)
		}
	case enum.GameActionType.Challenge:
		if extra == "" {
			return d1proto.ErrInvalidMsg
		}

		id, err := strconv.ParseInt(extra, 10, 32)
		if err != nil {
			return err
		}
		m.ActionChallenge.Id = int(id)
	default:
		return d1proto.ErrNotImplemented
	}

	return nil
}
