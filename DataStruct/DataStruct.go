package datastruct

import (
	"time"
)

type StaffInfor struct {
	MaNV        int     `json:"MaNV" gorm:"column:MaNV"`
	TenNV       string  `json:"TenNV" gorm:"column:TenNV"`
	SoDienThoai string  `json:"SoDienThoai" gorm:"column:SoDienThoai"`
	DiaChi      string  `json:"DiaChi" gorm:"column:DiaChi"`
	Email       string  `json:"Email" gorm:"column:Email"`
	Luong       float32 `json:"luong" gorm:"column:Luong"`
	MaBoPhan    string  `json:"BoPhan" gorm:"column:MaBP"`
}

type GetStaffInfor struct {
	MaNV   int    `json:"MaNV" gorm:"column:MaNV"`
	TenNV  string `json:"TenNV" gorm:"column:TenNV"`
	Email  string `json:"Email" gorm:"column:Email"`
	BoPhan string `json:"BoPhan" gorm:"column:BoPhan"`
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
	MaNV      int        `json:"MaNV" gorm:"column:MaNV"`
	MaBoPhan  string     `json:"BoPhan" gorm:"column:MaBP"`
	GioVaoLam *time.Time `json:"GioVaoLam" gorm:"column:CheckIn"`
	GioTanLam *time.Time `json:"GioTanLam" gorm:"column:CheckOut"`
	Muon      bool
}
