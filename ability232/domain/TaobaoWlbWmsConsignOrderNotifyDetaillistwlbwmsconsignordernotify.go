package domain


type TaobaoWlbWmsConsignOrderNotifyDetaillistwlbwmsconsignordernotify struct {
    /*
        发票信息     */
    ItemDetail  *TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify `json:"item_detail,omitempty" `

}

func (s *TaobaoWlbWmsConsignOrderNotifyDetaillistwlbwmsconsignordernotify) SetItemDetail(v TaobaoWlbWmsConsignOrderNotifyItemdetailwlbwmsconsignordernotify) *TaobaoWlbWmsConsignOrderNotifyDetaillistwlbwmsconsignordernotify {
    s.ItemDetail = &v
    return s
}
