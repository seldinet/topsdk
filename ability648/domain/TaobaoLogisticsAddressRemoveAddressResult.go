package domain

import (
        "topsdk/util"
    )

type TaobaoLogisticsAddressRemoveAddressResult struct {
    /*
        修改日期时间     */
    ModifyDate  *util.LocalTime `json:"modify_date,omitempty" `

}

func (s *TaobaoLogisticsAddressRemoveAddressResult) SetModifyDate(v util.LocalTime) *TaobaoLogisticsAddressRemoveAddressResult {
    s.ModifyDate = &v
    return s
}
