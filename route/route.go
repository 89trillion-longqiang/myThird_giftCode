package route

import (
	"giftCode/handle"
	"github.com/gin-gonic/gin"
)

func SetUpRount() *gin.Engine  {
	r := gin.Default()
	c1 := r.Group("/giftCode")

	c1.GET("/adminCreatGiftcode", adminCreatGiftcode)
	c1.GET("/admininquireGiftCode", adminInquireGiftCode)
	c1.GET("/client", client)
	return r
}


func adminCreatGiftcode(c *gin.Context){
	des := c.Query("des")
	GiftNum := c.Query("GN")
	ValidPeriod :=c.Query("VP")
	GiftContent :=c.Query("GC")
	CreatePer := c.Query("CP")
	retMap := handle.HandleAdminCreatGiftcode(des,GiftNum,ValidPeriod,GiftContent,CreatePer)
	c.JSON(200,gin.H{
		"condition":retMap["condition"],
		"GiftCode" : retMap["GiftCode"],
	})
}


func adminInquireGiftCode(c *gin.Context){
	GiftCode := c.Query("giftCode")
	retMap,infoMap := handle.HadnleAdminInquireGiftCode(GiftCode)
	c.JSON(200, gin.H{
		"condition": retMap["condition"],
		"GiftCode":  GiftCode,
		"data" : infoMap,
	})

}


func client(c *gin.Context)  {
	GiftCode := c.Query("giftCode")
	userName := c.Query("usr")
	ret := handle.HandleClient(GiftCode,userName)
	c.JSON(200,gin.H{
		"condition" :ret["condition"] ,
		"GiftContent" : ret["GiftContent"],
		"GiftCode" : ret["GiftCode"],
	})
}

