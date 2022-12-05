package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoDistributorProductsGetFenxiaoProduct struct {
    /*
        产品ID     */
    Pid  *int64 `json:"pid,omitempty" `

    /*
        产品名称     */
    Name  *string `json:"name,omitempty" `

    /*
        产品线ID     */
    ProductcatId  *int64 `json:"productcat_id,omitempty" `

    /*
        采购基准价，单位：元。     */
    StandardPrice  *string `json:"standard_price,omitempty" `

    /*
        零售基准价，单位：元     */
    StandardRetailPrice  *string `json:"standard_retail_price,omitempty" `

    /*
        最低零售价，单位：分。     */
    RetailPriceLow  *string `json:"retail_price_low,omitempty" `

    /*
        最高零售价，单位：分。     */
    RetailPriceHigh  *string `json:"retail_price_high,omitempty" `

    /*
        采购价格，单位：元。     */
    CostPrice  *string `json:"cost_price,omitempty" `

    /*
        经销采购价     */
    DealerCostPrice  *string `json:"dealer_cost_price,omitempty" `

    /*
        商家编码     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        产品库存     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        产品图片路径     */
    Pictures  *string `json:"pictures,omitempty" `

    /*
        产品描述路径，通过http请求获取     */
    DescPath  *string `json:"desc_path,omitempty" `

    /*
        类目id     */
    CategoryId  *string `json:"category_id,omitempty" `

    /*
        产品描述的内容     */
    Description  *string `json:"description,omitempty" `

    /*
        产品属性，格式为pid:vid;pid:vid     */
    Properties  *string `json:"properties,omitempty" `

    /*
        属性别名，格式为pid:vid:alias;pid:vid:alias     */
    PropertyAlias  *string `json:"property_alias,omitempty" `

    /*
        自定义属性，格式为pid:value;pid:value     */
    InputProperties  *string `json:"input_properties,omitempty" `

    /*
        所在地：省     */
    Prov  *string `json:"prov,omitempty" `

    /*
        所在地：市     */
    City  *string `json:"city,omitempty" `

    /*
        运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）     */
    PostageType  *string `json:"postage_type,omitempty" `

    /*
        运费模板ID     */
    PostageId  *int64 `json:"postage_id,omitempty" `

    /*
        平邮费用，单位：元     */
    PostageOrdinary  *string `json:"postage_ordinary,omitempty" `

    /*
        快递费用，单位：元     */
    PostageFast  *string `json:"postage_fast,omitempty" `

    /*
        ems费用，单位：元     */
    PostageEms  *string `json:"postage_ems,omitempty" `

    /*
        折扣ID（新增参数）     */
    DiscountId  *int64 `json:"discount_id,omitempty" `

    /*
        根据商品ID查询时，返回这个产品对应的商品ID，只对分销商查询接口有效     */
    QueryItemId  *int64 `json:"query_item_id,omitempty" `

    /*
        是否有发票，可选值：false（否）、true（是）     */
    HaveInvoice  *bool `json:"have_invoice,omitempty" `

    /*
        是否有保修，可选值：false（否）、true（是）     */
    HaveQuarantee  *bool `json:"have_quarantee,omitempty" `

    /*
        下载人数     */
    ItemsCount  *int64 `json:"items_count,omitempty" `

    /*
        累计采购次数     */
    OrdersCount  *int64 `json:"orders_count,omitempty" `

    /*
        创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        更新时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        铺货时间     */
    UpshelfTime  *util.LocalTime `json:"upshelf_time,omitempty" `

    /*
        发布状态，可选值：up（上架）、down（下架）     */
    Status  *string `json:"status,omitempty" `

    /*
        分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）     */
    TradeType  *string `json:"trade_type,omitempty" `

    /*
        sku列表     */
    Skus  *[]TaobaoFenxiaoDistributorProductsGetSkuList `json:"skus,omitempty" `

    /*
        产品图片     */
    Images  *[]TaobaoFenxiaoDistributorProductsGetProductImageList `json:"images,omitempty" `

    /*
        产品分销商信息     */
    Pdus  *[]TaobaoFenxiaoDistributorProductsGetPduList `json:"pdus,omitempty" `

    /*
        查询产品列表时，查询入参增加is_authz:yes|no yes:需要授权 no:不需要授权     */
    IsAuthz  *string `json:"is_authz,omitempty" `

    /*
        关联的后端商品id     */
    ScitemId  *int64 `json:"scitem_id,omitempty" `

    /*
        产品spu id     */
    SpuId  *int64 `json:"spu_id,omitempty" `

}

func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetPid(v int64) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Pid = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetName(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetProductcatId(v int64) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.ProductcatId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetStandardPrice(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.StandardPrice = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetStandardRetailPrice(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.StandardRetailPrice = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetRetailPriceLow(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.RetailPriceLow = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetRetailPriceHigh(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.RetailPriceHigh = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetCostPrice(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.CostPrice = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetDealerCostPrice(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.DealerCostPrice = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetOuterId(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.OuterId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetQuantity(v int64) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetPictures(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Pictures = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetDescPath(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.DescPath = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetCategoryId(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.CategoryId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetDescription(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Description = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetProperties(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Properties = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetPropertyAlias(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.PropertyAlias = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetInputProperties(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.InputProperties = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetProv(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Prov = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetCity(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.City = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetPostageType(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.PostageType = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetPostageId(v int64) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.PostageId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetPostageOrdinary(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.PostageOrdinary = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetPostageFast(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.PostageFast = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetPostageEms(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.PostageEms = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetDiscountId(v int64) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.DiscountId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetQueryItemId(v int64) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.QueryItemId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetHaveInvoice(v bool) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.HaveInvoice = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetHaveQuarantee(v bool) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.HaveQuarantee = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetItemsCount(v int64) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.ItemsCount = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetOrdersCount(v int64) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.OrdersCount = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetCreated(v util.LocalTime) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Created = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetModified(v util.LocalTime) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Modified = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetUpshelfTime(v util.LocalTime) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.UpshelfTime = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetStatus(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Status = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetTradeType(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.TradeType = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetSkus(v []TaobaoFenxiaoDistributorProductsGetSkuList) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Skus = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetImages(v []TaobaoFenxiaoDistributorProductsGetProductImageList) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Images = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetPdus(v []TaobaoFenxiaoDistributorProductsGetPduList) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.Pdus = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetIsAuthz(v string) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.IsAuthz = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetScitemId(v int64) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.ScitemId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct) SetSpuId(v int64) *TaobaoFenxiaoDistributorProductsGetFenxiaoProduct {
    s.SpuId = &v
    return s
}
