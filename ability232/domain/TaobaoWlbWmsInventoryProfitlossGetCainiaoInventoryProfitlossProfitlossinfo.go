package domain

import (
        "topsdk/util"
    )

type TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo struct {
    /*
        商品信息列表     */
    OrderItemList  *[]TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitemlist `json:"order_item_list,omitempty" `

    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        仓库订单编码     */
    CnOrderCode  *string `json:"cn_order_code,omitempty" `

    /*
        订单类型： 701 盘点出库 702 盘点入库     */
    OrderType  *int64 `json:"order_type,omitempty" `

    /*
        备注     */
    Remark  *string `json:"remark,omitempty" `

    /*
        单据生成时间     */
    CreatedTime  *util.LocalTime `json:"created_time,omitempty" `

}

func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo) SetOrderItemList(v []TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossOrderitemlist) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo {
    s.OrderItemList = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo) SetStoreCode(v string) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo) SetCnOrderCode(v string) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo {
    s.CnOrderCode = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo) SetOrderType(v int64) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo {
    s.OrderType = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo) SetRemark(v string) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo) SetCreatedTime(v util.LocalTime) *TaobaoWlbWmsInventoryProfitlossGetCainiaoInventoryProfitlossProfitlossinfo {
    s.CreatedTime = &v
    return s
}
