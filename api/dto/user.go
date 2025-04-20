package dto

type GetOtpRequest struct {
	MobileNumber string `json:"mobileNumber" binding:"required,mobile,min=11,max=11"`
}

type TokenDetail struct {
	AccessToken            string
	RefreshToken           string
	AccessTokenExpireTime  int64
	RefreshTokenExpireTime int64
}
