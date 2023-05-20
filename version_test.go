package bots_fw_telegram_models

import (
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	if strings.TrimSpace(Version) == "" {
		t.Error("Version is empty")
	}
}
