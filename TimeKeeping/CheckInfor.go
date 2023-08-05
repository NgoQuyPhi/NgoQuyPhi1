package timekeeping

import (
	datastruct "DACN/QLNV/DataStruct"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func CheckInfor(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		Manv, err := strconv.Atoi(c.Param("MaNV"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"loi": err.Error(),
			})
			return
		}

		var data []datastruct.CheckInCheckOut

		if err := db.Table("ChamCong").
			Where("MaNV = ?", Manv).
			Find(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
