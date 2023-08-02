package staff

import (
	datastruct "DACN/QLNV/DataStruct"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func UpdateStaffInfor(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data datastruct.StaffUpdateInfor

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi:": err.Error(),
			})
			return
		}

		MaNV, err := strconv.Atoi(c.Param("MaNV"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"loi": err.Error(),
			})
			return
		}

		if err := db.Table("NhanVien").
			Where("MaNV = ?", MaNV).
			Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"loi": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Sua Thong Tin Thanh Cong:": true,
		})
	}
}
