package request


type TaobaoLogisticsPartnersGetRequest struct {
    /*
        物流公司揽货地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取 defalutValue��0    */
    SourceId  *string `json:"source_id,omitempty" required:"false" `
    /*
        物流公司派送地地区码（必须是区、县一级的）.参考:http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201401/t20140116_501070.html或者调用 taobao.areas.get 获取 defalutValue��0    */
    TargetId  *string `json:"target_id,omitempty" required:"false" `
    /*
        服务类型，根据此参数可查出提供相应服务类型的物流公司信息(物流公司状态正常)，可选值：cod(货到付款)、online(在线下单)、 offline(自己联系)、limit(限时物流)。然后再根据source_id,target_id,goods_value这三个条件来过滤物流公司. 目前输入自己联系服务类型将会返回空，因为自己联系并没有具体的服务信息，所以不会有记录。     */
    ServiceType  *string `json:"service_type,omitempty" required:"false" `
    /*
        货物价格.只有当选择货到付款此参数才会有效 defalutValue��0    */
    GoodsValue  *string `json:"goods_value,omitempty" required:"false" `
    /*
        是否需要揽收资费信息，默认false。在此值为false时，返回内容中将无carriage。在未设置source_id或target_id的情况下，无法查询揽收资费信息。自己联系无揽收资费记录。 defalutValue��false    */
    IsNeedCarriage  *bool `json:"is_need_carriage,omitempty" required:"false" `
}

func (s *TaobaoLogisticsPartnersGetRequest) SetSourceId(v string) *TaobaoLogisticsPartnersGetRequest {
    s.SourceId = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetRequest) SetTargetId(v string) *TaobaoLogisticsPartnersGetRequest {
    s.TargetId = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetRequest) SetServiceType(v string) *TaobaoLogisticsPartnersGetRequest {
    s.ServiceType = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetRequest) SetGoodsValue(v string) *TaobaoLogisticsPartnersGetRequest {
    s.GoodsValue = &v
    return s
}
func (s *TaobaoLogisticsPartnersGetRequest) SetIsNeedCarriage(v bool) *TaobaoLogisticsPartnersGetRequest {
    s.IsNeedCarriage = &v
    return s
}

func (req *TaobaoLogisticsPartnersGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SourceId != nil) {
        paramMap["source_id"] = *req.SourceId
    }
    if(req.TargetId != nil) {
        paramMap["target_id"] = *req.TargetId
    }
    if(req.ServiceType != nil) {
        paramMap["service_type"] = *req.ServiceType
    }
    if(req.GoodsValue != nil) {
        paramMap["goods_value"] = *req.GoodsValue
    }
    if(req.IsNeedCarriage != nil) {
        paramMap["is_need_carriage"] = *req.IsNeedCarriage
    }
    return paramMap
}

func (req *TaobaoLogisticsPartnersGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}