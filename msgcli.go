package d1proto

import (
	"regexp"
	"strings"
)

type MsgCli interface {
	ProtocolId() MsgCliId
	Serialized() (string, error)
	Deserialize(string) error
}

type MsgCliId string

func MsgCliNameByID(id MsgCliId) (name string, ok bool) {
	name, ok = msgCliNameByID[id]
	if !ok {
		name = "Unknown"
	}
	return
}

func MsgCliIdByPkt(pkt string) (id MsgCliId, ok bool) {
	if pkt == "" {
		return
	}

	for _, v := range MsgCliIds {
		if strings.HasPrefix(pkt, string(v)) {
			ok = true
			if len(v) >= len(id) {
				id = v
			}
		}
	}

	if !ok {
		versionRx := regexp.MustCompile(`^\d+\.\d+\.\d+(\.\d+)?s?e?$`)
		if versionRx.MatchString(pkt) {
			ok = true
			id = AccountVersion
			return
		}

		credentialRx := regexp.MustCompile(`^[\w\d-@]{1,20}\n#\d[\w\d-_]+$`)
		if credentialRx.MatchString(pkt) {
			ok = true
			id = AccountCredential
			return
		}
	}

	return
}

const (
	AksPing      = MsgCliId("ping")
	AksQuickPing = MsgCliId("qping")
	AksRPong     = MsgCliId("rpong")

	BasicsAuthorizedCommand     = MsgCliId("BA")
	BasicsAuthorizedMoveCommand = MsgCliId("BaM")
	BasicsAuthorizedKickCommand = MsgCliId("BaK")
	BasicsWhoIs                 = MsgCliId("BW")
	BasicsKick                  = MsgCliId("BQ")
	BasicsAway                  = MsgCliId("BYA")
	BasicsInvisible             = MsgCliId("BYI")
	BasicsGetDate               = MsgCliId("BD")
	BasicsFileCheckAnswer       = MsgCliId("BC")
	BasicsSanctionMe            = MsgCliId("BK")
	BasicsRequestAveragePing    = MsgCliId("Bp")

	AccountVersion                  = MsgCliId("version")
	AccountCredential               = MsgCliId("credential")
	AccountSetNickname              = MsgCliId("nickname")
	AccountGetCharacters            = MsgCliId("AL")
	AccountGetCharactersForced      = MsgCliId("ALf")
	AccountGetServersList           = MsgCliId("Ax")
	AccountSetServer                = MsgCliId("AX")
	AccountSearchForFriend          = MsgCliId("AF")
	AccountSetCharacter             = MsgCliId("AS")
	AccountAddCharacter             = MsgCliId("AA")
	AccountDeleteCharacter          = MsgCliId("AD")
	AccountResetCharacter           = MsgCliId("AR")
	AccountBoost                    = MsgCliId("AB")
	AccountSendTicket               = MsgCliId("AT")
	AccountRequestRescue            = MsgCliId("Ar")
	AccountGetGifts                 = MsgCliId("Ag")
	AccountAttributeGiftToCharacter = MsgCliId("AG")
	AccountQueuePosition            = MsgCliId("Af")
	AccountGetRandomCharacterName   = MsgCliId("AP")
	AccountUseKey                   = MsgCliId("Ak")
	AccountRequestRegionalVersion   = MsgCliId("AV")
	AccountSendIdentity             = MsgCliId("Ai")
	AccountValidCharacterMigration  = MsgCliId("AM")
	AccountDeleteCharacterMigration = MsgCliId("AM-")
	AccountAskCharacterMigration    = MsgCliId("AM?")

	GameCreate                   = MsgCliId("GC")
	GameRequestLeave             = MsgCliId("GQ")
	GameSetPlayerPosition        = MsgCliId("Gp")
	GameRequestReady             = MsgCliId("GR")
	GameGetMapData               = MsgCliId("GD")
	GameGetExtraInformations     = MsgCliId("GI")
	GameTurnEnd                  = MsgCliId("Gt")
	GameTurnOk                   = MsgCliId("GT")
	GameAskDisablePVPMode        = MsgCliId("GP*")
	GameEnabledPVPMode           = MsgCliId("GP")
	GameFreeMySoul               = MsgCliId("GF")
	GameSetFlag                  = MsgCliId("Gf")
	GameShowFightChallengeTarget = MsgCliId("Gdi")

	GameActionsSendActions = MsgCliId("GA")
	GameActionAck          = MsgCliId("GKK")
	GameActionCancel       = MsgCliId("GKE")

	ChatSend                          = MsgCliId("BM")
	ChatReportMessage                 = MsgCliId("BR")
	ChatRequestSubscribeChannelAdd    = MsgCliId("cC+")
	ChatRequestSubscribeChannelRemove = MsgCliId("cC-")
	ChatUseSmiley                     = MsgCliId("BS")

	DialogBeginning    = MsgCliId("DB")
	DialogCreate       = MsgCliId("DC")
	DialogRequestLeave = MsgCliId("DV")
	DialogResponse     = MsgCliId("DR")

	InfosGetMaps        = MsgCliId("IM")
	InfosSendScreenInfo = MsgCliId("Ir")

	SpellsMoveToUsed = MsgCliId("SM")
	SpellsBoost      = MsgCliId("SB")
	SpellsForget     = MsgCliId("SF")

	ItemsRequestMovement = MsgCliId("OM")
	ItemsDrop            = MsgCliId("OD")
	ItemsDestroy         = MsgCliId("Od")
	ItemsUseConfirm      = MsgCliId("Ou")
	ItemsUseNoConfirm    = MsgCliId("OU")
	ItemsDissociate      = MsgCliId("Ox")
	ItemsSetSkin         = MsgCliId("Os")
	ItemsFeed            = MsgCliId("Of")

	FriendsGetFriendsList       = MsgCliId("FL")
	FriendsAddFriend            = MsgCliId("FA")
	FriendsRemoveFriend         = MsgCliId("FD")
	FriendsJoin                 = MsgCliId("FJ")
	FriendsJoinFriend           = MsgCliId("FJF")
	FriendsCompass              = MsgCliId("FJC")
	FriendsSetNotifyWhenConnect = MsgCliId("FO")

	EnemiesGetEnemiesList = MsgCliId("iL")
	EnemiesAddEnemy       = MsgCliId("iA")
	EnemiesRemoveEnemy    = MsgCliId("iD")

	KeyRequestLeave = MsgCliId("KV")
	KeySendKey      = MsgCliId("KK")

	JobChangeJobStats = MsgCliId("JO")

	ExchangeLeave                        = MsgCliId("EV")
	ExchangeRequest                      = MsgCliId("ER")
	ExchangeShop                         = MsgCliId("Es")
	ExchangeAccept                       = MsgCliId("EA")
	ExchangeRequestReady                 = MsgCliId("EK")
	ExchangeMovementItems                = MsgCliId("EMO")
	ExchangeMovementPay                  = MsgCliId("EP")
	ExchangeMovementKamas                = MsgCliId("EMG")
	ExchangeMovementSell                 = MsgCliId("ES")
	ExchangeMovementBuy                  = MsgCliId("EB")
	ExchangeOfflineExchange              = MsgCliId("EQ")
	ExchangeRequestAskOfflineExchange    = MsgCliId("Eq")
	ExchangeBigStoreType                 = MsgCliId("EHT")
	ExchangeBigStoreItemList             = MsgCliId("EHl")
	ExchangeBigStoreBuy                  = MsgCliId("EHB")
	ExchangeBigStoreSearch               = MsgCliId("EHS")
	ExchangeSetPublicMode                = MsgCliId("EW")
	ExchangeGetCrafterForJob             = MsgCliId("EJF")
	ExchangePutInShedFromInventory       = MsgCliId("Erp")
	ExchangePutInInventoryFromShed       = MsgCliId("Erg")
	ExchangePutInCertificateFromShed     = MsgCliId("Erc")
	ExchangePutInShedFromCertificate     = MsgCliId("ErC")
	ExchangePutInMountParkFromShed       = MsgCliId("Efp")
	ExchangePutInShedFromMountPark       = MsgCliId("Efg")
	ExchangeKillMountInPark              = MsgCliId("Eff")
	ExchangeKillMount                    = MsgCliId("Erf")
	ExchangeGetItemMiddlePriceInBigStore = MsgCliId("EHP")
	ExchangeReplayCraft                  = MsgCliId("EL")
	ExchangeRepeatCraft                  = MsgCliId("EMR")
	ExchangeStopRepeatCraft              = MsgCliId("EMr")

	HousesKick         = MsgCliId("hQ")
	HousesRequestLeave = MsgCliId("hV")
	HousesSell         = MsgCliId("hS")
	HousesBuy          = MsgCliId("hB")
	HousesState        = MsgCliId("hG") // get state or set rights
	HousesShare        = MsgCliId("hG+")
	HousesUnShare      = MsgCliId("hG-")

	EmotesUseEmote     = MsgCliId("eU")
	EmotesSetDirection = MsgCliId("eD")

	DocumentsRequestLeave = MsgCliId("dV")

	GuildCreate               = MsgCliId("gC")
	GuildRequestLeave         = MsgCliId("gV")
	GuildLeaveTaxInterface    = MsgCliId("gITV")
	GuildInvite               = MsgCliId("gJR")
	GuildAcceptInvitation     = MsgCliId("gJK")
	GuildRefuseInvitation     = MsgCliId("gJE")
	GuildGetInfosGeneral      = MsgCliId("gIG")
	GuildGetInfosMembers      = MsgCliId("gIM")
	GuildGetInfosBoosts       = MsgCliId("gIB")
	GuildGetInfosTaxCollector = MsgCliId("gIT")
	GuildGetInfosMountPark    = MsgCliId("gIF")
	GuildGetInfosGuildHouses  = MsgCliId("gIH")
	GuildBan                  = MsgCliId("gK")
	GuildChangeMemberProfile  = MsgCliId("gP")
	GuildBoostCharacteristic  = MsgCliId("gB")
	GuildBoostSpell           = MsgCliId("gb")
	GuildHireTaxCollector     = MsgCliId("gH")
	GuildJoinTaxCollector     = MsgCliId("gTJ")
	GuildLeaveTaxCollector    = MsgCliId("gTV")
	GuildRemoveTaxCollector   = MsgCliId("gF")
	GuildTeleportToGuildHouse = MsgCliId("gh")
	GuildTeleportToGuildFarm  = MsgCliId("gf")

	WaypointsRequestLeave = MsgCliId("WV")
	WaypointsUse          = MsgCliId("WU")

	SubwayRequestLeave      = MsgCliId("Wv")
	SubwayUse               = MsgCliId("Wu")
	SubwayRequestPrismLeave = MsgCliId("Ww")
	SubwayPrismUse          = MsgCliId("Wp")

	ConquestGetAlignedBonus = MsgCliId("CB")
	ConquestPrismInfosJoin  = MsgCliId("CIJ")
	ConquestPrismInfosLeave = MsgCliId("CIV")
	ConquestPrismFightJoin  = MsgCliId("CFJ")
	ConquestPrismFightLeave = MsgCliId("CFV")
	ConquestWorldInfosJoin  = MsgCliId("CWJ")
	ConquestWorldInfosLave  = MsgCliId("CWV")
	ConquestSwitchPlaces    = MsgCliId("CFS")
	ConquestRequestBalance  = MsgCliId("Cb")

	FightsGetList                = MsgCliId("fL")
	FightsGetDetails             = MsgCliId("fD")
	FightsBlockSpectators        = MsgCliId("fS")
	FightsBlockJoinerExceptParty = MsgCliId("fP")
	FightsBlockJoiner            = MsgCliId("fN")
	FightsNeedHelp               = MsgCliId("fH")

	TutorialEnd = MsgCliId("TV")

	QuestGetList = MsgCliId("QL")
	QuestGetStep = MsgCliId("QS")

	PartyInvite           = MsgCliId("PI")
	PartyRefuseInvitation = MsgCliId("PR")
	PartyAcceptInvitation = MsgCliId("PA")
	PartyRequestLeave     = MsgCliId("PV")
	PartyRequestFollow    = MsgCliId("PF")
	PartyWhere            = MsgCliId("PW")
	PartyFollowAll        = MsgCliId("PG")

	MountRename              = MsgCliId("Rn")
	MountKill                = MsgCliId("Rk")
	MountSetXP               = MsgCliId("Rx")
	MountRide                = MsgCliId("Rr")
	MountRequestData         = MsgCliId("Rd")
	MountParkMountData       = MsgCliId("Rp")
	MountRemoveObjectInPark  = MsgCliId("Ro")
	MountMountParkSell       = MsgCliId("Rs")
	MountRequestMountParkBuy = MsgCliId("Rb")
	MountRequestLeave        = MsgCliId("Rv")
	MountCastrate            = MsgCliId("Rc")
)

