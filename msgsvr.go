package d1proto

import "strings"

type MsgSvr interface {
	ProtocolId() (id MsgSvrId)
	Serialized() (extra string, err error)
	Deserialize(extra string) error
}

type MsgSvrId string

func MsgSvrNameByID(id MsgSvrId) (name string, ok bool) {
	name, ok = msgSvrNameByID[id]
	if !ok {
		name = "Unknown"
	}
	return
}

func MsgSvrIdByPkt(pkt string) (id MsgSvrId, ok bool) {
	if pkt == "" {
		return
	}

	for _, v := range MsgSvrIds {
		if strings.HasPrefix(pkt, string(v)) {
			ok = true
			if len(v) >= len(id) {
				id = v
			}
		}
	}
	return
}

const (
	AksHelloConnect         MsgSvrId = "HC"
	AksHelloGame            MsgSvrId = "HG"
	AksPong                 MsgSvrId = "p"
	AksQuickPong            MsgSvrId = "q"
	AksRPing                MsgSvrId = "rping"
	AksServerMessage        MsgSvrId = "M"
	AksServerWillDisconnect MsgSvrId = "k"

	BasicsNoticed                     MsgSvrId = "BN"
	BasicsAuthorizedCommandError      MsgSvrId = "BAE"
	BasicsAuthorizedCommandSuccess    MsgSvrId = "BAT"
	BasicsAuthorizedLine              MsgSvrId = "BAL"
	BasicsAuthorizedCommandPrompt     MsgSvrId = "BAP"
	BasicsAuthorizedCommandClear      MsgSvrId = "BAC"
	BasicsAuthorizedInterfaceOpen     MsgSvrId = "BAIO"
	BasicsAuthorizedInterfaceClose    MsgSvrId = "BAIC"
	BasicsTime                        MsgSvrId = "BT"
	BasicsDate                        MsgSvrId = "BD"
	BasicsWhoIsError                  MsgSvrId = "BWE"
	BasicsWhoIsSuccess                MsgSvrId = "BWK"
	BasicsSubscriberRestrictionAdd    MsgSvrId = "BP+"
	BasicsSubscriberRestrictionRemove MsgSvrId = "BP-"
	BasicsFileCheck                   MsgSvrId = "BC"
	BasicsAveragePing                 MsgSvrId = "Bp"

	AccountCommunity                     MsgSvrId = "Ac"
	AccountPseudo                        MsgSvrId = "Ad"
	AccountLoginSuccess                  MsgSvrId = "AlK"
	AccountLoginError                    MsgSvrId = "AlE"
	AccountCharactersListError           MsgSvrId = "ALE"
	AccountCharactersListSuccess         MsgSvrId = "ALK"
	AccountServersListError              MsgSvrId = "AxE"
	AccountServersListSuccess            MsgSvrId = "AxK"
	AccountCharacterAddError             MsgSvrId = "AAE"
	AccountCharacterAddSuccess           MsgSvrId = "AAK"
	AccountTicketResponseError           MsgSvrId = "ATE"
	AccountTicketResponseSuccess         MsgSvrId = "ATK"
	AccountSelectServerError             MsgSvrId = "AXE"
	AccountSelectServerSuccess           MsgSvrId = "AXK"
	AccountSelectServerPlainSuccess      MsgSvrId = "AYK"
	AccountCharacterSelectedError        MsgSvrId = "ASE"
	AccountCharacterSelectedSuccess      MsgSvrId = "ASK"
	AccountStats                         MsgSvrId = "As"
	AccountNewLevel                      MsgSvrId = "AN"
	AccountRestrictions                  MsgSvrId = "AR"
	AccountHosts                         MsgSvrId = "AH"
	AccountRescue                        MsgSvrId = "Ar"
	AccountGiftsList                     MsgSvrId = "Ag"
	AccountGiftStoredError               MsgSvrId = "AGE"
	AccountGiftStoredSuccess             MsgSvrId = "AGK"
	AccountQueue                         MsgSvrId = "Aq"
	AccountNewQueue                      MsgSvrId = "Af"
	AccountRegionalVersion               MsgSvrId = "AV"
	AccountCharacterNameGeneratedError   MsgSvrId = "APE"
	AccountCharacterNameGeneratedSuccess MsgSvrId = "APK"
	AccountKey                           MsgSvrId = "AK"
	AccountSecretQuestion                MsgSvrId = "AQ"
	AccountCharacterDeleteError          MsgSvrId = "ADE"
	AccountCharacterDeleteSuccess        MsgSvrId = "ADK"
	AccountCharacterMigrationAskConfirm  MsgSvrId = "AM?"
	AccountCharacterMigrationError       MsgSvrId = "AME"
	AccountCharacterMigrationSuccess     MsgSvrId = "AMK"
	AccountFriendServerList              MsgSvrId = "AF"
	AccountMiniClipInfo                  MsgSvrId = "Am"

	GameCreateError                 MsgSvrId = "GCE"
	GameCreateSuccess               MsgSvrId = "GCK"
	GameJoin                        MsgSvrId = "GJ"
	GamePositionStart               MsgSvrId = "GP"
	GameReady                       MsgSvrId = "GR"
	GameStartToPlay                 MsgSvrId = "GS"
	GameEnd                         MsgSvrId = "GE"
	GameMovementRemove              MsgSvrId = "GM|-"
	GameMovement                    MsgSvrId = "GM"
	GameChallenge                   MsgSvrId = "Gc"
	GameTeam                        MsgSvrId = "Gt"
	GameLeave                       MsgSvrId = "GV"
	GameFlag                        MsgSvrId = "Gf"
	GamePlayersCoordinates          MsgSvrId = "GIC"
	GameEffect                      MsgSvrId = "GIE"
	GameClearAllEffect              MsgSvrId = "GIe"
	GamePVP                         MsgSvrId = "GIP"
	GameMapData                     MsgSvrId = "GDM"
	GameMapLoaded                   MsgSvrId = "GDK"
	GameCellData                    MsgSvrId = "GDC"
	GameZoneData                    MsgSvrId = "GDZ"
	GameCellObject                  MsgSvrId = "GDO"
	GameFrameObject2                MsgSvrId = "GDF"
	GameFrameObjectExternal         MsgSvrId = "GDE"
	GameFightChallenge              MsgSvrId = "Gd"
	GameFightChallengeUpdateError   MsgSvrId = "GdO"
	GameFightChallengeUpdateSuccess MsgSvrId = "GdK"
	GameTurnStart                   MsgSvrId = "GTS"
	GameTurnFinish                  MsgSvrId = "GTF"
	GameTurnList                    MsgSvrId = "GTL"
	GameTurnMiddle                  MsgSvrId = "GTM"
	GameTurnReady                   MsgSvrId = "GTR"
	GameExtraClip                   MsgSvrId = "GX"
	GameFightOption                 MsgSvrId = "Go"
	GameGameOver                    MsgSvrId = "GO"

	GameActions       MsgSvrId = "GA"
	GameActionsStart  MsgSvrId = "GAS"
	GameActionsFinish MsgSvrId = "GAF"

	ChatMessageError           MsgSvrId = "cME"
	ChatMessageSuccess         MsgSvrId = "cMK"
	ChatServerMessage          MsgSvrId = "cs"
	ChatSmiley                 MsgSvrId = "cS"
	ChatSubscribeChannelAdd    MsgSvrId = "cC+"
	ChatSubscribeChannelRemove MsgSvrId = "cC-"

	DialogCustomAction  MsgSvrId = "DA"
	DialogCreateError   MsgSvrId = "DCE"
	DialogCreateSuccess MsgSvrId = "DCK"
	DialogQuestion      MsgSvrId = "DQ"
	DialogLeave         MsgSvrId = "DV"
	DialogPause         MsgSvrId = "DP"

	InfosInfoMaps                  MsgSvrId = "IM"
	InfosCompass                   MsgSvrId = "IC"
	InfosInfoCoordinatesPHighlight MsgSvrId = "IH"
	InfosMessage                   MsgSvrId = "Im"
	InfosQuantity                  MsgSvrId = "IQ"
	InfosLifeRestoreTimerStart     MsgSvrId = "ILS"
	InfosLifeRestoreTimerFinish    MsgSvrId = "ILF"

	SpellsList                MsgSvrId = "SL"
	SpellsChangeOption        MsgSvrId = "SLo"
	SpellsUpgradeSpellError   MsgSvrId = "SUE"
	SpellsUpgradeSpellSuccess MsgSvrId = "SUK"
	SpellsSpellBoost          MsgSvrId = "SB"
	SpellsSpellForgetShow     MsgSvrId = "SF+"
	SpellsSpellForgetClose    MsgSvrId = "SF-"

	ItemsAccessories      MsgSvrId = "Oa"
	ItemsDropError        MsgSvrId = "ODE"
	ItemsDropSuccess      MsgSvrId = "ODK"
	ItemsAddError         MsgSvrId = "OAE"
	ItemsAddSuccess       MsgSvrId = "OAK"
	ItemsChange           MsgSvrId = "OC"
	ItemsRemove           MsgSvrId = "OR"
	ItemsQuantity         MsgSvrId = "OQ"
	ItemsMovement         MsgSvrId = "OM"
	ItemsTool             MsgSvrId = "OT"
	ItemsWeight           MsgSvrId = "Ow"
	ItemsItemSetAdd       MsgSvrId = "OS+"
	ItemsItemSetRemove    MsgSvrId = "OS-"
	ItemsItemUseCondition MsgSvrId = "OK"
	ItemsItemFound        MsgSvrId = "OF"

	FriendsAddFriendError      MsgSvrId = "FAE"
	FriendsAddFriendSuccess    MsgSvrId = "FAK"
	FriendsRemoveFriendError   MsgSvrId = "FDE"
	FriendsRemoveFriendSuccess MsgSvrId = "FDK"
	FriendsFriendsList         MsgSvrId = "FL"
	FriendsSpouse              MsgSvrId = "FS"
	FriendsNotifyChange        MsgSvrId = "FO"

	EnemiesAddEnemyError      MsgSvrId = "iAE"
	EnemiesAddEnemySuccess    MsgSvrId = "iAK"
	EnemiesRemoveEnemyError   MsgSvrId = "iDE"
	EnemiesRemoveEnemySuccess MsgSvrId = "iDK"
	EnemiesEnemiesList        MsgSvrId = "iL"

	KeyCreate     MsgSvrId = "KC"
	KeyKeyError   MsgSvrId = "KKE"
	KeyKeySuccess MsgSvrId = "KKK"
	KeyLeave      MsgSvrId = "KL"

	JobSkills  MsgSvrId = "JS"
	JobXP      MsgSvrId = "JX"
	JobLevel   MsgSvrId = "JN"
	JobRemove  MsgSvrId = "JR"
	JobOptions MsgSvrId = "JO"

	ExchangeRequestError                      MsgSvrId = "ERE"
	ExchangeRequestSuccess                    MsgSvrId = "ERK"
	ExchangeReady                             MsgSvrId = "EK"
	ExchangeLeaveError                        MsgSvrId = "EVE"
	ExchangeLeaveSuccess                      MsgSvrId = "EVK"
	ExchangeCreateError                       MsgSvrId = "ECE"
	ExchangeCreateSuccess                     MsgSvrId = "ECK"
	ExchangeCraftError                        MsgSvrId = "EcE"
	ExchangeCraftSuccess                      MsgSvrId = "EcK"
	ExchangeLocalMovementError                MsgSvrId = "EME"
	ExchangeLocalMovementSuccess              MsgSvrId = "EMK"
	ExchangeLocalDistantError                 MsgSvrId = "EmE"
	ExchangeLocalDistantSuccess               MsgSvrId = "EmK"
	ExchangeCoopMovementError                 MsgSvrId = "ErE"
	ExchangeCoopMovementSuccess               MsgSvrId = "ErK"
	ExchangePayMovementError                  MsgSvrId = "EpE"
	ExchangePayMovementSuccess                MsgSvrId = "EpK"
	ExchangeStorageMovementError              MsgSvrId = "EsE"
	ExchangeStorageMovementSuccess            MsgSvrId = "EsK"
	ExchangePlayerShopMovementError           MsgSvrId = "EiE"
	ExchangePlayerShopMovementSuccess         MsgSvrId = "EiK"
	ExchangeCraftPublicMode                   MsgSvrId = "EW"
	ExchangeMountStorage                      MsgSvrId = "Ee"
	ExchangeMountPark                         MsgSvrId = "Ef"
	ExchangeMountPods                         MsgSvrId = "Ew"
	ExchangeList                              MsgSvrId = "EL"
	ExchangeSellError                         MsgSvrId = "ESE"
	ExchangeSellSuccess                       MsgSvrId = "ESK"
	ExchangeBuyError                          MsgSvrId = "EBE"
	ExchangeBuySuccess                        MsgSvrId = "EBK"
	ExchangeAskOfflineExchange                MsgSvrId = "Eq"
	ExchangeSearchError                       MsgSvrId = "EHSE"
	ExchangeSearchSuccess                     MsgSvrId = "EHSK"
	ExchangeBigStoreTypeItemsList             MsgSvrId = "EHL"
	ExchangeBigStoreTypeItemsMovementAdd      MsgSvrId = "EHM+"
	ExchangeBigStoreTypeItemsMovementRemove   MsgSvrId = "EHM-"
	ExchangeBigStoreItemsList                 MsgSvrId = "EHl"
	ExchangeBigStoreItemsMovementAdd          MsgSvrId = "EHm+"
	ExchangeBigStoreItemsMovementRemove       MsgSvrId = "EHm-"
	ExchangeBigStoreItemMiddlePriceInBigStore MsgSvrId = "EHP"
	ExchangeCrafterReferenceAdd               MsgSvrId = "EHj+"
	ExchangeCrafterReferenceRemove            MsgSvrId = "EHj-"
	ExchangeCraftLoop                         MsgSvrId = "EA"
	ExchangeCraftLoopEnd                      MsgSvrId = "Ea"

	HousesListAdd        MsgSvrId = "hL+"
	HousesListRemove     MsgSvrId = "hL-"
	HousesProperties     MsgSvrId = "hP"
	HousesLockedProperty MsgSvrId = "hX"
	HousesCreate         MsgSvrId = "hC"
	HousesSellError      MsgSvrId = "hSE"
	HousesSellSuccess    MsgSvrId = "hSK"
	HousesBuyError       MsgSvrId = "hBE"
	HousesBuySuccess     MsgSvrId = "hBK"
	HousesLeave          MsgSvrId = "hV"
	HousesGuildInfos     MsgSvrId = "hG"

	StoragesListAdd        MsgSvrId = "sL+"
	StoragesListRemove     MsgSvrId = "sL-"
	StoragesLockedProperty MsgSvrId = "sX"

	EmotesUseError   MsgSvrId = "eUE"
	EmotesUseSuccess MsgSvrId = "eUK"
	EmotesList       MsgSvrId = "eL"
	EmotesAdd        MsgSvrId = "eA"
	EmotesRemove     MsgSvrId = "eR"
	EmotesDirection  MsgSvrId = "eD"

	DocumentsCreateError   MsgSvrId = "dCE"
	DocumentsCreateSuccess MsgSvrId = "dCK"
	DocumentsLeave         MsgSvrId = "dV"

	GuildNew                         MsgSvrId = "gn"
	GuildCreateError                 MsgSvrId = "gCE"
	GuildCreateSuccess               MsgSvrId = "gCK"
	GuildStats                       MsgSvrId = "gS"
	GuildInfosGeneral                MsgSvrId = "gIG"
	GuildInfosMembers                MsgSvrId = "gIM"
	GuildInfosBoosts                 MsgSvrId = "gIB"
	GuildInfosMountPark              MsgSvrId = "gIF"
	GuildInfosTaxCollectorsMovement  MsgSvrId = "gITM"
	GuildInfosTaxCollectorsPlayers   MsgSvrId = "gITP"
	GuildInfosTaxCollectorsAttackers MsgSvrId = "gITp"
	GuildInfosHouses                 MsgSvrId = "gIH"
	GuildJoinError                   MsgSvrId = "gJE"
	GuildJoinSuccess                 MsgSvrId = "gJK"
	GuildJoinDistantSuccess          MsgSvrId = "gJC"
	GuildRequestLocal                MsgSvrId = "gJR"
	GuildRequestDistant              MsgSvrId = "gJr"
	GuildLeave                       MsgSvrId = "gV"
	GildBanError                     MsgSvrId = "gKE"
	GildBanSuccess                   MsgSvrId = "gKK"
	GildHireTaxCollectorError        MsgSvrId = "gHE"
	GildHireTaxCollectorSuccess      MsgSvrId = "gHK"
	GildTaxCollectorAttacked         MsgSvrId = "gA"
	GildTaxCollectorInfo             MsgSvrId = "gT"
	GildUserInterfaceOpen            MsgSvrId = "gU"

	WaypointsCreate   MsgSvrId = "WC"
	WaypointsLeave    MsgSvrId = "WV"
	WaypointsUseError MsgSvrId = "WU"

	SubwayCreate      MsgSvrId = "Wc"
	SubwayLeave       MsgSvrId = "Wv"
	SubwayUseError    MsgSvrId = "Wu"
	SubwayPrismCreate MsgSvrId = "Wp"
	SubwayPrismLeave  MsgSvrId = "Ww"

	SubareasList                  MsgSvrId = "al"
	SubareasAlignmentModification MsgSvrId = "am"

	ConquestAreaAlignmentChanged      MsgSvrId = "aM"
	ConquestPrismInfosJoined          MsgSvrId = "CIJ"
	ConquestPrismInfosClosing         MsgSvrId = "CIV"
	ConquestConquestBonus             MsgSvrId = "CB"
	ConquestPrismAttacked             MsgSvrId = "CA"
	ConquestPrismSurvived             MsgSvrId = "CS"
	ConquestPrismDead                 MsgSvrId = "CD"
	ConquestPrismFightAddPlayerAdd    MsgSvrId = "CP+"
	ConquestPrismFightAddPlayerRemove MsgSvrId = "CP-"
	ConquestPrismFightAddEnemyAdd     MsgSvrId = "Cp+"
	ConquestPrismFightAddEnemyRemove  MsgSvrId = "Cp-"
	ConquestWorldData                 MsgSvrId = "CW"
	ConquestConquestBalance           MsgSvrId = "Cb"

	SpecializationSet    MsgSvrId = "ZS"
	SpecializationChange MsgSvrId = "ZC"

	FightsCount   MsgSvrId = "fC"
	FightsList    MsgSvrId = "fL"
	FightsDetails MsgSvrId = "fD"

	TutorialCreate    MsgSvrId = "TC"
	TutorialShowTip   MsgSvrId = "TT"
	TutorialGameBegin MsgSvrId = "TB"

	QuestsList MsgSvrId = "QL"
	QuestsStep MsgSvrId = "QS"

	PartyInviteError   MsgSvrId = "PIE"
	PartyInviteSuccess MsgSvrId = "PIK"
	PartyLeader        MsgSvrId = "PL"
	PartyRefuse        MsgSvrId = "PR"
	PartyAccept        MsgSvrId = "PA"
	PartyCreateError   MsgSvrId = "PCE"
	PartyCreateSuccess MsgSvrId = "PCS"
	PartyLeave         MsgSvrId = "PV"
	PartyFollowError   MsgSvrId = "PFE"
	PartyFollowSuccess MsgSvrId = "PFK"
	PartyMovement      MsgSvrId = "PM"

	MountEquipError   MsgSvrId = "ReE"
	MountEquipAdd     MsgSvrId = "Re+"
	MountEquipRemove  MsgSvrId = "Re-"
	MountXP           MsgSvrId = "Rx"
	MountName         MsgSvrId = "Rn"
	MountData         MsgSvrId = "Rd"
	MountMountPark    MsgSvrId = "Rp"
	MountMountParkBuy MsgSvrId = "RD"
	MountLeave        MsgSvrId = "Rv"
	MountRidingState  MsgSvrId = "Rr"
)

