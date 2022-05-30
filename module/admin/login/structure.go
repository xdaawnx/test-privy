package login

type (

	//User struct
	User struct {
		UserID                   int    `json:"user_id" gorm:"column:user_id"`
		RBACRoleID               int    `json:"rbac_role-id" gorm:"column:rbac_role-id"`
		Username                 string `json:"username" gorm:"column:username"`
		Email                    string `json:"email" gorm:"column:email"`
		Password                 string `json:"password" gorm:"column:password"`
		Status                   string `json:"status" gorm:"column:status"`
		VerificationToken        string `json:"verification_token" gorm:"column:verification_token"`
		VerificationTokenExpTime string `json:"verif_token_exp_time" gorm:"column:verif_token_exp_time"`
		UpdatedDate              string `json:"updated_date" gorm:"column:updated_date"`
		UpdatedBy                int    `json:"updated_by" gorm:"column:updated_by"`
		UpdatedByName            string `json:"updated_by_name" gorm:"column:updated_by_name"`
		LastUpdatedPassword      string `json:"last_updated_password" gorm:"column:last_updated_password"`
	}
)

//TableName ResetPass
func (User) TableName() string {
	return "user"
}
