// TODO
package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
	"github.com/kralamoure/retroproto/enum"
	"github.com/kralamoure/retroproto/typ"
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

func NewGameActionsSendActions(extra string) (GameActionsSendActions, error) {
	var m GameActionsSendActions

	if err := m.Deserialize(extra); err != nil {
		return GameActionsSendActions{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameActionsSendActions) MessageId() retroproto.MsgCliId {
	return retroproto.GameActionsSendActions
}

func (m GameActionsSendActions) MessageName() string {
	return "GameActionsSendActions"
}

func (m GameActionsSendActions) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameActionsSendActions) Deserialize(extra string) error {
	if len(extra) < 3 {
		return retroproto.ErrInvalidMsg
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
			return retroproto.ErrInvalidMsg
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
			return retroproto.ErrInvalidMsg
		}

		id, err := strconv.ParseInt(extra, 10, 32)
		if err != nil {
			return err
		}
		m.ActionChallenge.ChallengedId = int(id)
	case enum.GameActionType.ChallengeAccept:
		if extra == "" {
			return retroproto.ErrInvalidMsg
		}

		id, err := strconv.ParseInt(extra, 10, 32)
		if err != nil {
			return err
		}
		m.ActionChallengeAccept.ChallengerId = int(id)
	case enum.GameActionType.ChallengeRefuse:
		if extra == "" {
			return retroproto.ErrInvalidMsg
		}

		id, err := strconv.ParseInt(extra, 10, 32)
		if err != nil {
			return err
		}
		m.ActionChallengeRefuse.ChallengerId = int(id)
	default:
		return retroproto.ErrNotImplemented
	}

	return nil
}
