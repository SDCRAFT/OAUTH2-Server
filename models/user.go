package models

type User struct {
	ID            int    `json:"id" gorm:"primarykey,unique"`
	Name          string `json:"name" gorm:"unique"`
	Email         string `json:"email" gorm:"unique"`
	Password      string `json:"password"`
	MinecraftUUID string `json:"uuid"`
}