var msgCliNameByID = map[MsgCliId]string{
	AksPing:      "AksPing",
	AksQuickPing: "AksQuickPing",
	AksRPong:     "AksRPong",

	BasicsAuthorizedCommand:     "BasicsAuthorizedCommand",
	BasicsAuthorizedMoveCommand: "BasicsAuthorizedMoveCommand",
	BasicsAuthorizedKickCommand: "BasicsAuthorizedKickCommand",
	BasicsWhoIs:                 "BasicsWhoIs",
	BasicsKick:                  "BasicsKick",
	BasicsAway:                  "BasicsAway",
	BasicsInvisible:             "BasicsInvisible",
	BasicsGetDate:               "BasicsGetDate",
	BasicsFileCheckAnswer:       "BasicsFileCheckAnswer",
	BasicsSanctionMe:            "BasicsSanctionMe",
	BasicsRequestAveragePing:    "BasicsRequestAveragePing",

	AccountVersion:                  "AccountVersion",
	AccountCredential:               "AccountCredential",
	AccountSetNickname:              "AccountSetNickname",
	AccountGetCharacters:            "AccountGetCharacters",
	AccountGetCharactersForced:      "AccountGetCharactersForced",
	AccountGetServersList:           "AccountGetServersList",
	AccountSetServer:                "AccountSetServer",
	AccountSearchForFriend:          "AccountSearchForFriend",
	AccountSetCharacter:             "AccountSetCharacter",
	AccountAddCharacter:             "AccountAddCharacter",
	AccountDeleteCharacter:          "AccountDeleteCharacter",
	AccountResetCharacter:           "AccountResetCharacter",
	AccountBoost:                    "AccountBoost",
	AccountSendTicket:               "AccountSendTicket",
	AccountRequestRescue:            "AccountRequestRescue",
	AccountGetGifts:                 "AccountGetGifts",
	AccountAttributeGiftToCharacter: "AccountAttributeGiftToCharacter",
	AccountQueuePosition:            "AccountQueuePosition",
	AccountGetRandomCharacterName:   "AccountGetRandomCharacterName",
	AccountUseKey:                   "AccountUseKey",
	AccountRequestRegionalVersion:   "AccountRequestRegionalVersion",
	AccountSendIdentity:             "AccountSendIdentity",
	AccountValidCharacterMigration:  "AccountValidCharacterMigration",
	AccountDeleteCharacterMigration: "AccountDeleteCharacterMigration",
	AccountAskCharacterMigration:    "AccountAskCharacterMigration",

	GameCreate:                   "GameCreate",
	GameRequestLeave:             "GameRequestLeave",
	GameSetPlayerPosition:        "GameSetPlayerPosition",
	GameRequestReady:             "GameRequestReady",
	GameGetMapData:               "GameGetMapData",
	GameGetExtraInformations:     "GameGetExtraInformations",
	GameTurnEnd:                  "GameTurnEnd",
	GameTurnOk:                   "GameTurnOk",
	GameAskDisablePVPMode:        "GameAskDisablePVPMode",
	GameEnabledPVPMode:           "GameEnabledPVPMode",
	GameFreeMySoul:               "GameFreeMySoul",
	GameSetFlag:                  "GameSetFlag",
	GameShowFightChallengeTarget: "GameShowFightChallengeTarget",

	GameActionsSendActions: "GameActionsSendActions",
	GameActionAck:          "GameActionAck",
	GameActionCancel:       "GameActionCancel",

	ChatSend:                          "ChatSend",
	ChatReportMessage:                 "ChatReportMessage",
	ChatRequestSubscribeChannelAdd:    "ChatRequestSubscribeChannelAdd",
	ChatRequestSubscribeChannelRemove: "ChatRequestSubscribeChannelRemove",
	ChatUseSmiley:                     "ChatUseSmiley",

	DialogBeginning:    "DialogBeginning",
	DialogCreate:       "DialogCreate",
	DialogRequestLeave: "DialogRequestLeave",
	DialogResponse:     "DialogResponse",

	InfosGetMaps:        "InfosGetMaps",
	InfosSendScreenInfo: "InfosSendScreenInfo",

	SpellsMoveToUsed: "SpellsMoveToUsed",
	SpellsBoost:      "SpellsBoost",
	SpellsForget:     "SpellsForget",

	ItemsRequestMovement: "ItemsRequestMovement",
	ItemsDrop:            "ItemsDrop",
	ItemsDestroy:         "ItemsDestroy",
	ItemsUseConfirm:      "ItemsUseConfirm",
	ItemsUseNoConfirm:    "ItemsUseNoConfirm",
	ItemsDissociate:      "ItemsDissociate",
	ItemsSetSkin:         "ItemsSetSkin",
	ItemsFeed:            "ItemsFeed",

	FriendsGetFriendsList:       "FriendsGetFriendsList",
	FriendsAddFriend:            "FriendsAddFriend",
	FriendsRemoveFriend:         "FriendsRemoveFriend",
	FriendsJoin:                 "FriendsJoin",
	FriendsJoinFriend:           "FriendsJoinFriend",
	FriendsCompass:              "FriendsCompass",
	FriendsSetNotifyWhenConnect: "FriendsSetNotifyWhenConnect",

	EnemiesGetEnemiesList: "EnemiesGetEnemiesList",
	EnemiesAddEnemy:       "EnemiesAddEnemy",
	EnemiesRemoveEnemy:    "EnemiesRemoveEnemy",

	KeyRequestLeave: "KeyRequestLeave",
	KeySendKey:      "KeySendKey",

	JobChangeJobStats: "JobChangeJobStats",

	ExchangeLeave:                        "ExchangeLeave",
	ExchangeRequest:                      "ExchangeRequest",
	ExchangeShop:                         "ExchangeShop",
	ExchangeAccept:                       "ExchangeAccept",
	ExchangeRequestReady:                 "ExchangeRequestReady",
	ExchangeMovementItems:                "ExchangeMovementItems",
	ExchangeMovementPay:                  "ExchangeMovementPay",
	ExchangeMovementKamas:                "ExchangeMovementKamas",
	ExchangeMovementSell:                 "ExchangeMovementSell",
	ExchangeMovementBuy:                  "ExchangeMovementBuy",
	ExchangeOfflineExchange:              "ExchangeOfflineExchange",
	ExchangeRequestAskOfflineExchange:    "ExchangeRequestAskOfflineExchange",
	ExchangeBigStoreType:                 "ExchangeBigStoreType",
	ExchangeBigStoreItemList:             "ExchangeBigStoreItemList",
	ExchangeBigStoreBuy:                  "ExchangeBigStoreBuy",
	ExchangeBigStoreSearch:               "ExchangeBigStoreSearch",
	ExchangeSetPublicMode:                "ExchangeSetPublicMode",
	ExchangeGetCrafterForJob:             "ExchangeGetCrafterForJob",
	ExchangePutInShedFromInventory:       "ExchangePutInShedFromInventory",
	ExchangePutInInventoryFromShed:       "ExchangePutInInventoryFromShed",
	ExchangePutInCertificateFromShed:     "ExchangePutInCertificateFromShed",
	ExchangePutInShedFromCertificate:     "ExchangePutInShedFromCertificate",
	ExchangePutInMountParkFromShed:       "ExchangePutInMountParkFromShed",
	ExchangePutInShedFromMountPark:       "ExchangePutInShedFromMountPark",
	ExchangeKillMountInPark:              "ExchangeKillMountInPark",
	ExchangeKillMount:                    "ExchangeKillMount",
	ExchangeGetItemMiddlePriceInBigStore: "ExchangeGetItemMiddlePriceInBigStore",
	ExchangeReplayCraft:                  "ExchangeReplayCraft",
	ExchangeRepeatCraft:                  "ExchangeRepeatCraft",
	ExchangeStopRepeatCraft:              "ExchangeStopRepeatCraft",

	HousesKick:         "HousesKick",
	HousesRequestLeave: "HousesRequestLeave",
	HousesSell:         "HousesSell",
	HousesBuy:          "HousesBuy",
	HousesState:        "HousesState",
	HousesShare:        "HousesShare",
	HousesUnShare:      "HousesUnShare",

	EmotesUseEmote:     "EmotesUseEmote",
	EmotesSetDirection: "EmotesSetDirection",

	DocumentsRequestLeave: "DocumentsRequestLeave",

	GuildCreate:               "GuildCreate",
	GuildRequestLeave:         "GuildRequestLeave",
	GuildLeaveTaxInterface:    "GuildLeaveTaxInterface",
	GuildInvite:               "GuildInvite",
	GuildAcceptInvitation:     "GuildAcceptInvitation",
	GuildRefuseInvitation:     "GuildRefuseInvitation",
	GuildGetInfosGeneral:      "GuildGetInfosGeneral",
	GuildGetInfosMembers:      "GuildGetInfosMembers",
	GuildGetInfosBoosts:       "GuildGetInfosBoosts",
	GuildGetInfosTaxCollector: "GuildGetInfosTaxCollector",
	GuildGetInfosMountPark:    "GuildGetInfosMountPark",
	GuildGetInfosGuildHouses:  "GuildGetInfosGuildHouses",
	GuildBan:                  "GuildBan",
	GuildChangeMemberProfile:  "GuildChangeMemberProfile",
	GuildBoostCharacteristic:  "GuildBoostCharacteristic",
	GuildBoostSpell:           "GuildBoostSpell",
	GuildHireTaxCollector:     "GuildHireTaxCollector",
	GuildJoinTaxCollector:     "GuildJoinTaxCollector",
	GuildLeaveTaxCollector:    "GuildLeaveTaxCollector",
	GuildRemoveTaxCollector:   "GuildRemoveTaxCollector",
	GuildTeleportToGuildHouse: "GuildTeleportToGuildHouse",
	GuildTeleportToGuildFarm:  "GuildTeleportToGuildFarm",

	WaypointsRequestLeave: "WaypointsRequestLeave",
	WaypointsUse:          "WaypointsUse",

	SubwayRequestLeave:      "SubwayRequestLeave",
	SubwayUse:               "SubwayUse",
	SubwayRequestPrismLeave: "SubwayRequestPrismLeave",
	SubwayPrismUse:          "SubwayPrismUse",

	ConquestGetAlignedBonus: "ConquestGetAlignedBonus",
	ConquestPrismInfosJoin:  "ConquestPrismInfosJoin",
	ConquestPrismInfosLeave: "ConquestPrismInfosLeave",
	ConquestPrismFightJoin:  "ConquestPrismFightJoin",
	ConquestPrismFightLeave: "ConquestPrismFightLeave",
	ConquestWorldInfosJoin:  "ConquestWorldInfosJoin",
	ConquestWorldInfosLave:  "ConquestWorldInfosLave",
	ConquestSwitchPlaces:    "ConquestSwitchPlaces",
	ConquestRequestBalance:  "ConquestRequestBalance",

	FightsGetList:                "FightsGetList",
	FightsGetDetails:             "FightsGetDetails",
	FightsBlockSpectators:        "FightsBlockSpectators",
	FightsBlockJoinerExceptParty: "FightsBlockJoinerExceptParty",
	FightsBlockJoiner:            "FightsBlockJoiner",
	FightsNeedHelp:               "FightsNeedHelp",

	TutorialEnd: "TutorialEnd",

	QuestGetList: "QuestGetList",
	QuestGetStep: "QuestGetStep",

	PartyInvite:           "PartyInvite",
	PartyRefuseInvitation: "PartyRefuseInvitation",
	PartyAcceptInvitation: "PartyAcceptInvitation",
	PartyRequestLeave:     "PartyRequestLeave",
	PartyRequestFollow:    "PartyRequestFollow",
	PartyWhere:            "PartyWhere",
	PartyFollowAll:        "PartyFollowAll",

	MountRename:              "MountRename",
	MountKill:                "MountKill",
	MountSetXP:               "MountSetXP",
	MountRide:                "MountRide",
	MountRequestData:         "MountRequestData",
	MountParkMountData:       "MountParkMountData",
	MountRemoveObjectInPark:  "MountRemoveObjectInPark",
	MountMountParkSell:       "MountMountParkSell",
	MountRequestMountParkBuy: "MountRequestMountParkBuy",
	MountRequestLeave:        "MountRequestLeave",
	MountCastrate:            "MountCastrate",
}

