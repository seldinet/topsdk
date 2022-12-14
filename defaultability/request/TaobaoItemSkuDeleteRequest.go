package request


type TaobaoItemSkuDeleteRequest struct {
    /*
        Sku所属商品数字id，可通过 taobao.item.get 获取。必选     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充     */
    Properties  *string `json:"properties" required:"true" `
    /*
        Sku文字的版本。可选值:zh_HK(繁体),zh_CN(简体);默认值:zh_CN defalutValue��zh_CN    */
    Lang  *string `json:"lang,omitempty" required:"false" `
    /*
        sku所属商品的价格。当用户删除sku，使商品价格不属于sku价格之间的时候，用于修改商品的价格，使sku能够删除成功     */
    ItemPrice  *string `json:"item_price,omitempty" required:"false" `
    /*
        sku所属商品的数量,大于0的整数。当用户删除sku，使商品数量不等于sku数量之和时候，用于修改商品的数量，使sku能够删除成功。特别是删除最后一个sku的时候，一定要设置商品数量到正常的值，否则删除失败     */
    ItemNum  *int64 `json:"item_num,omitempty" required:"false" `
    /*
        忽略警告提示.     */
    Ignorewarning  *string `json:"ignorewarning,omitempty" required:"false" `
}

func (s *TaobaoItemSkuDeleteRequest) SetNumIid(v int64) *TaobaoItemSkuDeleteRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemSkuDeleteRequest) SetProperties(v string) *TaobaoItemSkuDeleteRequest {
    s.Properties = &v
    return s
}
func (s *TaobaoItemSkuDeleteRequest) SetLang(v string) *TaobaoItemSkuDeleteRequest {
    s.Lang = &v
    return s
}
func (s *TaobaoItemSkuDeleteRequest) SetItemPrice(v string) *TaobaoItemSkuDeleteRequest {
    s.ItemPrice = &v
    return s
}
func (s *TaobaoItemSkuDeleteRequest) SetItemNum(v int64) *TaobaoItemSkuDeleteRequest {
    s.ItemNum = &v
    return s
}
func (s *TaobaoItemSkuDeleteRequest) SetIgnorewarning(v string) *TaobaoItemSkuDeleteRequest {
    s.Ignorewarning = &v
    return s
}

func (req *TaobaoItemSkuDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    if(req.Properties != nil) {
        paramMap["properties"] = *req.Properties
    }
    if(req.Lang != nil) {
        paramMap["lang"] = *req.Lang
    }
    if(req.ItemPrice != nil) {
        paramMap["item_price"] = *req.ItemPrice
    }
    if(req.ItemNum != nil) {
        paramMap["item_num"] = *req.ItemNum
    }
    if(req.Ignorewarning != nil) {
        paramMap["ignorewarning"] = *req.Ignorewarning
    }
    return paramMap
}

func (req *TaobaoItemSkuDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}