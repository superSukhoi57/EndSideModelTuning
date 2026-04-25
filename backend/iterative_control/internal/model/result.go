package model

import "time"

type Result struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Result    string    `gorm:"column:result;type:json;not null" json:"result"`
	UserID    int64     `gorm:"column:userid;not null" json:"userid"`
	MachineID int64     `gorm:"column:machineid;not null" json:"machineid"`
	CreateAt  time.Time `gorm:"column:createAt;autoCreateTime" json:"createAt"`
	UpdateAt  time.Time `gorm:"column:updateAt;autoUpdateTime" json:"updateAt"`
	Deleted   int8      `gorm:"column:deleted;default:0" json:"deleted"`
	Desc      string    `gorm:"column:desc;type:varchar(255)" json:"desc"`
}

func (Result) TableName() string {
	return "t_result"
}
