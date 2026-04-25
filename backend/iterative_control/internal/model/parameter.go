package model

import "time"

type Parameter struct {
	ID         int64     `gorm:"column:id;primaryKey" json:"id"`
	UserID     int64     `gorm:"column:userid;not null" json:"userid"`
	Parameters string    `gorm:"column:parameters;type:json" json:"parameters"`
	Script     string    `gorm:"column:script;type:mediumtext;not null" json:"script"`
	CreateAt   time.Time `gorm:"column:createAt;autoCreateTime" json:"createAt"`
	UpdateAt   time.Time `gorm:"column:updateAt;autoUpdateTime" json:"updateAt"`
	Deleted    int8      `gorm:"column:deleted;default:0" json:"deleted"`
	Desc       string    `gorm:"column:desc;type:varchar(255)" json:"desc"`
}

func (Parameter) TableName() string {
	return "t_parameter"
}
