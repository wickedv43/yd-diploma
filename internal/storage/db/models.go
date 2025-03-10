// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Bill struct {
	ID          int32
	OrderNumber string
	UserID      int32
	Sum         int32
	ProcessedAt time.Time
}

type Order struct {
	Number     string
	UserID     int32
	Status     string
	Accrual    int32
	UploadedAt time.Time
}

type User struct {
	ID               int32
	Login            string
	Password         string
	BalanceCurrent   int32
	BalanceWithdrawn int32
}