var msgSvrNameByID = map[MsgSvrId]string{
	AksHelloConnect:         "AksHelloConnect",
	AksHelloGame:            "AksHelloGame",
	AksPong:                 "AksPong",
	AksQuickPong:            "AksQuickPong",
	AksRPing:                "AksRPing",
	AksServerMessage:        "AksServerMessage",
	AksServerWillDisconnect: "AksServerWillDisconnect",

	BasicsNoticed:                     "BasicsNoticed",
	BasicsAuthorizedCommandError:      "BasicsAuthorizedCommandError",
	BasicsAuthorizedCommandSuccess:    "BasicsAuthorizedCommandSuccess",
	BasicsAuthorizedLine:              "BasicsAuthorizedLine",
	BasicsAuthorizedCommandPrompt:     "BasicsAuthorizedCommandPrompt",
	BasicsAuthorizedCommandClear:      "BasicsAuthorizedCommandClear",
	BasicsAuthorizedInterfaceOpen:     "BasicsAuthorizedInterfaceOpen",
	BasicsAuthorizedInterfaceClose:    "BasicsAuthorizedInterfaceClose",
	BasicsTime:                        "BasicsTime",
	BasicsDate:                        "BasicsDate",
	BasicsWhoIsError:                  "BasicsWhoIsError",
	BasicsWhoIsSuccess:                "BasicsWhoIsSuccess",
	BasicsSubscriberRestrictionAdd:    "BasicsSubscriberRestrictionAdd",
	BasicsSubscriberRestrictionRemove: "BasicsSubscriberRestrictionRemove",
	BasicsFileCheck:                   "BasicsFileCheck",
	BasicsAveragePing:                 "BasicsAveragePing",

	AccountCommunity:                     "AccountCommunity",
	AccountPseudo:                        "AccountPseudo",
	AccountLoginSuccess:                  "AccountLoginSuccess",
	AccountLoginError:                    "AccountLoginError",
	AccountCharactersListError:           "AccountCharactersListError",
	AccountCharactersListSuccess:         "AccountCharactersListSuccess",
	AccountServersListError:              "AccountServersListError",
	AccountServersListSuccess:            "AccountServersListSuccess",
	AccountCharacterAddError:             "AccountCharacterAddError",
	AccountCharacterAddSuccess:           "AccountCharacterAddSuccess",
	AccountTicketResponseError:           "AccountTicketResponseError",
	AccountTicketResponseSuccess:         "AccountTicketResponseSuccess",
	AccountSelectServerError:             "AccountSelectServerError",
	AccountSelectServerSuccess:           "AccountSelectServerSuccess",
	AccountSelectServerPlainSuccess:      "AccountSelectServerPlainSuccess",
	AccountCharacterSelectedError:        "AccountCharacterSelectedError",
	AccountCharacterSelectedSuccess:      "AccountCharacterSelectedSuccess",
	AccountStats:                         "AccountStats",
	AccountNewLevel:                      "AccountNewLevel",
	AccountRestrictions:                  "AccountRestrictions",
	AccountHosts:                         "AccountHosts",
	AccountRescue:                        "AccountRescue",
	AccountGiftsList:                     "AccountGiftsList",
	AccountGiftStoredError:               "AccountGiftStoredError",
	AccountGiftStoredSuccess:             "AccountGiftStoredSuccess",
	AccountQueue:                         "AccountQueue",
	AccountNewQueue:                      "AccountNewQueue",
	AccountRegionalVersion:               "AccountRegionalVersion",
	AccountCharacterNameGeneratedError:   "AccountCharacterNameGeneratedError",
	AccountCharacterNameGeneratedSuccess: "AccountCharacterNameGeneratedSuccess",
	AccountKey:                           "AccountKey",
	AccountSecretQuestion:                "AccountSecretQuestion",
	AccountCharacterDeleteError:          "AccountCharacterDeleteError",
	AccountCharacterDeleteSuccess:        "AccountCharacterDeleteSuccess",
	AccountCharacterMigrationAskConfirm:  "AccountCharacterMigrationAskConfirm",
	AccountCharacterMigrationError:       "AccountCharacterMigrationError",
	AccountCharacterMigrationSuccess:     "AccountCharacterMigrationSuccess",
	AccountFriendServerList:              "AccountFriendServerList",
	AccountMiniClipInfo:                  "AccountMiniClipInfo",

	GameCreateError:                 "GameCreateError",
	GameCreateSuccess:               "GameCreateSuccess",
	GameJoin:                        "GameJoin",
	GamePositionStart:               "GamePositionStart",
	GameReady:                       "GameReady",
	GameStartToPlay:                 "GameStartToPlay",
	GameEnd:                         "GameEnd",
	GameMovementRemove:              "GameMovementRemove",
	GameMovement:                    "GameMovement",
	GameChallenge:                   "GameChallenge",
	GameTeam:                        "GameTeam",
	GameLeave:                       "GameLeave",
	GameFlag:                        "GameFlag",
	GamePlayersCoordinates:          "GamePlayersCoordinates",
	GameEffect:                      "GameEffect",
	GameClearAllEffect:              "GameClearAllEffect",
	GamePVP:                         "GamePVP",
	GameMapData:                     "GameMapData",
	GameMapLoaded:                   "GameMapLoaded",
	GameCellData:                    "GameCellData",
	GameZoneData:                    "GameZoneData",
	GameCellObject:                  "GameCellObject",
	GameFrameObject2:                "GameFrameObject2",
	GameFrameObjectExternal:         "GameFrameObjectExternal",
	GameFightChallenge:              "GameFightChallenge",
	GameFightChallengeUpdateError:   "GameFightChallengeUpdateError",
	GameFightChallengeUpdateSuccess: "GameFightChallengeUpdateSuccess",
	GameTurnStart:                   "GameTurnStart",
	GameTurnFinish:                  "GameTurnFinish",
	GameTurnList:                    "GameTurnList",
	GameTurnMiddle:                  "GameTurnMiddle",
	GameTurnReady:                   "GameTurnReady",
	GameExtraClip:                   "GameExtraClip",
	GameFightOption:                 "GameFightOption",
	GameGameOver:                    "GameGameOver",

	GameActions:       "GameActions",
	GameActionsStart:  "GameActionsStart",
	GameActionsFinish: "GameActionsFinish",

	ChatMessageError:           "ChatMessageError",
	ChatMessageSuccess:         "ChatMessageSuccess",
	ChatServerMessage:          "ChatServerMessage",
	ChatSmiley:                 "ChatSmiley",
	ChatSubscribeChannelAdd:    "ChatSubscribeChannelAdd",
	ChatSubscribeChannelRemove: "ChatSubscribeChannelRemove",

	DialogCustomAction:  "DialogCustomAction",
	DialogCreateError:   "DialogCreateError",
	DialogCreateSuccess: "DialogCreateSuccess",
	DialogQuestion:      "DialogQuestion",
	DialogLeave:         "DialogLeave",
	DialogPause:         "DialogPause",

	InfosInfoMaps:                  "InfosInfoMaps",
	InfosCompass:                   "InfosCompass",
	InfosInfoCoordinatesPHighlight: "InfosInfoCoordinatesPHighlight",
	InfosMessage:                   "InfosMessage",
	InfosQuantity:                  "InfosQuantity",
	InfosLifeRestoreTimerStart:     "InfosLifeRestoreTimerStart",
	InfosLifeRestoreTimerFinish:    "InfosLifeRestoreTimerFinish",

	SpellsList:                "SpellsList",
	SpellsChangeOption:        "SpellsChangeOption",
	SpellsUpgradeSpellError:   "SpellsUpgradeSpellError",
	SpellsUpgradeSpellSuccess: "SpellsUpgradeSpellSuccess",
	SpellsSpellBoost:          "SpellsSpellBoost",
	SpellsSpellForgetShow:     "SpellsSpellForgetShow",
	SpellsSpellForgetClose:    "SpellsSpellForgetClose",

	ItemsAccessories:      "ItemsAccessories",
	ItemsDropError:        "ItemsDropError",
	ItemsDropSuccess:      "ItemsDropSuccess",
	ItemsAddError:         "ItemsAddError",
	ItemsAddSuccess:       "ItemsAddSuccess",
	ItemsChange:           "ItemsChange",
	ItemsRemove:           "ItemsRemove",
	ItemsQuantity:         "ItemsQuantity",
	ItemsMovement:         "ItemsMovement",
	ItemsTool:             "ItemsTool",
	ItemsWeight:           "ItemsWeight",
	ItemsItemSetAdd:       "ItemsItemSetAdd",
	ItemsItemSetRemove:    "ItemsItemSetRemove",
	ItemsItemUseCondition: "ItemsItemUseCondition",
	ItemsItemFound:        "ItemsItemFound",

	FriendsAddFriendError:      "FriendsAddFriendError",
	FriendsAddFriendSuccess:    "FriendsAddFriendSuccess",
	FriendsRemoveFriendError:   "FriendsRemoveFriendError",
	FriendsRemoveFriendSuccess: "FriendsRemoveFriendSuccess",
	FriendsFriendsList:         "FriendsFriendsList",
	FriendsSpouse:              "FriendsSpouse",
	FriendsNotifyChange:        "FriendsNotifyChange",

	EnemiesAddEnemyError:      "EnemiesAddEnemyError",
	EnemiesAddEnemySuccess:    "EnemiesAddEnemySuccess",
	EnemiesRemoveEnemyError:   "EnemiesRemoveEnemyError",
	EnemiesRemoveEnemySuccess: "EnemiesRemoveEnemySuccess",
	EnemiesEnemiesList:        "EnemiesEnemiesList",

	KeyCreate:     "KeyCreate",
	KeyKeyError:   "KeyKeyError",
	KeyKeySuccess: "KeyKeySuccess",
	KeyLeave:      "KeyLeave",

	JobSkills:  "JobSkills",
	JobXP:      "JobXP",
	JobLevel:   "JobLevel",
	JobRemove:  "JobRemove",
	JobOptions: "JobOptions",

	ExchangeRequestError:                      "ExchangeRequestError",
	ExchangeRequestSuccess:                    "ExchangeRequestSuccess",
	ExchangeReady:                             "ExchangeReady",
	ExchangeLeaveError:                        "ExchangeLeaveError",
	ExchangeLeaveSuccess:                      "ExchangeLeaveSuccess",
	ExchangeCreateError:                       "ExchangeCreateError",
	ExchangeCreateSuccess:                     "ExchangeCreateSuccess",
	ExchangeCraftError:                        "ExchangeCraftError",
	ExchangeCraftSuccess:                      "ExchangeCraftSuccess",
	ExchangeLocalMovementError:                "ExchangeLocalMovementError",
	ExchangeLocalMovementSuccess:              "ExchangeLocalMovementSuccess",
	ExchangeLocalDistantError:                 "ExchangeLocalDistantError",
	ExchangeLocalDistantSuccess:               "ExchangeLocalDistantSuccess",
	ExchangeCoopMovementError:                 "ExchangeCoopMovementError",
	ExchangeCoopMovementSuccess:               "ExchangeCoopMovementSuccess",
	ExchangePayMovementError:                  "ExchangePayMovementError",
	ExchangePayMovementSuccess:                "ExchangePayMovementSuccess",
	ExchangeStorageMovementError:              "ExchangeStorageMovementError",
	ExchangeStorageMovementSuccess:            "ExchangeStorageMovementSuccess",
	ExchangePlayerShopMovementError:           "ExchangePlayerShopMovementError",
	ExchangePlayerShopMovementSuccess:         "ExchangePlayerShopMovementSuccess",
	ExchangeCraftPublicMode:                   "ExchangeCraftPublicMode",
	ExchangeMountStorage:                      "ExchangeMountStorage",
	ExchangeMountPark:                         "ExchangeMountPark",
	ExchangeMountPods:                         "ExchangeMountPods",
	ExchangeList:                              "ExchangeList",
	ExchangeSellError:                         "ExchangeSellError",
	ExchangeSellSuccess:                       "ExchangeSellSuccess",
	ExchangeBuyError:                          "ExchangeBuyError",
	ExchangeBuySuccess:                        "ExchangeBuySuccess",
	ExchangeAskOfflineExchange:                "ExchangeAskOfflineExchange",
	ExchangeSearchError:                       "ExchangeSearchError",
	ExchangeSearchSuccess:                     "ExchangeSearchSuccess",
	ExchangeBigStoreTypeItemsList:             "ExchangeBigStoreTypeItemsList",
	ExchangeBigStoreTypeItemsMovementAdd:      "ExchangeBigStoreTypeItemsMovementAdd",
	ExchangeBigStoreTypeItemsMovementRemove:   "ExchangeBigStoreTypeItemsMovementRemove",
	ExchangeBigStoreItemsList:                 "ExchangeBigStoreItemsList",
	ExchangeBigStoreItemsMovementAdd:          "ExchangeBigStoreItemsMovementAdd",
	ExchangeBigStoreItemsMovementRemove:       "ExchangeBigStoreItemsMovementRemove",
	ExchangeBigStoreItemMiddlePriceInBigStore: "ExchangeBigStoreItemMiddlePriceInBigStore",
	ExchangeCrafterReferenceAdd:               "ExchangeCrafterReferenceAdd",
	ExchangeCrafterReferenceRemove:            "ExchangeCrafterReferenceRemove",
	ExchangeCraftLoop:                         "ExchangeCraftLoop",
	ExchangeCraftLoopEnd:                      "ExchangeCraftLoopEnd",

	HousesListAdd:        "HousesListAdd",
	HousesListRemove:     "HousesListRemove",
	HousesProperties:     "HousesProperties",
	HousesLockedProperty: "HousesLockedProperty",
	HousesCreate:         "HousesCreate",
	HousesSellError:      "HousesSellError",
	HousesSellSuccess:    "HousesSellSuccess",
	HousesBuyError:       "HousesBuyError",
	HousesBuySuccess:     "HousesBuySuccess",
	HousesLeave:          "HousesLeave",
	HousesGuildInfos:     "HousesGuildInfos",

	StoragesListAdd:        "StoragesListAdd",
	StoragesListRemove:     "StoragesListRemove",
	StoragesLockedProperty: "StoragesLockedProperty",

	EmotesUseError:   "EmotesUseError",
	EmotesUseSuccess: "EmotesUseSuccess",
	EmotesList:       "EmotesList",
	EmotesAdd:        "EmotesAdd",
	EmotesRemove:     "EmotesRemove",
	EmotesDirection:  "EmotesDirection",

	DocumentsCreateError:   "DocumentsCreateError",
	DocumentsCreateSuccess: "DocumentsCreateSuccess",
	DocumentsLeave:         "DocumentsLeave",

	GuildNew:                         "GuildNew",
	GuildCreateError:                 "GuildCreateError",
	GuildCreateSuccess:               "GuildCreateSuccess",
	GuildStats:                       "GuildStats",
	GuildInfosGeneral:                "GuildInfosGeneral",
	GuildInfosMembers:                "GuildInfosMembers",
	GuildInfosBoosts:                 "GuildInfosBoosts",
	GuildInfosMountPark:              "GuildInfosMountPark",
	GuildInfosTaxCollectorsMovement:  "GuildInfosTaxCollectorsMovement",
	GuildInfosTaxCollectorsPlayers:   "GuildInfosTaxCollectorsPlayers",
	GuildInfosTaxCollectorsAttackers: "GuildInfosTaxCollectorsAttackers",
	GuildInfosHouses:                 "GuildInfosHouses",
	GuildJoinError:                   "GuildJoinError",
	GuildJoinSuccess:                 "GuildJoinSuccess",
	GuildJoinDistantSuccess:          "GuildJoinDistantSuccess",
	GuildRequestLocal:                "GuildRequestLocal",
	GuildRequestDistant:              "GuildRequestDistant",
	GuildLeave:                       "GuildLeave",
	GildBanError:                     "GildBanError",
	GildBanSuccess:                   "GildBanSuccess",
	GildHireTaxCollectorError:        "GildHireTaxCollectorError",
	GildHireTaxCollectorSuccess:      "GildHireTaxCollectorSuccess",
	GildTaxCollectorAttacked:         "GildTaxCollectorAttacked",
	GildTaxCollectorInfo:             "GildTaxCollectorInfo",
	GildUserInterfaceOpen:            "GildUserInterfaceOpen",

	WaypointsCreate:   "WaypointsCreate",
	WaypointsLeave:    "WaypointsLeave",
	WaypointsUseError: "WaypointsUseError",

	SubwayCreate:      "SubwayCreate",
	SubwayLeave:       "SubwayLeave",
	SubwayUseError:    "SubwayUseError",
	SubwayPrismCreate: "SubwayPrismCreate",
	SubwayPrismLeave:  "SubwayPrismLeave",

	SubareasList:                  "SubareasList",
	SubareasAlignmentModification: "SubareasAlignmentModification",

	ConquestAreaAlignmentChanged:      "ConquestAreaAlignmentChanged",
	ConquestPrismInfosJoined:          "ConquestPrismInfosJoined",
	ConquestPrismInfosClosing:         "ConquestPrismInfosClosing",
	ConquestConquestBonus:             "ConquestConquestBonus",
	ConquestPrismAttacked:             "ConquestPrismAttacked",
	ConquestPrismSurvived:             "ConquestPrismSurvived",
	ConquestPrismDead:                 "ConquestPrismDead",
	ConquestPrismFightAddPlayerAdd:    "ConquestPrismFightAddPlayerAdd",
	ConquestPrismFightAddPlayerRemove: "ConquestPrismFightAddPlayerRemove",
	ConquestPrismFightAddEnemyAdd:     "ConquestPrismFightAddEnemyAdd",
	ConquestPrismFightAddEnemyRemove:  "ConquestPrismFightAddEnemyRemove",
	ConquestWorldData:                 "ConquestWorldData",
	ConquestConquestBalance:           "ConquestConquestBalance",

	SpecializationSet:    "SpecializationSet",
	SpecializationChange: "SpecializationChange",

	FightsCount:   "FightsCount",
	FightsList:    "FightsList",
	FightsDetails: "FightsDetails",

	TutorialCreate:    "TutorialCreate",
	TutorialShowTip:   "TutorialShowTip",
	TutorialGameBegin: "TutorialGameBegin",

	QuestsList: "QuestsList",
	QuestsStep: "QuestsStep",

	PartyInviteError:   "PartyInviteError",
	PartyInviteSuccess: "PartyInviteSuccess",
	PartyLeader:        "PartyLeader",
	PartyRefuse:        "PartyRefuse",
	PartyAccept:        "PartyAccept",
	PartyCreateError:   "PartyCreateError",
	PartyCreateSuccess: "PartyCreateSuccess",
	PartyLeave:         "PartyLeave",
	PartyFollowError:   "PartyFollowError",
	PartyFollowSuccess: "PartyFollowSuccess",
	PartyMovement:      "PartyMovement",

	MountEquipError:   "MountEquipError",
	MountEquipAdd:     "MountEquipAdd",
	MountEquipRemove:  "MountEquipRemove",
	MountXP:           "MountXP",
	MountName:         "MountName",
	MountData:         "MountData",
	MountMountPark:    "MountMountPark",
	MountMountParkBuy: "MountMountParkBuy",
	MountLeave:        "MountLeave",
	MountRidingState:  "MountRidingState",
}

