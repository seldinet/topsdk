package domain


type TmallChannelTradeOrderCreateTopChannelSubPurchaseOrderCreateParam struct {
    /*
        采购数量     */
    BuyQuantity  *int64 `json:"buy_quantity,omitempty" `

    /*
        采购货品SKU ID     */
    ProductSkuId  *int64 `json:"product_sku_id,omitempty" `

    /*
        采购货品ID     */
    ProductId  *int64 `json:"product_id,omitempty" `

}

func (s *TmallChannelTradeOrderCreateTopChannelSubPurchaseOrderCreateParam) SetBuyQuantity(v int64) *TmallChannelTradeOrderCreateTopChannelSubPurchaseOrderCreateParam {
    s.BuyQuantity = &v
    return s
}
func (s *TmallChannelTradeOrderCreateTopChannelSubPurchaseOrderCreateParam) SetProductSkuId(v int64) *TmallChannelTradeOrderCreateTopChannelSubPurchaseOrderCreateParam {
    s.ProductSkuId = &v
    return s
}
func (s *TmallChannelTradeOrderCreateTopChannelSubPurchaseOrderCreateParam) SetProductId(v int64) *TmallChannelTradeOrderCreateTopChannelSubPurchaseOrderCreateParam {
    s.ProductId = &v
    return s
}
