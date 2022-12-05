package ability250

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability250/request"
    "topsdk/ability250/response"
    "topsdk/util"
)

type Ability250 struct {
    Client *topsdk.TopClient
}

func NewAbility250(client *topsdk.TopClient) *Ability250{
    return &Ability250{client}
}

/*
    卖家店铺基础信息查询
*/
func (ability *Ability250) TaobaoShopSellerGet(req *request.TaobaoShopSellerGetRequest,session string) (*response.TaobaoShopSellerGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability250 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.shop.seller.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoShopSellerGetResponse{}
    if(err != nil){
        log.Println("taobaoShopSellerGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
