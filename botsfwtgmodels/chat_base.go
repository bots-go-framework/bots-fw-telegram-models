package botsfwtgmodels

import (
	"fmt"
	"github.com/bots-go-framework/bots-fw-store/botsfwmodels"
	"strconv"
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

func getChatID(tgBotID string, tgChatID int64) string {
	return fmt.Sprintf("%s:%d", tgBotID, tgChatID) // TODO: Should we migrate to format "id@bot"?
}

var _ TgChatData = (*TgChatBase)(nil)
var _ botsfwmodels.BotChat = (*TgChatBase)(nil)

// TgChatBase holds base properties of Telegram chat TgChatData
type TgChatBase struct {
	botsfwmodels.BotChatBaseData
	UserGroupID           string  `datastore:",index,omitempty" firestore:",omitempty" dalgo:",index,omitempty"` // Do index
	TelegramUserID        int64   `datastore:",noindex,omitempty" firestore:",noindex,omitempty"`
	TelegramUserIDs       []int64 `datastore:",noindex" firestore:",noindex"` // For groups
	LastProcessedUpdateID int     `datastore:",noindex,omitempty" firestore:",noindex,omitempty"`
	TgChatInstanceID      string  // !DO index! // TODO: document what is chat instance and why we need to keep id of it
}

func (data *TgChatBase) Base() *botsfwmodels.BotChatBaseData {
	return &data.BotChatBaseData
}

func (data *TgChatBase) BaseTgChatData() *TgChatBase {
	return data
}

// SetTgChatInstanceID is what it is
func (data *TgChatBase) SetTgChatInstanceID(v string) {
	data.TgChatInstanceID = v
}

// GetTgChatInstanceID is what it is
func (data *TgChatBase) GetTgChatInstanceID() string {
	return data.TgChatInstanceID
}

// GetPreferredLanguage returns preferred language for the chat
func (data *TgChatBase) GetPreferredLanguage() string {
	return data.PreferredLanguage
}

var _ botsfwmodels.BotChat = (*TgChatBase)(nil)

// NewTelegramChatBaseData create new telegram chat TgChatData
func NewTelegramChatBaseData() *TgChatBase {
	return &TgChatBase{
		BotChatBaseData: botsfwmodels.BotChatBaseData{
			BotBaseData: botsfwmodels.BotBaseData{
				//OwnedByUserWithID: user.NewOwnedByUserWithIntID(0, time.Now()),
			},
		},
	}
}

// SetAppUserID sets app user int ID
func (data *TgChatBase) SetAppUserID(appUserID string) {
	if data.IsGroup && appUserID != "" {
		panic("TgChatBase.IsGroup && id is not an empty string")
	}
	data.AppUserID = appUserID
}

// SetBotUserID sets bot user int ID
func (data *TgChatBase) SetBotUserID(id interface{}) {
	switch id := id.(type) {
	case string:
		var err error
		data.TelegramUserID, err = strconv.ParseInt(id, 10, 64)
		if err != nil {
			panic(err.Error())
		}
	case int:
		data.TelegramUserID = int64(id)
	case int64:
		data.TelegramUserID = id
	default:
		panic(fmt.Sprintf("Expected string, got: %T=%v", id, id))
	}
}

// Load loads Data from datastore
//func (data *TgChatBase) Load(ps []datastore.Property) error {
//	return datastore.LoadStruct(data, ps)
//}
//
//// Save saves Data to datastore
//func (data *TgChatBase) Save() (properties []datastore.Property, err error) {
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
//func (data *TgChatBase) CleanProperties(properties []datastore.Property) ([]datastore.Property, error) {
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
