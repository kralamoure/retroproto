// TODO
package msgcli

import (
	"strconv"
	"strings"

	"github.com/kralamoure/d1encoding"
	"github.com/kralamoure/d1encoding/enum"
	"github.com/kralamoure/d1encoding/typ"
)

type GameActionsSendActions struct {
	ActionType            int
	ActionMovement        GameActionsSendActionsActionMovement
	ActionChallenge       GameActionsSendActionsActionChallenge
	ActionChallengeAccept GameActionsSendActionsActionChallengeAccept
	ActionChallengeRefuse GameActionsSendActionsActionChallengeRefuse
}

type GameActionsSendActionsActionMovement struct {
	DirAndCells []typ.CommonDirAndCell
}

type GameActionsSendActionsActionChallenge struct {
	ChallengedId int
}

type GameActionsSendActionsActionChallengeAccept struct {
	ChallengerId int
}

type GameActionsSendActionsActionChallengeRefuse struct {
	ChallengerId int
}

func (m GameActionsSendActions) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameActionsSendActions
}

func (m GameActionsSendActions) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GameActionsSendActions) Deserialize(extra string) error {
	if len(extra) < 3 {
		return d1encoding.ErrInvalidMsg
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
			return d1encoding.ErrInvalidMsg
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
			return d1encoding.ErrInvalidMsg
		}

		id, err := strconv.ParseInt(extra, 10, 32)
		if err != nil {
			return err
		}
		m.ActionChallenge.ChallengedId = int(id)
	case enum.GameActionType.ChallengeAccept:
		if extra == "" {
			return d1encoding.ErrInvalidMsg
		}

		id, err := strconv.ParseInt(extra, 10, 32)
		if err != nil {
			return err
		}
		m.ActionChallengeAccept.ChallengerId = int(id)
	case enum.GameActionType.ChallengeRefuse:
		if extra == "" {
			return d1encoding.ErrInvalidMsg
		}

		id, err := strconv.ParseInt(extra, 10, 32)
		if err != nil {
			return err
		}
		m.ActionChallengeRefuse.ChallengerId = int(id)
	default:
		return d1encoding.ErrNotImplemented
	}

	return nil
}
