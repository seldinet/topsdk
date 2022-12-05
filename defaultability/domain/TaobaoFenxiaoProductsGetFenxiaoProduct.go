package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoProductsGetFenxiaoProduct struct {
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
        预扣库存     */
    ReservedQuantity  *int64 `json:"reserved_quantity,omitempty" `

    /*
        配额可用库存     */
    QuotaQuantity  *int64 `json:"quota_quantity,omitempty" `

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
        导入的商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" `

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
    Skus  *[]TaobaoFenxiaoProductsGetSkuList `json:"skus,omitempty" `

    /*
        产品分销商信息     */
    Pdus  *[]TaobaoFenxiaoProductsGetPduList `json:"pdus,omitempty" `

    /*
        查询产品列表时，查询入参增加is_authz:yes|no 
yes:需要授权 
no:不需要授权     */
    IsAuthz  *string `json:"is_authz,omitempty" `

    /*
        关联的后端商品id     */
    ScitemId  *int64 `json:"scitem_id,omitempty" `

    /*
        产品spu id     */
    SpuId  *int64 `json:"spu_id,omitempty" `

}

func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetPid(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Pid = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetName(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetProductcatId(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.ProductcatId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetStandardPrice(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.StandardPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetStandardRetailPrice(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.StandardRetailPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetRetailPriceLow(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.RetailPriceLow = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetRetailPriceHigh(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.RetailPriceHigh = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetCostPrice(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.CostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetDealerCostPrice(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.DealerCostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetOuterId(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.OuterId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetQuantity(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetReservedQuantity(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.ReservedQuantity = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetQuotaQuantity(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.QuotaQuantity = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetPictures(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Pictures = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetDescPath(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.DescPath = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetCategoryId(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.CategoryId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetDescription(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Description = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetProperties(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Properties = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetPropertyAlias(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.PropertyAlias = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetInputProperties(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.InputProperties = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetProv(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Prov = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetCity(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.City = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetPostageType(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.PostageType = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetPostageId(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.PostageId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetPostageOrdinary(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.PostageOrdinary = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetPostageFast(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.PostageFast = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetPostageEms(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.PostageEms = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetDiscountId(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.DiscountId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetItemId(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.ItemId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetHaveInvoice(v bool) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.HaveInvoice = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetHaveQuarantee(v bool) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.HaveQuarantee = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetItemsCount(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.ItemsCount = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetOrdersCount(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.OrdersCount = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetCreated(v util.LocalTime) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Created = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetModified(v util.LocalTime) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Modified = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetUpshelfTime(v util.LocalTime) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.UpshelfTime = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetStatus(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Status = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetTradeType(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.TradeType = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetSkus(v []TaobaoFenxiaoProductsGetSkuList) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Skus = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetPdus(v []TaobaoFenxiaoProductsGetPduList) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.Pdus = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetIsAuthz(v string) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.IsAuthz = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetScitemId(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.ScitemId = &v
    return s
}
func (s *TaobaoFenxiaoProductsGetFenxiaoProduct) SetSpuId(v int64) *TaobaoFenxiaoProductsGetFenxiaoProduct {
    s.SpuId = &v
    return s
}
