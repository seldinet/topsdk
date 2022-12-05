package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type TaobaoTradesSoldQueryRequest struct {
    /*
        查询条件列表，多个条件之间是OR关系，最多支持20个。receiver_name、receiver_mobile、receiver_phone至少有一个值不为空。     */
    QueryList  *[]domain.TaobaoTradesSoldQueryOrderQuery `json:"query_list" required:"true" `
    /*
        业务场景编码。不同场景，策略不同。请根据产品功能选择相应的场景编号。可选的场景：1001(客服咨询)、1002(售后服务)，<a href="https://open.taobao.com/doc.htm?docId=120186&docType=1" target="_blank">详情点击</a>     */
    Scene  *string `json:"scene,omitempty" required:"false" `
    /*
        订单类型，默认值为1，可选值为：交易(1)，分销(2)，换货(3) defalutValue��1    */
    OrderType  *string `json:"order_type,omitempty" required:"false" `
}

func (s *TaobaoTradesSoldQueryRequest) SetQueryList(v []domain.TaobaoTradesSoldQueryOrderQuery) *TaobaoTradesSoldQueryRequest {
    s.QueryList = &v
    return s
}
func (s *TaobaoTradesSoldQueryRequest) SetScene(v string) *TaobaoTradesSoldQueryRequest {
    s.Scene = &v
    return s
}
func (s *TaobaoTradesSoldQueryRequest) SetOrderType(v string) *TaobaoTradesSoldQueryRequest {
    s.OrderType = &v
    return s
}

func (req *TaobaoTradesSoldQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.QueryList != nil) {
        paramMap["query_list"] = util.ConvertStructList(*req.QueryList)
    }
    if(req.Scene != nil) {
        paramMap["scene"] = *req.Scene
    }
    if(req.OrderType != nil) {
        paramMap["order_type"] = *req.OrderType
    }
    return paramMap
}

func (req *TaobaoTradesSoldQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}