var MsgSvrIds = []MsgSvrId{
	AksHelloConnect,
	AksHelloGame,
	AksPong,
	AksQuickPong,
	AksRPing,
	AksServerMessage,
	AksServerWillDisconnect,

	BasicsNoticed,
	BasicsAuthorizedCommandError,
	BasicsAuthorizedCommandSuccess,
	BasicsAuthorizedLine,
	BasicsAuthorizedCommandPrompt,
	BasicsAuthorizedCommandClear,
	BasicsAuthorizedInterfaceOpen,
	BasicsAuthorizedInterfaceClose,
	BasicsTime,
	BasicsDate,
	BasicsWhoIsError,
	BasicsWhoIsSuccess,
	BasicsSubscriberRestrictionAdd,
	BasicsSubscriberRestrictionRemove,
	BasicsFileCheck,
	BasicsAveragePing,

	AccountCommunity,
	AccountPseudo,
	AccountLoginSuccess,
	AccountLoginError,
	AccountCharactersListError,
	AccountCharactersListSuccess,
	AccountServersListError,
	AccountServersListSuccess,
	AccountCharacterAddError,
	AccountCharacterAddSuccess,
	AccountTicketResponseError,
	AccountTicketResponseSuccess,
	AccountSelectServerError,
	AccountSelectServerSuccess,
	AccountSelectServerPlainSuccess,
	AccountCharacterSelectedError,
	AccountCharacterSelectedSuccess,
	AccountStats,
	AccountNewLevel,
	AccountRestrictions,
	AccountHosts,
	AccountRescue,
	AccountGiftsList,
	AccountGiftStoredError,
	AccountGiftStoredSuccess,
	AccountQueue,
	AccountNewQueue,
	AccountRegionalVersion,
	AccountCharacterNameGeneratedError,
	AccountCharacterNameGeneratedSuccess,
	AccountKey,
	AccountSecretQuestion,
	AccountCharacterDeleteError,
	AccountCharacterDeleteSuccess,
	AccountCharacterMigrationAskConfirm,
	AccountCharacterMigrationError,
	AccountCharacterMigrationSuccess,
	AccountFriendServerList,
	AccountMiniClipInfo,

	GameCreateError,
	GameCreateSuccess,
	GameJoin,
	GamePositionStart,
	GameReady,
	GameStartToPlay,
	GameEnd,
	GameMovementRemove,
	GameMovement,
	GameChallenge,
	GameTeam,
	GameLeave,
	GameFlag,
	GamePlayersCoordinates,
	GameEffect,
	GameClearAllEffect,
	GamePVP,
	GameMapData,
	GameMapLoaded,
	GameCellData,
	GameZoneData,
	GameCellObject,
	GameFrameObject2,
	GameFrameObjectExternal,
	GameFightChallenge,
	GameFightChallengeUpdateError,
	GameFightChallengeUpdateSuccess,
	GameTurnStart,
	GameTurnFinish,
	GameTurnList,
	GameTurnMiddle,
	GameTurnReady,
	GameExtraClip,
	GameFightOption,
	GameGameOver,

	GameActions,
	GameActionsStart,
	GameActionsFinish,

	ChatMessageError,
	ChatMessageSuccess,
	ChatServerMessage,
	ChatSmiley,
	ChatSubscribeChannelAdd,
	ChatSubscribeChannelRemove,

	DialogCustomAction,
	DialogCreateError,
	DialogCreateSuccess,
	DialogQuestion,
	DialogLeave,
	DialogPause,

	InfosInfoMaps,
	InfosCompass,
	InfosInfoCoordinatesPHighlight,
	InfosMessage,
	InfosQuantity,
	InfosLifeRestoreTimerStart,
	InfosLifeRestoreTimerFinish,

	SpellsList,
	SpellsChangeOption,
	SpellsUpgradeSpellError,
	SpellsUpgradeSpellSuccess,
	SpellsSpellBoost,
	SpellsSpellForgetShow,
	SpellsSpellForgetClose,

	ItemsAccessories,
	ItemsDropError,
	ItemsDropSuccess,
	ItemsAddError,
	ItemsAddSuccess,
	ItemsChange,
	ItemsRemove,
	ItemsQuantity,
	ItemsMovement,
	ItemsTool,
	ItemsWeight,
	ItemsItemSetAdd,
	ItemsItemSetRemove,
	ItemsItemUseCondition,
	ItemsItemFound,

	FriendsAddFriendError,
	FriendsAddFriendSuccess,
	FriendsRemoveFriendError,
	FriendsRemoveFriendSuccess,
	FriendsFriendsList,
	FriendsSpouse,
	FriendsNotifyChange,

	EnemiesAddEnemyError,
	EnemiesAddEnemySuccess,
	EnemiesRemoveEnemyError,
	EnemiesRemoveEnemySuccess,
	EnemiesEnemiesList,

	KeyCreate,
	KeyKeyError,
	KeyKeySuccess,
	KeyLeave,

	JobSkills,
	JobXP,
	JobLevel,
	JobRemove,
	JobOptions,

	ExchangeRequestError,
	ExchangeRequestSuccess,
	ExchangeReady,
	ExchangeLeaveError,
	ExchangeLeaveSuccess,
	ExchangeCreateError,
	ExchangeCreateSuccess,
	ExchangeCraftError,
	ExchangeCraftSuccess,
	ExchangeLocalMovementError,
	ExchangeLocalMovementSuccess,
	ExchangeLocalDistantError,
	ExchangeLocalDistantSuccess,
	ExchangeCoopMovementError,
	ExchangeCoopMovementSuccess,
	ExchangePayMovementError,
	ExchangePayMovementSuccess,
	ExchangeStorageMovementError,
	ExchangeStorageMovementSuccess,
	ExchangePlayerShopMovementError,
	ExchangePlayerShopMovementSuccess,
	ExchangeCraftPublicMode,
	ExchangeMountStorage,
	ExchangeMountPark,
	ExchangeMountPods,
	ExchangeList,
	ExchangeSellError,
	ExchangeSellSuccess,
	ExchangeBuyError,
	ExchangeBuySuccess,
	ExchangeAskOfflineExchange,
	ExchangeSearchError,
	ExchangeSearchSuccess,
	ExchangeBigStoreTypeItemsList,
	ExchangeBigStoreTypeItemsMovementAdd,
	ExchangeBigStoreTypeItemsMovementRemove,
	ExchangeBigStoreItemsList,
	ExchangeBigStoreItemsMovementAdd,
	ExchangeBigStoreItemsMovementRemove,
	ExchangeBigStoreItemMiddlePriceInBigStore,
	ExchangeCrafterReferenceAdd,
	ExchangeCrafterReferenceRemove,
	ExchangeCraftLoop,
	ExchangeCraftLoopEnd,

	HousesListAdd,
	HousesListRemove,
	HousesProperties,
	HousesLockedProperty,
	HousesCreate,
	HousesSellError,
	HousesSellSuccess,
	HousesBuyError,
	HousesBuySuccess,
	HousesLeave,
	HousesGuildInfos,

	StoragesListAdd,
	StoragesListRemove,
	StoragesLockedProperty,

	EmotesUseError,
	EmotesUseSuccess,
	EmotesList,
	EmotesAdd,
	EmotesRemove,
	EmotesDirection,

	DocumentsCreateError,
	DocumentsCreateSuccess,
	DocumentsLeave,

	GuildNew,
	GuildCreateError,
	GuildCreateSuccess,
	GuildStats,
	GuildInfosGeneral,
	GuildInfosMembers,
	GuildInfosBoosts,
	GuildInfosMountPark,
	GuildInfosTaxCollectorsMovement,
	GuildInfosTaxCollectorsPlayers,
	GuildInfosTaxCollectorsAttackers,
	GuildInfosHouses,
	GuildJoinError,
	GuildJoinSuccess,
	GuildJoinDistantSuccess,
	GuildRequestLocal,
	GuildRequestDistant,
	GuildLeave,
	GildBanError,
	GildBanSuccess,
	GildHireTaxCollectorError,
	GildHireTaxCollectorSuccess,
	GildTaxCollectorAttacked,
	GildTaxCollectorInfo,
	GildUserInterfaceOpen,

	WaypointsCreate,
	WaypointsLeave,
	WaypointsUseError,

	SubwayCreate,
	SubwayLeave,
	SubwayUseError,
	SubwayPrismCreate,
	SubwayPrismLeave,

	SubareasList,
	SubareasAlignmentModification,

	ConquestAreaAlignmentChanged,
	ConquestPrismInfosJoined,
	ConquestPrismInfosClosing,
	ConquestConquestBonus,
	ConquestPrismAttacked,
	ConquestPrismSurvived,
	ConquestPrismDead,
	ConquestPrismFightAddPlayerAdd,
	ConquestPrismFightAddPlayerRemove,
	ConquestPrismFightAddEnemyAdd,
	ConquestPrismFightAddEnemyRemove,
	ConquestWorldData,
	ConquestConquestBalance,

	SpecializationSet,
	SpecializationChange,

	FightsCount,
	FightsList,
	FightsDetails,

	TutorialCreate,
	TutorialShowTip,
	TutorialGameBegin,

	QuestsList,
	QuestsStep,

	PartyInviteError,
	PartyInviteSuccess,
	PartyLeader,
	PartyRefuse,
	PartyAccept,
	PartyCreateError,
	PartyCreateSuccess,
	PartyLeave,
	PartyFollowError,
	PartyFollowSuccess,
	PartyMovement,

	MountEquipError,
	MountEquipAdd,
	MountEquipRemove,
	MountXP,
	MountName,
	MountData,
	MountMountPark,
	MountMountParkBuy,
	MountLeave,
	MountRidingState,
}
