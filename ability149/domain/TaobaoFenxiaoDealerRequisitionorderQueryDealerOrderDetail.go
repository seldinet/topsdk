package domain


type TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail struct {
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
    Features  *[]TaobaoFenxiaoDealerRequisitionorderQueryFeature `json:"features,omitempty" `

}

func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetProductId(v int64) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetFinalPrice(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.FinalPrice = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetSkuSpec(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.SkuSpec = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetOriginalPrice(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.OriginalPrice = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetDealerDetailId(v int64) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.DealerDetailId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetSnapshotUrl(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.SnapshotUrl = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetIsDeleted(v bool) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.IsDeleted = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetProductTitle(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.ProductTitle = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetPriceCount(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.PriceCount = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetDealerOrderId(v int64) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.DealerOrderId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetQuantity(v int64) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetSkuId(v int64) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.SkuId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetSkuNumber(v string) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.SkuNumber = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail) SetFeatures(v []TaobaoFenxiaoDealerRequisitionorderQueryFeature) *TaobaoFenxiaoDealerRequisitionorderQueryDealerOrderDetail {
    s.Features = &v
    return s
}
