package response

import (
    "topsdk/ability232/domain"
)

type TaobaoWlbOrderJzpartnerQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        安装服务商列表
    */
    InstallList  []domain.TaobaoWlbOrderJzpartnerQueryPartnerNew `json:"install_list,omitempty" `
    /*
        物流配送服务商对象列表
    */
    ServerList  []domain.TaobaoWlbOrderJzpartnerQueryPartnerNew `json:"server_list,omitempty" `
    /*
        接口查询成功或者失败
    */
    IsSuccess  bool `json:"is_success,omitempty" `
    /*
        查询返回信息，如果失败，存储错误信息
    */
    ResultInfo  string `json:"result_info,omitempty" `
}
