package user

import (
	"reflect"
	"testing"
)

func TestMetricUnitClone(t *testing.T) {
	tests := []struct {
		name string
		pass CachedTelegramUser
		want CachedTelegramUser
	}{
		{
			name: "clone",
			pass: CachedTelegramUser{
				ID: 123, Nickname: "@Nickname1", FirstName: "FirstName1",
				LastName: "LastName1", IsBot: true, IsBanned: true,
			},
			want: CachedTelegramUser{
				ID: 123, Nickname: "@Nickname1", FirstName: "FirstName1",
				LastName: "LastName1", IsBot: true, IsBanned: true,
			},
		},
	}
	for _, testItem := range tests {
		test := testItem
		t.Run(test.name, func(t *testing.T) {
			if got := test.pass.Clone(); !reflect.DeepEqual(got, test.want) {
				t.Errorf("Clone() = %v, want %v", got, test.want)
			}
		})
	}
}
