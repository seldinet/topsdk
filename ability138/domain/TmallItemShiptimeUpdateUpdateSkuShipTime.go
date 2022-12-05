package domain


type TmallItemShiptimeUpdateUpdateSkuShipTime struct {
    /*
        被更新发货时间；格式和具体设置的发货时间格式相关。绝对发货时间填写yyyy-MM-dd;相对发货时间填写数字。     */
    ShipTime  *string `json:"ship_time,omitempty" `

    /*
        SKU的ID     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        Sku属性串。格式:pid:vid;pid:vid,如: 1627207:3232483;1630696:3284570,表示机身颜色:军绿色;手机套餐:一电一充     */
    Properties  *string `json:"properties,omitempty" `

    /*
        Sku的商家外部id；如：2015_07_23_D_123     */
    OuterId  *string `json:"outer_id,omitempty" `

}

func (s *TmallItemShiptimeUpdateUpdateSkuShipTime) SetShipTime(v string) *TmallItemShiptimeUpdateUpdateSkuShipTime {
    s.ShipTime = &v
    return s
}
func (s *TmallItemShiptimeUpdateUpdateSkuShipTime) SetSkuId(v int64) *TmallItemShiptimeUpdateUpdateSkuShipTime {
    s.SkuId = &v
    return s
}
func (s *TmallItemShiptimeUpdateUpdateSkuShipTime) SetProperties(v string) *TmallItemShiptimeUpdateUpdateSkuShipTime {
    s.Properties = &v
    return s
}
func (s *TmallItemShiptimeUpdateUpdateSkuShipTime) SetOuterId(v string) *TmallItemShiptimeUpdateUpdateSkuShipTime {
    s.OuterId = &v
    return s
}
