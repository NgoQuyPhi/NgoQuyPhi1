package datastruct

import (
	"time"
)

type StaffInfor struct {
	MaNV        int    `json:"MaNV" gorm:"column:MaNV"`
	TenNV       string `json:"TenNV" gorm:"column:TenNV"`
	SoDienThoai string `json:"SoDienThoai" gorm:"column:SoDienThoai"`
	DiaChi      string `json:"DiaChi" gorm:"column:DiaChi"`
	Email       string `json:"Email" gorm:"column:Email"`
}

type StaffUpdateInfor struct {
	TenNV       string `json:"TenNV" gorm:"column:TenNV"`
	SoDienThoai string `json:"SoDienThoai" gorm:"column:SoDienThoai"`
	DiaChi      string `json:"DiaChi" gorm:"column:DiaChi"`
	Email       string `json:"Email" gorm:"column:Email"`
}

type CheckInCheckOut struct {
	MaNV     int        `json:"MaNV" gorm:"column:MaNV"`
	CheckIn  *time.Time `json:"CheckIn" gorm:"column:CheckIn"`
	CheckOut *time.Time `json:"CheckOut" gorm:"column:CheckOut"`
}

type Giolam struct {
	GioVaoLam *time.Time `json:"GioVaoLam" gorm:"column:GioBatDau"`
	GioTanLam *time.Time `json:"GioTanLam" gorm:"column:GioKetThuc"`
}
