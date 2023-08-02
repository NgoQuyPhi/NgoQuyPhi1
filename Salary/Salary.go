package salary

import (
	datastruct "DACN/QLNV/DataStruct"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Tinhluong(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		MaNV, err := strconv.Atoi(c.Param("MaNV"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
			return
		}

		Thang, err := strconv.Atoi(c.Param("Thang"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
			return
		}

		var bangluong []datastruct.Giolam

		if err := db.Table("ChamCong").
			Select("CheckIn as GioVaoLam, CheckOut as GioTanLam").
			Where("MaNV = ? AND Month(CheckIn) = ?", MaNV, Thang).
			Find(&bangluong).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
			return
		}

		GioQuyDinh := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 8, 0, 0, 0, time.UTC)
		muonlam := 0

		for i := range bangluong {
			if bangluong[i].GioVaoLam.After(GioQuyDinh) {
				muonlam++
			}
		}

		GioVeQuyDinh := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 18, 0, 0, 0, time.UTC)
		vesom := 0
		for i := range bangluong {
			if bangluong[i].GioVaoLam.Before(GioVeQuyDinh) {
				vesom++
			}
		}

		var luongcoban float32

		if err := db.Table("NhanVien").
			Where("MaNV = ?", MaNV).
			Select("Luong").
			Find(&luongcoban).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
			return
		}

		var luong float32

		luong = (luongcoban - (luongcoban * (float32(muonlam) / 100)) - (luongcoban * (float32(vesom) / 100)))

		c.JSON(http.StatusOK, gin.H{
			"luong cua ban:": luong,
		})

	}
}
