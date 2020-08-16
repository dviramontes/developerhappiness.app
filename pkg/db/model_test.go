// +build integration

package db

import (
	"github.com/jinzhu/gorm"
	"testing"
)

const (
	truncateUsersTable = `TRUNCATE TABLE users;`
)

func TestDB_UserModel(t *testing.T) {
	database, err := Connect("postgres://postgres:postgres@10.254.254.254:5432/happydev_test?sslmode=disable")
	if err != nil {
		t.Fatalf("failed to connect to database from integration test, err: %v", err)
	}
	type fields struct {
		conn *gorm.DB
	}
	type args struct {
		user *User
	}
	var tests = []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "create a bot user",
			fields: fields{
				conn: database.Conn,
			},
			args: args{
				user: &User{
					Name:     "test-bot-user",
					Active:   false,
					IsBot:    true,
					Email:    "bot@bot.com",
					Timezone: "",
					ImgUrl:   "",
					IsAdmin:  false,
					IsOwner:  false,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				Conn: tt.fields.conn,
			}
			if err := db.CreateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			users, err := db.GetUsers()
			if err != nil {
				t.Fatalf("GetUsers() error = %v, wantErr %v", err, tt.wantErr)
			}

			if len(users) != 2 { // bot + seed user
				t.Errorf("GetUsers() failed to retrieve correct number of users got = %d, wanted %d", len(users), 2)

			}
		})
	}

	t.Cleanup(func() {
		_, err := database.Conn.DB().Exec(truncateUsersTable)
		if err != nil {
			t.Fatalf("failed to clean users table in advertising db, err: %v", err)
		}
	})

}
