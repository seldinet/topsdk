package domain

import (
        "topsdk/util"
    )

type TaobaoLogisticsAddressModifyAddressResult struct {
    /*
        修改日期时间     */
    ModifyDate  *util.LocalTime `json:"modify_date,omitempty" `

}

func (s *TaobaoLogisticsAddressModifyAddressResult) SetModifyDate(v util.LocalTime) *TaobaoLogisticsAddressModifyAddressResult {
    s.ModifyDate = &v
    return s
}
