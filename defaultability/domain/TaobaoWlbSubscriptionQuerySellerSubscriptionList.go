package domain

import (
        "topsdk/util"
    )

type TaobaoWlbSubscriptionQuerySellerSubscriptionList struct {
    /*
        定购ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        定购用户ID     */
    SubscriberUserId  *int64 `json:"subscriber_user_id,omitempty" `

    /*
        定购用户NICK     */
    SubscriberUserNick  *string `json:"subscriber_user_nick,omitempty" `

    /*
        服务商ID     */
    ProviderUserId  *int64 `json:"provider_user_id,omitempty" `

    /*
        服务ID     */
    ServiceId  *int64 `json:"service_id,omitempty" `

    /*
        定购有效开始日期     */
    StartDate  *util.LocalTime `json:"start_date,omitempty" `

    /*
        定购有效结束日期     */
    EndDate  *util.LocalTime `json:"end_date,omitempty" `

    /*
        服务编码     */
    ServiceCode  *string `json:"service_code,omitempty" `

    /*
        服务名     */
    ServiceName  *string `json:"service_name,omitempty" `

    /*
        服务类型，
STORE 1-仓储、
TMS 2-TMS、
PACKAGE 3-包装服务
SUPPLIER 4-供货
INSTALL 5-安装
COMPLEX_SERVICE 100-综合服务     */
    ServiceType  *string `json:"service_type,omitempty" `

    /*
        父定购ID
可通过该字段来得之服务上下级关系。
例定购了仓储服务，下又有TMS服务。
该字段保存仓储服务ID。     */
    ParentId  *int64 `json:"parent_id,omitempty" `

    /*
        状态
AUDITING 1-待审核
CANCEL 2-撤销
CHECKED 3-审核通过
FAILED 4-审核未通过
SYNCHRONIZING 5-同步中     */
    Status  *string `json:"status,omitempty" `

    /*
        创建时间     */
    GmtCreate  *util.LocalTime `json:"gmt_create,omitempty" `

    /*
        修改时间     */
    GmtModified  *util.LocalTime `json:"gmt_modified,omitempty" `

    /*
        备注     */
    Remark  *string `json:"remark,omitempty" `

    /*
        联系人地址信息     */
    WlbPartnerAddress  *TaobaoWlbSubscriptionQueryAddressInfo `json:"wlb_partner_address,omitempty" `

    /*
        联系人联系详情     */
    WlbPartnerContact  *TaobaoWlbSubscriptionQueryContactInfo `json:"wlb_partner_contact,omitempty" `

    /*
        判断该仓库是否是实体仓，还是虚拟仓，null是实体仓，10:代表虚拟仓     */
    IsOwnService  *int64 `json:"is_own_service,omitempty" `

    /*
        自有仓的别名，当当前订购记录为自有仓时才会有值     */
    ServiceAlias  *string `json:"service_alias,omitempty" `

    /*
        openuid     */
    Openuid  *string `json:"openuid,omitempty" `

}

func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetId(v int64) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.Id = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetSubscriberUserId(v int64) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.SubscriberUserId = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetSubscriberUserNick(v string) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.SubscriberUserNick = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetProviderUserId(v int64) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.ProviderUserId = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetServiceId(v int64) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.ServiceId = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetStartDate(v util.LocalTime) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.StartDate = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetEndDate(v util.LocalTime) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.EndDate = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetServiceCode(v string) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.ServiceCode = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetServiceName(v string) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.ServiceName = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetServiceType(v string) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.ServiceType = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetParentId(v int64) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.ParentId = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetStatus(v string) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.Status = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetGmtCreate(v util.LocalTime) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetGmtModified(v util.LocalTime) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.GmtModified = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetRemark(v string) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetWlbPartnerAddress(v TaobaoWlbSubscriptionQueryAddressInfo) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.WlbPartnerAddress = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetWlbPartnerContact(v TaobaoWlbSubscriptionQueryContactInfo) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.WlbPartnerContact = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetIsOwnService(v int64) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.IsOwnService = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetServiceAlias(v string) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.ServiceAlias = &v
    return s
}
func (s *TaobaoWlbSubscriptionQuerySellerSubscriptionList) SetOpenuid(v string) *TaobaoWlbSubscriptionQuerySellerSubscriptionList {
    s.Openuid = &v
    return s
}
