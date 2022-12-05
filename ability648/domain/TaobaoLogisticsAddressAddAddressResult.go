package domain

import (
        "topsdk/util"
    )

type TaobaoLogisticsAddressAddAddressResult struct {
    /*
        修改日期时间     */
    ModifyDate  *util.LocalTime `json:"modify_date,omitempty" `

}

func (s *TaobaoLogisticsAddressAddAddressResult) SetModifyDate(v util.LocalTime) *TaobaoLogisticsAddressAddAddressResult {
    s.ModifyDate = &v
    return s
}
