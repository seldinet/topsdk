package request


type TaobaoFenxiaoDiscountsGetRequest struct {
    /*
        折扣ID     */
    DiscountId  *int64 `json:"discount_id,omitempty" required:"false" `
    /*
        指定查询额外的信息，可选值：DETAIL（查询折扣详情），多个可选值用逗号分割。（只允许指定折扣ID情况下使用）     */
    ExtFields  *string `json:"ext_fields,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoDiscountsGetRequest) SetDiscountId(v int64) *TaobaoFenxiaoDiscountsGetRequest {
    s.DiscountId = &v
    return s
}
func (s *TaobaoFenxiaoDiscountsGetRequest) SetExtFields(v string) *TaobaoFenxiaoDiscountsGetRequest {
    s.ExtFields = &v
    return s
}

func (req *TaobaoFenxiaoDiscountsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DiscountId != nil) {
        paramMap["discount_id"] = *req.DiscountId
    }
    if(req.ExtFields != nil) {
        paramMap["ext_fields"] = *req.ExtFields
    }
    return paramMap
}

func (req *TaobaoFenxiaoDiscountsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}