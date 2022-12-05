package domain


type TaobaoTradesSoldGetServiceOrder struct {
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
        卖家昵称     */
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

}

func (s *TaobaoTradesSoldGetServiceOrder) SetOid(v int64) *TaobaoTradesSoldGetServiceOrder {
    s.Oid = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetItemOid(v int64) *TaobaoTradesSoldGetServiceOrder {
    s.ItemOid = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetServiceId(v int64) *TaobaoTradesSoldGetServiceOrder {
    s.ServiceId = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetServiceDetailUrl(v string) *TaobaoTradesSoldGetServiceOrder {
    s.ServiceDetailUrl = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetNum(v int64) *TaobaoTradesSoldGetServiceOrder {
    s.Num = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetPrice(v string) *TaobaoTradesSoldGetServiceOrder {
    s.Price = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetPayment(v string) *TaobaoTradesSoldGetServiceOrder {
    s.Payment = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetTitle(v string) *TaobaoTradesSoldGetServiceOrder {
    s.Title = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetTotalFee(v string) *TaobaoTradesSoldGetServiceOrder {
    s.TotalFee = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetBuyerNick(v string) *TaobaoTradesSoldGetServiceOrder {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetRefundId(v string) *TaobaoTradesSoldGetServiceOrder {
    s.RefundId = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetSellerNick(v string) *TaobaoTradesSoldGetServiceOrder {
    s.SellerNick = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetPicPath(v string) *TaobaoTradesSoldGetServiceOrder {
    s.PicPath = &v
    return s
}
func (s *TaobaoTradesSoldGetServiceOrder) SetTmserSpuCode(v string) *TaobaoTradesSoldGetServiceOrder {
    s.TmserSpuCode = &v
    return s
}
