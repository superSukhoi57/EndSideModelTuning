package model

import "time"

type Machine struct {
	ID      int64     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	IP      string    `gorm:"column:ip;type:varchar(50);not null" json:"ip"`
	Pwd     string    `gorm:"column:pwd;type:varchar(255)" json:"pwd"`
	UserID  int64     `gorm:"column:userid;not null" json:"userid"`
	CreateAt time.Time `gorm:"column:createAt;autoCreateTime" json:"createAt"`
	UpdateAt time.Time `gorm:"column:updateAt;autoUpdateTime" json:"updateAt"`
	Deleted int8      `gorm:"column:deleted;default:0" json:"deleted"`
	IsFinsh int8      `gorm:"column:isfinsh;default:0" json:"isfinsh"`
	ResultID *int64   `gorm:"column:resultid" json:"resultid"`
	Core    *int      `gorm:"column:core" json:"core"`
	RAM     *int      `gorm:"column:ram" json:"ram"`
	Memory  *int      `gorm:"column:memory" json:"memory"`
	OS      string    `gorm:"column:os;type:varchar(100)" json:"os"`
	Desc    string    `gorm:"column:desc;type:varchar(255)" json:"desc"`
}

func (Machine) TableName() string {
	return "t_machine"
}
