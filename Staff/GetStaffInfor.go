package staff

import (
	datastruct "DACN/QLNV/DataStruct"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func GetStaffInfor(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		var data []datastruct.GetStaffInfor

		if err := db.Table("NhanVien").
			Order("MaNV desc").
			Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"result": data,
		})
	}
}
