package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoLoginUserGetLoginUser struct {
    /*
        分销用户ID     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        会员NICK     */
    Nick  *string `json:"nick,omitempty" `

    /*
        分销用户类型(1:分销商，2:供应商，3:品牌商；2、3都表示有供货能力，只是身份不同)     */
    UserType  *string `json:"user_type,omitempty" `

    /*
        入驻时间     */
    CreateTime  *util.LocalTime `json:"create_time,omitempty" `

}

func (s *TaobaoFenxiaoLoginUserGetLoginUser) SetUserId(v int64) *TaobaoFenxiaoLoginUserGetLoginUser {
    s.UserId = &v
    return s
}
func (s *TaobaoFenxiaoLoginUserGetLoginUser) SetNick(v string) *TaobaoFenxiaoLoginUserGetLoginUser {
    s.Nick = &v
    return s
}
func (s *TaobaoFenxiaoLoginUserGetLoginUser) SetUserType(v string) *TaobaoFenxiaoLoginUserGetLoginUser {
    s.UserType = &v
    return s
}
func (s *TaobaoFenxiaoLoginUserGetLoginUser) SetCreateTime(v util.LocalTime) *TaobaoFenxiaoLoginUserGetLoginUser {
    s.CreateTime = &v
    return s
}
