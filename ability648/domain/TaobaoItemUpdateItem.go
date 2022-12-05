package domain

import (
        "topsdk/util"
    )

type TaobaoItemUpdateItem struct {
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

func (s *TaobaoItemUpdateItem) SetIid(v string) *TaobaoItemUpdateItem {
    s.Iid = &v
    return s
}
func (s *TaobaoItemUpdateItem) SetNumIid(v int64) *TaobaoItemUpdateItem {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemUpdateItem) SetModified(v util.LocalTime) *TaobaoItemUpdateItem {
    s.Modified = &v
    return s
}
