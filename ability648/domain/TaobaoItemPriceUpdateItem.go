package domain

import (
        "topsdk/util"
    )

type TaobaoItemPriceUpdateItem struct {
    /*
        商品iid     */
    Iid  *string `json:"iid,omitempty" `

    /*
        商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        商品修改时间（格式：yyyy-MM-dd HH:mm:ss）     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoItemPriceUpdateItem) SetIid(v string) *TaobaoItemPriceUpdateItem {
    s.Iid = &v
    return s
}
func (s *TaobaoItemPriceUpdateItem) SetNumIid(v int64) *TaobaoItemPriceUpdateItem {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemPriceUpdateItem) SetModified(v util.LocalTime) *TaobaoItemPriceUpdateItem {
    s.Modified = &v
    return s
}
