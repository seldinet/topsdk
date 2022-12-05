package request

import (
        "topsdk/ability149/domain"
        "topsdk/util"
    )

type TmallChannelTradeOrderCreateRequest struct {
    /*
        入参     */
    Param0  *domain.TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam `json:"param0" required:"true" `
}

func (s *TmallChannelTradeOrderCreateRequest) SetParam0(v domain.TmallChannelTradeOrderCreateTopChannelPurchaseOrderCreateParam) *TmallChannelTradeOrderCreateRequest {
    s.Param0 = &v
    return s
}

func (req *TmallChannelTradeOrderCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Param0 != nil) {
        paramMap["param0"] = util.ConvertStruct(*req.Param0)
    }
    return paramMap
}

func (req *TmallChannelTradeOrderCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}