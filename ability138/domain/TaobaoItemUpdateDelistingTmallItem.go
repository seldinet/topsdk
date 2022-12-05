package domain

import (
        "topsdk/util"
    )

type TaobaoItemUpdateDelistingTmallItem struct {
    /*
        商品id(注意：iid近期即将废弃，请用num_iid参数)     */
    Iid  *string `json:"iid,omitempty" `

    /*
        商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        商品修改时间（格式：yyyy-MM-dd HH:mm:ss）     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoItemUpdateDelistingTmallItem) SetIid(v string) *TaobaoItemUpdateDelistingTmallItem {
    s.Iid = &v
    return s
}
func (s *TaobaoItemUpdateDelistingTmallItem) SetNumIid(v int64) *TaobaoItemUpdateDelistingTmallItem {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemUpdateDelistingTmallItem) SetModified(v util.LocalTime) *TaobaoItemUpdateDelistingTmallItem {
    s.Modified = &v
    return s
}
