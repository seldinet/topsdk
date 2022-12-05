package ability138

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability138/request"
    "topsdk/ability138/response"
    "topsdk/util"
)

type Ability138 struct {
    Client *topsdk.TopClient
}

func NewAbility138(client *topsdk.TopClient) *Ability138{
    return &Ability138{client}
}

/*
    taobao.item.update.delisting.tmall
*/
func (ability *Ability138) TaobaoItemUpdateDelistingTmall(req *request.TaobaoItemUpdateDelistingTmallRequest,session string) (*response.TaobaoItemUpdateDelistingTmallResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.update.delisting.tmall",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemUpdateDelistingTmallResponse{}
    if(err != nil){
        log.Println("taobaoItemUpdateDelistingTmall error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫简化编辑商品
*/
func (ability *Ability138) TmallItemSimpleschemaUpdate(req *request.TmallItemSimpleschemaUpdateRequest,session string) (*response.TmallItemSimpleschemaUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.simpleschema.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemSimpleschemaUpdateResponse{}
    if(err != nil){
        log.Println("tmallItemSimpleschemaUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫商品/SKU库存更新接口
*/
func (ability *Ability138) TmallItemQuantityUpdate(req *request.TmallItemQuantityUpdateRequest,session string) (*response.TmallItemQuantityUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.quantity.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemQuantityUpdateResponse{}
    if(err != nil){
        log.Println("tmallItemQuantityUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    taobao.item.update.listing天猫分流
*/
func (ability *Ability138) TaobaoItemUpdateListingTmall(req *request.TaobaoItemUpdateListingTmallRequest,session string) (*response.TaobaoItemUpdateListingTmallResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.update.listing.tmall",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemUpdateListingTmallResponse{}
    if(err != nil){
        log.Println("taobaoItemUpdateListingTmall error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品描述模块信息获取
*/
func (ability *Ability138) TmallItemDescModulesGet(req *request.TmallItemDescModulesGetRequest) (*response.TmallItemDescModulesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("tmall.item.desc.modules.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TmallItemDescModulesGetResponse{}
    if(err != nil){
        log.Println("tmallItemDescModulesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取天猫商品尺码表模板列表
*/
func (ability *Ability138) TmallItemSizemappingTemplatesList(req *request.TmallItemSizemappingTemplatesListRequest,session string) (*response.TmallItemSizemappingTemplatesListResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.sizemapping.templates.list",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemSizemappingTemplatesListResponse{}
    if(err != nil){
        log.Println("tmallItemSizemappingTemplatesList error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取天猫商品尺码表模板
*/
func (ability *Ability138) TmallItemSizemappingTemplateGet(req *request.TmallItemSizemappingTemplateGetRequest,session string) (*response.TmallItemSizemappingTemplateGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.sizemapping.template.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemSizemappingTemplateGetResponse{}
    if(err != nil){
        log.Println("tmallItemSizemappingTemplateGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除天猫商品尺码表模板
*/
func (ability *Ability138) TmallItemSizemappingTemplateDelete(req *request.TmallItemSizemappingTemplateDeleteRequest,session string) (*response.TmallItemSizemappingTemplateDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.sizemapping.template.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemSizemappingTemplateDeleteResponse{}
    if(err != nil){
        log.Println("tmallItemSizemappingTemplateDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新天猫商品尺码表模板
*/
func (ability *Ability138) TmallItemSizemappingTemplateUpdate(req *request.TmallItemSizemappingTemplateUpdateRequest,session string) (*response.TmallItemSizemappingTemplateUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.sizemapping.template.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemSizemappingTemplateUpdateResponse{}
    if(err != nil){
        log.Println("tmallItemSizemappingTemplateUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新增天猫商品尺码表模板
*/
func (ability *Ability138) TmallItemSizemappingTemplateCreate(req *request.TmallItemSizemappingTemplateCreateRequest,session string) (*response.TmallItemSizemappingTemplateCreateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.sizemapping.template.create",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemSizemappingTemplateCreateResponse{}
    if(err != nil){
        log.Println("tmallItemSizemappingTemplateCreate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新商品发货时间
*/
func (ability *Ability138) TmallItemShiptimeUpdate(req *request.TmallItemShiptimeUpdateRequest,session string) (*response.TmallItemShiptimeUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.shiptime.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemShiptimeUpdateResponse{}
    if(err != nil){
        log.Println("tmallItemShiptimeUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫简化发布商品
*/
func (ability *Ability138) TmallItemSimpleschemaAdd(req *request.TmallItemSimpleschemaAddRequest,session string) (*response.TmallItemSimpleschemaAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.simpleschema.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemSimpleschemaAddResponse{}
    if(err != nil){
        log.Println("tmallItemSimpleschemaAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    发品资质校验
*/
func (ability *Ability138) TaobaoItemPermitCheck(req *request.TaobaoItemPermitCheckRequest,session string) (*response.TaobaoItemPermitCheckResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability138 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.permit.check",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemPermitCheckResponse{}
    if(err != nil){
        log.Println("taobaoItemPermitCheck error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
