package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	ProjectName string      `json:"project_name"`
	Task        []TaskModel `gorm:"foreignKey:ProjectID"`
}

type TaskModel struct {
	gorm.Model
	TaskName    string `json:"task_name"`
	Status      string `json:"status" gorm:"default:in-queue"`
	Description string `json:"description"`
	ProjectID   uint   `json:"project_id"`
}
