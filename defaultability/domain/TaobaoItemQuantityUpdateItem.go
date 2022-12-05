package domain

import (
        "topsdk/util"
    )

type TaobaoItemQuantityUpdateItem struct {
    /*
        商品id(注意：iid近期即将废弃，请用num_iid参数)     */
    Iid  *string `json:"iid,omitempty" `

    /*
        商品数字id     */
    NumIid  *int64 `json:"num_iid,omitempty" `

    /*
        商品修改时间（格式：yyyy-MM-dd HH:mm:ss）     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        商品数量     */
    Num  *int64 `json:"num,omitempty" `

    /*
        Sku列表。fields中只设置sku可以返回Sku结构体中所有字段，如果设置为sku.sku_id、sku.properties、sku.quantity等形式就只会返回相应的字段     */
    Skus  *[]TaobaoItemQuantityUpdateSku `json:"skus,omitempty" `

}

func (s *TaobaoItemQuantityUpdateItem) SetIid(v string) *TaobaoItemQuantityUpdateItem {
    s.Iid = &v
    return s
}
func (s *TaobaoItemQuantityUpdateItem) SetNumIid(v int64) *TaobaoItemQuantityUpdateItem {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemQuantityUpdateItem) SetModified(v util.LocalTime) *TaobaoItemQuantityUpdateItem {
    s.Modified = &v
    return s
}
func (s *TaobaoItemQuantityUpdateItem) SetNum(v int64) *TaobaoItemQuantityUpdateItem {
    s.Num = &v
    return s
}
func (s *TaobaoItemQuantityUpdateItem) SetSkus(v []TaobaoItemQuantityUpdateSku) *TaobaoItemQuantityUpdateItem {
    s.Skus = &v
    return s
}
