package table

type User struct {
	ID        int64  `json:"id"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
	Name      string `json:"name"`
	Gender    int8   `json:"gender"`
	Age       int8   `json:"age"`
}

func (t *User) TableName() string {
	return "user"
}

func (t *User) Show() (*User, error) {
	var user User
	err := DB.First(&user, t.ID).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
