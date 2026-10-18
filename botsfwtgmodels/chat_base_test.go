package botsfwtgmodels

import (
	"testing"
)

func TestTelegramChat(t *testing.T) {
	c := NewTelegramChatBaseData()
	if c == nil {
		t.Fatal("expected non-nil chat base data")
	}

	if c.Base() != &c.ChatBaseData {
		t.Error("expected Base() to return &c.ChatBaseData")
	}

	if c.BaseTgChatData() != c {
		t.Error("expected BaseTgChatData() to return c")
	}

	c.SetTgChatInstanceID("inst-123")
	if c.GetTgChatInstanceID() != "inst-123" {
		t.Errorf("got %q, want 'inst-123'", c.GetTgChatInstanceID())
	}

	c.PreferredLanguage = "en"
	if c.GetPreferredLanguage() != "en" {
		t.Errorf("got %q, want 'en'", c.GetPreferredLanguage())
	}

	// SetAppUserID
	c.IsGroup = false
	c.SetAppUserID("app-user-1")
	if c.AppUserID != "app-user-1" {
		t.Errorf("got %q, want 'app-user-1'", c.AppUserID)
	}

	// SetAppUserID panic on group
	c.IsGroup = true
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic setting app user id on group")
		}
	}()
	c.SetAppUserID("app-user-2")
}

func TestTgChatBaseData_SetBotUserID(t *testing.T) {
	t.Run("non-group", func(t *testing.T) {
		c := NewTelegramChatBaseData()
		c.IsGroup = false

		c.SetBotUserID("bot-str")
		if len(c.BotUserIDs) != 1 || c.BotUserIDs[0] != "bot-str" {
			t.Errorf("unexpected BotUserIDs: %v", c.BotUserIDs)
		}

		c.SetBotUserID(42)
		if len(c.BotUserIDs) != 1 || c.BotUserIDs[0] != "42" {
			t.Errorf("unexpected BotUserIDs: %v", c.BotUserIDs)
		}

		c.SetBotUserID(int64(100))
		if len(c.BotUserIDs) != 1 || c.BotUserIDs[0] != "100" {
			t.Errorf("unexpected BotUserIDs: %v", c.BotUserIDs)
		}
	})

	t.Run("group", func(t *testing.T) {
		c := NewTelegramChatBaseData()
		c.IsGroup = true

		c.SetBotUserID("bot1")
		c.SetBotUserID("bot1") // Duplicate, should not be re-added
		c.SetBotUserID(123)

		if len(c.BotUserIDs) != 2 || c.BotUserIDs[0] != "bot1" || c.BotUserIDs[1] != "123" {
			t.Errorf("unexpected BotUserIDs: %v", c.BotUserIDs)
		}
	})

	t.Run("invalid_type_panics", func(t *testing.T) {
		c := NewTelegramChatBaseData()
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic for invalid type")
			}
		}()
		c.SetBotUserID(3.14)
	})
}
