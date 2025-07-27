package home_page

import (
	"log"
	"net/http"

	"api/internal/database"
	"api/internal/models/home_page"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetCarouselByType 根据类型获取公司或学校轮播数据
func GetCarouselByType(c *gin.Context) {
	carouselType := c.Query("type") // 获取 URL 查询参数 ?type=...

	var partners []home_page.PartnerCarousel
	var result *gorm.DB

	// 根据 type 参数动态选择查询的表
	if carouselType == "companies" {
		result = database.DB.Table("partner_carousel_companies").Find(&partners)
	} else if carouselType == "schools" {
		result = database.DB.Table("partner_carousel_schools").Find(&partners)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效或缺少类型参数"})
		return
	}

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法检索合作伙伴轮播数据"})
		return
	}

	log.Printf("Found %d partners for type %s", len(partners), carouselType)

	// 为每个 partner 的 logo 构建完整的 URL
	for i := range partners {
		log.Printf("Original logo: %s", partners[i].Logo)

		partners[i].Logo = "/static/images/home_page/PartnerCarousel/" + partners[i].Logo
		log.Printf("Constructed logo URL: %s", partners[i].Logo)
	}

	c.JSON(http.StatusOK, partners)
}
