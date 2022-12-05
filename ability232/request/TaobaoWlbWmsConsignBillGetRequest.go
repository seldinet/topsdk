package request


type TaobaoWlbWmsConsignBillGetRequest struct {
    /*
        菜鸟订单编码,cnOrderCode与orderCode必须有一个值不可为空     */
    CnOrderCode  *string `json:"cn_order_code,omitempty" required:"false" `
    /*
        ERP订单编码,cnOrderCode与orderCode必须有一个值不可为空     */
    OrderCode  *string `json:"order_code,omitempty" required:"false" `
}

func (s *TaobaoWlbWmsConsignBillGetRequest) SetCnOrderCode(v string) *TaobaoWlbWmsConsignBillGetRequest {
    s.CnOrderCode = &v
    return s
}
func (s *TaobaoWlbWmsConsignBillGetRequest) SetOrderCode(v string) *TaobaoWlbWmsConsignBillGetRequest {
    s.OrderCode = &v
    return s
}

func (req *TaobaoWlbWmsConsignBillGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CnOrderCode != nil) {
        paramMap["cn_order_code"] = *req.CnOrderCode
    }
    if(req.OrderCode != nil) {
        paramMap["order_code"] = *req.OrderCode
    }
    return paramMap
}

func (req *TaobaoWlbWmsConsignBillGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}