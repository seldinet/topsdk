package ability181

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability181/request"
    "topsdk/ability181/response"
    "topsdk/util"
)

type Ability181 struct {
    Client *topsdk.TopClient
}

func NewAbility181(client *topsdk.TopClient) *Ability181{
    return &Ability181{client}
}

/*
    查询商品类目属性变更
*/
func (ability *Ability181) TaobaoItemCatpropsModificationGet(req *request.TaobaoItemCatpropsModificationGetRequest,session string) (*response.TaobaoItemCatpropsModificationGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability181 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.catprops.modification.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemCatpropsModificationGetResponse{}
    if(err != nil){
        log.Println("taobaoItemCatpropsModificationGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
