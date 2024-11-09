package repositories

import (
	"CENT_Notes/cmd/models"
	"CENT_Notes/cmd/storage"
)

// CreateUser inserts a new user into the database
func CreateUser(user models.User) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// GetUser retrieves a user from the database by ID
func GetUser(id int) (models.User, error) {
	db := storage.GetDB()
	var user models.User
	sqlStatement := `SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1`
	err := db.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUser updates an existing user in the database
func UpdateUser(user models.User) (models.User, error) {
	db := storage.GetDB()
	sqlStatement := `UPDATE users SET name = $1, email = $2, password = $3, updated_at = NOW() WHERE id = $4`
	_, err := db.Exec(sqlStatement, user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		return user, err
	}
	return user, nil
}

// DeleteUser deletes a user from the database by ID
func DeleteUser(id int) error {
	db := storage.GetDB()
	sqlStatement := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	return nil
}
