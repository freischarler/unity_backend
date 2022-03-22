// pkg/post/post.go
package role

import (
	"time"
)

// Role of the system.
type Role struct {
	User_id   uint      `json:"user_id,omitempty"`
	Role_id   uint      `json:"role_id,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// Role of the system.
type Roles struct {
	Id        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
