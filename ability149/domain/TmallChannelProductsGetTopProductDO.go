package domain

import (
        "topsdk/util"
    )

type TmallChannelProductsGetTopProductDO struct {
    /*
        所在地：市     */
    City  *string `json:"city,omitempty" `

    /*
        代销采购价格，单位：元。     */
    CostPriceYuan  *string `json:"cost_price_yuan,omitempty" `

    /*
        产品描述路径，通过http请求获取     */
    DescPath  *string `json:"desc_path,omitempty" `

    /*
        创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        更新时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        是否有发票，可选值：false（否）、true（是）     */
    HaveInvoice  *int64 `json:"have_invoice,omitempty" `

    /*
        是否有保修，可选值：false（否）、true（是）     */
    HaveQuarantee  *int64 `json:"have_quarantee,omitempty" `

    /*
        关联的前台宝贝id     */
    AuctionId  *int64 `json:"auction_id,omitempty" `

    /*
        下载人数     */
    ItemsCount  *int64 `json:"items_count,omitempty" `

    /*
        ems费用，单位：元     */
    PostageEms  *string `json:"postage_ems,omitempty" `

    /*
        快递费用，单位：元     */
    PostageFast  *string `json:"postage_fast,omitempty" `

    /*
        运费模板ID     */
    PostageId  *int64 `json:"postage_id,omitempty" `

    /*
        平邮费用，单位：元     */
    PostageOrdinary  *string `json:"postage_ordinary,omitempty" `

    /*
        运费类型：1（供应商承担运费）、2（分销商承担运费）可选值：seller（供应商承担运费）、buyer（分销商承担运费）     */
    PostageType  *int64 `json:"postage_type,omitempty" `

    /*
        产品ID     */
    Pid  *int64 `json:"pid,omitempty" `

    /*
        产品所属产品线id     */
    ProductLineId  *int64 `json:"product_line_id,omitempty" `

    /*
        商家编码     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        所在地：省     */
    Prov  *string `json:"prov,omitempty" `

    /*
        产品库存     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        最高零售价，单位：元。     */
    RetailPriceHigh  *string `json:"retail_price_high,omitempty" `

    /*
        最低零售价，单位：元。     */
    RetailPriceLow  *string `json:"retail_price_low,omitempty" `

    /*
        scItemId     */
    ScItemId  *int64 `json:"sc_item_id,omitempty" `

    /*
        分销产品SKU列表     */
    Skus  *[]TmallChannelProductsGetProductSkuDo `json:"skus,omitempty" `

    /*
        spuId     */
    SpuId  *int64 `json:"spu_id,omitempty" `

    /*
        市场价：单位元     */
    StandardPrice  *string `json:"standard_price,omitempty" `

    /*
        发布状态： 1 上架，2 下架，3 删除     */
    Status  *int64 `json:"status,omitempty" `

    /*
        产品名称     */
    Name  *string `json:"name,omitempty" `

    /*
        累计采购次数     */
    OrdersCount  *int64 `json:"orders_count,omitempty" `

    /*
        产品图片路径列表，用“,”分隔     */
    Pictures  *string `json:"pictures,omitempty" `

}

func (s *TmallChannelProductsGetTopProductDO) SetCity(v string) *TmallChannelProductsGetTopProductDO {
    s.City = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetCostPriceYuan(v string) *TmallChannelProductsGetTopProductDO {
    s.CostPriceYuan = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetDescPath(v string) *TmallChannelProductsGetTopProductDO {
    s.DescPath = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetCreated(v util.LocalTime) *TmallChannelProductsGetTopProductDO {
    s.Created = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetModified(v util.LocalTime) *TmallChannelProductsGetTopProductDO {
    s.Modified = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetHaveInvoice(v int64) *TmallChannelProductsGetTopProductDO {
    s.HaveInvoice = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetHaveQuarantee(v int64) *TmallChannelProductsGetTopProductDO {
    s.HaveQuarantee = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetAuctionId(v int64) *TmallChannelProductsGetTopProductDO {
    s.AuctionId = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetItemsCount(v int64) *TmallChannelProductsGetTopProductDO {
    s.ItemsCount = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetPostageEms(v string) *TmallChannelProductsGetTopProductDO {
    s.PostageEms = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetPostageFast(v string) *TmallChannelProductsGetTopProductDO {
    s.PostageFast = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetPostageId(v int64) *TmallChannelProductsGetTopProductDO {
    s.PostageId = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetPostageOrdinary(v string) *TmallChannelProductsGetTopProductDO {
    s.PostageOrdinary = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetPostageType(v int64) *TmallChannelProductsGetTopProductDO {
    s.PostageType = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetPid(v int64) *TmallChannelProductsGetTopProductDO {
    s.Pid = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetProductLineId(v int64) *TmallChannelProductsGetTopProductDO {
    s.ProductLineId = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetOuterId(v string) *TmallChannelProductsGetTopProductDO {
    s.OuterId = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetProv(v string) *TmallChannelProductsGetTopProductDO {
    s.Prov = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetQuantity(v int64) *TmallChannelProductsGetTopProductDO {
    s.Quantity = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetRetailPriceHigh(v string) *TmallChannelProductsGetTopProductDO {
    s.RetailPriceHigh = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetRetailPriceLow(v string) *TmallChannelProductsGetTopProductDO {
    s.RetailPriceLow = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetScItemId(v int64) *TmallChannelProductsGetTopProductDO {
    s.ScItemId = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetSkus(v []TmallChannelProductsGetProductSkuDo) *TmallChannelProductsGetTopProductDO {
    s.Skus = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetSpuId(v int64) *TmallChannelProductsGetTopProductDO {
    s.SpuId = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetStandardPrice(v string) *TmallChannelProductsGetTopProductDO {
    s.StandardPrice = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetStatus(v int64) *TmallChannelProductsGetTopProductDO {
    s.Status = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetName(v string) *TmallChannelProductsGetTopProductDO {
    s.Name = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetOrdersCount(v int64) *TmallChannelProductsGetTopProductDO {
    s.OrdersCount = &v
    return s
}
func (s *TmallChannelProductsGetTopProductDO) SetPictures(v string) *TmallChannelProductsGetTopProductDO {
    s.Pictures = &v
    return s
}
