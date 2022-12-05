package request

import (
        "topsdk/util"
    )

type TaobaoLogisticsCompaniesGetRequest struct {
    /*
        需返回的字段列表。可选值:LogisticCompany 结构中的所有字段;多个字段间用","逗号隔开.
如:id,code,name,reg_mail_no
<br><font color='red'>说明：</font>
<br>id：物流公司ID
<br>code：物流公司code
<br>name：物流公司名称
<br>reg_mail_no：物流公司对应的运单规则     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        是否查询推荐物流公司.可选值:true,false.如果不提供此参数,将会返回所有支持电话联系的物流公司.     */
    IsRecommended  *bool `json:"is_recommended,omitempty" required:"false" `
    /*
        推荐物流公司的下单方式.可选值:offline(电话联系/自己联系),online(在线下单),all(即电话联系又在线下单). 此参数仅仅用于is_recommended 为ture时。就是说对于推荐物流公司才可用.如果不选择此参数将会返回推荐物流中支持电话联系的物流公司.     */
    OrderMode  *string `json:"order_mode,omitempty" required:"false" `
}

func (s *TaobaoLogisticsCompaniesGetRequest) SetFields(v []string) *TaobaoLogisticsCompaniesGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoLogisticsCompaniesGetRequest) SetIsRecommended(v bool) *TaobaoLogisticsCompaniesGetRequest {
    s.IsRecommended = &v
    return s
}
func (s *TaobaoLogisticsCompaniesGetRequest) SetOrderMode(v string) *TaobaoLogisticsCompaniesGetRequest {
    s.OrderMode = &v
    return s
}

func (req *TaobaoLogisticsCompaniesGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.IsRecommended != nil) {
        paramMap["is_recommended"] = *req.IsRecommended
    }
    if(req.OrderMode != nil) {
        paramMap["order_mode"] = *req.OrderMode
    }
    return paramMap
}

func (req *TaobaoLogisticsCompaniesGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}