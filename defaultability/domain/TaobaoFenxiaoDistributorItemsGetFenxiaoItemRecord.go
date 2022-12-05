package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord struct {
    /*
        分销商ID     */
    DistributorId  *int64 `json:"distributor_id,omitempty" `

    /*
        商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        产品ID     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        下载时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        更新时间（系统时间，无业务意义）     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）     */
    TradeType  *string `json:"trade_type,omitempty" `

}

func (s *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord) SetDistributorId(v int64) *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord {
    s.DistributorId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord) SetItemId(v int64) *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord {
    s.ItemId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord) SetProductId(v int64) *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord) SetCreated(v util.LocalTime) *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord {
    s.Created = &v
    return s
}
func (s *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord) SetModified(v util.LocalTime) *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord {
    s.Modified = &v
    return s
}
func (s *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord) SetTradeType(v string) *TaobaoFenxiaoDistributorItemsGetFenxiaoItemRecord {
    s.TradeType = &v
    return s
}
