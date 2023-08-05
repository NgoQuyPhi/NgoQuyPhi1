package salary

import (
	datastruct "DACN/QLNV/DataStruct"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type salarystaffinfor struct {
	MaNV  int     `json:"MaNV" gorm:"column:MaNV"`
	Luong float32 `json:"luong" gorm:"column:Luong"`
}

func SALARY(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		MaNV, err := strconv.Atoi(c.Param("MaNV"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
			return
		}

		Month, err := strconv.Atoi(c.Param("month"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
			return
		}

		var salaryinfor salarystaffinfor

		if err := db.Table("NhanVien").
			Where("MaNV = ?", MaNV).
			Find(&salaryinfor).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
			return
		}

		var WorkingCalendar []datastruct.CheckInCheckOut

		if err := db.Table("ChamCong").
			Where("MaNV = ? and MONTH(CheckIn) = ?", MaNV, Month).
			Find(&WorkingCalendar).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Loi": err.Error(),
			})
			return
		}
		count := 0
		for i := range WorkingCalendar {
			if WorkingCalendar[i].CheckIn.Hour() > 8 {
				count++
			}
		}

		totalsalary := salaryinfor.Luong - (salaryinfor.Luong * 2 * float32(count) / 100)

		c.JSON(http.StatusOK, gin.H{
			"Luong:": totalsalary,
		})
	}
}
