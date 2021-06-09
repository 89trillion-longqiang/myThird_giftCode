package route

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"gift"
	"setUpRedis"
)

func SetUpRount() *gin.Engine  {
	r := gin.Default()
	c1 := r.Group("/giftCode")

	c1.GET("/adminCreatGiftcode", adminCreatGiftcode)
	c1.GET("/admininquireGiftCode",adminInquireGiftCode)
	c1.GET("/client",client)
	return r
}


func adminCreatGiftcode(c *gin.Context){
	GiftCode := GetRandCode(8)
	des := c.Query("des")
	GiftNum := c.Query("GN")
	ValidPeriod :=c.Query("VP")
	GiftContent :=c.Query("GC")
	CreatePer := c.Query("CP")
	CreatTime := time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05")
	err := setUpRedis.HashSet(gift.Gift{
		GiftCode,
		des,
		GiftNum,
		ValidPeriod,
		GiftContent,
		CreatePer,
		CreatTime,
		"0",
		"",
	})
	if err != nil {

		c.JSON(200,gin.H{
			"condition":"error",
			"GiftCode" : err,
		})

	}else {
		c.JSON(200,gin.H{
			"condition":"success",
			"GiftCode" : GiftCode,
		})
	}

}
func adminInquireGiftCode(c *gin.Context){
	GiftCode := c.Query("giftCode")
	if len(GiftCode) == 0 {
		c.JSON(200,gin.H{
			"condition":"error",
			"giftCode" : "GiftCode is empty",
		})
		return
	}
	if setUpRedis.ExistsKey(GiftCode) {
		ret, err := setUpRedis.HashGetAll(GiftCode)
		if err != nil {
			c.JSON(200, gin.H{
				"condition": "error",
				"GiftCode":  err,
			})
		} else if len(ret) != 0 {
			c.JSON(200, gin.H{
				"condition": "success",
				"data":      ret,
			})
		}
	}else {
		c.JSON(200,gin.H{
			"condition":"error",
			"giftCode" : "GiftCode is error",
		})
	}
}
func client(c *gin.Context)  {
	var errString string
	flagCondition := false
	GiftCode := c.Query("giftCode")
	if GiftCode == ""{
		c.JSON(200,gin.H{
			"condition":"error",
			"GiftCode" : "input GiftCode",
		})
		return
	}
	userName := c.Query("usr")
	if userName == ""{
		c.JSON(200,gin.H{
			"condition":"error",
			"GiftCode" : "input usr",
		})
		return
	}

	ret , err := setUpRedis.HashGetAll(GiftCode)
	ret["GiftCode"] = GiftCode
	if err != nil{
		c.JSON(200,gin.H{
			"condition":"error",
			"GiftCode" : err,
		})
	}
	creatTime,_:=time.Parse("2006-01-02",ret["CreatTime"])
	curTime ,_:=time.Parse("2006-01-02",time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
	d:=creatTime.Sub(curTime)
	validFlo, _ := strconv.ParseFloat(ret["ValidPeriod"], 64)

	if validFlo > d.Hours()/24{
		flagCondition = true
	}else {
		errString = "Expired"
	}
	Claim:= ret["ClaimList"]
	ClaimList := strings.Fields(Claim)
	fmt.Println("ClaimList")
	fmt.Println(ClaimList)
	for _, s := range ClaimList {
		if userName == s {
			c.JSON(200,gin.H{
				"condition":"error",
				"GiftCode" : "User has received",
			})
			return
		}
	}
	avaNum , _:=strconv.Atoi(ret["AvailableNum"])
	giftNum,_ :=strconv.Atoi(ret["GiftNum"])
	if avaNum+1 <= giftNum {
		flagCondition = true
		outString := "{ " + userName + " "+ time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")+"}"
		avanumS := strconv.Itoa(avaNum+1)
		ret["AvailableNum"] = avanumS
		fmt.Println("==============")
		fmt.Println(ret["AvailableNum"])

		ret["ClaimList"] = string(ret["ClaimList"]) +" " + outString + ";"
		fmt.Println("==============")
		fmt.Println(ret["ClaimList"])
		err := setUpRedis.HashSetMap(ret)
		if err != nil {
			fmt.Printf("%s",err)
		}
	}else {

		errString = "Insufficient quantity ；" + errString
	}

	if flagCondition {

		c.JSON(200,gin.H{
			"condition":"success",
			"GiftContent" : ret["GiftContent"],
		})
	}else {
		c.JSON(200,gin.H{
			"condition":    "error",
			"GiftContent" : errString,
		})
	}


}
func GetRandCode(codeLen int) string {
	// 1. 定义原始字符串
	rawStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	// 2. 定义一个buf，并且将buf交给bytes往buf中写数据
	buf := make([]byte, 0, codeLen)
	b := bytes.NewBuffer(buf)
	// 随机从中获取
	rand.Seed(time.Now().UnixNano())
	for rawStrLen := len(rawStr);codeLen > 0; codeLen-- {
		randNum := rand.Intn(rawStrLen)
		b.WriteByte(rawStr[randNum])
	}
	return b.String()
}