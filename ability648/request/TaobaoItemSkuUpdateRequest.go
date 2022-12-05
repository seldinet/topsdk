package request


type TaobaoItemSkuUpdateRequest struct {
    /*
        Sku所属商品数字id，可通过 taobao.item.get 获取     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充。
如果包含自定义属性，则格式为pid:vid;pid2:vid2;$pText:vText , 其中$pText:vText为自定义属性。限制：其中$pText的’$’前缀不能少，且pText和vText文本中不可以存在 冒号:和分号;以及逗号，     */
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

func (s *TaobaoItemSkuUpdateRequest) SetNumIid(v int64) *TaobaoItemSkuUpdateRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemSkuUpdateRequest) SetProperties(v string) *TaobaoItemSkuUpdateRequest {
    s.Properties = &v
    return s
}
func (s *TaobaoItemSkuUpdateRequest) SetQuantity(v int64) *TaobaoItemSkuUpdateRequest {
    s.Quantity = &v
    return s
}
func (s *TaobaoItemSkuUpdateRequest) SetPrice(v string) *TaobaoItemSkuUpdateRequest {
    s.Price = &v
    return s
}
func (s *TaobaoItemSkuUpdateRequest) SetOuterId(v string) *TaobaoItemSkuUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoItemSkuUpdateRequest) SetLang(v string) *TaobaoItemSkuUpdateRequest {
    s.Lang = &v
    return s
}
func (s *TaobaoItemSkuUpdateRequest) SetItemPrice(v string) *TaobaoItemSkuUpdateRequest {
    s.ItemPrice = &v
    return s
}
func (s *TaobaoItemSkuUpdateRequest) SetIgnorewarning(v string) *TaobaoItemSkuUpdateRequest {
    s.Ignorewarning = &v
    return s
}

func (req *TaobaoItemSkuUpdateRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoItemSkuUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}