var MsgCliIds = []MsgCliId{
	AksPing,
	AksQuickPing,
	AksRPong,

	BasicsAuthorizedCommand,
	BasicsAuthorizedMoveCommand,
	BasicsAuthorizedKickCommand,
	BasicsWhoIs,
	BasicsKick,
	BasicsAway,
	BasicsInvisible,
	BasicsGetDate,
	BasicsFileCheckAnswer,
	BasicsSanctionMe,
	BasicsRequestAveragePing,

	AccountVersion,
	AccountCredential,
	AccountSetNickname,
	AccountGetCharacters,
	AccountGetCharactersForced,
	AccountGetServersList,
	AccountSetServer,
	AccountSearchForFriend,
	AccountSetCharacter,
	AccountAddCharacter,
	AccountDeleteCharacter,
	AccountResetCharacter,
	AccountBoost,
	AccountSendTicket,
	AccountRequestRescue,
	AccountGetGifts,
	AccountAttributeGiftToCharacter,
	AccountQueuePosition,
	AccountGetRandomCharacterName,
	AccountUseKey,
	AccountRequestRegionalVersion,
	AccountSendIdentity,
	AccountValidCharacterMigration,
	AccountDeleteCharacterMigration,
	AccountAskCharacterMigration,

	GameCreate,
	GameRequestLeave,
	GameSetPlayerPosition,
	GameRequestReady,
	GameGetMapData,
	GameGetExtraInformations,
	GameTurnEnd,
	GameTurnOk,
	GameAskDisablePVPMode,
	GameEnabledPVPMode,
	GameFreeMySoul,
	GameSetFlag,
	GameShowFightChallengeTarget,

	GameActionsSendActions,
	GameActionAck,
	GameActionCancel,

	ChatSend,
	ChatReportMessage,
	ChatRequestSubscribeChannelAdd,
	ChatRequestSubscribeChannelRemove,
	ChatUseSmiley,

	DialogBeginning,
	DialogCreate,
	DialogRequestLeave,
	DialogResponse,

	InfosGetMaps,
	InfosSendScreenInfo,

	SpellsMoveToUsed,
	SpellsBoost,
	SpellsForget,

	ItemsRequestMovement,
	ItemsDrop,
	ItemsDestroy,
	ItemsUseConfirm,
	ItemsUseNoConfirm,
	ItemsDissociate,
	ItemsSetSkin,
	ItemsFeed,

	FriendsGetFriendsList,
	FriendsAddFriend,
	FriendsRemoveFriend,
	FriendsJoin,
	FriendsJoinFriend,
	FriendsCompass,
	FriendsSetNotifyWhenConnect,

	EnemiesGetEnemiesList,
	EnemiesAddEnemy,
	EnemiesRemoveEnemy,

	KeyRequestLeave,
	KeySendKey,

	JobChangeJobStats,

	ExchangeLeave,
	ExchangeRequest,
	ExchangeShop,
	ExchangeAccept,
	ExchangeRequestReady,
	ExchangeMovementItems,
	ExchangeMovementPay,
	ExchangeMovementKamas,
	ExchangeMovementSell,
	ExchangeMovementBuy,
	ExchangeOfflineExchange,
	ExchangeRequestAskOfflineExchange,
	ExchangeBigStoreType,
	ExchangeBigStoreItemList,
	ExchangeBigStoreBuy,
	ExchangeBigStoreSearch,
	ExchangeSetPublicMode,
	ExchangeGetCrafterForJob,
	ExchangePutInShedFromInventory,
	ExchangePutInInventoryFromShed,
	ExchangePutInCertificateFromShed,
	ExchangePutInShedFromCertificate,
	ExchangePutInMountParkFromShed,
	ExchangePutInShedFromMountPark,
	ExchangeKillMountInPark,
	ExchangeKillMount,
	ExchangeGetItemMiddlePriceInBigStore,
	ExchangeReplayCraft,
	ExchangeRepeatCraft,
	ExchangeStopRepeatCraft,

	HousesKick,
	HousesRequestLeave,
	HousesSell,
	HousesBuy,
	HousesState,
	HousesShare,
	HousesUnShare,

	EmotesUseEmote,
	EmotesSetDirection,

	DocumentsRequestLeave,

	GuildCreate,
	GuildRequestLeave,
	GuildLeaveTaxInterface,
	GuildInvite,
	GuildAcceptInvitation,
	GuildRefuseInvitation,
	GuildGetInfosGeneral,
	GuildGetInfosMembers,
	GuildGetInfosBoosts,
	GuildGetInfosTaxCollector,
	GuildGetInfosMountPark,
	GuildGetInfosGuildHouses,
	GuildBan,
	GuildChangeMemberProfile,
	GuildBoostCharacteristic,
	GuildBoostSpell,
	GuildHireTaxCollector,
	GuildJoinTaxCollector,
	GuildLeaveTaxCollector,
	GuildRemoveTaxCollector,
	GuildTeleportToGuildHouse,
	GuildTeleportToGuildFarm,

	WaypointsRequestLeave,
	WaypointsUse,

	SubwayRequestLeave,
	SubwayUse,
	SubwayRequestPrismLeave,
	SubwayPrismUse,

	ConquestGetAlignedBonus,
	ConquestPrismInfosJoin,
	ConquestPrismInfosLeave,
	ConquestPrismFightJoin,
	ConquestPrismFightLeave,
	ConquestWorldInfosJoin,
	ConquestWorldInfosLave,
	ConquestSwitchPlaces,
	ConquestRequestBalance,

	FightsGetList,
	FightsGetDetails,
	FightsBlockSpectators,
	FightsBlockJoinerExceptParty,
	FightsBlockJoiner,
	FightsNeedHelp,

	TutorialEnd,

	QuestGetList,
	QuestGetStep,

	PartyInvite,
	PartyRefuseInvitation,
	PartyAcceptInvitation,
	PartyRequestLeave,
	PartyRequestFollow,
	PartyWhere,
	PartyFollowAll,

	MountRename,
	MountKill,
	MountSetXP,
	MountRide,
	MountRequestData,
	MountParkMountData,
	MountRemoveObjectInPark,
	MountMountParkSell,
	MountRequestMountParkBuy,
	MountRequestLeave,
	MountCastrate,
}
