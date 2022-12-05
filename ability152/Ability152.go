package ability152

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability152/request"
    "topsdk/ability152/response"
    "topsdk/util"
)

type Ability152 struct {
    Client *topsdk.TopClient
}

func NewAbility152(client *topsdk.TopClient) *Ability152{
    return &Ability152{client}
}

/*
    任务列表获取接口
*/
func (ability *Ability152) TaobaoVmarketEticketTasksGet(req *request.TaobaoVmarketEticketTasksGetRequest,session string) (*response.TaobaoVmarketEticketTasksGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.tasks.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketTasksGetResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketTasksGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    主动发起通知接口
*/
func (ability *Ability152) TaobaoVmarketEticketManageNotify(req *request.TaobaoVmarketEticketManageNotifyRequest,session string) (*response.TaobaoVmarketEticketManageNotifyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.manage.notify",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketManageNotifyResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketManageNotify error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取电子凭证预约门店信息
*/
func (ability *Ability152) TaobaoVmarketEticketStoreGet(req *request.TaobaoVmarketEticketStoreGetRequest,session string) (*response.TaobaoVmarketEticketStoreGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.store.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketStoreGetResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketStoreGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    码商二维码图片上传
*/
func (ability *Ability152) TaobaoVmarketEticketQrcodeUpload(req *request.TaobaoVmarketEticketQrcodeUploadRequest,session string) (*response.TaobaoVmarketEticketQrcodeUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.qrcode.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketQrcodeUploadResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketQrcodeUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子凭证冲正接口
*/
func (ability *Ability152) TaobaoEticketMerchantMaReverse(req *request.TaobaoEticketMerchantMaReverseRequest,session string) (*response.TaobaoEticketMerchantMaReverseResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.eticket.merchant.ma.reverse",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoEticketMerchantMaReverseResponse{}
    if(err != nil){
        log.Println("taobaoEticketMerchantMaReverse error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    码商上传二维码图片
*/
func (ability *Ability152) TaobaoEticketMerchantImgUpload(req *request.TaobaoEticketMerchantImgUploadRequest,session string) (*response.TaobaoEticketMerchantImgUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.eticket.merchant.img.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoEticketMerchantImgUploadResponse{}
    if(err != nil){
        log.Println("taobaoEticketMerchantImgUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子凭证核销前校验接口
*/
func (ability *Ability152) TaobaoEticketMerchantMaAvailable(req *request.TaobaoEticketMerchantMaAvailableRequest,session string) (*response.TaobaoEticketMerchantMaAvailableResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.eticket.merchant.ma.available",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoEticketMerchantMaAvailableResponse{}
    if(err != nil){
        log.Println("taobaoEticketMerchantMaAvailable error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子凭证核销接口
*/
func (ability *Ability152) TaobaoEticketMerchantMaConsume(req *request.TaobaoEticketMerchantMaConsumeRequest,session string) (*response.TaobaoEticketMerchantMaConsumeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.eticket.merchant.ma.consume",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoEticketMerchantMaConsumeResponse{}
    if(err != nil){
        log.Println("taobaoEticketMerchantMaConsume error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子凭证重发回调接口
*/
func (ability *Ability152) TaobaoEticketMerchantMaResend(req *request.TaobaoEticketMerchantMaResendRequest,session string) (*response.TaobaoEticketMerchantMaResendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.eticket.merchant.ma.resend",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoEticketMerchantMaResendResponse{}
    if(err != nil){
        log.Println("taobaoEticketMerchantMaResend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    码商发码失败回调接口
*/
func (ability *Ability152) TaobaoEticketMerchantMaFailsend(req *request.TaobaoEticketMerchantMaFailsendRequest,session string) (*response.TaobaoEticketMerchantMaFailsendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.eticket.merchant.ma.failsend",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoEticketMerchantMaFailsendResponse{}
    if(err != nil){
        log.Println("taobaoEticketMerchantMaFailsend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    码商发码成功回调接口
*/
func (ability *Ability152) TaobaoEticketMerchantMaSend(req *request.TaobaoEticketMerchantMaSendRequest,session string) (*response.TaobaoEticketMerchantMaSendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.eticket.merchant.ma.send",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoEticketMerchantMaSendResponse{}
    if(err != nil){
        log.Println("taobaoEticketMerchantMaSend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    业务重新触发发码短信
*/
func (ability *Ability152) TaobaoVmarketEticketFlowResend(req *request.TaobaoVmarketEticketFlowResendRequest,session string) (*response.TaobaoVmarketEticketFlowResendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.flow.resend",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketFlowResendResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketFlowResend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    无交易类凭证核销
*/
func (ability *Ability152) TaobaoVmarketEticketFlowConsume(req *request.TaobaoVmarketEticketFlowConsumeRequest,session string) (*response.TaobaoVmarketEticketFlowConsumeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.flow.consume",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketFlowConsumeResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketFlowConsume error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    核销放行的查询接口
*/
func (ability *Ability152) TaobaoVmarketEticketAuthBeforeconsume(req *request.TaobaoVmarketEticketAuthBeforeconsumeRequest,session string) (*response.TaobaoVmarketEticketAuthBeforeconsumeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.auth.beforeconsume",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketAuthBeforeconsumeResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketAuthBeforeconsume error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    核销放行的核销接口
*/
func (ability *Ability152) TaobaoVmarketEticketAuthConsume(req *request.TaobaoVmarketEticketAuthConsumeRequest,session string) (*response.TaobaoVmarketEticketAuthConsumeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.auth.consume",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketAuthConsumeResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketAuthConsume error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    无法发码回调
*/
func (ability *Ability152) TaobaoVmarketEticketFailsend(req *request.TaobaoVmarketEticketFailsendRequest,session string) (*response.TaobaoVmarketEticketFailsendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.failsend",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketFailsendResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketFailsend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    订单延时接口
*/
func (ability *Ability152) TaobaoVmarketEticketTimeExpand(req *request.TaobaoVmarketEticketTimeExpandRequest) (*response.TaobaoVmarketEticketTimeExpandResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.vmarket.eticket.time.expand",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoVmarketEticketTimeExpandResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketTimeExpand error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子凭证验码前置确认
*/
func (ability *Ability152) TaobaoVmarketEticketBeforeconsume(req *request.TaobaoVmarketEticketBeforeconsumeRequest,session string) (*response.TaobaoVmarketEticketBeforeconsumeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.beforeconsume",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketBeforeconsumeResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketBeforeconsume error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子票券消费通知
*/
func (ability *Ability152) TaobaoVmarketEticketConsume(req *request.TaobaoVmarketEticketConsumeRequest,session string) (*response.TaobaoVmarketEticketConsumeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.consume",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketConsumeResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketConsume error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    外部合作商家重发电子凭证回调接口
*/
func (ability *Ability152) TaobaoVmarketEticketResend(req *request.TaobaoVmarketEticketResendRequest,session string) (*response.TaobaoVmarketEticketResendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.resend",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketResendResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketResend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商家电子凭证发码成功回调接口
*/
func (ability *Ability152) TaobaoVmarketEticketSend(req *request.TaobaoVmarketEticketSendRequest,session string) (*response.TaobaoVmarketEticketSendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.send",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketSendResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketSend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子凭证冲正接口
*/
func (ability *Ability152) TaobaoVmarketEticketReverse(req *request.TaobaoVmarketEticketReverseRequest,session string) (*response.TaobaoVmarketEticketReverseResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.reverse",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketReverseResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketReverse error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子凭证码列表查询
*/
func (ability *Ability152) TaobaoVmarketEticketCodesGet(req *request.TaobaoVmarketEticketCodesGetRequest,session string) (*response.TaobaoVmarketEticketCodesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.codes.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketCodesGetResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketCodesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    电子凭证操作日志查询
*/
func (ability *Ability152) TaobaoVmarketEticketOplogsGet(req *request.TaobaoVmarketEticketOplogsGetRequest,session string) (*response.TaobaoVmarketEticketOplogsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability152 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.vmarket.eticket.oplogs.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoVmarketEticketOplogsGetResponse{}
    if(err != nil){
        log.Println("taobaoVmarketEticketOplogsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
