// build+ integration

package db

import (
	"github.com/dviramontes/developerhappiness.app/internal/config"
	"github.com/jinzhu/gorm"
	"testing"
)

const (
	deleteTestUser = `DELETE FROM users WHERE email = 'bot@bot.com'`
)

func TestDB_CreateUser(t *testing.T) {
	config := config.Read("../../config.yaml", nil)
	connStr := config.GetString("connStr")
	database, err := Connect(connStr)
	if err != nil {
		t.Fatalf("failed to connect to database from test, err: %v", err)
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
		})
	}

	t.Cleanup(func() {
		_, err := database.Conn.DB().Exec(deleteTestUser)
		if err != nil {
			t.Fatalf("failed to clean users table in advertising db, err: %v", err)
		}
	})

}
