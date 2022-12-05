package domain

import (
        "topsdk/util"
    )

type TaobaoItemAddItem struct {
    /*
        商品iid     */
    Iid  *string `json:"iid,omitempty" `

    /*
        商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        Item的发布时间，目前仅供taobao.item.add和taobao.item.get可用     */
    Created  *util.LocalTime `json:"created,omitempty" `

}

func (s *TaobaoItemAddItem) SetIid(v string) *TaobaoItemAddItem {
    s.Iid = &v
    return s
}
func (s *TaobaoItemAddItem) SetNumIid(v int64) *TaobaoItemAddItem {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemAddItem) SetCreated(v util.LocalTime) *TaobaoItemAddItem {
    s.Created = &v
    return s
}
