package domain

import (
        "topsdk/util"
    )

type TaobaoProductUpdateProduct struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        修改时间.格式:yyyy-mm-dd hh:mm:ss     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoProductUpdateProduct) SetProductId(v int64) *TaobaoProductUpdateProduct {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductUpdateProduct) SetModified(v util.LocalTime) *TaobaoProductUpdateProduct {
    s.Modified = &v
    return s
}
