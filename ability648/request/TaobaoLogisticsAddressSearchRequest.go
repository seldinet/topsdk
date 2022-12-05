package request


type TaobaoLogisticsAddressSearchRequest struct {
    /*
        可选，参数列表如下：<br><font color='red'>no_def:查询非默认地址<br>get_def:查询默认取货地址，也即对应卖家后台地址库中发货地址（send_def暂不起作用）<br>cancel_def:查询默认退货地址<br>缺省此参数时，查询所有当前用户的地址库</font>     */
    Rdef  *string `json:"rdef,omitempty" required:"false" `
}

func (s *TaobaoLogisticsAddressSearchRequest) SetRdef(v string) *TaobaoLogisticsAddressSearchRequest {
    s.Rdef = &v
    return s
}

func (req *TaobaoLogisticsAddressSearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Rdef != nil) {
        paramMap["rdef"] = *req.Rdef
    }
    return paramMap
}

func (req *TaobaoLogisticsAddressSearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}