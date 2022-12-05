package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoDealerRequisitionorderCreateRequest struct {
    /*
        采购清单，存放多个采购明细，每个采购明细内部以‘:’隔开，多个采购明细之间以‘,’隔开. 例(分销产品id:skuid:购买数量:申请单价,分销产品id:skuid:购买数量:申请单价)，申请单价的单位为分。不存在sku请留空skuid，如（分销产品id::购买数量:申请单价）     */
    OrderDetail  *[]string `json:"order_detail" required:"true" `
    /*
        配送方式。SELF_PICKUP：自提；LOGISTICS：仓库发货     */
    LogisticsType  *string `json:"logistics_type,omitempty" required:"false" `
    /*
        收货人所在省份     */
    Province  *string `json:"province" required:"true" `
    /*
        收货人所在市     */
    City  *string `json:"city,omitempty" required:"false" `
    /*
        收货人所在区     */
    District  *string `json:"district,omitempty" required:"false" `
    /*
        收货人所在街道地址     */
    Address  *string `json:"address" required:"true" `
    /*
        收货人所在地区邮政编码     */
    PostCode  *string `json:"post_code,omitempty" required:"false" `
    /*
        买家联系电话（此字段和mobile字段至少填写一个）     */
    Phone  *string `json:"phone,omitempty" required:"false" `
    /*
        买家的手机号码（1、此字段与phone字段至少填写一个。2、自提方式下此字段必填，保存提货人联系电话）     */
    Mobile  *string `json:"mobile,omitempty" required:"false" `
    /*
        买家姓名（自提方式填写提货人姓名）     */
    BuyerName  *string `json:"buyer_name" required:"true" `
    /*
        身份证号（自提方式必填，填写提货人身份证号码，提货时用于确认提货人身份）     */
    IdCardNumber  *string `json:"id_card_number,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetOrderDetail(v []string) *TaobaoFenxiaoDealerRequisitionorderCreateRequest {
    s.OrderDetail = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetLogisticsType(v string) *TaobaoFenxiaoDealerRequisitionorderCreateRequest {
    s.LogisticsType = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetProvince(v string) *TaobaoFenxiaoDealerRequisitionorderCreateRequest {
    s.Province = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetCity(v string) *TaobaoFenxiaoDealerRequisitionorderCreateRequest {
    s.City = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetDistrict(v string) *TaobaoFenxiaoDealerRequisitionorderCreateRequest {
    s.District = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetAddress(v string) *TaobaoFenxiaoDealerRequisitionorderCreateRequest {
    s.Address = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetPostCode(v string) *TaobaoFenxiaoDealerRequisitionorderCreateRequest {
    s.PostCode = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetPhone(v string) *TaobaoFenxiaoDealerRequisitionorderCreateRequest {
    s.Phone = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetMobile(v string) *TaobaoFenxiaoDealerRequisitionorderCreateRequest {
    s.Mobile = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetBuyerName(v string) *TaobaoFenxiaoDealerRequisitionorderCreateRequest {
    s.BuyerName = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderCreateRequest) SetIdCardNumber(v string) *TaobaoFenxiaoDealerRequisitionorderCreateRequest {
    s.IdCardNumber = &v
    return s
}

func (req *TaobaoFenxiaoDealerRequisitionorderCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.OrderDetail != nil) {
        paramMap["order_detail"] = util.ConvertBasicList(*req.OrderDetail)
    }
    if(req.LogisticsType != nil) {
        paramMap["logistics_type"] = *req.LogisticsType
    }
    if(req.Province != nil) {
        paramMap["province"] = *req.Province
    }
    if(req.City != nil) {
        paramMap["city"] = *req.City
    }
    if(req.District != nil) {
        paramMap["district"] = *req.District
    }
    if(req.Address != nil) {
        paramMap["address"] = *req.Address
    }
    if(req.PostCode != nil) {
        paramMap["post_code"] = *req.PostCode
    }
    if(req.Phone != nil) {
        paramMap["phone"] = *req.Phone
    }
    if(req.Mobile != nil) {
        paramMap["mobile"] = *req.Mobile
    }
    if(req.BuyerName != nil) {
        paramMap["buyer_name"] = *req.BuyerName
    }
    if(req.IdCardNumber != nil) {
        paramMap["id_card_number"] = *req.IdCardNumber
    }
    return paramMap
}

func (req *TaobaoFenxiaoDealerRequisitionorderCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}