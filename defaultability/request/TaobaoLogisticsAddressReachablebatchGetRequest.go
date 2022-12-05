package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type TaobaoLogisticsAddressReachablebatchGetRequest struct {
    /*
        筛单用户输入地址结构     */
    AddressList  *[]domain.TaobaoLogisticsAddressReachablebatchGetAddressReachable `json:"address_list,omitempty" required:"false" `
}

func (s *TaobaoLogisticsAddressReachablebatchGetRequest) SetAddressList(v []domain.TaobaoLogisticsAddressReachablebatchGetAddressReachable) *TaobaoLogisticsAddressReachablebatchGetRequest {
    s.AddressList = &v
    return s
}

func (req *TaobaoLogisticsAddressReachablebatchGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AddressList != nil) {
        paramMap["address_list"] = util.ConvertStructList(*req.AddressList)
    }
    return paramMap
}

func (req *TaobaoLogisticsAddressReachablebatchGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}