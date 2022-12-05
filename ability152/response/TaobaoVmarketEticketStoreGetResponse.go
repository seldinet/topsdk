package response

import (
)

type TaobaoVmarketEticketStoreGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商户id
    */
    StoreId  int64 `json:"store_id,omitempty" `
    /*
        商户地址
    */
    Address  string `json:"address,omitempty" `
    /*
        商户名称
    */
    Name  string `json:"name,omitempty" `
    /*
        区
    */
    District  string `json:"district,omitempty" `
    /*
        所在城市
    */
    City  string `json:"city,omitempty" `
    /*
        省份
    */
    Province  string `json:"province,omitempty" `
    /*
        联系电话
    */
    Contract  string `json:"contract,omitempty" `
    /*
        自有卖家导入门店的时候，可以把自己系统门店信息的主键或者唯一key传入，用于快速匹配
    */
    Selfcode  string `json:"selfcode,omitempty" `
}
