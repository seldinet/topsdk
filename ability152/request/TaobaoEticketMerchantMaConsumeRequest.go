package request

import (
        "topsdk/ability152/domain"
        "topsdk/util"
    )

type TaobaoEticketMerchantMaConsumeRequest struct {
    /*
        业务类型 defalutValue��3001    */
    BizType  *int64 `json:"biz_type,omitempty" required:"false" `
    /*
        需要被核销的码     */
    Code  *string `json:"code" required:"true" `
    /*
        核销份数     */
    ConsumeNum  *int64 `json:"consume_num" required:"true" `
    /*
        核销后换码的码列表     */
    IsvMaList  *[]domain.TaobaoEticketMerchantMaConsumeIsvMa `json:"isv_ma_list,omitempty" required:"false" `
    /*
        业务id（订单号）     */
    OuterId  *string `json:"outer_id" required:"true" `
    /*
        机具编号     */
    PosId  *string `json:"pos_id,omitempty" required:"false" `
    /*
        核销序列号，需要保证唯一     */
    SerialNum  *string `json:"serial_num" required:"true" `
    /*
        需要跟发码通知获取到的参数一致     */
    Token  *string `json:"token" required:"true" `
}

func (s *TaobaoEticketMerchantMaConsumeRequest) SetBizType(v int64) *TaobaoEticketMerchantMaConsumeRequest {
    s.BizType = &v
    return s
}
func (s *TaobaoEticketMerchantMaConsumeRequest) SetCode(v string) *TaobaoEticketMerchantMaConsumeRequest {
    s.Code = &v
    return s
}
func (s *TaobaoEticketMerchantMaConsumeRequest) SetConsumeNum(v int64) *TaobaoEticketMerchantMaConsumeRequest {
    s.ConsumeNum = &v
    return s
}
func (s *TaobaoEticketMerchantMaConsumeRequest) SetIsvMaList(v []domain.TaobaoEticketMerchantMaConsumeIsvMa) *TaobaoEticketMerchantMaConsumeRequest {
    s.IsvMaList = &v
    return s
}
func (s *TaobaoEticketMerchantMaConsumeRequest) SetOuterId(v string) *TaobaoEticketMerchantMaConsumeRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoEticketMerchantMaConsumeRequest) SetPosId(v string) *TaobaoEticketMerchantMaConsumeRequest {
    s.PosId = &v
    return s
}
func (s *TaobaoEticketMerchantMaConsumeRequest) SetSerialNum(v string) *TaobaoEticketMerchantMaConsumeRequest {
    s.SerialNum = &v
    return s
}
func (s *TaobaoEticketMerchantMaConsumeRequest) SetToken(v string) *TaobaoEticketMerchantMaConsumeRequest {
    s.Token = &v
    return s
}

func (req *TaobaoEticketMerchantMaConsumeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.ConsumeNum != nil) {
        paramMap["consume_num"] = *req.ConsumeNum
    }
    if(req.IsvMaList != nil) {
        paramMap["isv_ma_list"] = util.ConvertStructList(*req.IsvMaList)
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.PosId != nil) {
        paramMap["pos_id"] = *req.PosId
    }
    if(req.SerialNum != nil) {
        paramMap["serial_num"] = *req.SerialNum
    }
    if(req.Token != nil) {
        paramMap["token"] = *req.Token
    }
    return paramMap
}

func (req *TaobaoEticketMerchantMaConsumeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}