package domain


type TaobaoFulfillmentOrderAssembleOrderGroup struct {
    /*
        淘宝交易子订单id     */
    TaobaoSubOrderId  *string `json:"taobao_sub_order_id,omitempty" `

    /*
        order_id的类型，0:淘宝交易订单,1: 换货单,2:分销单，3:补货单，4:代发单 defalutValue:0    */
    OrderType  *int64 `json:"order_type,omitempty" `

    /*
        商品类型, 0:下单货品，1:赠品，2:其他 defalutValue:0    */
    ItemType  *int64 `json:"item_type,omitempty" `

    /*
        淘宝单号可以是交易订单、换货单、补货单、代发单或分销单等，当 order_type=0时，order_id =  taobao_parent_order_id。     */
    OrderId  *string `json:"order_id,omitempty" `

    /*
        erp系统内的订单id     */
    ErpOrderId  *string `json:"erp_order_id,omitempty" `

    /*
        淘宝交易主订单id     */
    TaobaoParentOrderId  *string `json:"taobao_parent_order_id,omitempty" `

}

func (s *TaobaoFulfillmentOrderAssembleOrderGroup) SetTaobaoSubOrderId(v string) *TaobaoFulfillmentOrderAssembleOrderGroup {
    s.TaobaoSubOrderId = &v
    return s
}
func (s *TaobaoFulfillmentOrderAssembleOrderGroup) SetOrderType(v int64) *TaobaoFulfillmentOrderAssembleOrderGroup {
    s.OrderType = &v
    return s
}
func (s *TaobaoFulfillmentOrderAssembleOrderGroup) SetItemType(v int64) *TaobaoFulfillmentOrderAssembleOrderGroup {
    s.ItemType = &v
    return s
}
func (s *TaobaoFulfillmentOrderAssembleOrderGroup) SetOrderId(v string) *TaobaoFulfillmentOrderAssembleOrderGroup {
    s.OrderId = &v
    return s
}
func (s *TaobaoFulfillmentOrderAssembleOrderGroup) SetErpOrderId(v string) *TaobaoFulfillmentOrderAssembleOrderGroup {
    s.ErpOrderId = &v
    return s
}
func (s *TaobaoFulfillmentOrderAssembleOrderGroup) SetTaobaoParentOrderId(v string) *TaobaoFulfillmentOrderAssembleOrderGroup {
    s.TaobaoParentOrderId = &v
    return s
}
