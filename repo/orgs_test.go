package repo

import (
	"errors"
	"reflect"
	"saturday/domain"
	"saturday/interfaces"
	mi "saturday/mocks"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
)

func TestOrgsRepo_SaveComment(t *testing.T) {
	any := gomock.Any()
	var intf interface{}

	type fields struct {
		DB interfaces.Database
	}
	type args struct {
		orgsName string
		comment  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(mock *mi.MockDatabase)
		wantErr bool
	}{
		{
			name: "error case, build query",
			args: args{
				orgsName: "xendit",
				comment:  "amaze",
			},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any, any).Return("", intf, errors.New("errors"))
			},
			wantErr: true,
		},
		{
			name: "error case, exec query",
			args: args{
				orgsName: "xendit",
				comment:  "amaze",
			},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Exec(any, any).Return(sqlmock.NewResult(0, 0), errors.New("errors"))
			},
			wantErr: true,
		},
		{
			name: "error case, row affected after query",
			args: args{
				orgsName: "xendit",
				comment:  "amaze",
			},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Exec(any, any).Return(sqlmock.NewErrorResult(errors.New("errors")), nil)
			},
			wantErr: true,
		},
		{
			name: "error case, no row affected",
			args: args{
				orgsName: "xendit",
				comment:  "amaze",
			},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Exec(any, any).Return(sqlmock.NewResult(0, 0), nil)
			},
			wantErr: true,
		},
		{
			name: "success case",
			args: args{
				orgsName: "xendit",
				comment:  "amaze",
			},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Exec(any, any).Return(sqlmock.NewResult(0, 1), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			or := &OrgsRepo{
				DB: tt.fields.DB,
			}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockDB := mi.NewMockDatabase(mockCtrl)

			tt.mock(mockDB)

			or.DB = mockDB

			if err := or.SaveComment(tt.args.orgsName, tt.args.comment); (err != nil) != tt.wantErr {
				t.Errorf("OrgsRepo.SaveComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrgsRepo_GetComment(t *testing.T) {
	any := gomock.Any()
	var intf interface{}
	comments := make([]*domain.OrgsCommentBody, 0)

	type fields struct {
		DB interfaces.Database
	}
	type args struct {
		orgsName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(mock *mi.MockDatabase)
		want    []*domain.OrgsCommentBody
		wantErr bool
	}{
		{
			name: "error case, build query",
			args: args{orgsName: "xendit"},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any).Return("", intf, errors.New("errors"))
			},
			want:    comments,
			wantErr: true,
		},
		{
			name: "error case, select query",
			args: args{orgsName: "xendit"},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Select(any, any, any).Return(errors.New("errors"))
			},
			want:    comments,
			wantErr: true,
		},
		{
			name: "error case, select query",
			args: args{orgsName: "xendit"},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Select(any, any, any).Return(nil)
			},
			want: comments,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			or := &OrgsRepo{
				DB: tt.fields.DB,
			}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockDB := mi.NewMockDatabase(mockCtrl)

			tt.mock(mockDB)

			or.DB = mockDB

			got, err := or.GetComment(tt.args.orgsName)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrgsRepo.GetComment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrgsRepo.GetComment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrgsRepo_DeleteComment(t *testing.T) {
	any := gomock.Any()
	var intf interface{}

	type fields struct {
		DB interfaces.Database
	}
	type args struct {
		orgsName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(mock *mi.MockDatabase)
		wantErr bool
	}{
		{
			name: "error case, build query",
			args: args{orgsName: "xendit"},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any).Return("", intf, errors.New("errors"))
			},
			wantErr: true,
		},
		{
			name: "error case, exec query",
			args: args{orgsName: "xendit"},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Exec(any, any).Return(sqlmock.NewErrorResult(errors.New("errors")), errors.New("errors"))
			},
			wantErr: true,
		},
		{
			name: "error case, row affected query",
			args: args{orgsName: "xendit"},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Exec(any, any).Return(sqlmock.NewErrorResult(errors.New("errors")), nil)
			},
			wantErr: true,
		},
		{
			name: "success case, no row affected",
			args: args{orgsName: "xendit"},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Exec(any, any).Return(sqlmock.NewResult(0, 0), nil)
			},
		},
		{
			name: "success case, row affected",
			args: args{orgsName: "xendit"},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Exec(any, any).Return(sqlmock.NewResult(0, 1), nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			or := &OrgsRepo{
				DB: tt.fields.DB,
			}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockDB := mi.NewMockDatabase(mockCtrl)

			tt.mock(mockDB)

			or.DB = mockDB

			if err := or.DeleteComment(tt.args.orgsName); (err != nil) != tt.wantErr {
				t.Errorf("OrgsRepo.DeleteComment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOrgsRepo_GetMember(t *testing.T) {
	any := gomock.Any()
	var intf interface{}
	members := make([]*domain.GetMemberResult, 0)

	type fields struct {
		DB interfaces.Database
	}
	type args struct {
		orgsName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		mock    func(mock *mi.MockDatabase)
		want    []*domain.GetMemberResult
		wantErr bool
	}{
		{
			name: "error case, build query",
			args: args{orgsName: "xendit"},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any).Return("", intf, errors.New("errors"))
			},
			want:    members,
			wantErr: true,
		},
		{
			name: "error case, select query",
			args: args{orgsName: "xendit"},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Select(any, any, any).Return(errors.New("errors"))
			},
			want:    members,
			wantErr: true,
		},
		{
			name: "error case, select query",
			args: args{orgsName: "xendit"},
			mock: func(mock *mi.MockDatabase) {
				mock.EXPECT().In(any, any).Return("", intf, nil)
				mock.EXPECT().Rebind(any)
				mock.EXPECT().Select(any, any, any).Return(nil)
			},
			want: members,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			or := &OrgsRepo{
				DB: tt.fields.DB,
			}

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockDB := mi.NewMockDatabase(mockCtrl)

			tt.mock(mockDB)

			or.DB = mockDB

			got, err := or.GetMember(tt.args.orgsName)
			if (err != nil) != tt.wantErr {
				t.Errorf("OrgsRepo.GetMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrgsRepo.GetMember() = %v, want %v", got, tt.want)
			}
		})
	}
}
