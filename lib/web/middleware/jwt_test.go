package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"reflect"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "generateToken",
			args:    args{username: "admin"},
			want:    "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiaXNzIjoic21hcnQuY29tLmNuIiwiZXhwIjoxNjc4OTU1NzA1fQ._wogxHd5dlbtJ-ZG3417eXOC78yLWjtUdIzK33LxZIg",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateToken(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    *Claims
		wantErr bool
	}{
		{
			name: "parseToken",
			args: args{tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwiaXNzIjoic21hcnQuY29tLmNuIiwiZXhwIjoxNjc4OTU1NzA1fQ._wogxHd5dlbtJ-ZG3417eXOC78yLWjtUdIzK33LxZIg"},
			want: &Claims{
				Username:         "admin",
				RegisteredClaims: jwt.RegisteredClaims{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Username, tt.want.Username) {
				t.Errorf("ParseToken() got = %v, want %v", got, tt.want)
			}
			fmt.Println(got)
		})
	}
}
