package ability648

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability648/request"
    "topsdk/ability648/response"
    "topsdk/util"
)

type Ability648 struct {
    Client *topsdk.TopClient
}

func NewAbility648(client *topsdk.TopClient) *Ability648{
    return &Ability648{client}
}

/*
    更新商品信息
*/
func (ability *Ability648) TaobaoItemUpdate(req *request.TaobaoItemUpdateRequest,session string) (*response.TaobaoItemUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemUpdateResponse{}
    if(err != nil){
        log.Println("taobaoItemUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    添加SKU
*/
func (ability *Ability648) TaobaoItemSkuAdd(req *request.TaobaoItemSkuAddRequest,session string) (*response.TaobaoItemSkuAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.sku.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkuAddResponse{}
    if(err != nil){
        log.Println("taobaoItemSkuAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新SKU信息
*/
func (ability *Ability648) TaobaoItemSkuUpdate(req *request.TaobaoItemSkuUpdateRequest,session string) (*response.TaobaoItemSkuUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.sku.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkuUpdateResponse{}
    if(err != nil){
        log.Println("taobaoItemSkuUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询卖家已卖出的交易数据（根据创建时间）
*/
func (ability *Ability648) TaobaoTradesSoldGet(req *request.TaobaoTradesSoldGetRequest,session string) (*response.TaobaoTradesSoldGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trades.sold.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradesSoldGetResponse{}
    if(err != nil){
        log.Println("taobaoTradesSoldGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取单笔交易的部分信息(性能高)
*/
func (ability *Ability648) TaobaoTradeGet(req *request.TaobaoTradeGetRequest,session string) (*response.TaobaoTradeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeGetResponse{}
    if(err != nil){
        log.Println("taobaoTradeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    对一笔交易添加备注
*/
func (ability *Ability648) TaobaoTradeMemoAdd(req *request.TaobaoTradeMemoAddRequest,session string) (*response.TaobaoTradeMemoAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.memo.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeMemoAddResponse{}
    if(err != nil){
        log.Println("taobaoTradeMemoAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改交易备注
*/
func (ability *Ability648) TaobaoTradeMemoUpdate(req *request.TaobaoTradeMemoUpdateRequest,session string) (*response.TaobaoTradeMemoUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.memo.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeMemoUpdateResponse{}
    if(err != nil){
        log.Println("taobaoTradeMemoUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取单笔交易的详细信息
*/
func (ability *Ability648) TaobaoTradeFullinfoGet(req *request.TaobaoTradeFullinfoGetRequest,session string) (*response.TaobaoTradeFullinfoGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.fullinfo.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeFullinfoGetResponse{}
    if(err != nil){
        log.Println("taobaoTradeFullinfoGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询地址区域
*/
func (ability *Ability648) TaobaoAreasGet(req *request.TaobaoAreasGetRequest) (*response.TaobaoAreasGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.areas.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoAreasGetResponse{}
    if(err != nil){
        log.Println("taobaoAreasGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取交易确认收货费用
*/
func (ability *Ability648) TaobaoTradeConfirmfeeGet(req *request.TaobaoTradeConfirmfeeGetRequest,session string) (*response.TaobaoTradeConfirmfeeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.confirmfee.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeConfirmfeeGetResponse{}
    if(err != nil){
        log.Println("taobaoTradeConfirmfeeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询物流公司信息
*/
func (ability *Ability648) TaobaoLogisticsCompaniesGet(req *request.TaobaoLogisticsCompaniesGetRequest) (*response.TaobaoLogisticsCompaniesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.logistics.companies.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoLogisticsCompaniesGetResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsCompaniesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新交易的销售属性
*/
func (ability *Ability648) TaobaoTradeOrderskuUpdate(req *request.TaobaoTradeOrderskuUpdateRequest,session string) (*response.TaobaoTradeOrderskuUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.ordersku.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeOrderskuUpdateResponse{}
    if(err != nil){
        log.Println("taobaoTradeOrderskuUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    交易帐务查询
*/
func (ability *Ability648) TaobaoTradeAmountGet(req *request.TaobaoTradeAmountGetRequest,session string) (*response.TaobaoTradeAmountGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.amount.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeAmountGetResponse{}
    if(err != nil){
        log.Println("taobaoTradeAmountGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更改交易的收货地址
*/
func (ability *Ability648) TaobaoTradeShippingaddressUpdate(req *request.TaobaoTradeShippingaddressUpdateRequest,session string) (*response.TaobaoTradeShippingaddressUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.shippingaddress.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeShippingaddressUpdateResponse{}
    if(err != nil){
        log.Println("taobaoTradeShippingaddressUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询支持起始地到目的地范围的物流公司
*/
func (ability *Ability648) TaobaoLogisticsPartnersGet(req *request.TaobaoLogisticsPartnersGetRequest) (*response.TaobaoLogisticsPartnersGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.logistics.partners.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoLogisticsPartnersGetResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsPartnersGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    延长交易收货时间
*/
func (ability *Ability648) TaobaoTradeReceivetimeDelay(req *request.TaobaoTradeReceivetimeDelayRequest,session string) (*response.TaobaoTradeReceivetimeDelayResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.receivetime.delay",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradeReceivetimeDelayResponse{}
    if(err != nil){
        log.Println("taobaoTradeReceivetimeDelay error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询卖家地址库
*/
func (ability *Ability648) TaobaoLogisticsAddressSearch(req *request.TaobaoLogisticsAddressSearchRequest,session string) (*response.TaobaoLogisticsAddressSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.logistics.address.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLogisticsAddressSearchResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsAddressSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除卖家地址库
*/
func (ability *Ability648) TaobaoLogisticsAddressRemove(req *request.TaobaoLogisticsAddressRemoveRequest,session string) (*response.TaobaoLogisticsAddressRemoveResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.logistics.address.remove",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLogisticsAddressRemoveResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsAddressRemove error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    卖家地址库修改
*/
func (ability *Ability648) TaobaoLogisticsAddressModify(req *request.TaobaoLogisticsAddressModifyRequest,session string) (*response.TaobaoLogisticsAddressModifyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.logistics.address.modify",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLogisticsAddressModifyResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsAddressModify error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    卖家地址库新增接口
*/
func (ability *Ability648) TaobaoLogisticsAddressAdd(req *request.TaobaoLogisticsAddressAddRequest,session string) (*response.TaobaoLogisticsAddressAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.logistics.address.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLogisticsAddressAddResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsAddressAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    取消物流订单接口
*/
func (ability *Ability648) TaobaoLogisticsOnlineCancel(req *request.TaobaoLogisticsOnlineCancelRequest,session string) (*response.TaobaoLogisticsOnlineCancelResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.logistics.online.cancel",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLogisticsOnlineCancelResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsOnlineCancel error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    无需物流（虚拟）发货处理
*/
func (ability *Ability648) TaobaoLogisticsDummySend(req *request.TaobaoLogisticsDummySendRequest,session string) (*response.TaobaoLogisticsDummySendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.logistics.dummy.send",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLogisticsDummySendResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsDummySend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    限时打折详情查询
*/
func (ability *Ability648) TaobaoPromotionLimitdiscountDetailGet(req *request.TaobaoPromotionLimitdiscountDetailGetRequest,session string) (*response.TaobaoPromotionLimitdiscountDetailGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.promotion.limitdiscount.detail.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPromotionLimitdiscountDetailGetResponse{}
    if(err != nil){
        log.Println("taobaoPromotionLimitdiscountDetailGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    定向优惠活动名称与描述违禁词检查
*/
func (ability *Ability648) TaobaoMarketingPromotionKfc(req *request.TaobaoMarketingPromotionKfcRequest,session string) (*response.TaobaoMarketingPromotionKfcResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.marketing.promotion.kfc",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoMarketingPromotionKfcResponse{}
    if(err != nil){
        log.Println("taobaoMarketingPromotionKfc error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改交易邮费价格
*/
func (ability *Ability648) TaobaoTradePostageUpdate(req *request.TaobaoTradePostageUpdateRequest,session string) (*response.TaobaoTradePostageUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trade.postage.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradePostageUpdateResponse{}
    if(err != nil){
        log.Println("taobaoTradePostageUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取用户指定运费模板信息
*/
func (ability *Ability648) TaobaoDeliveryTemplateGet(req *request.TaobaoDeliveryTemplateGetRequest) (*response.TaobaoDeliveryTemplateGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.delivery.template.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoDeliveryTemplateGetResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取用户下所有模板
*/
func (ability *Ability648) TaobaoDeliveryTemplatesGet(req *request.TaobaoDeliveryTemplatesGetRequest,session string) (*response.TaobaoDeliveryTemplatesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.templates.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplatesGetResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplatesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除运费模板
*/
func (ability *Ability648) TaobaoDeliveryTemplateDelete(req *request.TaobaoDeliveryTemplateDeleteRequest,session string) (*response.TaobaoDeliveryTemplateDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.template.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplateDeleteResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新增运费模板
*/
func (ability *Ability648) TaobaoDeliveryTemplateAdd(req *request.TaobaoDeliveryTemplateAddRequest,session string) (*response.TaobaoDeliveryTemplateAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.template.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplateAddResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改运费模板
*/
func (ability *Ability648) TaobaoDeliveryTemplateUpdate(req *request.TaobaoDeliveryTemplateUpdateRequest,session string) (*response.TaobaoDeliveryTemplateUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.delivery.template.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoDeliveryTemplateUpdateResponse{}
    if(err != nil){
        log.Println("taobaoDeliveryTemplateUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新商品价格
*/
func (ability *Ability648) TaobaoItemPriceUpdate(req *request.TaobaoItemPriceUpdateRequest,session string) (*response.TaobaoItemPriceUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.price.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemPriceUpdateResponse{}
    if(err != nil){
        log.Println("taobaoItemPriceUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新商品SKU的价格
*/
func (ability *Ability648) TaobaoItemSkuPriceUpdate(req *request.TaobaoItemSkuPriceUpdateRequest,session string) (*response.TaobaoItemSkuPriceUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability648 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.sku.price.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkuPriceUpdateResponse{}
    if(err != nil){
        log.Println("taobaoItemSkuPriceUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
