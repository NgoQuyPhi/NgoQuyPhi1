package timekeeping

import (
	datastruct "DACN/QLNV/DataStruct"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func CheckOut(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		Manv, err := strconv.Atoi(c.Param("MaNV"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"loi": err.Error(),
			})
			return
		}

		now := time.Now().UTC()

		data := datastruct.CheckInCheckOut{
			MaNV:     Manv,
			CheckOut: &now,
		}

		if err := db.Table("ChamCong").
			Where("MaNV = ? and CheckOut is null", Manv).
			Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"loi": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"CheckOut": true,
		})
	}
}
