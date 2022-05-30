package cake

type (
	CakeListRes struct {
		Id          int     `json:"id"`
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Image       string  `json:"image"`
		Rating      float32 `json:"rating"`
		CreatedAt   string  `json:"created_at"`
		UpdatedAt   *string `json:"updated_at"`
	}
	CakeDetailRes struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Image       string  `json:"image"`
		Rating      float32 `json:"rating"`
		CreatedBy   int     `json:"created_by"`
		CreatedAt   string  `json:"created_at"`
		UpdatedBy   *int    `json:"updated_by"`
		UpdatedAt   *string `json:"updated_at"`
	}
)
