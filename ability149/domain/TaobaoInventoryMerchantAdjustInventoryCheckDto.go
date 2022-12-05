package domain


type TaobaoInventoryMerchantAdjustInventoryCheckDto struct {
    /*
        1:全量更新   2: 出入库盘盈盘亏     */
    CheckMode  *int64 `json:"check_mode,omitempty" `

    /*
        2： 仓库类型   6：门店类型     */
    InvStoreType  *int64 `json:"inv_store_type,omitempty" `

    /*
        调整明细     */
    DetailList  *[]TaobaoInventoryMerchantAdjustInventoryCheckDetailDto `json:"detail_list,omitempty" `

    /*
        仓库code或者门店id     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        调整单据号     */
    OrderId  *string `json:"order_id,omitempty" `

}

func (s *TaobaoInventoryMerchantAdjustInventoryCheckDto) SetCheckMode(v int64) *TaobaoInventoryMerchantAdjustInventoryCheckDto {
    s.CheckMode = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustInventoryCheckDto) SetInvStoreType(v int64) *TaobaoInventoryMerchantAdjustInventoryCheckDto {
    s.InvStoreType = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustInventoryCheckDto) SetDetailList(v []TaobaoInventoryMerchantAdjustInventoryCheckDetailDto) *TaobaoInventoryMerchantAdjustInventoryCheckDto {
    s.DetailList = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustInventoryCheckDto) SetStoreCode(v string) *TaobaoInventoryMerchantAdjustInventoryCheckDto {
    s.StoreCode = &v
    return s
}
func (s *TaobaoInventoryMerchantAdjustInventoryCheckDto) SetOrderId(v string) *TaobaoInventoryMerchantAdjustInventoryCheckDto {
    s.OrderId = &v
    return s
}
