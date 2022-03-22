package data

import (
	"context"
)

// UserRepository manages the operations with the database that
// correspond to the user model.
type RoleRepository struct {
	Data *Data
}

// GetRole returns all users.
func (rr *RoleRepository) GetRole(ctx context.Context, id uint) ([]string, error) {
	//	SELECT user_id, role_id, name
	q := `
	SELECT name
	FROM user_roles
	INNER JOIN roles 
	ON user_roles.role_id=roles.id
	WHERE
	user_id=$1;
	`

	rows, err := rr.Data.DB.QueryContext(ctx, q, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var roles []string
	var data string

	for rows.Next() {
		rows.Scan(&data)
		roles = append(roles, data)
	}

	return roles, nil
}
