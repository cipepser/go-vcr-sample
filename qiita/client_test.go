package qiita

import (
	"github.com/dnaeon/go-vcr/recorder"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_FetchUser(t *testing.T) {
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
			r, err := recorder.New("../fixtures/qiita")
			if err != nil {
				t.Errorf("got err when constructs go-vcr recorder. error:%v\n", err)
			}
			defer r.Stop()

			c := NewClient(&http.Client{Transport: r})
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
