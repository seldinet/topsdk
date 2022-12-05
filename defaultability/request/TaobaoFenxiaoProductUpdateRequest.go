package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoProductUpdateRequest struct {
    /*
        产品ID     */
    Pid  *int64 `json:"pid" required:"true" `
    /*
        产品名称，长度不超过60个字节。     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        采购基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。     */
    StandardPrice  *string `json:"standard_price,omitempty" required:"false" `
    /*
        零售基准价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。     */
    StandardRetailPrice  *string `json:"standard_retail_price,omitempty" required:"false" `
    /*
        最低零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间。     */
    RetailPriceLow  *string `json:"retail_price_low,omitempty" required:"false" `
    /*
        最高零售价，单位：元。例：“10.56”。必须在0.01元到10000000元之间，最高零售价必须大于最低零售价。     */
    RetailPriceHigh  *string `json:"retail_price_high,omitempty" required:"false" `
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
    Quantity  *int64 `json:"quantity,omitempty" required:"false" `
    /*
        产品描述，长度为5到25000字符。     */
    Desc  *string `json:"desc,omitempty" required:"false" `
    /*
        所属类目id，参考Taobao.itemcats.get，不支持成人等类目，输入成人类目id保存提示类目属性错误。     */
    CategoryId  *int64 `json:"category_id,omitempty" required:"false" `
    /*
        产品属性     */
    Properties  *string `json:"properties,omitempty" required:"false" `
    /*
        属性别名     */
    PropertyAlias  *string `json:"property_alias,omitempty" required:"false" `
    /*
        自定义属性。格式为pid:value;pid:value     */
    InputProperties  *string `json:"input_properties,omitempty" required:"false" `
    /*
        产品主图图片空间相对路径或绝对路径     */
    PicPath  *string `json:"pic_path,omitempty" required:"false" `
    /*
        主图图片，如果pic_path参数不空，则优先使用pic_path，忽略该参数     */
    Image  *[]byte `json:"image,omitempty" required:"false" `
    /*
        所在地：省，例：“浙江”     */
    Prov  *string `json:"prov,omitempty" required:"false" `
    /*
        所在地：市，例：“杭州”     */
    City  *string `json:"city,omitempty" required:"false" `
    /*
        运费类型，可选值：seller（供应商承担运费）、buyer（分销商承担运费）。     */
    PostageType  *string `json:"postage_type,omitempty" required:"false" `
    /*
        运费模板ID，参考taobao.postages.get。更新时必须指定运费类型为 buyer，否则不更新。     */
    PostageId  *int64 `json:"postage_id,omitempty" required:"false" `
    /*
        平邮费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。     */
    PostageOrdinary  *string `json:"postage_ordinary,omitempty" required:"false" `
    /*
        快递费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。     */
    PostageFast  *string `json:"postage_fast,omitempty" required:"false" `
    /*
        ems费用，单位：元。例：“10.56”。大小为0.01元到999999元之间。更新时必须指定运费类型为buyer，否则不更新。     */
    PostageEms  *string `json:"postage_ems,omitempty" required:"false" `
    /*
        是否有发票，可选值：false（否）、true（是），默认false。     */
    HaveInvoice  *string `json:"have_invoice,omitempty" required:"false" `
    /*
        是否有保修，可选值：false（否）、true（是），默认false。     */
    HaveQuarantee  *string `json:"have_quarantee,omitempty" required:"false" `
    /*
        发布状态，可选值：up（上架）、down（下架）、delete（删除），输入非法字符则忽略。     */
    Status  *string `json:"status,omitempty" required:"false" `
    /*
        sku id列表，例：1001,1002,1003。如果传入sku_properties将忽略此参数。     */
    SkuIds  *string `json:"sku_ids,omitempty" required:"false" `
    /*
        sku采购基准价，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。     */
    SkuStandardPrices  *string `json:"sku_standard_prices,omitempty" required:"false" `
    /*
        sku采购价格，单位元，例："10.50,11.00,20.50"，字段必须和上面的sku_ids或sku_properties保持一致。     */
    SkuCostPrices  *string `json:"sku_cost_prices,omitempty" required:"false" `
    /*
        sku的经销采购价。如果多个，用逗号分隔，并与其他sku信息保持相同顺序。其中每个值的单位：元。例：“10.56,12.3”。必须在0.01元到10000000元之间。     */
    SkuDealerCostPrices  *string `json:"sku_dealer_cost_prices,omitempty" required:"false" `
    /*
        sku库存，单位元，例："10,20,30"，字段必须和sku_ids或sku_properties保持一致。     */
    SkuQuantitys  *string `json:"sku_quantitys,omitempty" required:"false" `
    /*
        sku商家编码 ，单位元，例："S1000,S1002,S1003"，字段必须和上面的id或sku_properties保持一致，如果没有可以写成",,"     */
    SkuOuterIds  *string `json:"sku_outer_ids,omitempty" required:"false" `
    /*
        sku属性。格式:pid:vid;pid:vid,表示一组属性如:1627207:3232483;1630696:3284570,表示一组:机身颜色:军绿色;手机套餐:一电一充。多组之间用逗号“,”区分。(属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid)
通过此字段可新增和更新sku。若传入此值将忽略sku_ids字段。sku其他字段与此值保持一致。     */
    SkuProperties  *string `json:"sku_properties,omitempty" required:"false" `
    /*
        根据sku属性删除sku信息。需要按组删除属性。     */
    SkuPropertiesDel  *string `json:"sku_properties_del,omitempty" required:"false" `
    /*
        折扣ID     */
    DiscountId  *int64 `json:"discount_id,omitempty" required:"false" `
    /*
        产品是否需要授权isAuthz:yes|no 
yes:需要授权 
no:不需要授权     */
    IsAuthz  *string `json:"is_authz,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoProductUpdateRequest) SetPid(v int64) *TaobaoFenxiaoProductUpdateRequest {
    s.Pid = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetName(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetStandardPrice(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.StandardPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetStandardRetailPrice(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.StandardRetailPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetRetailPriceLow(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.RetailPriceLow = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetRetailPriceHigh(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.RetailPriceHigh = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetCostPrice(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.CostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetDealerCostPrice(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.DealerCostPrice = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetOuterId(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetQuantity(v int64) *TaobaoFenxiaoProductUpdateRequest {
    s.Quantity = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetDesc(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.Desc = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetCategoryId(v int64) *TaobaoFenxiaoProductUpdateRequest {
    s.CategoryId = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetProperties(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.Properties = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetPropertyAlias(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.PropertyAlias = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetInputProperties(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.InputProperties = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetPicPath(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.PicPath = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetImage(v []byte) *TaobaoFenxiaoProductUpdateRequest {
    s.Image = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetProv(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.Prov = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetCity(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.City = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetPostageType(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.PostageType = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetPostageId(v int64) *TaobaoFenxiaoProductUpdateRequest {
    s.PostageId = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetPostageOrdinary(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.PostageOrdinary = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetPostageFast(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.PostageFast = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetPostageEms(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.PostageEms = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetHaveInvoice(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.HaveInvoice = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetHaveQuarantee(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.HaveQuarantee = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetStatus(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.Status = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetSkuIds(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.SkuIds = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetSkuStandardPrices(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.SkuStandardPrices = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetSkuCostPrices(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.SkuCostPrices = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetSkuDealerCostPrices(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.SkuDealerCostPrices = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetSkuQuantitys(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.SkuQuantitys = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetSkuOuterIds(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.SkuOuterIds = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetSkuProperties(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.SkuProperties = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetSkuPropertiesDel(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.SkuPropertiesDel = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetDiscountId(v int64) *TaobaoFenxiaoProductUpdateRequest {
    s.DiscountId = &v
    return s
}
func (s *TaobaoFenxiaoProductUpdateRequest) SetIsAuthz(v string) *TaobaoFenxiaoProductUpdateRequest {
    s.IsAuthz = &v
    return s
}

func (req *TaobaoFenxiaoProductUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Pid != nil) {
        paramMap["pid"] = *req.Pid
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
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
    if(req.CategoryId != nil) {
        paramMap["category_id"] = *req.CategoryId
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
    if(req.PicPath != nil) {
        paramMap["pic_path"] = *req.PicPath
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
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.SkuIds != nil) {
        paramMap["sku_ids"] = *req.SkuIds
    }
    if(req.SkuStandardPrices != nil) {
        paramMap["sku_standard_prices"] = *req.SkuStandardPrices
    }
    if(req.SkuCostPrices != nil) {
        paramMap["sku_cost_prices"] = *req.SkuCostPrices
    }
    if(req.SkuDealerCostPrices != nil) {
        paramMap["sku_dealer_cost_prices"] = *req.SkuDealerCostPrices
    }
    if(req.SkuQuantitys != nil) {
        paramMap["sku_quantitys"] = *req.SkuQuantitys
    }
    if(req.SkuOuterIds != nil) {
        paramMap["sku_outer_ids"] = *req.SkuOuterIds
    }
    if(req.SkuProperties != nil) {
        paramMap["sku_properties"] = *req.SkuProperties
    }
    if(req.SkuPropertiesDel != nil) {
        paramMap["sku_properties_del"] = *req.SkuPropertiesDel
    }
    if(req.DiscountId != nil) {
        paramMap["discount_id"] = *req.DiscountId
    }
    if(req.IsAuthz != nil) {
        paramMap["is_authz"] = *req.IsAuthz
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}