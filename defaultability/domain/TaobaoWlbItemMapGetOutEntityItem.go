package domain


type TaobaoWlbItemMapGetOutEntityItem struct {
    /*
        外部实体类型：
IC_ITEM--ic商品
IC_SKU--ic销售单元     */
    EntityType  *string `json:"entity_type,omitempty" `

    /*
        entity_type对应的实体类型的id：
当entity_type为IC_ITEM时，entity_id为ic的商品id     */
    EntityId  *string `json:"entity_id,omitempty" `

}

func (s *TaobaoWlbItemMapGetOutEntityItem) SetEntityType(v string) *TaobaoWlbItemMapGetOutEntityItem {
    s.EntityType = &v
    return s
}
func (s *TaobaoWlbItemMapGetOutEntityItem) SetEntityId(v string) *TaobaoWlbItemMapGetOutEntityItem {
    s.EntityId = &v
    return s
}
