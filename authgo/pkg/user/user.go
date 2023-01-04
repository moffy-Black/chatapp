package user

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type ChatUser struct {
	Id        int    `json:"id"`
	Name      string `json:"name" form:"name"`
	Password  string `json:"password" form:"password"`
	CreatedBy time.Time
}

func (cu ChatUser) Get(db *sql.DB) (chatuser ChatUser, err error) {
	row := db.QueryRow("SELECT * FROM users WHERE name=$1 AND password=$2", cu.Name, cu.Password)
	err = row.Scan(&chatuser.Id, &chatuser.Name, &chatuser.Password, &chatuser.CreatedBy)
	if err != nil {
		return
	}
	return
}

func (cu ChatUser) Add(db *sql.DB) (chatuser ChatUser, err error) {
	row := db.QueryRow("INSERT INTO users(name, password, created_by) VALUES ($1, $2, $3) RETURNING id", cu.Name, cu.Password, cu.CreatedBy)
	err = row.Scan(&cu.Id)
	if err != nil {
		return
	}
	chatuser = cu
	return
}
