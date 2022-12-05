package request


type TaobaoRpReturngoodsAgreeRequest struct {
    /*
        退款编号     */
    RefundId  *int64 `json:"refund_id" required:"true" `
    /*
        售中：onsale，售后：aftersale，天猫退款为必填项。     */
    RefundPhase  *string `json:"refund_phase,omitempty" required:"false" `
    /*
        退款版本号，天猫退款为必填项。     */
    RefundVersion  *int64 `json:"refund_version,omitempty" required:"false" `
    /*
        卖家退货留言，天猫退款为必填项。     */
    Remark  *string `json:"remark,omitempty" required:"false" `
    /*
        卖家提供的退货地址，淘宝退款为必填项。     */
    Address  *string `json:"address,omitempty" required:"false" `
    /*
        卖家收货地址编号，天猫淘宝退款都为必填项。     */
    SellerAddressId  *int64 `json:"seller_address_id,omitempty" required:"false" `
    /*
        卖家手机，淘宝退款为必填项。     */
    Mobile  *string `json:"mobile,omitempty" required:"false" `
    /*
        卖家提供的退货地址的邮编，淘宝退款为必填项。     */
    Post  *string `json:"post,omitempty" required:"false" `
    /*
        卖家姓名，淘宝退款为必填项。     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        卖家座机，淘宝退款为必填项。     */
    Tel  *string `json:"tel,omitempty" required:"false" `
    /*
        邮费承担方，买家承担值为1，卖家承担值为0     */
    PostFeeBearRole  *int64 `json:"post_fee_bear_role,omitempty" required:"false" `
    /*
        是否虚拟退货，可选项 defalutValue��false    */
    VirtualReturnGoods  *bool `json:"virtual_return_goods,omitempty" required:"false" `
}

func (s *TaobaoRpReturngoodsAgreeRequest) SetRefundId(v int64) *TaobaoRpReturngoodsAgreeRequest {
    s.RefundId = &v
    return s
}
func (s *TaobaoRpReturngoodsAgreeRequest) SetRefundPhase(v string) *TaobaoRpReturngoodsAgreeRequest {
    s.RefundPhase = &v
    return s
}
func (s *TaobaoRpReturngoodsAgreeRequest) SetRefundVersion(v int64) *TaobaoRpReturngoodsAgreeRequest {
    s.RefundVersion = &v
    return s
}
func (s *TaobaoRpReturngoodsAgreeRequest) SetRemark(v string) *TaobaoRpReturngoodsAgreeRequest {
    s.Remark = &v
    return s
}
func (s *TaobaoRpReturngoodsAgreeRequest) SetAddress(v string) *TaobaoRpReturngoodsAgreeRequest {
    s.Address = &v
    return s
}
func (s *TaobaoRpReturngoodsAgreeRequest) SetSellerAddressId(v int64) *TaobaoRpReturngoodsAgreeRequest {
    s.SellerAddressId = &v
    return s
}
func (s *TaobaoRpReturngoodsAgreeRequest) SetMobile(v string) *TaobaoRpReturngoodsAgreeRequest {
    s.Mobile = &v
    return s
}
func (s *TaobaoRpReturngoodsAgreeRequest) SetPost(v string) *TaobaoRpReturngoodsAgreeRequest {
    s.Post = &v
    return s
}
func (s *TaobaoRpReturngoodsAgreeRequest) SetName(v string) *TaobaoRpReturngoodsAgreeRequest {
    s.Name = &v
    return s
}
func (s *TaobaoRpReturngoodsAgreeRequest) SetTel(v string) *TaobaoRpReturngoodsAgreeRequest {
    s.Tel = &v
    return s
}
func (s *TaobaoRpReturngoodsAgreeRequest) SetPostFeeBearRole(v int64) *TaobaoRpReturngoodsAgreeRequest {
    s.PostFeeBearRole = &v
    return s
}
func (s *TaobaoRpReturngoodsAgreeRequest) SetVirtualReturnGoods(v bool) *TaobaoRpReturngoodsAgreeRequest {
    s.VirtualReturnGoods = &v
    return s
}

func (req *TaobaoRpReturngoodsAgreeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.RefundId != nil) {
        paramMap["refund_id"] = *req.RefundId
    }
    if(req.RefundPhase != nil) {
        paramMap["refund_phase"] = *req.RefundPhase
    }
    if(req.RefundVersion != nil) {
        paramMap["refund_version"] = *req.RefundVersion
    }
    if(req.Remark != nil) {
        paramMap["remark"] = *req.Remark
    }
    if(req.Address != nil) {
        paramMap["address"] = *req.Address
    }
    if(req.SellerAddressId != nil) {
        paramMap["seller_address_id"] = *req.SellerAddressId
    }
    if(req.Mobile != nil) {
        paramMap["mobile"] = *req.Mobile
    }
    if(req.Post != nil) {
        paramMap["post"] = *req.Post
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.Tel != nil) {
        paramMap["tel"] = *req.Tel
    }
    if(req.PostFeeBearRole != nil) {
        paramMap["post_fee_bear_role"] = *req.PostFeeBearRole
    }
    if(req.VirtualReturnGoods != nil) {
        paramMap["virtual_return_goods"] = *req.VirtualReturnGoods
    }
    return paramMap
}

func (req *TaobaoRpReturngoodsAgreeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}