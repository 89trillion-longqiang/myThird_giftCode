package gift

type Gift struct {
	GiftCode string      ///礼包码
	Description string	 ///礼包描述
	GiftNum string		 ///可领取数
	ValidPeriod string	 ///有效期
	GiftContent string   ///礼包内容
	CreatePer string     ///创建用户
	CreatTime string     ///创建时间

	AvailableNum string  ///已领取次数
	ClaimList string     ///领取列表
}

func (g Gift)IsEmpty() bool{
	if len(g.GiftCode) == 0 {return true}
	if len(g.Description) == 0 {return true}
	if len(g.GiftNum) == 0 {return true}
	if len(g.ValidPeriod) == 0 {return true}
	if len(g.GiftContent) == 0 {return true}
	if len(g.CreatePer) == 0 {return true}
	if len(g.CreatTime) == 0 {return true}
	return false
}