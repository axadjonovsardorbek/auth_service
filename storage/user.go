package storage

import (
	m "auth-service/models"
	"database/sql"
	"github.com/google/uuid"
)

type UserManager struct {
	Conn *sql.DB
}

func NewUserManager(db *sql.DB) *UserManager {
	return &UserManager{Conn: db}
}

func (u *UserManager) Register(req m.RegisterReq) (*m.User, error) {
	id := uuid.NewString()

	row := u.Conn.QueryRow("INSERT INTO users (id, username, password, email) VALUES ($1, $2, $3, $4) RETURNING *", id, req.Username, req.Password, req.Email)
	user := m.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserManager) Login(req m.LoginReq) error {
	// Retrieve the user's information from the database based on the provided username
	row := u.Conn.QueryRow("SELECT id, username, password FROM users WHERE username = $1", req.Username)
	var (
		id       uint32
		username string
		password string
	)
	err := row.Scan(&id, &username, &password)

	//------------------------------------

	return err
}

func (u *UserManager) Profile(id string) (*m.UserRes, error) {
	row := u.Conn.QueryRow("SELECT id, username FROM users WHERE id = $1", id)
	var (
		user_id  uint32
		username string
	)
	err := row.Scan(&user_id, &username)
	if err != nil {
		return nil, err
	}

	return &m.UserRes{
		ID:       id,
		Username: username,
	}, nil
}
