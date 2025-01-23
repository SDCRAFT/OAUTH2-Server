package models

type User struct {
	ID            string `json:"id" gorm:"primarykey,unique"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	MinecraftUUID string `json:"uuid"`
}
