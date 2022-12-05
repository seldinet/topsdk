package request


type TaobaoItemSkuAddRequest struct {
    /*
        Sku所属商品数字id。必选     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        Sku属性串。格式:pid:vid;pid:vid,如:1627207:3232483;1630696:3284570,表示:机身颜色:军绿色;手机套餐:一电一充。     */
    Properties  *string `json:"properties" required:"true" `
    /*
        Sku的库存数量。sku的总数量应该小于等于商品总数量(Item的NUM)。取值范围:大于零的整数     */
    Quantity  *int64 `json:"quantity" required:"true" `
    /*
        Sku的销售价格。商品的价格要在商品所有的sku的价格之间。精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    Price  *string `json:"price" required:"true" `
    /*
        Sku的商家外部id     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN defalutValue��zh_CN    */
    Lang  *string `json:"lang,omitempty" required:"false" `
    /*
        sku所属商品的价格。当用户新增sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够添加成功     */
    ItemPrice  *string `json:"item_price,omitempty" required:"false" `
    /*
        忽略警告提示.     */
    Ignorewarning  *string `json:"ignorewarning,omitempty" required:"false" `
}

func (s *TaobaoItemSkuAddRequest) SetNumIid(v int64) *TaobaoItemSkuAddRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemSkuAddRequest) SetProperties(v string) *TaobaoItemSkuAddRequest {
    s.Properties = &v
    return s
}
func (s *TaobaoItemSkuAddRequest) SetQuantity(v int64) *TaobaoItemSkuAddRequest {
    s.Quantity = &v
    return s
}
func (s *TaobaoItemSkuAddRequest) SetPrice(v string) *TaobaoItemSkuAddRequest {
    s.Price = &v
    return s
}
func (s *TaobaoItemSkuAddRequest) SetOuterId(v string) *TaobaoItemSkuAddRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoItemSkuAddRequest) SetLang(v string) *TaobaoItemSkuAddRequest {
    s.Lang = &v
    return s
}
func (s *TaobaoItemSkuAddRequest) SetItemPrice(v string) *TaobaoItemSkuAddRequest {
    s.ItemPrice = &v
    return s
}
func (s *TaobaoItemSkuAddRequest) SetIgnorewarning(v string) *TaobaoItemSkuAddRequest {
    s.Ignorewarning = &v
    return s
}

func (req *TaobaoItemSkuAddRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoItemSkuAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}