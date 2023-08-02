package staff

import (
	datastruct "DACN/QLNV/DataStruct"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func CreateNewStaffInfor(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data datastruct.StaffInfor

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"Loi": err,
			})
			return
		}

		if err := db.Table("NhanVien").
			Create(&data).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"Loi": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Them Thong Tin": "Thanh Cong",
		})
	}
}
