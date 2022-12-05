package domain

import (
        "topsdk/util"
    )

type TaobaoItemUpdateListingTmallItem struct {
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

func (s *TaobaoItemUpdateListingTmallItem) SetIid(v string) *TaobaoItemUpdateListingTmallItem {
    s.Iid = &v
    return s
}
func (s *TaobaoItemUpdateListingTmallItem) SetNumIid(v int64) *TaobaoItemUpdateListingTmallItem {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemUpdateListingTmallItem) SetModified(v util.LocalTime) *TaobaoItemUpdateListingTmallItem {
    s.Modified = &v
    return s
}
