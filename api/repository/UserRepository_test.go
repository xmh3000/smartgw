package repository

import (
	"github.com/boltdb/bolt"
	"reflect"
	"smartgw/api/domain"
	"smartgw/lib/database"
	"testing"
)

func TestNewUserRepository(t *testing.T) {
	db, _ := database.NewDatabase()
	defer db.Close()

	type args struct {
		db *bolt.DB
	}
	tests := []struct {
		name string
		args args
		want UserRepository
	}{
		{
			name: "NewUserRepository",
			args: args{
				db: db,
			},
			want: NewUserRepository(db),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Delete(t *testing.T) {
	db, _ := database.NewDatabase()
	defer db.Close()

	type fields struct {
		db *bolt.DB
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "",
			fields:  fields{db: db},
			args:    args{username: "xxx"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userRepository{
				db: tt.fields.db,
			}
			if err := u.Delete(tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_Find(t *testing.T) {
	db, _ := database.NewDatabase()
	defer db.Close()

	type fields struct {
		db *bolt.DB
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.User
		wantErr bool
	}{
		{
			name:   "admin",
			fields: fields{db: db},
			args:   args{username: "admin"},
			want: domain.User{
				Username: "admin",
				Password: "admin",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userRepository{
				db: tt.fields.db,
			}
			got, err := u.Find(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_FindAll(t *testing.T) {
	db, _ := database.NewDatabase()
	defer db.Close()

	type fields struct {
		db *bolt.DB
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.User
		wantErr bool
	}{
		{
			name:   "",
			fields: fields{db: db},
			want: []domain.User{
				domain.User{
					Username: "admin",
					Password: "admin",
				},
				domain.User{
					Username: "bxy",
					Password: "bxy",
				},
				domain.User{
					Username: "张三",
					Password: "张三",
				},
				domain.User{
					Username: "李四",
					Password: "李四",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userRepository{
				db: tt.fields.db,
			}
			got, err := u.FindAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userRepository_Migrate(t *testing.T) {
	db, _ := database.NewDatabase()
	defer db.Close()

	type fields struct {
		db *bolt.DB
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "migrate",
			fields: fields{
				db: db,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userRepository{
				db: tt.fields.db,
			}
			if err := u.Migrate(); (err != nil) != tt.wantErr {
				t.Errorf("Migrate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_Save(t *testing.T) {
	db, _ := database.NewDatabase()
	defer db.Close()

	type fields struct {
		db *bolt.DB
	}
	type args struct {
		user *domain.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "bxy",
			fields: fields{db: db},
			args: args{user: &domain.User{
				Username: "bxy",
				Password: "bxy",
			}},
			wantErr: false,
		},
		{
			name:   "张三",
			fields: fields{db: db},
			args: args{user: &domain.User{
				Username: "张三",
				Password: "张三",
			}},
			wantErr: false,
		},
		{
			name:   "李四",
			fields: fields{db: db},
			args: args{user: &domain.User{
				Username: "李四",
				Password: "李四",
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userRepository{
				db: tt.fields.db,
			}
			if err := u.Save(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
