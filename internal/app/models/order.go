package models

import (
	"go-assignment-bootcamp/internal/app/models/types"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID             uint              `gorm:"primaryKey"`
	AgencyId       uint              `gorm:"type:bigint;not null;index:idx_orders_agency_id"`
	Agency         Agency            `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Status         types.OrderStatus `gorm:"type:varchar(255);not null;default:'waiting';check:status in ('canceled', 'completed', 'paid', 'prepayment', 'waiting')"`
	IsChecked      bool              `gorm:"type:boolean;not null;default:false"`
	IsConfirmed    bool              `gorm:"type:boolean;not null;default:false"`
	RentalDate     *time.Time        `gorm:"type:timestamp(0)"`
	GuestCount     int               `gorm:"type:integer;not null"`
	TransportCount int               `gorm:"type:integer;not null"`
	UserName       *string           `gorm:"type:varchar(255)"`
	Email          string            `gorm:"type:varchar(255);not null"`
	Phone          string            `gorm:"type:varchar(255);not null"`
	Note           *string           `gorm:"type:text"`
	AdminNote      *string           `gorm:"type:text"`
	ConfirmedAt    *time.Time        `gorm:"type:timestamp(0)"`
	CreatedAt      *time.Time        `gorm:"type:timestamp(0)"`
	UpdatedAt      *time.Time        `gorm:"type:timestamp(0)"`
	DeletedAt      *gorm.DeletedAt   `gorm:"type:timestamp(0)"`
}

const (
	TableOrders               = "orders"
	ColumnIDTableOrders       = "id"
	RelationAgencyTableOrders = "Agency"
)
