package role

type Role struct {
	ID	int	`json:"id" gorm:"column:id;primary_key;auto_increment;"`

	Name	string	`json:"name" gorm:"column:name"`
	Description	string	`json:"description" gorm:"column:description"`

	CreatedAt	time.Time	`json:"created_at" gorm:"column:created_at"`
	UpdatedAt	time.Time	`json:"updated_at" gorm:"column:updated_at"`
	DeletedAt	*time.Time	`json:"-" gorm:"column:deleted_at"`
}














