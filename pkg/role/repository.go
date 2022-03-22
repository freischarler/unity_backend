package role

import (
	"context"
)

// Repository handle the CRUD operations with Roles.
type Repository interface {
	GetRole(ctx context.Context, id uint) ([]string, error)
}
