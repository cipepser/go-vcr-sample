package qiita

import (
	"net/http"
	"reflect"
	"testing"
)

func TestClient_FetchUser(t *testing.T) {
	type fields struct {
		Client *http.Client
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name: "Location",
			args: args{
				id: "cipepser",
			},
			want: &User{
				ID:       "cipepser",
				Location: "Tokyo, Japan",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient(http.DefaultClient)
			got, err := c.FetchUser(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}
