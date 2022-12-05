package domain


type TaobaoEticketMerchantMaConsumeConsumeMaCallbackResp struct {
    /*
        系统自动生成     */
    AttributeMap  *string `json:"attribute_map,omitempty" `

}

func (s *TaobaoEticketMerchantMaConsumeConsumeMaCallbackResp) SetAttributeMap(v string) *TaobaoEticketMerchantMaConsumeConsumeMaCallbackResp {
    s.AttributeMap = &v
    return s
}
