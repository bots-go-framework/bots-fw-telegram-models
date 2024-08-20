package botsfwtgmodels

import (
	"fmt"
	"github.com/bots-go-framework/bots-fw-store/botsfwmodels"
	"slices"
)

// TgChatRecord holds base properties of Telegram chat TgChatData
//type TgChatRecord struct { // TODO: Do we need this struct at all?
//	record.WithID[string]
//	//Data *TgChatBase
//}

// SetID sets ID
//func (v *TgChatRecord) SetID(tgBotID string, tgChatID int64) {
//	v.ID = getChatID(tgBotID, tgChatID) // TODO: Should we migrate to format "id@bot"?
//	v.Key = dal.NewKeyWithID(TgChatCollection, v.ID)
//}

//func getChatID(tgBotID string, tgChatID int64) string {
//	return fmt.Sprintf("%s:%d", tgBotID, tgChatID) // TODO: Should we migrate to format "id@bot"?
//}

var _ TgChatData = (*TgChatBaseData)(nil)
var _ botsfwmodels.BotChatData = (*TgChatBaseData)(nil)

// TgChatBaseData holds base properties of Telegram chat TgChatData
type TgChatBaseData struct {
	botsfwmodels.ChatBaseData

	// UserGroupID TODO: needs documentation what is it and intended usage
	UserGroupID string `firestore:"userGroupID,omitempty"` // Do index

	LastProcessedUpdateID int `firestore:"lastProcessedUpdateID,omitempty"`

	TgChatInstanceID string `firestore:"tgChatInstanceID,omitempty"` // !DO index! TODO: document why needed
}

func (data *TgChatBaseData) Base() *botsfwmodels.ChatBaseData {
	return &data.ChatBaseData
}

//func (data *TgChatBaseData) ChatKey() botsfwmodels.ChatKey {
//	return data.ChatBaseData.ChatKey
//}

func (data *TgChatBaseData) BaseTgChatData() *TgChatBaseData {
	return data
}

// SetTgChatInstanceID is what it is
func (data *TgChatBaseData) SetTgChatInstanceID(v string) {
	data.TgChatInstanceID = v
}

// GetTgChatInstanceID is what it is
func (data *TgChatBaseData) GetTgChatInstanceID() string {
	return data.TgChatInstanceID
}

// GetPreferredLanguage returns preferred language for the chat
func (data *TgChatBaseData) GetPreferredLanguage() string {
	return data.PreferredLanguage
}

var _ botsfwmodels.BotChatData = (*TgChatBaseData)(nil)

// NewTelegramChatBaseData create new telegram chat TgChatData
func NewTelegramChatBaseData() *TgChatBaseData {
	return &TgChatBaseData{
		ChatBaseData: botsfwmodels.ChatBaseData{
			BotBaseData: botsfwmodels.BotBaseData{
				//OwnedByUserWithID: user.NewOwnedByUserWithIntID(0, time.Now()),
			},
		},
	}
}

// SetAppUserID sets app user int ID
func (data *TgChatBaseData) SetAppUserID(appUserID string) {
	if data.IsGroup && appUserID != "" {
		panic("TgChatBase.IsGroup && id is not an empty string")
	}
	data.AppUserID = appUserID
}

// SetBotUserID sets bot user int ID
func (data *TgChatBaseData) SetBotUserID(id interface{}) {
	switch id := id.(type) {
	case string:
		data.BotUserID = id
	case int, int64:
		data.BotUserID = fmt.Sprintf("%d", id)
	default:
		panic(fmt.Sprintf("Expected string or int, got: %T=%v", id, id))
	}
	if data.IsGroup && !slices.Contains(data.BotUserIDs, data.BotUserID) {
		data.BotUserIDs = append(data.BotUserIDs, data.BotUserID)
	}
}

// Load loads Data from datastore
//func (data *TgChatBaseData) Load(ps []datastore.Property) error {
//	return datastore.LoadStruct(data, ps)
//}
//
//// Save saves Data to datastore
//func (data *TgChatBaseData) Save() (properties []datastore.Property, err error) {
//	if properties, err = datastore.SaveStruct(data); err != nil {
//		return
//	}
//	if properties, err = data.CleanProperties(properties); err != nil {
//		return
//	}
//	return
//}
//
//// CleanProperties cleans properties
//func (data *TgChatBaseData) CleanProperties(properties []datastore.Property) ([]datastore.Property, error) {
//	if data.IsGroup && data.AppUserIntID != 0 {
//		for _, userID := range data.AppUserIntIDs {
//			if userID == data.AppUserIntID {
//				goto found
//			}
//		}
//		data.AppUserIntIDs = append(data.AppUserIntIDs, data.AppUserIntID)
//		data.AppUserIntID = 0
//	found:
//	}
//
//	for i, userID := range data.AppUserIntIDs {
//		if userID == 0 {
//			panic(fmt.Sprintf("*TgChatBase.AppUserIntIDs[%d] == 0", i))
//		}
//	}
//
//	var err error
//	//if properties, err = gaedb.CleanProperties(properties, map[string]gaedb.IsOkToRemove{
//	//	"AppUserIntID":          gaedb.IsZeroInt,
//	//	"AccessGranted":         gaedb.IsFalse,
//	//	"AwaitingReplyTo":       gaedb.IsEmptyString,
//	//	"DtForbidden":           gaedb.IsZeroTime,
//	//	"DtForbiddenLast":       gaedb.IsZeroTime,
//	//	"GaClientID":            gaedb.IsEmptyByteArray,
//	//	"TelegramUserID":        gaedb.IsZeroInt,
//	//	"LastProcessedUpdateID": gaedb.IsZeroInt,
//	//	"PreferredLanguage":     gaedb.IsEmptyString,
//	//	"Title":                 gaedb.IsEmptyString, // TODO: Is it obsolete?
//	//	"Type":                  gaedb.IsEmptyString, // TODO: Is it obsolete?
//	//}); err != nil {
//	//	return properties, err
//	//}
//	return properties, err
//}
