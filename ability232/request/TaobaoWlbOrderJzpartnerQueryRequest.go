package request


type TaobaoWlbOrderJzpartnerQueryRequest struct {
    /*
        淘宝交易订单号，如果不填写Tid则必须填写serviceType。如果填写Tid，则表明只需要查询对应订单的服务商。     */
    TaobaoTradeId  *int64 `json:"taobao_trade_id,omitempty" required:"false" `
    /*
        serviceType表示查询所有的支持服务类型的服务商。 家装干线服务     11 家装干支服务     12 家装干支装服务   13 卫浴大件干线     14 卫浴大件干支     15 卫浴大件安装     16 地板干线         17 地板干支         18 地板安装         19 灯具安装         20 卫浴小件安装     21 （注：同一个服务商针对不同类型的serviceType是具有不同的tpCode的）     */
    ServiceType  *int64 `json:"service_type,omitempty" required:"false" `
}

func (s *TaobaoWlbOrderJzpartnerQueryRequest) SetTaobaoTradeId(v int64) *TaobaoWlbOrderJzpartnerQueryRequest {
    s.TaobaoTradeId = &v
    return s
}
func (s *TaobaoWlbOrderJzpartnerQueryRequest) SetServiceType(v int64) *TaobaoWlbOrderJzpartnerQueryRequest {
    s.ServiceType = &v
    return s
}

func (req *TaobaoWlbOrderJzpartnerQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TaobaoTradeId != nil) {
        paramMap["taobao_trade_id"] = *req.TaobaoTradeId
    }
    if(req.ServiceType != nil) {
        paramMap["service_type"] = *req.ServiceType
    }
    return paramMap
}

func (req *TaobaoWlbOrderJzpartnerQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}