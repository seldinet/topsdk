package request


type AlibabaItemPublishSubmitRequest struct {
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
    /*
        商品schema信息，通过alibaba.item.publish.props.get获取并补全后提交     */
    Schema  *string `json:"schema" required:"true" `
}

func (s *AlibabaItemPublishSubmitRequest) SetBizType(v string) *AlibabaItemPublishSubmitRequest {
    s.BizType = &v
    return s
}
func (s *AlibabaItemPublishSubmitRequest) SetMarket(v string) *AlibabaItemPublishSubmitRequest {
    s.Market = &v
    return s
}
func (s *AlibabaItemPublishSubmitRequest) SetCatId(v int64) *AlibabaItemPublishSubmitRequest {
    s.CatId = &v
    return s
}
func (s *AlibabaItemPublishSubmitRequest) SetSpuId(v int64) *AlibabaItemPublishSubmitRequest {
    s.SpuId = &v
    return s
}
func (s *AlibabaItemPublishSubmitRequest) SetBarcode(v string) *AlibabaItemPublishSubmitRequest {
    s.Barcode = &v
    return s
}
func (s *AlibabaItemPublishSubmitRequest) SetSchema(v string) *AlibabaItemPublishSubmitRequest {
    s.Schema = &v
    return s
}

func (req *AlibabaItemPublishSubmitRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
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
    if(req.Schema != nil) {
        paramMap["schema"] = *req.Schema
    }
    return paramMap
}

func (req *AlibabaItemPublishSubmitRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}