package staff

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func DeleteStaff(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		MaNV, err := strconv.Atoi(c.Param("MaNV"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"loi": err.Error(),
			})
			return
		}

		if err := db.Table("NhanVien").
			Where("MaNV = ?", MaNV).
			Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"loi": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Xoa Nhan Vien:": true,
		})

	}
}
