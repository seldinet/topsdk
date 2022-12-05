package request


type TaobaoFenxiaoOrderConfirmPaidRequest struct {
    /*
        采购单编号。     */
    PurchaseOrderId  *int64 `json:"purchase_order_id" required:"true" `
    /*
        确认支付信息（字数小于100）     */
    ConfirmRemark  *string `json:"confirm_remark,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoOrderConfirmPaidRequest) SetPurchaseOrderId(v int64) *TaobaoFenxiaoOrderConfirmPaidRequest {
    s.PurchaseOrderId = &v
    return s
}
func (s *TaobaoFenxiaoOrderConfirmPaidRequest) SetConfirmRemark(v string) *TaobaoFenxiaoOrderConfirmPaidRequest {
    s.ConfirmRemark = &v
    return s
}

func (req *TaobaoFenxiaoOrderConfirmPaidRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.PurchaseOrderId != nil) {
        paramMap["purchase_order_id"] = *req.PurchaseOrderId
    }
    if(req.ConfirmRemark != nil) {
        paramMap["confirm_remark"] = *req.ConfirmRemark
    }
    return paramMap
}

func (req *TaobaoFenxiaoOrderConfirmPaidRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}