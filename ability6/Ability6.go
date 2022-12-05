package ability6

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability6/request"
    "topsdk/ability6/response"
    "topsdk/util"
)

type Ability6 struct {
    Client *topsdk.TopClient
}

func NewAbility6(client *topsdk.TopClient) *Ability6{
    return &Ability6{client}
}

/*
    天猫逆向纠纷查询
*/
func (ability *Ability6) TmallDisputeReceiveGet(req *request.TmallDisputeReceiveGetRequest,session string) (*response.TmallDisputeReceiveGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability6 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.dispute.receive.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallDisputeReceiveGetResponse{}
    if(err != nil){
        log.Println("tmallDisputeReceiveGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    审核退款单
*/
func (ability *Ability6) TaobaoRpRefundReview(req *request.TaobaoRpRefundReviewRequest,session string) (*response.TaobaoRpRefundReviewResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability6 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.rp.refund.review",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRpRefundReviewResponse{}
    if(err != nil){
        log.Println("taobaoRpRefundReview error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    卖家回填物流信息
*/
func (ability *Ability6) TaobaoRpReturngoodsRefill(req *request.TaobaoRpReturngoodsRefillRequest,session string) (*response.TaobaoRpReturngoodsRefillResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability6 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.rp.returngoods.refill",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRpReturngoodsRefillResponse{}
    if(err != nil){
        log.Println("taobaoRpReturngoodsRefill error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    卖家拒绝退货
*/
func (ability *Ability6) TaobaoRpReturngoodsRefuse(req *request.TaobaoRpReturngoodsRefuseRequest,session string) (*response.TaobaoRpReturngoodsRefuseResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability6 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.rp.returngoods.refuse",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRpReturngoodsRefuseResponse{}
    if(err != nil){
        log.Println("taobaoRpReturngoodsRefuse error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取拒绝原因列表
*/
func (ability *Ability6) TaobaoRefundRefusereasonGet(req *request.TaobaoRefundRefusereasonGetRequest,session string) (*response.TaobaoRefundRefusereasonGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability6 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refund.refusereason.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundRefusereasonGetResponse{}
    if(err != nil){
        log.Println("taobaoRefundRefusereasonGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    特殊部分退纠纷单查询
*/
func (ability *Ability6) TaobaoSpecialRefundGet(req *request.TaobaoSpecialRefundGetRequest,session string) (*response.TaobaoSpecialRefundGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability6 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.special.refund.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSpecialRefundGetResponse{}
    if(err != nil){
        log.Println("taobaoSpecialRefundGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    特殊退款类型的纠纷单列表查询
*/
func (ability *Ability6) TaobaoSpecialRefundsReceiveGet(req *request.TaobaoSpecialRefundsReceiveGetRequest,session string) (*response.TaobaoSpecialRefundsReceiveGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability6 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.special.refunds.receive.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSpecialRefundsReceiveGetResponse{}
    if(err != nil){
        log.Println("taobaoSpecialRefundsReceiveGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
