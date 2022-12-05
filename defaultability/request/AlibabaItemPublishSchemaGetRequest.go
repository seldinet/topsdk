package request

import (
        "topsdk/util"
    )

type AlibabaItemPublishSchemaGetRequest struct {
    /*
        商品主图链接，最多5张，传入完整URL     */
    Images  *[]string `json:"images,omitempty" required:"false" `
    /*
        商品类型。b:一口价  a:拍卖  默认值b一口价 defalutValue��b    */
    ItemType  *string `json:"item_type,omitempty" required:"false" `
    /*
        业务扩展参数，需与平台约定好     */
    BizType  *string `json:"biz_type,omitempty" required:"false" `
    /*
        商品发布的市场。taobao:淘宝,tmall:天猫,litetao:淘宝特价版     */
    Market  *string `json:"market" required:"true" `
    /*
        商品类目ID     */
    CatId  *int64 `json:"cat_id" required:"true" `
    /*
        产品ID，天猫市场(market=tmall)时必填     */
    SpuId  *int64 `json:"spu_id,omitempty" required:"false" `
    /*
        商品条码     */
    Barcode  *string `json:"barcode,omitempty" required:"false" `
}

func (s *AlibabaItemPublishSchemaGetRequest) SetImages(v []string) *AlibabaItemPublishSchemaGetRequest {
    s.Images = &v
    return s
}
func (s *AlibabaItemPublishSchemaGetRequest) SetItemType(v string) *AlibabaItemPublishSchemaGetRequest {
    s.ItemType = &v
    return s
}
func (s *AlibabaItemPublishSchemaGetRequest) SetBizType(v string) *AlibabaItemPublishSchemaGetRequest {
    s.BizType = &v
    return s
}
func (s *AlibabaItemPublishSchemaGetRequest) SetMarket(v string) *AlibabaItemPublishSchemaGetRequest {
    s.Market = &v
    return s
}
func (s *AlibabaItemPublishSchemaGetRequest) SetCatId(v int64) *AlibabaItemPublishSchemaGetRequest {
    s.CatId = &v
    return s
}
func (s *AlibabaItemPublishSchemaGetRequest) SetSpuId(v int64) *AlibabaItemPublishSchemaGetRequest {
    s.SpuId = &v
    return s
}
func (s *AlibabaItemPublishSchemaGetRequest) SetBarcode(v string) *AlibabaItemPublishSchemaGetRequest {
    s.Barcode = &v
    return s
}

func (req *AlibabaItemPublishSchemaGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Images != nil) {
        paramMap["images"] = util.ConvertBasicList(*req.Images)
    }
    if(req.ItemType != nil) {
        paramMap["item_type"] = *req.ItemType
    }
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    if(req.Market != nil) {
        paramMap["market"] = *req.Market
    }
    if(req.CatId != nil) {
        paramMap["cat_id"] = *req.CatId
    }
    if(req.SpuId != nil) {
        paramMap["spu_id"] = *req.SpuId
    }
    if(req.Barcode != nil) {
        paramMap["barcode"] = *req.Barcode
    }
    return paramMap
}

func (req *AlibabaItemPublishSchemaGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}