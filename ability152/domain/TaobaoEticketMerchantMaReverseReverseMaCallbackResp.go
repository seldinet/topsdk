package domain


type TaobaoEticketMerchantMaReverseReverseMaCallbackResp struct {
    /*
        业务参数KV     */
    AttributeMap  *string `json:"attribute_map,omitempty" `

}

func (s *TaobaoEticketMerchantMaReverseReverseMaCallbackResp) SetAttributeMap(v string) *TaobaoEticketMerchantMaReverseReverseMaCallbackResp {
    s.AttributeMap = &v
    return s
}
