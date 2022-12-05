package domain


type TaobaoInventoryWarehouseManageResult struct {
    /*
        data     */
    Data  *bool `json:"data,omitempty" `

    /*
        errorCode     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        errorMsg     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoInventoryWarehouseManageResult) SetData(v bool) *TaobaoInventoryWarehouseManageResult {
    s.Data = &v
    return s
}
func (s *TaobaoInventoryWarehouseManageResult) SetErrorCode(v string) *TaobaoInventoryWarehouseManageResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoInventoryWarehouseManageResult) SetErrorMsg(v string) *TaobaoInventoryWarehouseManageResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoInventoryWarehouseManageResult) SetSuccess(v bool) *TaobaoInventoryWarehouseManageResult {
    s.Success = &v
    return s
}
