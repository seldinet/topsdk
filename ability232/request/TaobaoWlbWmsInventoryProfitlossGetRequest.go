package request


type TaobaoWlbWmsInventoryProfitlossGetRequest struct {
    /*
        菜鸟订单编码     */
    CnOrderCode  *string `json:"cn_order_code" required:"true" `
}

func (s *TaobaoWlbWmsInventoryProfitlossGetRequest) SetCnOrderCode(v string) *TaobaoWlbWmsInventoryProfitlossGetRequest {
    s.CnOrderCode = &v
    return s
}

func (req *TaobaoWlbWmsInventoryProfitlossGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CnOrderCode != nil) {
        paramMap["cn_order_code"] = *req.CnOrderCode
    }
    return paramMap
}

func (req *TaobaoWlbWmsInventoryProfitlossGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}