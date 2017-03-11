package store

import (
	"time"
)

type UserModel struct {
	ID        int
	Name      string
	CreatedAt time.Time `db:"created_at"`
}

func (db *DB) NewUser() UserModel {
	return UserModel{}
}

func (db *DB) GetUsers() ([]UserModel, error) {
	users := []UserModel{}
	err := db.conn.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil

}

func (db *DB) GetUser(id int) (UserModel, error) {
	user := UserModel{}
	err := db.conn.Get(&user, "SELECT * FROM users where id = $1", id)
	if err != nil {
		return user, err
	}
	return user, nil

}

func (db *DB) SaveUser(model UserModel) {
	tx := db.conn.MustBegin()
	//tx.NamedExec("UPDATE users SET name = :name where id = :id", &model)
	tx.NamedExec("INSERT INTO users (name) VALUES (:name)", &model)
	tx.Commit()
}
