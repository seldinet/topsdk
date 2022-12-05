package domain


type TaobaoFenxiaoTradePrepayOfflineAddResultTopDo struct {
    /*
        success     */
    Success  *bool `json:"success,omitempty" `

}

func (s *TaobaoFenxiaoTradePrepayOfflineAddResultTopDo) SetSuccess(v bool) *TaobaoFenxiaoTradePrepayOfflineAddResultTopDo {
    s.Success = &v
    return s
}
