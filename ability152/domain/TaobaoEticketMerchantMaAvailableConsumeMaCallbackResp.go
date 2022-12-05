package domain


type TaobaoEticketMerchantMaAvailableConsumeMaCallbackResp struct {
    /*
        业务回复KV     */
    AttributeMap  *string `json:"attribute_map,omitempty" `

}

func (s *TaobaoEticketMerchantMaAvailableConsumeMaCallbackResp) SetAttributeMap(v string) *TaobaoEticketMerchantMaAvailableConsumeMaCallbackResp {
    s.AttributeMap = &v
    return s
}
