package stafffault

import (
	datastruct "DACN/QLNV/DataStruct"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListStaffLateForWork(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var ListTimekeeping []datastruct.Workingtime

		err := db.Table("ChamCong").
			Joins("LEFT JOIN NhanVien on ChamCong.MaNV = NhanVien.MaNV"). //lay danh sach cham cong truyen vao slice danhsachchamcong
			Select("ChamCong.MaNV,NhanVien.MaBP, ChamCong.CheckIn, ChamCong.CheckOut").
			Find(&ListTimekeeping).Error

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
			return
		}

		var LateForWork []datastruct.Workingtime

		for i := range ListTimekeeping {
			if ListTimekeeping[i].GioVaoLam.Hour() > 8 { //kiem tra tung phan tu trong danh sach cham cong
				LateForWork = append(LateForWork, ListTimekeeping[i]) //neu phan tu nao co gio checkin > 8h thi di muon
			}
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"Danh sach nhan vien muon lam:": LateForWork, // xuat ra danh sach nhan vien di muon
		})
	}
}
