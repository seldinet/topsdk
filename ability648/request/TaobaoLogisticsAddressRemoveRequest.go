package request


type TaobaoLogisticsAddressRemoveRequest struct {
    /*
        地址库ID     */
    ContactId  *int64 `json:"contact_id" required:"true" `
}

func (s *TaobaoLogisticsAddressRemoveRequest) SetContactId(v int64) *TaobaoLogisticsAddressRemoveRequest {
    s.ContactId = &v
    return s
}

func (req *TaobaoLogisticsAddressRemoveRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ContactId != nil) {
        paramMap["contact_id"] = *req.ContactId
    }
    return paramMap
}

func (req *TaobaoLogisticsAddressRemoveRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}