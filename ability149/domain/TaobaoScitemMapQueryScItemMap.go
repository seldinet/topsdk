package domain


type TaobaoScitemMapQueryScItemMap struct {
    /*
        后端商品所有者名称     */
    RelUserNick  *string `json:"rel_user_nick,omitempty" `

    /*
        后端商品ID     */
    RelItemId  *int64 `json:"rel_item_id,omitempty" `

    /*
        map_type=1时，item_id为IC商品id

map_type=7时，item_id为分销商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        后端商品所有者id     */
    RelUserId  *int64 `json:"rel_user_id,omitempty" `

    /*
        1:前台和后台关系
7:分销和后台关系     */
    MapType  *int64 `json:"map_type,omitempty" `

    /*
        Ic商家id(分销商id)     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        当宝贝下没SKU时该字段为空     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        Ic商家nick(分销商nick)     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        后端商品所有者商家编码     */
    RelOuterCode  *string `json:"rel_outer_code,omitempty" `

}

func (s *TaobaoScitemMapQueryScItemMap) SetRelUserNick(v string) *TaobaoScitemMapQueryScItemMap {
    s.RelUserNick = &v
    return s
}
func (s *TaobaoScitemMapQueryScItemMap) SetRelItemId(v int64) *TaobaoScitemMapQueryScItemMap {
    s.RelItemId = &v
    return s
}
func (s *TaobaoScitemMapQueryScItemMap) SetItemId(v int64) *TaobaoScitemMapQueryScItemMap {
    s.ItemId = &v
    return s
}
func (s *TaobaoScitemMapQueryScItemMap) SetRelUserId(v int64) *TaobaoScitemMapQueryScItemMap {
    s.RelUserId = &v
    return s
}
func (s *TaobaoScitemMapQueryScItemMap) SetMapType(v int64) *TaobaoScitemMapQueryScItemMap {
    s.MapType = &v
    return s
}
func (s *TaobaoScitemMapQueryScItemMap) SetUserId(v int64) *TaobaoScitemMapQueryScItemMap {
    s.UserId = &v
    return s
}
func (s *TaobaoScitemMapQueryScItemMap) SetSkuId(v int64) *TaobaoScitemMapQueryScItemMap {
    s.SkuId = &v
    return s
}
func (s *TaobaoScitemMapQueryScItemMap) SetUserNick(v string) *TaobaoScitemMapQueryScItemMap {
    s.UserNick = &v
    return s
}
func (s *TaobaoScitemMapQueryScItemMap) SetRelOuterCode(v string) *TaobaoScitemMapQueryScItemMap {
    s.RelOuterCode = &v
    return s
}
