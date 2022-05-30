package cake

var (
	Column = []interface{}{"title", "rating", "description"}
)

type (
	Cake struct {
		Id          int      `gorm:"column:id"`
		Title       string   `gorm:"column:title"`
		Description string   `gorm:"column:description"`
		Image       string   `gorm:"column:image"`
		Rating      *float32 `gorm:"column:rating"`
		CreatedBy   int      `gorm:"column:created_by"`
		CreatedAt   string   `gorm:"column:created_at"`
		UpdatedBy   *int     `gorm:"column:updated_by"`
		UpdatedAt   *string  `gorm:"column:updated_at"`
		DeletedBy   *int     `gorm:"column:deleted_by"`
		DeletedAt   *string  `gorm:"column:deleted_at"`
	}
	ListCakes struct {
		Cakes []Cake
	}
)
