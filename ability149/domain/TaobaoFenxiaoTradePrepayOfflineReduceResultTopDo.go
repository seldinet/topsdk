package domain


type TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo struct {
    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo) SetSuccess(v bool) *TaobaoFenxiaoTradePrepayOfflineReduceResultTopDo {
    s.Success = &v
    return s
}
