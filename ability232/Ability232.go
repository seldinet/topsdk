package ability232

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability232/request"
    "topsdk/ability232/response"
    "topsdk/util"
)

type Ability232 struct {
    Client *topsdk.TopClient
}

func NewAbility232(client *topsdk.TopClient) *Ability232{
    return &Ability232{client}
}

/*
    查询单据列表
*/
func (ability *Ability232) TaobaoWlbWmsCainiaoBillQuery(req *request.TaobaoWlbWmsCainiaoBillQueryRequest,session string) (*response.TaobaoWlbWmsCainiaoBillQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wms.cainiao.bill.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWmsCainiaoBillQueryResponse{}
    if(err != nil){
        log.Println("taobaoWlbWmsCainiaoBillQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取销售订单发货信息
*/
func (ability *Ability232) TaobaoWlbWmsConsignBillGet(req *request.TaobaoWlbWmsConsignBillGetRequest,session string) (*response.TaobaoWlbWmsConsignBillGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wms.consign.bill.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWmsConsignBillGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbWmsConsignBillGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    通过订单列表批量获取库存损益单信息
*/
func (ability *Ability232) TaobaoWlbWmsInventoryProfitlossGet(req *request.TaobaoWlbWmsInventoryProfitlossGetRequest,session string) (*response.TaobaoWlbWmsInventoryProfitlossGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wms.inventory.profitloss.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWmsInventoryProfitlossGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbWmsInventoryProfitlossGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询组合商品的组合关系
*/
func (ability *Ability232) TaobaoWlbWmsItemCombinationGet(req *request.TaobaoWlbWmsItemCombinationGetRequest,session string) (*response.TaobaoWlbWmsItemCombinationGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wms.item.combination.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWmsItemCombinationGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbWmsItemCombinationGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    菜鸟商品库存查询
*/
func (ability *Ability232) TaobaoWlbWmsInventoryQuery(req *request.TaobaoWlbWmsInventoryQueryRequest,session string) (*response.TaobaoWlbWmsInventoryQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wms.inventory.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWmsInventoryQueryResponse{}
    if(err != nil){
        log.Println("taobaoWlbWmsInventoryQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询家装服务商列表
*/
func (ability *Ability232) TaobaoWlbOrderJzpartnerQuery(req *request.TaobaoWlbOrderJzpartnerQueryRequest,session string) (*response.TaobaoWlbOrderJzpartnerQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.order.jzpartner.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbOrderJzpartnerQueryResponse{}
    if(err != nil){
        log.Println("taobaoWlbOrderJzpartnerQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    驱动保税交易订单发货
*/
func (ability *Ability232) CainiaoBimTradeorderConsign(req *request.CainiaoBimTradeorderConsignRequest,session string) (*response.CainiaoBimTradeorderConsignResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("cainiao.bim.tradeorder.consign",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.CainiaoBimTradeorderConsignResponse{}
    if(err != nil){
        log.Println("cainiaoBimTradeorderConsign error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品信息查询
*/
func (ability *Ability232) TaobaoWlbWmsSkuGet(req *request.TaobaoWlbWmsSkuGetRequest,session string) (*response.TaobaoWlbWmsSkuGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wms.sku.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWmsSkuGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbWmsSkuGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    按照日期范围查询物流订单详情
*/
func (ability *Ability232) TaobaoWlbOrderdetailDateGet(req *request.TaobaoWlbOrderdetailDateGetRequest,session string) (*response.TaobaoWlbOrderdetailDateGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.orderdetail.date.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbOrderdetailDateGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbOrderdetailDateGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询单据序列号信息
*/
func (ability *Ability232) TaobaoWlbWmsSnInfoQuery(req *request.TaobaoWlbWmsSnInfoQueryRequest,session string) (*response.TaobaoWlbWmsSnInfoQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wms.sn.info.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWmsSnInfoQueryResponse{}
    if(err != nil){
        log.Println("taobaoWlbWmsSnInfoQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    入库通知单
*/
func (ability *Ability232) TaobaoWlbWmsStockInOrderNotify(req *request.TaobaoWlbWmsStockInOrderNotifyRequest,session string) (*response.TaobaoWlbWmsStockInOrderNotifyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wms.stock.in.order.notify",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWmsStockInOrderNotifyResponse{}
    if(err != nil){
        log.Println("taobaoWlbWmsStockInOrderNotify error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    发货订单通知
*/
func (ability *Ability232) TaobaoWlbWmsConsignOrderNotify(req *request.TaobaoWlbWmsConsignOrderNotifyRequest,session string) (*response.TaobaoWlbWmsConsignOrderNotifyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wms.consign.order.notify",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWmsConsignOrderNotifyResponse{}
    if(err != nil){
        log.Println("taobaoWlbWmsConsignOrderNotify error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    单据取消接口
*/
func (ability *Ability232) TaobaoWlbWmsOrderCancelNotify(req *request.TaobaoWlbWmsOrderCancelNotifyRequest,session string) (*response.TaobaoWlbWmsOrderCancelNotifyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wms.order.cancel.notify",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWmsOrderCancelNotifyResponse{}
    if(err != nil){
        log.Println("taobaoWlbWmsOrderCancelNotify error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品同步
*/
func (ability *Ability232) TaobaoWlbWmsSkuCreate(req *request.TaobaoWlbWmsSkuCreateRequest,session string) (*response.TaobaoWlbWmsSkuCreateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability232 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wms.sku.create",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWmsSkuCreateResponse{}
    if(err != nil){
        log.Println("taobaoWlbWmsSkuCreate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
