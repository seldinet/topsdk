package domain


type TaobaoEticketMerchantMaSendSendMaCallbackResp struct {
    /*
        业务参数KV     */
    AttributeMap  *string `json:"attribute_map,omitempty" `

}

func (s *TaobaoEticketMerchantMaSendSendMaCallbackResp) SetAttributeMap(v string) *TaobaoEticketMerchantMaSendSendMaCallbackResp {
    s.AttributeMap = &v
    return s
}
