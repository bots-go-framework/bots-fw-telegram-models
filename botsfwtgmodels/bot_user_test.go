package botsfwtgmodels

import (
	"testing"

	"github.com/bots-go-framework/bots-fw-store/botsfwmodels"
)

func TestTelegramUser(t *testing.T) {
	u := &TgPlatformUserBaseDbo{
		PlatformUserBaseDbo: botsfwmodels.PlatformUserBaseDbo{
			FirstName: "John",
			LastName:  "Doe",
			UserName:  "johndoe",
		},
	}

	if u.TgPlatformUserBaseDbo() != u {
		t.Error("expected TgPlatformUserBaseDbo to return receiver")
	}

	if email := u.GetEmail(); email != "" {
		t.Errorf("expected empty email, got %q", email)
	}

	if u.IsEmailConfirmed() {
		t.Error("expected IsEmailConfirmed to be false")
	}

	names := u.GetNames()
	if names.FirstName != "John" || names.LastName != "Doe" || names.UserName != "johndoe" {
		t.Errorf("unexpected names: %+v", names)
	}

	// Name variations
	if got := u.Name(); got != "@johndoe - John Doe" {
		t.Errorf("got %q, want '@johndoe - John Doe'", got)
	}

	uNoUser := &TgPlatformUserBaseDbo{
		PlatformUserBaseDbo: botsfwmodels.PlatformUserBaseDbo{
			FirstName: "John",
			LastName:  "Doe",
		},
	}
	if got := uNoUser.Name(); got != "John Doe" {
		t.Errorf("got %q, want 'John Doe'", got)
	}

	uOnlyLast := &TgPlatformUserBaseDbo{
		PlatformUserBaseDbo: botsfwmodels.PlatformUserBaseDbo{
			LastName: "Doe",
		},
	}
	if got := uOnlyLast.Name(); got != "Doe" {
		t.Errorf("got %q, want 'Doe'", got)
	}

	uOnlyUser := &TgPlatformUserBaseDbo{
		PlatformUserBaseDbo: botsfwmodels.PlatformUserBaseDbo{
			UserName: "johndoe",
		},
	}
	if got := uOnlyUser.Name(); got != "@johndoe" {
		t.Errorf("got %q, want '@johndoe'", got)
	}
}
