package botsfwtgmodels

import "testing"

func TestChatInstanceBaseData_GetPreferredLanguage(t *testing.T) {
	type fields struct {
		TgChatID          int64
		PreferredLanguage string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "empty", fields: fields{}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &TgChatInstanceBaseData{
				TgChatID:          tt.fields.TgChatID,
				PreferredLanguage: tt.fields.PreferredLanguage,
			}
			if got := v.GetPreferredLanguage(); got != tt.want {
				t.Errorf("GetPreferredLanguage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatInstanceBaseData_GetTgChatID(t *testing.T) {
	type fields struct {
		TgChatID          int64
		PreferredLanguage string
	}
	tests := []struct {
		name   string
		fields fields
		want   int64
	}{
		{"empty", fields{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &TgChatInstanceBaseData{
				TgChatID:          tt.fields.TgChatID,
				PreferredLanguage: tt.fields.PreferredLanguage,
			}
			if got := v.GetTgChatID(); got != tt.want {
				t.Errorf("GetTgChatID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatInstanceBaseData_SetPreferredLanguage(t *testing.T) {
	type fields struct {
		TgChatID          int64
		PreferredLanguage string
	}
	type args struct {
		code string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{name: "empty", fields: fields{}, args: args{""}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &TgChatInstanceBaseData{
				TgChatID:          tt.fields.TgChatID,
				PreferredLanguage: tt.fields.PreferredLanguage,
			}
			v.SetPreferredLanguage(tt.args.code)
		})
	}
}

func TestChatInstanceBaseData_Validate(t *testing.T) {
	type fields struct {
		TgChatID          int64
		PreferredLanguage string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "empty", fields: fields{}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &TgChatInstanceBaseData{
				TgChatID:          tt.fields.TgChatID,
				PreferredLanguage: tt.fields.PreferredLanguage,
			}
			if err := v.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
