package model

import "time"

type Task struct {
	ID             int64     `gorm:"column:id;primaryKey" json:"id"`
	ParameterID    int64     `gorm:"column:paramterid;not null" json:"paramterid"`
	UserID         int64     `gorm:"column:userid;not null" json:"userid"`
	CreateAt       time.Time `gorm:"column:createAt;autoCreateTime" json:"createAt"`
	UpdateAt       time.Time `gorm:"column:updateAt;autoUpdateTime" json:"updateAt"`
	Deleted        int8      `gorm:"column:deleted;default:0" json:"deleted"`
	Desc           string    `gorm:"column:desc;type:varchar(255)" json:"desc"`
	MachineID      int       `gorm:"column:machineid;default:0;not null" json:"machineid"`
	MemoryPercent  int       `gorm:"column:memory_percent;default:50;not null" json:"memory_percent"`
	CPUPercent     int       `gorm:"column:cpu_percent;default:50;not null" json:"cpu_percent"`
	CompletionTime *int      `gorm:"column:completion_time" json:"completion_time"`
	Limit          *int      `gorm:"column:limit" json:"limit"`
}

func (Task) TableName() string {
	return "t_task"
}
