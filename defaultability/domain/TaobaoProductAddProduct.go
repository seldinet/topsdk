package domain

import (
        "topsdk/util"
    )

type TaobaoProductAddProduct struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        创建时间.格式:yyyy-mm-dd hh:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

}

func (s *TaobaoProductAddProduct) SetProductId(v int64) *TaobaoProductAddProduct {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductAddProduct) SetCreated(v util.LocalTime) *TaobaoProductAddProduct {
    s.Created = &v
    return s
}
