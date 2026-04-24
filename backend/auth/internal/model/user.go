package model

import "time"

type User struct {
	ID          int64      `gorm:"column:id;primaryKey;autoIncrement:false" json:"id"`
	LarkUnionID string     `gorm:"column:lark_union_id;uniqueIndex;type:varchar(64);not null" json:"lark_union_id"`
	LarkAvatar  string     `gorm:"column:lark_avatar_url;type:varchar(512)" json:"lark_avatar_url"`
	Username    string     `gorm:"column:username;uniqueIndex;type:varchar(50);not null" json:"username"`
	CreateAt    time.Time  `gorm:"column:createAt;autoCreateTime" json:"create_at"`
	UpdateAt    time.Time  `gorm:"column:updateAt;autoUpdateTime" json:"update_at"`
	Role        string     `gorm:"column:role;type:json" json:"role"`
	Deleted     int8       `gorm:"column:deleted;default:0" json:"deleted"`
}

func (User) TableName() string {
	return "t_user"
}
