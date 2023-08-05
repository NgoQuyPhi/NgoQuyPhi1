package main

import (
	dbconnection "DACN/QLNV/DBconnect"
	Salary "DACN/QLNV/Salary"
	Staff "DACN/QLNV/Staff"
	stafffault "DACN/QLNV/StaffFault"
	timekeeping "DACN/QLNV/TimeKeeping"

	"github.com/gin-gonic/gin"
)

func main() {
	db := dbconnection.ConnectDB("testuser:testpass@tcp(127.0.0.1:3306)/QLNV?charset=utf8mb4&parseTime=True&loc=Local")
	r := gin.Default()

	NV := r.Group("/NV")
	{
		NV.POST("/newstaff", Staff.CreateNewStaffInfor(db))
		NV.GET("/:MaNV", Staff.GetStaffInforByID(db))
		NV.GET("/allstaff", Staff.GetStaffInfor(db))
		NV.PUT("/update/:MaNV", Staff.UpdateStaffInfor(db))
		NV.DELETE("/del/:MaNV", Staff.DeleteStaff(db))
	}

	ChamCong := r.Group("/CC")
	{
		ChamCong.POST("/CheckIn/:MaNV", timekeeping.CheckIn(db))
		ChamCong.PUT("/CheckOut/:MaNV", timekeeping.CheckOut(db))
		ChamCong.GET("/CheckInfor/:MaNV", timekeeping.CheckInfor(db))
	}
	TinhLuong := r.Group("/luong")
	{
		TinhLuong.GET("/:month/:MaNV", Salary.SALARY(db))
		TinhLuong.GET("/Danhsachphamloi", stafffault.ListStaffLateForWork(db))
	}
	r.Run()

}
