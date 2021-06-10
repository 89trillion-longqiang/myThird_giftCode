package handle

import (
	"fmt"
	"giftCode/gift"
	"strconv"
	"strings"
	"time"

	"giftCode/setUpRedis"

)
func HandleAdminCreatGiftcode(des string,GiftNum string,ValidPeriod string,GiftContent string,CreatePer string)  map[string]string{
	retMap := make(map[string]string,2)
	GiftCode := gift.GetRandCode(8)
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
		retMap["condition"]="error"
		retMap["GiftCode" ] =  err.Error()
	}else {
		retMap["condition"]="success"
		retMap["GiftCode" ] =  GiftCode
	}

	return retMap
}
func HadnleAdminInquireGiftCode(GiftCode string) (map[string]string,map[string]string){
	retMap := make(map[string]string,2)
	if len(GiftCode) == 0 {
		retMap["condition"]="error"
		retMap["giftCode"] = "GiftCode is empty"
		return retMap , nil
	}

	if setUpRedis.ExistsKey(GiftCode) {
		ret, err := setUpRedis.HashGetAll(GiftCode)
		if err != nil {
			retMap["condition"]="error"
			retMap["giftCode"] =  err.Error()
		} else if len(ret) != 0 {
			retMap["condition"]="success"

			return retMap,ret
		}
	}else {
		retMap["condition"]="error"
		retMap["giftCode"] = "GiftCode is error"
	}
	return retMap , nil
}
func HandleClient(GiftCode string,userName string) map[string]string{
	var errString string
	var flagCondition bool
	retMap := make(map[string]string,3)
	retMap["GiftCode"] = GiftCode
	if GiftCode == ""{
		retMap["condition"]="error"
		retMap["GiftCode" ]="input GiftCode"
		return retMap
	}
	if userName == ""{
		retMap["condition"]="error"
		retMap["GiftCode" ]="input usr"
		return retMap
	}

	ret , err := setUpRedis.HashGetAll(GiftCode)
	ret["GiftCode"] = GiftCode
	if err != nil{
		retMap["condition"]="error"
		retMap["GiftCode" ]= err.Error()
		return retMap
	}

	creatTime,_:=time.Parse("2006-01-02",ret["CreatTime"])
	curTime ,_:=time.Parse("2006-01-02",time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
	d:=creatTime.Sub(curTime)
	validFlo, _ := strconv.ParseFloat(ret["ValidPeriod"], 64)

	if validFlo > d.Hours()/24{
		flagCondition =true
	}else {
		errString = "Expired"
	}
	Claim:= ret["ClaimList"]
	ClaimList := strings.Fields(Claim)
	for _, s := range ClaimList {
		if userName == s {

			retMap["condition"]="error"
			retMap["GiftCode"] = "User has received"
			return retMap
		}
	}

	avaNum , _:=strconv.Atoi(ret["AvailableNum"])
	giftNum,_ :=strconv.Atoi(ret["GiftNum"])
	if avaNum+1 <= giftNum {
		flagCondition = true
		outString := "{ " + userName + " "+ time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")+"}"
		avanumS := strconv.Itoa(avaNum+1)
		ret["AvailableNum"] = avanumS
		ret["ClaimList"] = string(ret["ClaimList"]) +" " + outString + ";"

		err := setUpRedis.HashSetMap(ret)
		if err != nil {
			fmt.Printf("%s",err)
		}
	}else {
		errString = ""
		errString = "Insufficient quantity"
		flagCondition = false
	}
	if flagCondition {
		retMap["condition"]="success"
		retMap["GiftContent"] = ret["GiftContent"]
	}else {

		retMap["condition"]   = "error"
		retMap["GiftContent"] = errString
	}
	return retMap
}