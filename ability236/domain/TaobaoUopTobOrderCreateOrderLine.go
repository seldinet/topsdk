package domain

import (
        "topsdk/util"
    )

type TaobaoUopTobOrderCreateOrderLine struct {
    /*
        库存类型，ZP=正品、CC=残次     */
    InventoryType  *string `json:"inventory_type,omitempty" `

    /*
        原交易单，供销平台交易单号、分销平台单号     */
    SourceOrderCode  *string `json:"source_order_code,omitempty" `

    /*
        子交易单号     */
    SubSourceOrderCode  *string `json:"sub_source_order_code,omitempty" `

    /*
        批次编码     */
    BatchCode  *string `json:"batch_code,omitempty" `

    /*
        生产日期，生产日期(YYYY-MM-DD)     */
    ProductDate  *util.LocalTime `json:"product_date,omitempty" `

    /*
        过期日期，生产日期(YYYY-MM-DD)     */
    ExpireDate  *util.LocalTime `json:"expire_date,omitempty" `

    /*
        生产批号     */
    ProduceCode  *string `json:"produce_code,omitempty" `

    /*
        商品数量     */
    ItemQuantity  *int64 `json:"item_quantity,omitempty" `

    /*
        商品编码     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        商品名称     */
    ItemName  *string `json:"item_name,omitempty" `

    /*
        商品ID     */
    ItemId  *string `json:"item_id,omitempty" `

    /*
        订单行号     */
    OrderLineNo  *string `json:"order_line_no,omitempty" `

}

func (s *TaobaoUopTobOrderCreateOrderLine) SetInventoryType(v string) *TaobaoUopTobOrderCreateOrderLine {
    s.InventoryType = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderLine) SetSourceOrderCode(v string) *TaobaoUopTobOrderCreateOrderLine {
    s.SourceOrderCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderLine) SetSubSourceOrderCode(v string) *TaobaoUopTobOrderCreateOrderLine {
    s.SubSourceOrderCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderLine) SetBatchCode(v string) *TaobaoUopTobOrderCreateOrderLine {
    s.BatchCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderLine) SetProductDate(v util.LocalTime) *TaobaoUopTobOrderCreateOrderLine {
    s.ProductDate = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderLine) SetExpireDate(v util.LocalTime) *TaobaoUopTobOrderCreateOrderLine {
    s.ExpireDate = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderLine) SetProduceCode(v string) *TaobaoUopTobOrderCreateOrderLine {
    s.ProduceCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderLine) SetItemQuantity(v int64) *TaobaoUopTobOrderCreateOrderLine {
    s.ItemQuantity = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderLine) SetItemCode(v string) *TaobaoUopTobOrderCreateOrderLine {
    s.ItemCode = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderLine) SetItemName(v string) *TaobaoUopTobOrderCreateOrderLine {
    s.ItemName = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderLine) SetItemId(v string) *TaobaoUopTobOrderCreateOrderLine {
    s.ItemId = &v
    return s
}
func (s *TaobaoUopTobOrderCreateOrderLine) SetOrderLineNo(v string) *TaobaoUopTobOrderCreateOrderLine {
    s.OrderLineNo = &v
    return s
}
