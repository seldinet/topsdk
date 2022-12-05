package request

import (
        "topsdk/ability152/domain"
        "topsdk/util"
    )

type TaobaoEticketMerchantMaSendRequest struct {
    /*
        业务类型 defalutValue��3001    */
    BizType  *int64 `json:"biz_type,omitempty" required:"false" `
    /*
        需要发送的码列表     */
    IsvMaList  *[]domain.TaobaoEticketMerchantMaSendIsvMa `json:"isv_ma_list" required:"true" `
    /*
        业务id（订单号）     */
    OuterId  *string `json:"outer_id" required:"true" `
    /*
        需要跟发码通知获取到的参数一致     */
    Token  *string `json:"token" required:"true" `
}

func (s *TaobaoEticketMerchantMaSendRequest) SetBizType(v int64) *TaobaoEticketMerchantMaSendRequest {
    s.BizType = &v
    return s
}
func (s *TaobaoEticketMerchantMaSendRequest) SetIsvMaList(v []domain.TaobaoEticketMerchantMaSendIsvMa) *TaobaoEticketMerchantMaSendRequest {
    s.IsvMaList = &v
    return s
}
func (s *TaobaoEticketMerchantMaSendRequest) SetOuterId(v string) *TaobaoEticketMerchantMaSendRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoEticketMerchantMaSendRequest) SetToken(v string) *TaobaoEticketMerchantMaSendRequest {
    s.Token = &v
    return s
}

func (req *TaobaoEticketMerchantMaSendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.BizType != nil) {
        paramMap["biz_type"] = *req.BizType
    }
    if(req.IsvMaList != nil) {
        paramMap["isv_ma_list"] = util.ConvertStructList(*req.IsvMaList)
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Token != nil) {
        paramMap["token"] = *req.Token
    }
    return paramMap
}

func (req *TaobaoEticketMerchantMaSendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}