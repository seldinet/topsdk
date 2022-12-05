package request

import (
        "topsdk/ability149/domain"
        "topsdk/util"
    )

type TaobaoFenxiaoTradePrepayOfflineAddRequest struct {
    /*
        增加流水     */
    OfflineAddPrepayParam  *domain.TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto `json:"offline_add_prepay_param" required:"true" `
}

func (s *TaobaoFenxiaoTradePrepayOfflineAddRequest) SetOfflineAddPrepayParam(v domain.TaobaoFenxiaoTradePrepayOfflineAddTopOfflineAddPrepayDto) *TaobaoFenxiaoTradePrepayOfflineAddRequest {
    s.OfflineAddPrepayParam = &v
    return s
}

func (req *TaobaoFenxiaoTradePrepayOfflineAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OfflineAddPrepayParam != nil) {
        paramMap["offline_add_prepay_param"] = util.ConvertStruct(*req.OfflineAddPrepayParam)
    }
    return paramMap
}

func (req *TaobaoFenxiaoTradePrepayOfflineAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}