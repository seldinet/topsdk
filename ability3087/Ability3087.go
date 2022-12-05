package ability3087

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability3087/request"
    "topsdk/ability3087/response"
    "topsdk/util"
)

type Ability3087 struct {
    Client *topsdk.TopClient
}

func NewAbility3087(client *topsdk.TopClient) *Ability3087{
    return &Ability3087{client}
}

/*
    网关一次性token获取
*/
func (ability *Ability3087) TaobaoTopOnceTokenGet(req *request.TaobaoTopOnceTokenGetRequest,session string) (*response.TaobaoTopOnceTokenGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability3087 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.top.once.token.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTopOnceTokenGetResponse{}
    if(err != nil){
        log.Println("taobaoTopOnceTokenGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
