package domain


type TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail struct {
    /*
        产品id     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        最终价格(精确到2位小数;单位:元。如:200.07，表示:200元7分 )     */
    FinalPrice  *string `json:"final_price,omitempty" `

    /*
        产品规格     */
    SkuSpec  *string `json:"sku_spec,omitempty" `

    /*
        原始价格(精确到2位小数;单位:元。如:200.07，表示:200元7分 )     */
    OriginalPrice  *string `json:"original_price,omitempty" `

    /*
        经销采购单明细id     */
    DealerDetailId  *int64 `json:"dealer_detail_id,omitempty" `

    /*
        产品快照url     */
    SnapshotUrl  *string `json:"snapshot_url,omitempty" `

    /*
        该条明细是否已删除。true：已删除；false：未删除。     */
    IsDeleted  *bool `json:"is_deleted,omitempty" `

    /*
        产品标题     */
    ProductTitle  *string `json:"product_title,omitempty" `

    /*
        金额小计(采购数量乘以最终价格。精确到2位小数;单位:元。如:200.07，表示:200元7分 )     */
    PriceCount  *string `json:"price_count,omitempty" `

    /*
        经销采购单编号     */
    DealerOrderId  *int64 `json:"dealer_order_id,omitempty" `

    /*
        采购数量     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        sku id     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        商家编码，是sku则为sku的商家编码，否则是产品的商家编码     */
    SkuNumber  *string `json:"sku_number,omitempty" `

    /*
        属性信息列表，key-value形式。如attr_key为storeCode，attr_value则为仓库编码。     */
    Features  *[]TaobaoFenxiaoDealerRequisitionorderGetFeature `json:"features,omitempty" `

}

func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetProductId(v int64) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetFinalPrice(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.FinalPrice = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetSkuSpec(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.SkuSpec = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetOriginalPrice(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.OriginalPrice = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetDealerDetailId(v int64) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.DealerDetailId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetSnapshotUrl(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.SnapshotUrl = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetIsDeleted(v bool) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.IsDeleted = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetProductTitle(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.ProductTitle = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetPriceCount(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.PriceCount = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetDealerOrderId(v int64) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.DealerOrderId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetQuantity(v int64) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetSkuId(v int64) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.SkuId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetSkuNumber(v string) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.SkuNumber = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail) SetFeatures(v []TaobaoFenxiaoDealerRequisitionorderGetFeature) *TaobaoFenxiaoDealerRequisitionorderGetDealerOrderDetail {
    s.Features = &v
    return s
}
