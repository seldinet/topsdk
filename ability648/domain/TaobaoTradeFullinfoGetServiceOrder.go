package domain


type TaobaoTradeFullinfoGetServiceOrder struct {
    /*
        虚拟服务子订单订单号     */
    Oid  *int64 `json:"oid,omitempty" `

    /*
        服务所属的交易订单号。如果服务为一年包换，则item_oid这笔订单享受改服务的保护     */
    ItemOid  *int64 `json:"item_oid,omitempty" `

    /*
        服务数字id     */
    ServiceId  *int64 `json:"service_id,omitempty" `

    /*
        服务详情的URL地址     */
    ServiceDetailUrl  *string `json:"service_detail_url,omitempty" `

    /*
        购买数量，取值范围为大于0的整数     */
    Num  *int64 `json:"num,omitempty" `

    /*
        服务价格，精确到小数点后两位：单位:元     */
    Price  *string `json:"price,omitempty" `

    /*
        子订单实付金额。精确到2位小数，单位:元。如:200.07，表示:200元7分。     */
    Payment  *string `json:"payment,omitempty" `

    /*
        商品名称     */
    Title  *string `json:"title,omitempty" `

    /*
        服务子订单总费用     */
    TotalFee  *string `json:"total_fee,omitempty" `

    /*
        买家昵称     */
    BuyerNick  *string `json:"buyer_nick,omitempty" `

    /*
        最近退款的id     */
    RefundId  *string `json:"refund_id,omitempty" `

    /*
        卖家昵称     */
    SellerNick  *string `json:"seller_nick,omitempty" `

    /*
        服务图片地址     */
    PicPath  *string `json:"pic_path,omitempty" `

    /*
        支持家装类物流的类型     */
    TmserSpuCode  *string `json:"tmser_spu_code,omitempty" `

    /*
        天猫汽车预约服务时间     */
    EtSerTime  *string `json:"et_ser_time,omitempty" `

    /*
        电子凭证预约门店地址     */
    EtShopName  *string `json:"et_shop_name,omitempty" `

    /*
        电子凭证核销门店地址     */
    EtVerifiedShopName  *string `json:"et_verified_shop_name,omitempty" `

    /*
        车牌号     */
    EtPlateNumber  *string `json:"et_plate_number,omitempty" `

    /*
        虚拟服务子订单订单号(String类型)     */
    OidStr  *string `json:"oid_str,omitempty" `

    /*
        appleCareEmail     */
    AppleCareEmail  *string `json:"apple_care_email,omitempty" `

    /*
        appleCareMPN     */
    AppleCareMpn  *string `json:"apple_care_mpn,omitempty" `

    /*
        服务供应链-服务商外部编码标     */
    ServiceOuterId  *string `json:"service_outer_id,omitempty" `

    /*
        服务供应链-服务订单类型,1:主子挂载；2：双主挂载；3：单独售卖     */
    ServiceOrderType  *string `json:"service_order_type,omitempty" `

    /*
        服务订单与实物订单关联关系     */
    ExtServiceBizId  *string `json:"ext_service_biz_id,omitempty" `

}

func (s *TaobaoTradeFullinfoGetServiceOrder) SetOid(v int64) *TaobaoTradeFullinfoGetServiceOrder {
    s.Oid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetItemOid(v int64) *TaobaoTradeFullinfoGetServiceOrder {
    s.ItemOid = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetServiceId(v int64) *TaobaoTradeFullinfoGetServiceOrder {
    s.ServiceId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetServiceDetailUrl(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.ServiceDetailUrl = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetNum(v int64) *TaobaoTradeFullinfoGetServiceOrder {
    s.Num = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetPrice(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.Price = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetPayment(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.Payment = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetTitle(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.Title = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetTotalFee(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetBuyerNick(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetRefundId(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.RefundId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetSellerNick(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetPicPath(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetTmserSpuCode(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.TmserSpuCode = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetEtSerTime(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.EtSerTime = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetEtShopName(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.EtShopName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetEtVerifiedShopName(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.EtVerifiedShopName = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetEtPlateNumber(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.EtPlateNumber = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetOidStr(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.OidStr = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetAppleCareEmail(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.AppleCareEmail = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetAppleCareMpn(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.AppleCareMpn = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetServiceOuterId(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.ServiceOuterId = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetServiceOrderType(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.ServiceOrderType = &v
    return s
}
func (s *TaobaoTradeFullinfoGetServiceOrder) SetExtServiceBizId(v string) *TaobaoTradeFullinfoGetServiceOrder {
    s.ExtServiceBizId = &v
    return s
}
