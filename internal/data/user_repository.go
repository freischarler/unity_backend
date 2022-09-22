package data

import (
	"context"
	"time"

	"github.com/martinpaz/restfulapi/pkg/user"
)

// UserRepository manages the operations with the database that
// correspond to the user model.
type UserRepository struct {
	Data *Data
}

// GetAll returns all users.
func (ur *UserRepository) GetAll(ctx context.Context) ([]user.User, error) {
	q := `
	SELECT id, username, email, picture,
		authorized, created_at, updated_at
		FROM users;
	`

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		rows.Scan(&u.ID, &u.Username,
			&u.Email, &u.Picture, &u.Authorized, &u.CreatedAt, &u.UpdatedAt)
		users = append(users, u)
	}

	return users, nil
}

// GetOne returns one user by id.
func (ur *UserRepository) GetOne(ctx context.Context, id uint) (user.User, error) {
	q := `
	SELECT id, username, email, picture,
		created_at, updated_at
		FROM users WHERE id = $1;
	`

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var u user.User
	err := row.Scan(&u.ID, &u.Username, &u.Email,
		&u.Picture, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return user.User{}, err
	}

	return u, nil
}

// GetByUsername returns one user by username.
func (ur *UserRepository) GetByUsername(ctx context.Context, username string) (user.User, error) {
	q := `
	SELECT id, username, email, picture,
		password, created_at, updated_at
		FROM users WHERE username = $1;
	`
	row := ur.Data.DB.QueryRowContext(ctx, q, username)

	var u user.User
	err := row.Scan(&u.ID, &u.Username,
		&u.Email, &u.Picture, &u.PasswordHash, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return user.User{}, err
	}

	q = `
	SELECT name
	FROM user_roles
	INNER JOIN roles 
	ON user_roles.role_id=roles.id
	WHERE
	user_id=$1;
	`

	rows, err := ur.Data.DB.QueryContext(ctx, q, u.ID)
	if err != nil {
		return user.User{}, err
	}

	var roles []string
	var data string

	for rows.Next() {
		rows.Scan(&data)
		roles = append(roles, data)
	}

	u.Roles = roles

	return u, nil
}

// Create adds a new user.
func (ur *UserRepository) Create(ctx context.Context, u *user.User) error {
	q := `
	INSERT INTO users (username, email, picture, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`

	if u.Picture == "" {
		u.Picture = "https://placekitten.com/g/300/300"
	}

	if err := u.HashPassword(); err != nil {
		return err
	}

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, u.Username, u.Email,
		u.Picture, u.PasswordHash, time.Now(), time.Now(),
	)

	err = row.Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

// Update updates a user by id.
func (ur *UserRepository) Update(ctx context.Context, id uint, u user.User) error {
	q := `
	UPDATE users set , email=$3, picture=$4, updated_at=$5
		WHERE id=$6;
	`

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, u.Email,
		u.Picture, time.Now(), id,
	)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes a user by id.
func (ur *UserRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM users WHERE id=$1;`

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
