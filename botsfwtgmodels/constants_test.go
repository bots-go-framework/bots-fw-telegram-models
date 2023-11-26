package botsfwtgmodels

import (
	"strings"
	"testing"
)

func TestConstants(t *testing.T) {
	t.Run("BotUserCollection", func(t *testing.T) {
		if strings.TrimSpace(BotUserCollection) == "" {
			t.Error("BotUserCollection is empty")
		}
	})
	t.Run("TgChatCollection", func(t *testing.T) {
		if strings.TrimSpace(TgChatCollection) == "" {
			t.Error("TgChatCollection is empty")
		}
	})
}
