package request


type TaobaoItemSkuPriceUpdateRequest struct {
    /*
        Sku所属商品数字id，可通过 taobao.item.get 获取     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充     */
    Properties  *string `json:"properties" required:"true" `
    /*
        Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)，sku数量变化后item的总数量也会随着变化。取值范围:大于等于零的整数     */
    Quantity  *int64 `json:"quantity,omitempty" required:"false" `
    /*
        Sku的销售价格。精确到2位小数;单位:元。如:200.07，表示:200元7分。修改后的sku价格要保证商品的价格在所有sku价格所形成的价格区间内（例如：商品价格为6元，sku价格有5元、10元两种，如果要修改5元sku的价格，那么修改的范围只能是0-6元之间；如果要修改10元的sku，那么修改的范围只能是6到无穷大的区间中）     */
    Price  *string `json:"price,omitempty" required:"false" `
    /*
        Sku的商家外部id     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN defalutValue��zh_CN    */
    Lang  *string `json:"lang,omitempty" required:"false" `
    /*
        sku所属商品的价格。当用户更新sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够更新成功     */
    ItemPrice  *string `json:"item_price,omitempty" required:"false" `
    /*
        忽略警告提示.     */
    Ignorewarning  *string `json:"ignorewarning,omitempty" required:"false" `
}

func (s *TaobaoItemSkuPriceUpdateRequest) SetNumIid(v int64) *TaobaoItemSkuPriceUpdateRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemSkuPriceUpdateRequest) SetProperties(v string) *TaobaoItemSkuPriceUpdateRequest {
    s.Properties = &v
    return s
}
func (s *TaobaoItemSkuPriceUpdateRequest) SetQuantity(v int64) *TaobaoItemSkuPriceUpdateRequest {
    s.Quantity = &v
    return s
}
func (s *TaobaoItemSkuPriceUpdateRequest) SetPrice(v string) *TaobaoItemSkuPriceUpdateRequest {
    s.Price = &v
    return s
}
func (s *TaobaoItemSkuPriceUpdateRequest) SetOuterId(v string) *TaobaoItemSkuPriceUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoItemSkuPriceUpdateRequest) SetLang(v string) *TaobaoItemSkuPriceUpdateRequest {
    s.Lang = &v
    return s
}
func (s *TaobaoItemSkuPriceUpdateRequest) SetItemPrice(v string) *TaobaoItemSkuPriceUpdateRequest {
    s.ItemPrice = &v
    return s
}
func (s *TaobaoItemSkuPriceUpdateRequest) SetIgnorewarning(v string) *TaobaoItemSkuPriceUpdateRequest {
    s.Ignorewarning = &v
    return s
}

func (req *TaobaoItemSkuPriceUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    if(req.Properties != nil) {
        paramMap["properties"] = *req.Properties
    }
    if(req.Quantity != nil) {
        paramMap["quantity"] = *req.Quantity
    }
    if(req.Price != nil) {
        paramMap["price"] = *req.Price
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Lang != nil) {
        paramMap["lang"] = *req.Lang
    }
    if(req.ItemPrice != nil) {
        paramMap["item_price"] = *req.ItemPrice
    }
    if(req.Ignorewarning != nil) {
        paramMap["ignorewarning"] = *req.Ignorewarning
    }
    return paramMap
}

func (req *TaobaoItemSkuPriceUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}