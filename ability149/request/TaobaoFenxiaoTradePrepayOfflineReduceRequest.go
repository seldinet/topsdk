package request

import (
        "topsdk/ability149/domain"
        "topsdk/util"
    )

type TaobaoFenxiaoTradePrepayOfflineReduceRequest struct {
    /*
        减少流水     */
    OfflineReducePrepayParam  *domain.TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto `json:"offline_reduce_prepay_param" required:"true" `
}

func (s *TaobaoFenxiaoTradePrepayOfflineReduceRequest) SetOfflineReducePrepayParam(v domain.TaobaoFenxiaoTradePrepayOfflineReduceTopOfflineReducePrepayDto) *TaobaoFenxiaoTradePrepayOfflineReduceRequest {
    s.OfflineReducePrepayParam = &v
    return s
}

func (req *TaobaoFenxiaoTradePrepayOfflineReduceRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OfflineReducePrepayParam != nil) {
        paramMap["offline_reduce_prepay_param"] = util.ConvertStruct(*req.OfflineReducePrepayParam)
    }
    return paramMap
}

func (req *TaobaoFenxiaoTradePrepayOfflineReduceRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}