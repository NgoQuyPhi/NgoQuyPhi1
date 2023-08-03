package stafffault

import (
	datastruct "DACN/QLNV/DataStruct"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DanhSachNhanVienDiMuonTheoBoPhan(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var danhsachchamcong []datastruct.Giolam

		err := db.Table("ChamCong").
			Joins("LEFT JOIN NhanVien on ChamCong.MaNV = NhanVien.MaNV"). //lay danh sach cham cong truyen vao slice danhsachchamcong
			Select("ChamCong.MaNV,NhanVien.MaBP, ChamCong.CheckIn, ChamCong.CheckOut").
			Find(&danhsachchamcong).Error

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
			return
		}

		var muonlam []datastruct.Giolam

		for i := range danhsachchamcong {
			if danhsachchamcong[i].GioVaoLam.Hour() > 8 { //kiem tra tung phan tu trong danh sach cham cong
				danhsachchamcong[i].Muon = true //neu phan tu nao co gio checkin > 8h thi di muon
				muonlam = append(muonlam, danhsachchamcong[i])
			} else {
				danhsachchamcong[i].Muon = false
			}

		}

		c.JSON(http.StatusBadRequest, gin.H{
			"Danh sach nhan vien muon lam:": muonlam, // xuat ra danh sach nhan vien di muon
		})
	}
}
