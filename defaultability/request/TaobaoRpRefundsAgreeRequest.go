package request


type TaobaoRpRefundsAgreeRequest struct {
    /*
        短信验证码，如果退款金额达到一定的数量，后端会返回调用失败，并同时往卖家的手机发送一条短信验证码。接下来用收到的短信验证码再次发起API调用即可完成退款操作。     */
    Code  *string `json:"code,omitempty" required:"false" `
    /*
        退款信息，格式：refund_id|amount|version|phase，其中refund_id为退款编号，amount为退款金额（以分为单位），version为退款最后更新时间（时间戳格式），phase为退款阶段（可选值为：onsale, aftersale，天猫退款必值，淘宝退款不需要传），多个退款以半角逗号分隔。     */
    RefundInfos  *string `json:"refund_infos" required:"true" `
}

func (s *TaobaoRpRefundsAgreeRequest) SetCode(v string) *TaobaoRpRefundsAgreeRequest {
    s.Code = &v
    return s
}
func (s *TaobaoRpRefundsAgreeRequest) SetRefundInfos(v string) *TaobaoRpRefundsAgreeRequest {
    s.RefundInfos = &v
    return s
}

func (req *TaobaoRpRefundsAgreeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.RefundInfos != nil) {
        paramMap["refund_infos"] = *req.RefundInfos
    }
    return paramMap
}

func (req *TaobaoRpRefundsAgreeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}