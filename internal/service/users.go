package service

import (
	"database/sql"
	"errors"
)

type User struct {
	id       int
	name     string
	password string
}

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// Retorna lista de Usuários
func (s *UserService) Getuser() ([]User, error) {
	query := "SELECT * FROM user"
	rows, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.id, &user.name, &user.password); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// Retorna um Usuário pelo seu ID.
func (s *UserService) GetUserById(id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := s.db.QueryRow(query, id)

	var user User
	if err := row.Scan(&user.id, &user.name, &user.password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// Cria um Usuário
func (s *UserService) CreateUser(user *User) error {
	query := "INSERT INTO users (name, password) VALUES (?, ?)"
	result, err := s.db.Exec(query, user.name, user.password)

	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.id = int(lastInsertID)

	return nil
}

// Atualiza um Usuário
func (s *UserService) UpdateUser(user *User) error {
	query := "UPDATE users SET name = ?, password = ? WHERE id = ?"
	_, err := s.db.Exec(query, user.name, user.password)
	return err
}

// Deleta um Usuário
func (s *UserService) DeleteUser(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := s.db.Exec(query, id)
	return err
}
