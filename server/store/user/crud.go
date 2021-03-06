package user

import (
	"database/sql"

	"dg-server/core"
	"dg-server/store/base/db"
)

// New returns a new UserStore.
func New(db *db.DB) core.UserStore {
	return &userStore{db}
}

type userStore struct {
	db *db.DB
}

// Create persists a new user to the datastore.
func (s *userStore) Create(user *core.User) error {
	return s.db.Lock(func(execer db.Execer, binder db.Binder) error {
		params := toParams(user)
		stmt, args, err := binder.BindNamed(stmtInsert, params)
		if err != nil {
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			return err
		}
		user.ID, err = res.LastInsertId()
		return err
	})
}

// Count returns a count of active users.
func (s *userStore) Count() (int64, error) {
	var out int64
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		return queryer.QueryRow(queryCount).Scan(&out)
	})
	return out, err
}

// FindLogin returns a user from the datastore by username.
func (s *userStore) FindLogin(login string) (*core.User, error) {
	out := &core.User{Login: login}
	err := s.db.View(func(queryer db.Queryer, binder db.Binder) error {
		params := toParams(out)
		query, args, err := binder.BindNamed(queryLogin, params)
		if err != nil {
			return err
		}
		row := queryer.QueryRow(query, args...)
		err = scanRow(row, out)
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	})
	return out, err
}

const queryBase = `
SELECT
 user_id
,user_login
,user_password
,user_email
,user_admin
,user_active
,user_avatar
,user_created
,user_updated
,user_last_login
`

const queryCount = `
SELECT COUNT(*)
FROM users
`

const queryLogin = queryBase + `
FROM users
WHERE user_login = :user_login
`

const stmtInsert = `
INSERT INTO users (
 user_login
,user_password
,user_email
,user_admin
,user_active
,user_avatar
,user_created
,user_updated
,user_last_login
) VALUES (
 :user_login
,:user_password
,:user_email
,:user_admin
,:user_active
,:user_avatar
,:user_created
,:user_updated
,:user_last_login
)
`
