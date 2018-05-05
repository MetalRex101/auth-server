package models

type Client struct {
	ID uint `gorm:"AUTO_INCREMENT"`
	Name *string `gorm:"size:200"`
	Status bool `gorm:"default:false;index"`
	ClientID *string `gorm:"size:200;index;not null"`
	ClientSecret *string `gorm:"size:200;index;not null"`
	ClientTemplate *string `gorm:"size:50;column:template"`
	IP *string `gorm:"size:96;not null"`
	Url *string `gorm:"size:255;not null"`
	Scope *string `gorm:"size:255"`
}

func (Client) TableName() string {
	return "oauth_clients"
}
