package domain


type TaobaoEticketMerchantMaResendSendMaCallbackResp struct {
    /*
        回复业务KV     */
    AttributeMap  *string `json:"attribute_map,omitempty" `

}

func (s *TaobaoEticketMerchantMaResendSendMaCallbackResp) SetAttributeMap(v string) *TaobaoEticketMerchantMaResendSendMaCallbackResp {
    s.AttributeMap = &v
    return s
}
