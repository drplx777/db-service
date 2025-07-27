package model

import (
	"time"
)

type Task struct {
	ID            string     `db:"id" json:"id"`
	Title         string     `db:"title" json:"title"`
	Description   string     `db:"description" json:"description"`
	Status        string     `db:"status" json:"status"`
	reporterD     string     `db:"reporterD" json:"reporterD"`
	assignerID    string     `db:"assignerID" json:"assignerID"`
	reviewerID    string     `db:"reviewerID" json:"reviewerID"`
	approverID    string     `db:"approverID" json:"approverID"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
	StartedAt     time.Time  `db:"started_At" json:"started_At"`
	DoneAt        *time.Time `db:"done_at" json:"done_at,omitempty"`
	DeadLine      time.Time  `db:"deadline" json:"deadline"`
	dashboardID   string     `db:"dashboardID" json:"dashboardID"`
	blockedBy     string     `db:"blockedBy" json"blockedBy"`
	approveStatus bool       `db:"approveStatus" json"approveStatus"`
	Completed     bool       `db:"completed" json:"completed"`
}

type User struct {
	ID         int    `db:"id" json:"ID"`
	Name       string `db:"name" json:"user"`
	Sumame     string `db:"sumame" json:"reviewerID"`
	Middlename string `db:"middlename" json:"middlename"`
	Login      string `db:"login" json:"login"`
	RoleID     int    `db:"roleID" json:"roleID"`
	Password   string `db:"password" json:"password"`
	Token      string `db:"token" json:"token"`
}
