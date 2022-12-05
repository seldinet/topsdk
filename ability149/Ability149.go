package ability149

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability149/request"
    "topsdk/ability149/response"
    "topsdk/util"
)

type Ability149 struct {
    Client *topsdk.TopClient
}

func NewAbility149(client *topsdk.TopClient) *Ability149{
    return &Ability149{client}
}

/*
    渠道中心-查询产品列表
*/
func (ability *Ability149) TmallChannelProductsQuery(req *request.TmallChannelProductsQueryRequest,session string) (*response.TmallChannelProductsQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.channel.products.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallChannelProductsQueryResponse{}
    if(err != nil){
        log.Println("tmallChannelProductsQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品库存修改
*/
func (ability *Ability149) TaobaoFenxiaoProductQuantityUpdate(req *request.TaobaoFenxiaoProductQuantityUpdateRequest,session string) (*response.TaobaoFenxiaoProductQuantityUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.quantity.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductQuantityUpdateResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductQuantityUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据outerCode查询商品
*/
func (ability *Ability149) TaobaoScitemOutercodeGet(req *request.TaobaoScitemOutercodeGetRequest,session string) (*response.TaobaoScitemOutercodeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.scitem.outercode.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoScitemOutercodeGetResponse{}
    if(err != nil){
        log.Println("taobaoScitemOutercodeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据id查询商品
*/
func (ability *Ability149) TaobaoScitemGet(req *request.TaobaoScitemGetRequest,session string) (*response.TaobaoScitemGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.scitem.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoScitemGetResponse{}
    if(err != nil){
        log.Println("taobaoScitemGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询后端商品
*/
func (ability *Ability149) TaobaoScitemQuery(req *request.TaobaoScitemQueryRequest,session string) (*response.TaobaoScitemQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.scitem.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoScitemQueryResponse{}
    if(err != nil){
        log.Println("taobaoScitemQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    发布后端商品
*/
func (ability *Ability149) TaobaoScitemAdd(req *request.TaobaoScitemAddRequest,session string) (*response.TaobaoScitemAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.scitem.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoScitemAddResponse{}
    if(err != nil){
        log.Println("taobaoScitemAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    失效指定用户的商品与后端商品的映射关系
*/
func (ability *Ability149) TaobaoScitemMapDelete(req *request.TaobaoScitemMapDeleteRequest,session string) (*response.TaobaoScitemMapDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.scitem.map.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoScitemMapDeleteResponse{}
    if(err != nil){
        log.Println("taobaoScitemMapDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查找IC商品或分销商品与后端商品的关联信息
*/
func (ability *Ability149) TaobaoScitemMapQuery(req *request.TaobaoScitemMapQueryRequest) (*response.TaobaoScitemMapQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.scitem.map.query",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoScitemMapQueryResponse{}
    if(err != nil){
        log.Println("taobaoScitemMapQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    创建IC商品与后端商品的映射关系
*/
func (ability *Ability149) TaobaoScitemMapAdd(req *request.TaobaoScitemMapAddRequest,session string) (*response.TaobaoScitemMapAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.scitem.map.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoScitemMapAddResponse{}
    if(err != nil){
        log.Println("taobaoScitemMapAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    交易库存调整单
*/
func (ability *Ability149) TaobaoInventoryAdjustTrade(req *request.TaobaoInventoryAdjustTradeRequest,session string) (*response.TaobaoInventoryAdjustTradeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.inventory.adjust.trade",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoInventoryAdjustTradeResponse{}
    if(err != nil){
        log.Println("taobaoInventoryAdjustTrade error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    非交易库存调整单
*/
func (ability *Ability149) TaobaoInventoryAdjustExternal(req *request.TaobaoInventoryAdjustExternalRequest,session string) (*response.TaobaoInventoryAdjustExternalResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.inventory.adjust.external",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoInventoryAdjustExternalResponse{}
    if(err != nil){
        log.Println("taobaoInventoryAdjustExternal error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    库存初始化
*/
func (ability *Ability149) TaobaoInventoryInitial(req *request.TaobaoInventoryInitialRequest,session string) (*response.TaobaoInventoryInitialResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.inventory.initial",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoInventoryInitialResponse{}
    if(err != nil){
        log.Println("taobaoInventoryInitial error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询商品库存信息
*/
func (ability *Ability149) TaobaoInventoryQuery(req *request.TaobaoInventoryQueryRequest,session string) (*response.TaobaoInventoryQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.inventory.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoInventoryQueryResponse{}
    if(err != nil){
        log.Println("taobaoInventoryQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    合作申请查询
*/
func (ability *Ability149) TaobaoFenxiaoRequisitionsGet(req *request.TaobaoFenxiaoRequisitionsGetRequest,session string) (*response.TaobaoFenxiaoRequisitionsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.requisitions.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoRequisitionsGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoRequisitionsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    等级折扣查询
*/
func (ability *Ability149) TaobaoFenxiaoProductGradepriceGet(req *request.TaobaoFenxiaoProductGradepriceGetRequest,session string) (*response.TaobaoFenxiaoProductGradepriceGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.gradeprice.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductGradepriceGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductGradepriceGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除产品线
*/
func (ability *Ability149) TaobaoFenxiaoProductcatDelete(req *request.TaobaoFenxiaoProductcatDeleteRequest,session string) (*response.TaobaoFenxiaoProductcatDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.productcat.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductcatDeleteResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductcatDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改产品线
*/
func (ability *Ability149) TaobaoFenxiaoProductcatUpdate(req *request.TaobaoFenxiaoProductcatUpdateRequest,session string) (*response.TaobaoFenxiaoProductcatUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.productcat.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductcatUpdateResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductcatUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新增产品线
*/
func (ability *Ability149) TaobaoFenxiaoProductcatAdd(req *request.TaobaoFenxiaoProductcatAddRequest,session string) (*response.TaobaoFenxiaoProductcatAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.productcat.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductcatAddResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductcatAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    导入商品生成产品
*/
func (ability *Ability149) TaobaoFenxiaoProductImportFromAuction(req *request.TaobaoFenxiaoProductImportFromAuctionRequest,session string) (*response.TaobaoFenxiaoProductImportFromAuctionResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.import.from.auction",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductImportFromAuctionResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductImportFromAuction error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品库存初始化
*/
func (ability *Ability149) TaobaoInventoryInitialItem(req *request.TaobaoInventoryInitialItemRequest,session string) (*response.TaobaoInventoryInitialItemResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.inventory.initial.item",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoInventoryInitialItemResponse{}
    if(err != nil){
        log.Println("taobaoInventoryInitialItem error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品导入到渠道
*/
func (ability *Ability149) TaobaoFenxiaoProductToChannelImport(req *request.TaobaoFenxiaoProductToChannelImportRequest,session string) (*response.TaobaoFenxiaoProductToChannelImportResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.to.channel.import",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductToChannelImportResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductToChannelImport error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询供应商的产品数据
*/
func (ability *Ability149) TmallChannelProductsGet(req *request.TmallChannelProductsGetRequest,session string) (*response.TmallChannelProductsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.channel.products.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallChannelProductsGetResponse{}
    if(err != nil){
        log.Println("tmallChannelProductsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    创建经销采购申请
*/
func (ability *Ability149) TaobaoFenxiaoDealerRequisitionorderCreate(req *request.TaobaoFenxiaoDealerRequisitionorderCreateRequest,session string) (*response.TaobaoFenxiaoDealerRequisitionorderCreateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.dealer.requisitionorder.create",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDealerRequisitionorderCreateResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDealerRequisitionorderCreate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    地点关联关系增量编辑
*/
func (ability *Ability149) TaobaoLocationRelationEdit(req *request.TaobaoLocationRelationEditRequest,session string) (*response.TaobaoLocationRelationEditResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.location.relation.edit",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLocationRelationEditResponse{}
    if(err != nil){
        log.Println("taobaoLocationRelationEdit error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    地点关联关系查询
*/
func (ability *Ability149) TaobaoLocationRelationQuery(req *request.TaobaoLocationRelationQueryRequest,session string) (*response.TaobaoLocationRelationQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.location.relation.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLocationRelationQueryResponse{}
    if(err != nil){
        log.Println("taobaoLocationRelationQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    货品库存商家端调整
*/
func (ability *Ability149) TaobaoInventoryMerchantAdjust(req *request.TaobaoInventoryMerchantAdjustRequest,session string) (*response.TaobaoInventoryMerchantAdjustResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.inventory.merchant.adjust",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoInventoryMerchantAdjustResponse{}
    if(err != nil){
        log.Println("taobaoInventoryMerchantAdjust error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询采购单退款信息
*/
func (ability *Ability149) TaobaoFenxiaoRefundGet(req *request.TaobaoFenxiaoRefundGetRequest,session string) (*response.TaobaoFenxiaoRefundGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.refund.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoRefundGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoRefundGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分销商查询产品信息
*/
func (ability *Ability149) TaobaoFenxiaoDistributorProductsGet(req *request.TaobaoFenxiaoDistributorProductsGetRequest,session string) (*response.TaobaoFenxiaoDistributorProductsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.distributor.products.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDistributorProductsGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDistributorProductsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    批量查询采购退款
*/
func (ability *Ability149) TaobaoFenxiaoRefundQuery(req *request.TaobaoFenxiaoRefundQueryRequest,session string) (*response.TaobaoFenxiaoRefundQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.refund.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoRefundQueryResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoRefundQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    供应商/分销商关闭采购申请/经销采购单
*/
func (ability *Ability149) TaobaoFenxiaoDealerRequisitionorderClose(req *request.TaobaoFenxiaoDealerRequisitionorderCloseRequest,session string) (*response.TaobaoFenxiaoDealerRequisitionorderCloseResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.dealer.requisitionorder.close",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDealerRequisitionorderCloseResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDealerRequisitionorderClose error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    供应商/分销商通过采购申请/经销采购单申请
*/
func (ability *Ability149) TaobaoFenxiaoDealerRequisitionorderAgree(req *request.TaobaoFenxiaoDealerRequisitionorderAgreeRequest,session string) (*response.TaobaoFenxiaoDealerRequisitionorderAgreeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.dealer.requisitionorder.agree",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDealerRequisitionorderAgreeResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDealerRequisitionorderAgree error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    批量查询采购申请/经销采购单
*/
func (ability *Ability149) TaobaoFenxiaoDealerRequisitionorderGet(req *request.TaobaoFenxiaoDealerRequisitionorderGetRequest,session string) (*response.TaobaoFenxiaoDealerRequisitionorderGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.dealer.requisitionorder.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDealerRequisitionorderGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDealerRequisitionorderGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    按编号查询采购申请/经销采购单
*/
func (ability *Ability149) TaobaoFenxiaoDealerRequisitionorderQuery(req *request.TaobaoFenxiaoDealerRequisitionorderQueryRequest,session string) (*response.TaobaoFenxiaoDealerRequisitionorderQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.dealer.requisitionorder.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDealerRequisitionorderQueryResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDealerRequisitionorderQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分页查询商家仓信息
*/
func (ability *Ability149) TaobaoInventoryWarehouseQuery(req *request.TaobaoInventoryWarehouseQueryRequest,session string) (*response.TaobaoInventoryWarehouseQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.inventory.warehouse.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoInventoryWarehouseQueryResponse{}
    if(err != nil){
        log.Println("taobaoInventoryWarehouseQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    创建商家仓或者更新商家仓信息
*/
func (ability *Ability149) TaobaoInventoryWarehouseManage(req *request.TaobaoInventoryWarehouseManageRequest,session string) (*response.TaobaoInventoryWarehouseManageResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.inventory.warehouse.manage",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoInventoryWarehouseManageResponse{}
    if(err != nil){
        log.Println("taobaoInventoryWarehouseManage error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询商品销售区域
*/
func (ability *Ability149) TaobaoRegionSaleQuery(req *request.TaobaoRegionSaleQueryRequest,session string) (*response.TaobaoRegionSaleQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.region.sale.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRegionSaleQueryResponse{}
    if(err != nil){
        log.Println("taobaoRegionSaleQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    创建渠道分销单
*/
func (ability *Ability149) TmallChannelTradeOrderCreate(req *request.TmallChannelTradeOrderCreateRequest,session string) (*response.TmallChannelTradeOrderCreateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.channel.trade.order.create",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallChannelTradeOrderCreateResponse{}
    if(err != nil){
        log.Println("tmallChannelTradeOrderCreate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    线下预存款流水增加
*/
func (ability *Ability149) TaobaoFenxiaoTradePrepayOfflineAdd(req *request.TaobaoFenxiaoTradePrepayOfflineAddRequest,session string) (*response.TaobaoFenxiaoTradePrepayOfflineAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.trade.prepay.offline.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoTradePrepayOfflineAddResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoTradePrepayOfflineAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    渠道分销供应商上传线下流水预存款（减少）
*/
func (ability *Ability149) TaobaoFenxiaoTradePrepayOfflineReduce(req *request.TaobaoFenxiaoTradePrepayOfflineReduceRequest,session string) (*response.TaobaoFenxiaoTradePrepayOfflineReduceResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.trade.prepay.offline.reduce",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoTradePrepayOfflineReduceResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoTradePrepayOfflineReduce error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据sku设置折扣价
*/
func (ability *Ability149) TaobaoFenxiaoProductGradepriceUpdate(req *request.TaobaoFenxiaoProductGradepriceUpdateRequest,session string) (*response.TaobaoFenxiaoProductGradepriceUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.gradeprice.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductGradepriceUpdateResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductGradepriceUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分销商查询产品库存
*/
func (ability *Ability149) TaobaoFenxiaoDistributorProductQuantityGet(req *request.TaobaoFenxiaoDistributorProductQuantityGetRequest,session string) (*response.TaobaoFenxiaoDistributorProductQuantityGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.distributor.product.quantity.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDistributorProductQuantityGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDistributorProductQuantityGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改经销采购单备注
*/
func (ability *Ability149) TaobaoFenxiaoDealerRequisitionorderRemarkUpdate(req *request.TaobaoFenxiaoDealerRequisitionorderRemarkUpdateRequest,session string) (*response.TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability149 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.dealer.requisitionorder.remark.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDealerRequisitionorderRemarkUpdateResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDealerRequisitionorderRemarkUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
