package domain


type TaobaoInventoryMerchantAdjustInventoryCheckDetailDto struct {
    /*
        如果是门店类型,则为必填。 ONLINE_INVENTORY  线上可售库存，  SHARE_INVENTORY 线下可售库存     */
    InvBizCode  *string `json:"inv_biz_code,omitempty" `

    /*
        调整数量，正数盘盈，负数盘亏     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        若为货品仓库存，则此处是货品ID 若为商品直接设置仓库存，则此处是商品ID， 若商品带SKU，还需要补充skuId     */
    ScItemId  *int64 `json:"sc_item_id,omitempty" `

    /*
        每个货品的调整子单据号，作为业务调整依据，处理时会根据此单据号作幂     */
    SubOrderId  *string `json:"sub_order_id,omitempty" `

    /*
        调整商品对应的SKUID，如果商品为货品，则为0     */
    SkuId  *int64 `json:"sku_id,omitempty" `

}

func (s *TaobaoInventoryMerchantAdjustInventoryCheckDetailDto) SetInvBizCode(v string) *TaobaoInventoryMerchantAdjustInventoryCheckDetailDto {
    s.InvBizCode = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustInventoryCheckDetailDto) SetQuantity(v int64) *TaobaoInventoryMerchantAdjustInventoryCheckDetailDto {
    s.Quantity = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustInventoryCheckDetailDto) SetScItemId(v int64) *TaobaoInventoryMerchantAdjustInventoryCheckDetailDto {
    s.ScItemId = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustInventoryCheckDetailDto) SetSubOrderId(v string) *TaobaoInventoryMerchantAdjustInventoryCheckDetailDto {
    s.SubOrderId = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustInventoryCheckDetailDto) SetSkuId(v int64) *TaobaoInventoryMerchantAdjustInventoryCheckDetailDto {
    s.SkuId = &v
    return s
}
