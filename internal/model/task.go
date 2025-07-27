package model

import (
	"time"
)

type Task struct {
	ID            string     `db:"id" json:"id"`
	Title         string     `db:"title" json:"title"`
	Description   string     `db:"description" json:"description"`
	Status        string     `db:"status" json:"status"`
	ReporterD     string     `db:"reporterD" json:"reporterD"`
	AssignerID    string     `db:"assignerID" json:"assignerID"`
	ReviewerID    string     `db:"reviewerID" json:"reviewerID"`
	ApproverID    string     `db:"approverID" json:"approverID"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
	StartedAt     time.Time  `db:"started_At" json:"started_At"`
	DoneAt        *time.Time `db:"done_at" json:"done_at,omitempty"`
	DeadLine      time.Time  `db:"deadline" json:"deadline"`
	DashboardID   string     `db:"dashboardID" json:"dashboardID"`
	DlockedBy     string     `db:"blockedBy" json:"blockedBy"`
	ApproveStatus bool       `db:"approveStatus" json:"approveStatus"`
	Completed     bool       `db:"completed" json:"completed"`
}

type User struct {
	ID         int    `db:"id" json:"ID"`
	Name       string `db:"name" json:"user"`
	Surname    string `db:"surname" json:"reviewerID"`
	Middlename string `db:"middlename" json:"middlename"`
	Login      string `db:"login" json:"login"`
	RoleID     int    `db:"roleID" json:"roleID"`
	Password   string `db:"password" json:"password"`
	Token      string `db:"token" json:"token"`
}
