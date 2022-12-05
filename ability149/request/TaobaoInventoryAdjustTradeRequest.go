package request

import (
        "topsdk/util"
    )

type TaobaoInventoryAdjustTradeRequest struct {
    /*
        订单类型：B2C、B2B     */
    TbOrderType  *string `json:"tb_order_type" required:"true" `
    /*
        业务操作时间     */
    OperateTime  *util.LocalTime `json:"operate_time" required:"true" `
    /*
        商家外部定单号     */
    BizUniqueCode  *string `json:"biz_unique_code" required:"true" `
    /*
        商品初始库存信息： [{ "TBOrderCode”:”淘宝交易号”,"TBSubOrderCode ":"淘宝子交易单号,赠品可以不填","”isGift”:”TRUE或者FALSE,是否赠品”,storeCode":"商家仓库编码"," scItemId ":"商品后端ID","scItemCode":"商品商家编码"," originScItemId ":"原商品ID","inventoryType":"","quantity":"111","isComplete":"TRUE或者FALSE，是否全部确认出库"}]     */
    Items  *string `json:"items" required:"true" `
}

func (s *TaobaoInventoryAdjustTradeRequest) SetTbOrderType(v string) *TaobaoInventoryAdjustTradeRequest {
    s.TbOrderType = &v
    return s
}
func (s *TaobaoInventoryAdjustTradeRequest) SetOperateTime(v util.LocalTime) *TaobaoInventoryAdjustTradeRequest {
    s.OperateTime = &v
    return s
}
func (s *TaobaoInventoryAdjustTradeRequest) SetBizUniqueCode(v string) *TaobaoInventoryAdjustTradeRequest {
    s.BizUniqueCode = &v
    return s
}
func (s *TaobaoInventoryAdjustTradeRequest) SetItems(v string) *TaobaoInventoryAdjustTradeRequest {
    s.Items = &v
    return s
}

func (req *TaobaoInventoryAdjustTradeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TbOrderType != nil) {
        paramMap["tb_order_type"] = *req.TbOrderType
    }
    if(req.OperateTime != nil) {
        paramMap["operate_time"] = *req.OperateTime
    }
    if(req.BizUniqueCode != nil) {
        paramMap["biz_unique_code"] = *req.BizUniqueCode
    }
    if(req.Items != nil) {
        paramMap["items"] = *req.Items
    }
    return paramMap
}

func (req *TaobaoInventoryAdjustTradeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}