package request


type TaobaoVmarketEticketTasksGetRequest struct {
    /*
        卖家家ID(信任卖家不必传，码商可选)     */
    SellerId  *int64 `json:"seller_id,omitempty" required:"false" `
    /*
        返回结果类型:
1:返回通知失败的订单
2.返回通知成功回调失败的订单     */
    Type  *int64 `json:"type" required:"true" `
    /*
        页码。取值范围:大于零的整数; 默认值:1 defalutValue��1    */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页获取条数。默认值40，最小值1，最大值100。 defalutValue��40    */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        码商ID，如果是码商，必须传，如果是信任卖家，不需要传     */
    CodemerchantId  *int64 `json:"codemerchant_id,omitempty" required:"false" `
}

func (s *TaobaoVmarketEticketTasksGetRequest) SetSellerId(v int64) *TaobaoVmarketEticketTasksGetRequest {
    s.SellerId = &v
    return s
}
func (s *TaobaoVmarketEticketTasksGetRequest) SetType(v int64) *TaobaoVmarketEticketTasksGetRequest {
    s.Type = &v
    return s
}
func (s *TaobaoVmarketEticketTasksGetRequest) SetPageNo(v int64) *TaobaoVmarketEticketTasksGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoVmarketEticketTasksGetRequest) SetPageSize(v int64) *TaobaoVmarketEticketTasksGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoVmarketEticketTasksGetRequest) SetCodemerchantId(v int64) *TaobaoVmarketEticketTasksGetRequest {
    s.CodemerchantId = &v
    return s
}

func (req *TaobaoVmarketEticketTasksGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SellerId != nil) {
        paramMap["seller_id"] = *req.SellerId
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.CodemerchantId != nil) {
        paramMap["codemerchant_id"] = *req.CodemerchantId
    }
    return paramMap
}

func (req *TaobaoVmarketEticketTasksGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}