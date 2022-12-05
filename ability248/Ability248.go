package ability248

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability248/request"
    "topsdk/ability248/response"
    "topsdk/util"
)

type Ability248 struct {
    Client *topsdk.TopClient
}

func NewAbility248(client *topsdk.TopClient) *Ability248{
    return &Ability248{client}
}

/*
    获取文件上传token
*/
func (ability *Ability248) TaobaoPictureToken(req *request.TaobaoPictureTokenRequest,session string) (*response.TaobaoPictureTokenResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability248 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.token",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureTokenResponse{}
    if(err != nil){
        log.Println("taobaoPictureToken error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
