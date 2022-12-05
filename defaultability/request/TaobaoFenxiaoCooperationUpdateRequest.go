package request


type TaobaoFenxiaoCooperationUpdateRequest struct {
    /*
        分销商ID     */
    DistributorId  *int64 `json:"distributor_id" required:"true" `
    /*
        等级ID(0代表取消)     */
    GradeId  *int64 `json:"grade_id" required:"true" `
    /*
        分销方式(新增)： AGENT(代销)、DEALER(经销)(默认为代销)     */
    TradeType  *string `json:"trade_type,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoCooperationUpdateRequest) SetDistributorId(v int64) *TaobaoFenxiaoCooperationUpdateRequest {
    s.DistributorId = &v
    return s
}
func (s *TaobaoFenxiaoCooperationUpdateRequest) SetGradeId(v int64) *TaobaoFenxiaoCooperationUpdateRequest {
    s.GradeId = &v
    return s
}
func (s *TaobaoFenxiaoCooperationUpdateRequest) SetTradeType(v string) *TaobaoFenxiaoCooperationUpdateRequest {
    s.TradeType = &v
    return s
}

func (req *TaobaoFenxiaoCooperationUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DistributorId != nil) {
        paramMap["distributor_id"] = *req.DistributorId
    }
    if(req.GradeId != nil) {
        paramMap["grade_id"] = *req.GradeId
    }
    if(req.TradeType != nil) {
        paramMap["trade_type"] = *req.TradeType
    }
    return paramMap
}

func (req *TaobaoFenxiaoCooperationUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}