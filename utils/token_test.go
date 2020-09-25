package utils

import "testing"

func TestValidateToken(t *testing.T) {
	type args struct {
		authID string
		auth   string
		key    string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "data nil",
			want:    false,
			wantErr: true,
		},
		{
			name: "empty authID",
			args: args{
				authID: "lala",
				auth:   "lalala",
				key:    "get-member",
			},
			want: false,
		},
		{
			name: "token doen't match",
			args: args{
				authID: "get-member",
				auth:   "lalala",
				key:    "get-member",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateToken(tt.args.authID, tt.args.auth, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
