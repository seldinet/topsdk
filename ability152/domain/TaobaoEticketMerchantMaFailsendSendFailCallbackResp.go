package domain


type TaobaoEticketMerchantMaFailsendSendFailCallbackResp struct {
    /*
        回复业务KV     */
    AttributeMap  *string `json:"attribute_map,omitempty" `

}

func (s *TaobaoEticketMerchantMaFailsendSendFailCallbackResp) SetAttributeMap(v string) *TaobaoEticketMerchantMaFailsendSendFailCallbackResp {
    s.AttributeMap = &v
    return s
}
