package data

import (
	"testing"
	"database/sql"
)

func Test_UserCreate(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	if users[0].Id == 0 {
		t.Error("No id or created_at in user.")
	}
	u, err := UserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "User not created.")
	}
	if users[0].Email != u.Email {
		t.Error("User retrieved is not the same as the one created.")
	}
}

func Test_UserDelete(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Cannot create user.")
	}
	if err := users[0].Delete(); err != nil {
		t.Error(err, "- Cannot delete user")
	}
	if _, err := UserByEmail(users[0].Email); err != sql.ErrNoRows {
		t.Error(err, "- User not deleted.")
	}
}
