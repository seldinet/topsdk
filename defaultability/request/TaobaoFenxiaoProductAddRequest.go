package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoProductAddRequest struct {
    /*
        所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。     */
    CategoryId  *int64 `json:"category_id" required:"true" `
    /*
        产品名称，长度不超过60个字节。     */
    Name  *string `json:"name" required:"true" `
    /*
        产品线ID     */
    ProductcatId  *int64 `json:"productcat_id" required:"true" `
    /*
        采购基准价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。     */
    StandardPrice  *string `json:"standard_price" required:"true" `
    /*
        零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。     */
    StandardRetailPrice  *string `json:"standard_retail_price,omitempty" required:"false" `
    /*
        最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。     */
    RetailPriceLow  *string `json:"retail_price_low" required:"true" `
    /*
        最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。     */
    RetailPriceHigh  *string `json:"retail_price_high" required:"true" `
    /*
        代销采购价格，单位：元。例：“10.56”。必须在0.01元到10000000元之间。     */
    CostPrice  *string `json:"cost_price,omitempty" required:"false" `
    /*
        经销采购价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。     */
    DealerCostPrice  *string `json:"dealer_cost_price,omitempty" required:"false" `
    /*
        商家编码，长度不能超过60个字节。     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        产品库存必须是1到999999。     */
    Quantity  *int64 `json:"quantity" required:"true" `
    /*
        产品描述，长度为5到25000字符。     */
    Desc  *string `json:"desc" required:"true" `
    /*
        所在地：省，例：“浙江”     */
    Prov  *string `json:"prov" required:"true" `
    /*
        所在地：市，例：“杭州”     */
    City  *string `json:"city" required:"true" `
    /*
        运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）,默认seller。     */
    PostageType  *string `json:"postage_type,omitempty" required:"false" `
    /*
        运费模板ID，参考taobao.postages.get。     */
    PostageId  *int64 `json:"postage_id,omitempty" required:"false" `
    /*
        平邮费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。     */
    PostageOrdinary  *string `json:"postage_ordinary,omitempty" required:"false" `
    /*
        快递费用，单位：元。例：“10.56”。 大小为0.01元到999999元之间。     */
    PostageFast  *string `json:"postage_fast,omitempty" required:"false" `
    /*
        ems费用，单位：元。例：“10.56”。 大小为0.00元到999999元之间。     */
    PostageEms  *string `json:"postage_ems,omitempty" required:"false" `
    /*
        是否有发票，可选值：false（否）、true（是），默认false。     */
    HaveInvoice  *string `json:"have_invoice,omitempty" required:"false" `
    /*
        是否有保修，可选值：false（否）、true（是），默认false。     */
    HaveQuarantee  *string `json:"have_quarantee,omitempty" required:"false" `
    /*
        折扣ID     */
    DiscountId  *int64 `json:"discount_id,omitempty" required:"false" `
    /*
        分销方式：AGENT（只做代销，默认值）、DEALER（只做经销）、ALL（代销和经销都做）     */
    TradeType  *string `json:"trade_type,omitempty" required:"false" `
    /*
        添加产品时，添加入参isAuthz:yes|no 
yes:需要授权 
no:不需要授权 
默认是需要授权 defalutValue��yes    */
    IsAuthz  *string `json:"is_authz,omitempty" required:"false" `
    /*
        产品主图图片空间相对路径或绝对路径     */
    PicPath  *string `json:"pic_path,omitempty" required:"false" `
    /*
        产品主图，大小不超过500k，格式为gif,jpg,jpeg,png,bmp等图片     */
    Image  *[]byte `json:"image,omitempty" required:"false" `
    /*
        产品属性，格式为pid:vid;pid:vid     */
    Properties  *string `json:"properties,omitempty" required:"false" `
    /*
        属性别名，格式为：pid:vid:alias;pid:vid:alias（alias为别名）     */
    PropertyAlias  *string `json:"property_alias,omitempty" required:"false" `
    /*
        自定义属性。格式为pid:value;pid:value     */
    InputProperties  *string `json:"input_properties,omitempty" required:"false" `
    /*
        sku的采购基准价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序     */
    SkuStandardPrices  *string `json:"sku_standard_prices,omitempty" required:"false" `
    /*
        sku的采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序     */
    SkuCostPrices  *string `json:"sku_cost_prices,omitempty" required:"false" `
    /*
        sku的商家编码。如果多个，用逗号分隔，并与其他sku信息保持相同顺序     */
    SkuOuterIds  *string `json:"sku_outer_ids,omitempty" required:"false" `
    /*
        sku的库存。如果多个，用逗号分隔，并与其他sku信息保持相同顺序     */
    SkuQuantitys  *string `json:"sku_quantitys,omitempty" required:"false" `
    /*
        sku的属性。如果多个，用逗号分隔，并与其他sku信息保持相同顺序     */
    SkuProperties  *string `json:"sku_properties,omitempty" required:"false" `
    /*
        sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。     */
    SkuDealerCostPrices  *string `json:"sku_dealer_cost_prices,omitempty" required:"false" `
    /*
        导入的商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" required:"false" `
    /*
        产品spuID，达尔文产品必须要传spuID，否则不能发布。其他非达尔文产品，看情况传     */
    SpuId  *int64 `json:"spu_id,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoProductAddRequest) SetCategoryId(v int64) *TaobaoFenxiaoProductAddRequest {
    s.CategoryId = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetName(v string) *TaobaoFenxiaoProductAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetProductcatId(v int64) *TaobaoFenxiaoProductAddRequest {
    s.ProductcatId = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetStandardPrice(v string) *TaobaoFenxiaoProductAddRequest {
    s.StandardPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetStandardRetailPrice(v string) *TaobaoFenxiaoProductAddRequest {
    s.StandardRetailPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetRetailPriceLow(v string) *TaobaoFenxiaoProductAddRequest {
    s.RetailPriceLow = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetRetailPriceHigh(v string) *TaobaoFenxiaoProductAddRequest {
    s.RetailPriceHigh = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetCostPrice(v string) *TaobaoFenxiaoProductAddRequest {
    s.CostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetDealerCostPrice(v string) *TaobaoFenxiaoProductAddRequest {
    s.DealerCostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetOuterId(v string) *TaobaoFenxiaoProductAddRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetQuantity(v int64) *TaobaoFenxiaoProductAddRequest {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetDesc(v string) *TaobaoFenxiaoProductAddRequest {
    s.Desc = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetProv(v string) *TaobaoFenxiaoProductAddRequest {
    s.Prov = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetCity(v string) *TaobaoFenxiaoProductAddRequest {
    s.City = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetPostageType(v string) *TaobaoFenxiaoProductAddRequest {
    s.PostageType = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetPostageId(v int64) *TaobaoFenxiaoProductAddRequest {
    s.PostageId = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetPostageOrdinary(v string) *TaobaoFenxiaoProductAddRequest {
    s.PostageOrdinary = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetPostageFast(v string) *TaobaoFenxiaoProductAddRequest {
    s.PostageFast = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetPostageEms(v string) *TaobaoFenxiaoProductAddRequest {
    s.PostageEms = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetHaveInvoice(v string) *TaobaoFenxiaoProductAddRequest {
    s.HaveInvoice = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetHaveQuarantee(v string) *TaobaoFenxiaoProductAddRequest {
    s.HaveQuarantee = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetDiscountId(v int64) *TaobaoFenxiaoProductAddRequest {
    s.DiscountId = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetTradeType(v string) *TaobaoFenxiaoProductAddRequest {
    s.TradeType = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetIsAuthz(v string) *TaobaoFenxiaoProductAddRequest {
    s.IsAuthz = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetPicPath(v string) *TaobaoFenxiaoProductAddRequest {
    s.PicPath = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetImage(v []byte) *TaobaoFenxiaoProductAddRequest {
    s.Image = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetProperties(v string) *TaobaoFenxiaoProductAddRequest {
    s.Properties = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetPropertyAlias(v string) *TaobaoFenxiaoProductAddRequest {
    s.PropertyAlias = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetInputProperties(v string) *TaobaoFenxiaoProductAddRequest {
    s.InputProperties = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetSkuStandardPrices(v string) *TaobaoFenxiaoProductAddRequest {
    s.SkuStandardPrices = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetSkuCostPrices(v string) *TaobaoFenxiaoProductAddRequest {
    s.SkuCostPrices = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetSkuOuterIds(v string) *TaobaoFenxiaoProductAddRequest {
    s.SkuOuterIds = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetSkuQuantitys(v string) *TaobaoFenxiaoProductAddRequest {
    s.SkuQuantitys = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetSkuProperties(v string) *TaobaoFenxiaoProductAddRequest {
    s.SkuProperties = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetSkuDealerCostPrices(v string) *TaobaoFenxiaoProductAddRequest {
    s.SkuDealerCostPrices = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetItemId(v int64) *TaobaoFenxiaoProductAddRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoFenxiaoProductAddRequest) SetSpuId(v int64) *TaobaoFenxiaoProductAddRequest {
    s.SpuId = &v
    return s
}

func (req *TaobaoFenxiaoProductAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.ProductcatId != nil) {
        paramMap["productcat_id"] = *req.ProductcatId
    }
    if(req.StandardPrice != nil) {
        paramMap["standard_price"] = *req.StandardPrice
    }
    if(req.StandardRetailPrice != nil) {
        paramMap["standard_retail_price"] = *req.StandardRetailPrice
    }
    if(req.RetailPriceLow != nil) {
        paramMap["retail_price_low"] = *req.RetailPriceLow
    }
    if(req.RetailPriceHigh != nil) {
        paramMap["retail_price_high"] = *req.RetailPriceHigh
    }
    if(req.CostPrice != nil) {
        paramMap["cost_price"] = *req.CostPrice
    }
    if(req.DealerCostPrice != nil) {
        paramMap["dealer_cost_price"] = *req.DealerCostPrice
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Quantity != nil) {
        paramMap["quantity"] = *req.Quantity
    }
    if(req.Desc != nil) {
        paramMap["desc"] = *req.Desc
    }
    if(req.Prov != nil) {
        paramMap["prov"] = *req.Prov
    }
    if(req.City != nil) {
        paramMap["city"] = *req.City
    }
    if(req.PostageType != nil) {
        paramMap["postage_type"] = *req.PostageType
    }
    if(req.PostageId != nil) {
        paramMap["postage_id"] = *req.PostageId
    }
    if(req.PostageOrdinary != nil) {
        paramMap["postage_ordinary"] = *req.PostageOrdinary
    }
    if(req.PostageFast != nil) {
        paramMap["postage_fast"] = *req.PostageFast
    }
    if(req.PostageEms != nil) {
        paramMap["postage_ems"] = *req.PostageEms
    }
    if(req.HaveInvoice != nil) {
        paramMap["have_invoice"] = *req.HaveInvoice
    }
    if(req.HaveQuarantee != nil) {
        paramMap["have_quarantee"] = *req.HaveQuarantee
    }
    if(req.DiscountId != nil) {
        paramMap["discount_id"] = *req.DiscountId
    }
    if(req.TradeType != nil) {
        paramMap["trade_type"] = *req.TradeType
    }
    if(req.IsAuthz != nil) {
        paramMap["is_authz"] = *req.IsAuthz
    }
    if(req.PicPath != nil) {
        paramMap["pic_path"] = *req.PicPath
    }
    if(req.Properties != nil) {
        paramMap["properties"] = *req.Properties
    }
    if(req.PropertyAlias != nil) {
        paramMap["property_alias"] = *req.PropertyAlias
    }
    if(req.InputProperties != nil) {
        paramMap["input_properties"] = *req.InputProperties
    }
    if(req.SkuStandardPrices != nil) {
        paramMap["sku_standard_prices"] = *req.SkuStandardPrices
    }
    if(req.SkuCostPrices != nil) {
        paramMap["sku_cost_prices"] = *req.SkuCostPrices
    }
    if(req.SkuOuterIds != nil) {
        paramMap["sku_outer_ids"] = *req.SkuOuterIds
    }
    if(req.SkuQuantitys != nil) {
        paramMap["sku_quantitys"] = *req.SkuQuantitys
    }
    if(req.SkuProperties != nil) {
        paramMap["sku_properties"] = *req.SkuProperties
    }
    if(req.SkuDealerCostPrices != nil) {
        paramMap["sku_dealer_cost_prices"] = *req.SkuDealerCostPrices
    }
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.SpuId != nil) {
        paramMap["spu_id"] = *req.SpuId
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}