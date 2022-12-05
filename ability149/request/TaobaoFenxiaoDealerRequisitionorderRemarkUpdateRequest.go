package request


type TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest struct {
    /*
        经销采购单ID     */
    DealerOrderId  *int64 `json:"dealer_order_id" required:"true" `
    /*
        备注留言，可为空     */
    SupplierMemo  *string `json:"supplier_memo,omitempty" required:"false" `
    /*
        旗子的标记，必选。
1-5之间的数字。
非1-5之间，都采用1作为默认。
1:红色
2:黄色
3:绿色
4:蓝色
5:粉红色     */
    SupplierMemoFlag  *int64 `json:"supplier_memo_flag" required:"true" `
}

func (s *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) SetDealerOrderId(v int64) *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest {
    s.DealerOrderId = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) SetSupplierMemo(v string) *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest {
    s.SupplierMemo = &v
    return s
}
func (s *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) SetSupplierMemoFlag(v int64) *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest {
    s.SupplierMemoFlag = &v
    return s
}

func (req *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DealerOrderId != nil) {
        paramMap["dealer_order_id"] = *req.DealerOrderId
    }
    if(req.SupplierMemo != nil) {
        paramMap["supplier_memo"] = *req.SupplierMemo
    }
    if(req.SupplierMemoFlag != nil) {
        paramMap["supplier_memo_flag"] = *req.SupplierMemoFlag
    }
    return paramMap
}

func (req *TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}