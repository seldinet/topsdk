package domain

import (
        "topsdk/util"
    )

type TaobaoItemBarcodeUpdateItem struct {
    /*
        商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        商品修改时间（格式：yyyy-MM-dd HH:mm:ss）     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoItemBarcodeUpdateItem) SetNumIid(v int64) *TaobaoItemBarcodeUpdateItem {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemBarcodeUpdateItem) SetModified(v util.LocalTime) *TaobaoItemBarcodeUpdateItem {
    s.Modified = &v
    return s
}
