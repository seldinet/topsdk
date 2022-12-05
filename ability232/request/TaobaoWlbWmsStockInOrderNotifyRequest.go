package request

import (
        "topsdk/ability232/domain"
        "topsdk/util"
    )

type TaobaoWlbWmsStockInOrderNotifyRequest struct {
    /*
        入库单据编码     */
    OrderCode  *string `json:"order_code" required:"true" `
    /*
        仓库编码     */
    StoreCode  *string `json:"store_code" required:"true" `
    /*
        单据类型 601普通入库单、501销退入库单、302 调拨入库单、904其他入库单、306 B2B入库     */
    OrderType  *int64 `json:"order_type" required:"true" `
    /*
        可选择性文本透传至WMS，比如加工归还、委外归还、借出归还、内部归还等     */
    InboundTypeDesc  *string `json:"inbound_type_desc,omitempty" required:"false" `
    /*
        订单标记以逗号分隔：  9:上门退货入库 13: 退货时是否收取发票，默认不收取（即没13为多选项，如1,2,8,9）     */
    OrderFlag  *string `json:"order_flag,omitempty" required:"false" `
    /*
        单据创建时间     */
    OrderCreateTime  *util.LocalTime `json:"order_create_time" required:"true" `
    /*
        供应商编码，往来单位编码     */
    SupplierCode  *string `json:"supplier_code,omitempty" required:"false" `
    /*
        供应商名称 ，往来单位名称     */
    SupplierName  *string `json:"supplier_name,omitempty" required:"false" `
    /*
        配送公司编码，拒收（非妥投）的销退订单，由ERP填充原单配送公司编码     */
    TmsServiceCode  *string `json:"tms_service_code,omitempty" required:"false" `
    /*
        快递公司名称     */
    TmsServiceName  *string `json:"tms_service_name,omitempty" required:"false" `
    /*
        运单号&托运单号 1)	入库单支持多运单号录入 2)	销退场景下如果是拒收（非妥投运单）由ERP填充原运单号     */
    TmsOrderCode  *string `json:"tms_order_code,omitempty" required:"false" `
    /*
        来源单据号，如销售退货时填充原销售订单号     */
    PrevOrderCode  *string `json:"prev_order_code,omitempty" required:"false" `
    /*
        销退时请提供退货的原因     */
    ReturnReason  *string `json:"return_reason,omitempty" required:"false" `
    /*
        预期送达开始时间     */
    ExpectStartTime  *string `json:"expect_start_time,omitempty" required:"false" `
    /*
        预期送达结束时间     */
    ExpectEndTime  *string `json:"expect_end_time,omitempty" required:"false" `
    /*
        扩展属性, key-value结构，格式要求： 以英文分号“;”分隔每组key-value，以英文冒号“:”分隔key与value。如果value中带有分号，需要转成下划线“_”，如果带有冒号，需要转成中划线“-”     */
    ExtendFields  *string `json:"extend_fields,omitempty" required:"false" `
    /*
        备注，销退入库订单需要留言备注填充到此字段     */
    Remark  *string `json:"remark,omitempty" required:"false" `
    /*
        系统自动生成     */
    SenderInfo  *domain.TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl `json:"sender_info,omitempty" required:"false" `
    /*
        系统自动生成     */
    ReceiverInfo  *domain.TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl `json:"receiver_info,omitempty" required:"false" `
    /*
        系统自动生成     */
    OrderItemList  *[]domain.TaobaoWlbWmsStockInOrderNotifyOrderitemlistwlbwmsstockinordernotifywl `json:"order_item_list" required:"true" `
}

func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderCode(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.OrderCode = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetStoreCode(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.StoreCode = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderType(v int64) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.OrderType = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetInboundTypeDesc(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.InboundTypeDesc = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderFlag(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.OrderFlag = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderCreateTime(v util.LocalTime) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.OrderCreateTime = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetSupplierCode(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.SupplierCode = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetSupplierName(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.SupplierName = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetTmsServiceCode(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.TmsServiceCode = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetTmsServiceName(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.TmsServiceName = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetTmsOrderCode(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.TmsOrderCode = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetPrevOrderCode(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.PrevOrderCode = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetReturnReason(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.ReturnReason = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetExpectStartTime(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.ExpectStartTime = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetExpectEndTime(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.ExpectEndTime = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetExtendFields(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.ExtendFields = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetRemark(v string) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetSenderInfo(v domain.TaobaoWlbWmsStockInOrderNotifySenderinfowlbwmsstockinordernotifywl) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.SenderInfo = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetReceiverInfo(v domain.TaobaoWlbWmsStockInOrderNotifyReceiverinfowlbwmsstockinordernotifywl) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.ReceiverInfo = &v
    return s
}
func (s *TaobaoWlbWmsStockInOrderNotifyRequest) SetOrderItemList(v []domain.TaobaoWlbWmsStockInOrderNotifyOrderitemlistwlbwmsstockinordernotifywl) *TaobaoWlbWmsStockInOrderNotifyRequest {
    s.OrderItemList = &v
    return s
}

func (req *TaobaoWlbWmsStockInOrderNotifyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderCode != nil) {
        paramMap["order_code"] = *req.OrderCode
    }
    if(req.StoreCode != nil) {
        paramMap["store_code"] = *req.StoreCode
    }
    if(req.OrderType != nil) {
        paramMap["order_type"] = *req.OrderType
    }
    if(req.InboundTypeDesc != nil) {
        paramMap["inbound_type_desc"] = *req.InboundTypeDesc
    }
    if(req.OrderFlag != nil) {
        paramMap["order_flag"] = *req.OrderFlag
    }
    if(req.OrderCreateTime != nil) {
        paramMap["order_create_time"] = *req.OrderCreateTime
    }
    if(req.SupplierCode != nil) {
        paramMap["supplier_code"] = *req.SupplierCode
    }
    if(req.SupplierName != nil) {
        paramMap["supplier_name"] = *req.SupplierName
    }
    if(req.TmsServiceCode != nil) {
        paramMap["tms_service_code"] = *req.TmsServiceCode
    }
    if(req.TmsServiceName != nil) {
        paramMap["tms_service_name"] = *req.TmsServiceName
    }
    if(req.TmsOrderCode != nil) {
        paramMap["tms_order_code"] = *req.TmsOrderCode
    }
    if(req.PrevOrderCode != nil) {
        paramMap["prev_order_code"] = *req.PrevOrderCode
    }
    if(req.ReturnReason != nil) {
        paramMap["return_reason"] = *req.ReturnReason
    }
    if(req.ExpectStartTime != nil) {
        paramMap["expect_start_time"] = *req.ExpectStartTime
    }
    if(req.ExpectEndTime != nil) {
        paramMap["expect_end_time"] = *req.ExpectEndTime
    }
    if(req.ExtendFields != nil) {
        paramMap["extend_fields"] = *req.ExtendFields
    }
    if(req.Remark != nil) {
        paramMap["remark"] = *req.Remark
    }
    if(req.SenderInfo != nil) {
        paramMap["sender_info"] = util.ConvertStruct(*req.SenderInfo)
    }
    if(req.ReceiverInfo != nil) {
        paramMap["receiver_info"] = util.ConvertStruct(*req.ReceiverInfo)
    }
    if(req.OrderItemList != nil) {
        paramMap["order_item_list"] = util.ConvertStructList(*req.OrderItemList)
    }
    return paramMap
}

func (req *TaobaoWlbWmsStockInOrderNotifyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}