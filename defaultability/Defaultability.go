package defaultability

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/defaultability/request"
    "topsdk/defaultability/response"
    "topsdk/util"
)

type Defaultability struct {
    Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability{
    return &Defaultability{client}
}

/*
    获取一个产品的信息
*/
func (ability *Defaultability) TaobaoProductGet(req *request.TaobaoProductGetRequest) (*response.TaobaoProductGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.product.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoProductGetResponse{}
    if(err != nil){
        log.Println("taobaoProductGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    搜索产品信息
*/
func (ability *Defaultability) TaobaoProductsSearch(req *request.TaobaoProductsSearchRequest) (*response.TaobaoProductsSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.products.search",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoProductsSearchResponse{}
    if(err != nil){
        log.Println("taobaoProductsSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    上传一个产品，不包括产品非主图和属性图片
*/
func (ability *Defaultability) TaobaoProductAdd(req *request.TaobaoProductAddRequest,session string) (*response.TaobaoProductAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.product.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoProductAddResponse{}
    if(err != nil){
        log.Println("taobaoProductAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    上传单张产品非主图，如果需要传多张，可调多次
*/
func (ability *Defaultability) TaobaoProductImgUpload(req *request.TaobaoProductImgUploadRequest,session string) (*response.TaobaoProductImgUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.product.img.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoProductImgUploadResponse{}
    if(err != nil){
        log.Println("taobaoProductImgUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    上传单张产品属性图片，如果需要传多张，可调多次
*/
func (ability *Defaultability) TaobaoProductPropimgUpload(req *request.TaobaoProductPropimgUploadRequest,session string) (*response.TaobaoProductPropimgUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.product.propimg.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoProductPropimgUploadResponse{}
    if(err != nil){
        log.Println("taobaoProductPropimgUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改一个产品，可以修改主图，不能修改子图片
*/
func (ability *Defaultability) TaobaoProductUpdate(req *request.TaobaoProductUpdateRequest) (*response.TaobaoProductUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.product.update",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoProductUpdateResponse{}
    if(err != nil){
        log.Println("taobaoProductUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取产品列表
*/
func (ability *Defaultability) TaobaoProductsGet(req *request.TaobaoProductsGetRequest) (*response.TaobaoProductsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.products.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoProductsGetResponse{}
    if(err != nil){
        log.Println("taobaoProductsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取标准类目属性值
*/
func (ability *Defaultability) TaobaoItempropvaluesGet(req *request.TaobaoItempropvaluesGetRequest) (*response.TaobaoItempropvaluesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.itempropvalues.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoItempropvaluesGetResponse{}
    if(err != nil){
        log.Println("taobaoItempropvaluesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    批量判定服务是否可达
*/
func (ability *Defaultability) TaobaoLogisticsAddressReachablebatchGet(req *request.TaobaoLogisticsAddressReachablebatchGetRequest) (*response.TaobaoLogisticsAddressReachablebatchGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.logistics.address.reachablebatch.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoLogisticsAddressReachablebatchGetResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsAddressReachablebatchGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取当前会话用户出售中的商品列表
*/
func (ability *Defaultability) TaobaoItemsOnsaleGet(req *request.TaobaoItemsOnsaleGetRequest,session string) (*response.TaobaoItemsOnsaleGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.items.onsale.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemsOnsaleGetResponse{}
    if(err != nil){
        log.Println("taobaoItemsOnsaleGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    添加一个商品
*/
func (ability *Defaultability) TaobaoItemAdd(req *request.TaobaoItemAddRequest,session string) (*response.TaobaoItemAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemAddResponse{}
    if(err != nil){
        log.Println("taobaoItemAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    添加商品图片
*/
func (ability *Defaultability) TaobaoItemImgUpload(req *request.TaobaoItemImgUploadRequest,session string) (*response.TaobaoItemImgUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.img.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemImgUploadResponse{}
    if(err != nil){
        log.Println("taobaoItemImgUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除商品图片
*/
func (ability *Defaultability) TaobaoItemImgDelete(req *request.TaobaoItemImgDeleteRequest,session string) (*response.TaobaoItemImgDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.img.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemImgDeleteResponse{}
    if(err != nil){
        log.Println("taobaoItemImgDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除属性图片
*/
func (ability *Defaultability) TaobaoItemPropimgDelete(req *request.TaobaoItemPropimgDeleteRequest,session string) (*response.TaobaoItemPropimgDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.propimg.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemPropimgDeleteResponse{}
    if(err != nil){
        log.Println("taobaoItemPropimgDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取用户已开通消息
*/
func (ability *Defaultability) TaobaoTmcUserGet(req *request.TaobaoTmcUserGetRequest) (*response.TaobaoTmcUserGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.user.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcUserGetResponse{}
    if(err != nil){
        log.Println("taobaoTmcUserGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    添加或修改属性图片
*/
func (ability *Defaultability) TaobaoItemPropimgUpload(req *request.TaobaoItemPropimgUploadRequest,session string) (*response.TaobaoItemPropimgUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.propimg.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemPropimgUploadResponse{}
    if(err != nil){
        log.Println("taobaoItemPropimgUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取SKU
*/
func (ability *Defaultability) TaobaoItemSkuGet(req *request.TaobaoItemSkuGetRequest,session string) (*response.TaobaoItemSkuGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.sku.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkuGetResponse{}
    if(err != nil){
        log.Println("taobaoItemSkuGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据商品ID列表获取SKU信息
*/
func (ability *Defaultability) TaobaoItemSkusGet(req *request.TaobaoItemSkusGetRequest,session string) (*response.TaobaoItemSkusGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.skus.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkusGetResponse{}
    if(err != nil){
        log.Println("taobaoItemSkusGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取单个商品详细信息
*/
func (ability *Defaultability) TaobaoItemSellerGet(req *request.TaobaoItemSellerGetRequest,session string) (*response.TaobaoItemSellerGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.seller.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSellerGetResponse{}
    if(err != nil){
        log.Println("taobaoItemSellerGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    批量获取商品详细信息
*/
func (ability *Defaultability) TaobaoItemsSellerListGet(req *request.TaobaoItemsSellerListGetRequest,session string) (*response.TaobaoItemsSellerListGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.items.seller.list.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemsSellerListGetResponse{}
    if(err != nil){
        log.Println("taobaoItemsSellerListGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询买家申请的退款列表
*/
func (ability *Defaultability) TaobaoRefundsApplyGet(req *request.TaobaoRefundsApplyGetRequest,session string) (*response.TaobaoRefundsApplyGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refunds.apply.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundsApplyGetResponse{}
    if(err != nil){
        log.Println("taobaoRefundsApplyGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取单笔退款详情
*/
func (ability *Defaultability) TaobaoRefundGet(req *request.TaobaoRefundGetRequest,session string) (*response.TaobaoRefundGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refund.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundGetResponse{}
    if(err != nil){
        log.Println("taobaoRefundGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取前台展示的店铺类目
*/
func (ability *Defaultability) TaobaoShopcatsListGet(req *request.TaobaoShopcatsListGetRequest) (*response.TaobaoShopcatsListGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.shopcats.list.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoShopcatsListGetResponse{}
    if(err != nil){
        log.Println("taobaoShopcatsListGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取前台展示的店铺内卖家自定义商品类目
*/
func (ability *Defaultability) TaobaoSellercatsListGet(req *request.TaobaoSellercatsListGetRequest,session string) (*response.TaobaoSellercatsListGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercats.list.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercatsListGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercatsListGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取标准商品类目属性
*/
func (ability *Defaultability) TaobaoItempropsGet(req *request.TaobaoItempropsGetRequest) (*response.TaobaoItempropsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.itemprops.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoItempropsGetResponse{}
    if(err != nil){
        log.Println("taobaoItempropsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取后台供卖家发布商品的标准商品类目
*/
func (ability *Defaultability) TaobaoItemcatsGet(req *request.TaobaoItemcatsGetRequest) (*response.TaobaoItemcatsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.itemcats.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoItemcatsGetResponse{}
    if(err != nil){
        log.Println("taobaoItemcatsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询退款留言/凭证列表
*/
func (ability *Defaultability) TaobaoRefundMessagesGet(req *request.TaobaoRefundMessagesGetRequest,session string) (*response.TaobaoRefundMessagesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refund.messages.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundMessagesGetResponse{}
    if(err != nil){
        log.Println("taobaoRefundMessagesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    创建退款留言/凭证
*/
func (ability *Defaultability) TaobaoRefundMessageAdd(req *request.TaobaoRefundMessageAddRequest,session string) (*response.TaobaoRefundMessageAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.refund.message.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRefundMessageAddResponse{}
    if(err != nil){
        log.Println("taobaoRefundMessageAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取图片分类信息
*/
func (ability *Defaultability) TaobaoPictureCategoryGet(req *request.TaobaoPictureCategoryGetRequest,session string) (*response.TaobaoPictureCategoryGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.category.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureCategoryGetResponse{}
    if(err != nil){
        log.Println("taobaoPictureCategoryGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    服务平台营销链接生成接口
*/
func (ability *Defaultability) TaobaoFuwuSaleLinkGen(req *request.TaobaoFuwuSaleLinkGenRequest) (*response.TaobaoFuwuSaleLinkGenResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.fuwu.sale.link.gen",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoFuwuSaleLinkGenResponse{}
    if(err != nil){
        log.Println("taobaoFuwuSaleLinkGen error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取图片信息
*/
func (ability *Defaultability) TaobaoPictureGet(req *request.TaobaoPictureGetRequest,session string) (*response.TaobaoPictureGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureGetResponse{}
    if(err != nil){
        log.Println("taobaoPictureGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    上传单张图片
*/
func (ability *Defaultability) TaobaoPictureUpload(req *request.TaobaoPictureUploadRequest,session string) (*response.TaobaoPictureUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureUploadResponse{}
    if(err != nil){
        log.Println("taobaoPictureUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    关键词过滤匹配
*/
func (ability *Defaultability) TaobaoKfcKeywordSearch(req *request.TaobaoKfcKeywordSearchRequest,session string) (*response.TaobaoKfcKeywordSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.kfc.keyword.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoKfcKeywordSearchResponse{}
    if(err != nil){
        log.Println("taobaoKfcKeywordSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品关联子图
*/
func (ability *Defaultability) TaobaoItemJointImg(req *request.TaobaoItemJointImgRequest,session string) (*response.TaobaoItemJointImgResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.joint.img",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemJointImgResponse{}
    if(err != nil){
        log.Println("taobaoItemJointImg error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品关联属性图
*/
func (ability *Defaultability) TaobaoItemJointPropimg(req *request.TaobaoItemJointPropimgRequest,session string) (*response.TaobaoItemJointPropimgResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.joint.propimg",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemJointPropimgResponse{}
    if(err != nil){
        log.Println("taobaoItemJointPropimg error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询商家被授权品牌列表和类目列表
*/
func (ability *Defaultability) TaobaoItemcatsAuthorizeGet(req *request.TaobaoItemcatsAuthorizeGetRequest,session string) (*response.TaobaoItemcatsAuthorizeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.itemcats.authorize.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemcatsAuthorizeGetResponse{}
    if(err != nil){
        log.Println("taobaoItemcatsAuthorizeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    得到当前会话用户库存中的商品列表
*/
func (ability *Defaultability) TaobaoItemsInventoryGet(req *request.TaobaoItemsInventoryGetRequest,session string) (*response.TaobaoItemsInventoryGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.items.inventory.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemsInventoryGetResponse{}
    if(err != nil){
        log.Println("taobaoItemsInventoryGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据外部ID取商品SKU
*/
func (ability *Defaultability) TaobaoSkusCustomGet(req *request.TaobaoSkusCustomGetRequest,session string) (*response.TaobaoSkusCustomGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.skus.custom.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSkusCustomGetResponse{}
    if(err != nil){
        log.Println("taobaoSkusCustomGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户的子账号简易信息列表
*/
func (ability *Defaultability) TaobaoSubusersGet(req *request.TaobaoSubusersGetRequest,session string) (*response.TaobaoSubusersGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subusers.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubusersGetResponse{}
    if(err != nil){
        log.Println("taobaoSubusersGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户子账号的详细信息
*/
func (ability *Defaultability) TaobaoSubuserFullinfoGet(req *request.TaobaoSubuserFullinfoGetRequest,session string) (*response.TaobaoSubuserFullinfoGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.fullinfo.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserFullinfoGetResponse{}
    if(err != nil){
        log.Println("taobaoSubuserFullinfoGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户的所有部门列表
*/
func (ability *Defaultability) TaobaoSubuserDepartmentsGet(req *request.TaobaoSubuserDepartmentsGetRequest,session string) (*response.TaobaoSubuserDepartmentsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.departments.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserDepartmentsGetResponse{}
    if(err != nil){
        log.Println("taobaoSubuserDepartmentsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定账户的所有职务信息列表
*/
func (ability *Defaultability) TaobaoSubuserDutysGet(req *request.TaobaoSubuserDutysGetRequest,session string) (*response.TaobaoSubuserDutysGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.dutys.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserDutysGetResponse{}
    if(err != nil){
        log.Println("taobaoSubuserDutysGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改指定账户子账号的基本信息
*/
func (ability *Defaultability) TaobaoSubuserInfoUpdate(req *request.TaobaoSubuserInfoUpdateRequest,session string) (*response.TaobaoSubuserInfoUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subuser.info.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubuserInfoUpdateResponse{}
    if(err != nil){
        log.Println("taobaoSubuserInfoUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    新增图片分类信息
*/
func (ability *Defaultability) TaobaoPictureCategoryAdd(req *request.TaobaoPictureCategoryAddRequest,session string) (*response.TaobaoPictureCategoryAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.category.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureCategoryAddResponse{}
    if(err != nil){
        log.Println("taobaoPictureCategoryAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫增量更新商品规则获取
*/
func (ability *Defaultability) TmallItemIncrementUpdateSchemaGet(req *request.TmallItemIncrementUpdateSchemaGetRequest,session string) (*response.TmallItemIncrementUpdateSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.increment.update.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemIncrementUpdateSchemaGetResponse{}
    if(err != nil){
        log.Println("tmallItemIncrementUpdateSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫根据规则增量更新商品
*/
func (ability *Defaultability) TmallItemSchemaIncrementUpdate(req *request.TmallItemSchemaIncrementUpdateRequest,session string) (*response.TmallItemSchemaIncrementUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.schema.increment.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemSchemaIncrementUpdateResponse{}
    if(err != nil){
        log.Println("tmallItemSchemaIncrementUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品编辑增量更新
*/
func (ability *Defaultability) AlibabaItemEditFastupdate(req *request.AlibabaItemEditFastupdateRequest,session string) (*response.AlibabaItemEditFastupdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.edit.fastupdate",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemEditFastupdateResponse{}
    if(err != nil){
        log.Println("alibabaItemEditFastupdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询appstore应用订购关系
*/
func (ability *Defaultability) TaobaoAppstoreSubscribeGet(req *request.TaobaoAppstoreSubscribeGetRequest) (*response.TaobaoAppstoreSubscribeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.appstore.subscribe.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoAppstoreSubscribeGetResponse{}
    if(err != nil){
        log.Println("taobaoAppstoreSubscribeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除SKU
*/
func (ability *Defaultability) TaobaoItemSkuDelete(req *request.TaobaoItemSkuDeleteRequest,session string) (*response.TaobaoItemSkuDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.sku.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemSkuDeleteResponse{}
    if(err != nil){
        log.Println("taobaoItemSkuDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    添加产品
*/
func (ability *Defaultability) TaobaoFenxiaoProductAdd(req *request.TaobaoFenxiaoProductAddRequest,session string) (*response.TaobaoFenxiaoProductAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductAddResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新产品
*/
func (ability *Defaultability) TaobaoFenxiaoProductUpdate(req *request.TaobaoFenxiaoProductUpdateRequest,session string) (*response.TaobaoFenxiaoProductUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductUpdateResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询产品线列表
*/
func (ability *Defaultability) TaobaoFenxiaoProductcatsGet(req *request.TaobaoFenxiaoProductcatsGetRequest,session string) (*response.TaobaoFenxiaoProductcatsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.productcats.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductcatsGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductcatsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询产品列表
*/
func (ability *Defaultability) TaobaoFenxiaoProductsGet(req *request.TaobaoFenxiaoProductsGetRequest,session string) (*response.TaobaoFenxiaoProductsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.products.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductsGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新合作关系等级
*/
func (ability *Defaultability) TaobaoFenxiaoCooperationUpdate(req *request.TaobaoFenxiaoCooperationUpdateRequest,session string) (*response.TaobaoFenxiaoCooperationUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.cooperation.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoCooperationUpdateResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoCooperationUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分销商等级查询
*/
func (ability *Defaultability) TaobaoFenxiaoGradesGet(req *request.TaobaoFenxiaoGradesGetRequest,session string) (*response.TaobaoFenxiaoGradesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.grades.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoGradesGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoGradesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取折扣信息
*/
func (ability *Defaultability) TaobaoFenxiaoDiscountsGet(req *request.TaobaoFenxiaoDiscountsGetRequest,session string) (*response.TaobaoFenxiaoDiscountsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.discounts.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDiscountsGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDiscountsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    订购关系查询
*/
func (ability *Defaultability) TaobaoVasSubscribeGet(req *request.TaobaoVasSubscribeGetRequest) (*response.TaobaoVasSubscribeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.vas.subscribe.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoVasSubscribeGetResponse{}
    if(err != nil){
        log.Println("taobaoVasSubscribeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    订单记录导出
*/
func (ability *Defaultability) TaobaoVasOrderSearch(req *request.TaobaoVasOrderSearchRequest) (*response.TaobaoVasOrderSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.vas.order.search",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoVasOrderSearchResponse{}
    if(err != nil){
        log.Println("taobaoVasOrderSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    订购记录导出
*/
func (ability *Defaultability) TaobaoVasSubscSearch(req *request.TaobaoVasSubscSearchRequest) (*response.TaobaoVasSubscSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.vas.subsc.search",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoVasSubscSearchResponse{}
    if(err != nil){
        log.Println("taobaoVasSubscSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    宝贝/SKU库存修改
*/
func (ability *Defaultability) TaobaoItemQuantityUpdate(req *request.TaobaoItemQuantityUpdateRequest,session string) (*response.TaobaoItemQuantityUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.quantity.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemQuantityUpdateResponse{}
    if(err != nil){
        log.Println("taobaoItemQuantityUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    物流宝订单已发货通知接口
*/
func (ability *Defaultability) TaobaoWlbOrderConsign(req *request.TaobaoWlbOrderConsignRequest,session string) (*response.TaobaoWlbOrderConsignResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.order.consign",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbOrderConsignResponse{}
    if(err != nil){
        log.Println("taobaoWlbOrderConsign error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    物流宝通知消息查询接口
*/
func (ability *Defaultability) TaobaoWlbNotifyMessagePageGet(req *request.TaobaoWlbNotifyMessagePageGetRequest,session string) (*response.TaobaoWlbNotifyMessagePageGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.notify.message.page.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbNotifyMessagePageGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbNotifyMessagePageGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新商品条形码信息
*/
func (ability *Defaultability) TaobaoItemBarcodeUpdate(req *request.TaobaoItemBarcodeUpdateRequest,session string) (*response.TaobaoItemBarcodeUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.barcode.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemBarcodeUpdateResponse{}
    if(err != nil){
        log.Println("taobaoItemBarcodeUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    通过物流订单编号查询物流信息
*/
func (ability *Defaultability) TaobaoWlbTmsorderQuery(req *request.TaobaoWlbTmsorderQueryRequest,session string) (*response.TaobaoWlbTmsorderQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.tmsorder.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbTmsorderQueryResponse{}
    if(err != nil){
        log.Println("taobaoWlbTmsorderQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据物流宝商品ID查询商品映射关系
*/
func (ability *Defaultability) TaobaoWlbItemMapGet(req *request.TaobaoWlbItemMapGetRequest,session string) (*response.TaobaoWlbItemMapGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.item.map.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbItemMapGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbItemMapGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据商品id查询商品组合关系
*/
func (ability *Defaultability) TaobaoWlbItemCombinationGet(req *request.TaobaoWlbItemCombinationGetRequest,session string) (*response.TaobaoWlbItemCombinationGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.item.combination.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbItemCombinationGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbItemCombinationGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据商品ID查询所有库存变更记录
*/
func (ability *Defaultability) TaobaoWlbInventorylogQuery(req *request.TaobaoWlbInventorylogQueryRequest,session string) (*response.TaobaoWlbInventorylogQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.inventorylog.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbInventorylogQueryResponse{}
    if(err != nil){
        log.Println("taobaoWlbInventorylogQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    物流宝商品修改
*/
func (ability *Defaultability) TaobaoWlbItemUpdate(req *request.TaobaoWlbItemUpdateRequest,session string) (*response.TaobaoWlbItemUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.item.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbItemUpdateResponse{}
    if(err != nil){
        log.Println("taobaoWlbItemUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询商家定购的所有服务
*/
func (ability *Defaultability) TaobaoWlbSubscriptionQuery(req *request.TaobaoWlbSubscriptionQueryRequest,session string) (*response.TaobaoWlbSubscriptionQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.subscription.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbSubscriptionQueryResponse{}
    if(err != nil){
        log.Println("taobaoWlbSubscriptionQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分页查询物流宝订单
*/
func (ability *Defaultability) TaobaoWlbOrderPageGet(req *request.TaobaoWlbOrderPageGetRequest,session string) (*response.TaobaoWlbOrderPageGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.order.page.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbOrderPageGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbOrderPageGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分页查询物流宝订单商品详情
*/
func (ability *Defaultability) TaobaoWlbOrderitemPageGet(req *request.TaobaoWlbOrderitemPageGetRequest,session string) (*response.TaobaoWlbOrderitemPageGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.orderitem.page.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbOrderitemPageGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbOrderitemPageGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    物流宝订单流转状态查询
*/
func (ability *Defaultability) TaobaoWlbOrderstatusGet(req *request.TaobaoWlbOrderstatusGetRequest,session string) (*response.TaobaoWlbOrderstatusGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.orderstatus.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbOrderstatusGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbOrderstatusGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    取消物流宝订单
*/
func (ability *Defaultability) TaobaoWlbOrderCancel(req *request.TaobaoWlbOrderCancelRequest,session string) (*response.TaobaoWlbOrderCancelResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.order.cancel",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbOrderCancelResponse{}
    if(err != nil){
        log.Println("taobaoWlbOrderCancel error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据商品ID获取商品信息
*/
func (ability *Defaultability) TaobaoWlbItemGet(req *request.TaobaoWlbItemGetRequest,session string) (*response.TaobaoWlbItemGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.item.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbItemGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbItemGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据交易号获取物流宝订单
*/
func (ability *Defaultability) TaobaoWlbTradeorderGet(req *request.TaobaoWlbTradeorderGetRequest,session string) (*response.TaobaoWlbTradeorderGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.tradeorder.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbTradeorderGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbTradeorderGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询库存详情
*/
func (ability *Defaultability) TaobaoWlbInventoryDetailGet(req *request.TaobaoWlbInventoryDetailGetRequest,session string) (*response.TaobaoWlbInventoryDetailGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.inventory.detail.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbInventoryDetailGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbInventoryDetailGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    创建物流宝订单
*/
func (ability *Defaultability) TaobaoWlbOrderCreate(req *request.TaobaoWlbOrderCreateRequest,session string) (*response.TaobaoWlbOrderCreateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.order.create",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbOrderCreateResponse{}
    if(err != nil){
        log.Println("taobaoWlbOrderCreate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分页查询商品
*/
func (ability *Defaultability) TaobaoWlbItemQuery(req *request.TaobaoWlbItemQueryRequest,session string) (*response.TaobaoWlbItemQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.item.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbItemQueryResponse{}
    if(err != nil){
        log.Println("taobaoWlbItemQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据物流宝订单编号查询物流宝订单概要信息
*/
func (ability *Defaultability) TaobaoWlbWlborderGet(req *request.TaobaoWlbWlborderGetRequest,session string) (*response.TaobaoWlbWlborderGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.wlborder.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbWlborderGetResponse{}
    if(err != nil){
        log.Println("taobaoWlbWlborderGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    添加单个物流宝商品
*/
func (ability *Defaultability) TaobaoWlbItemAdd(req *request.TaobaoWlbItemAddRequest,session string) (*response.TaobaoWlbItemAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.item.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbItemAddResponse{}
    if(err != nil){
        log.Println("taobaoWlbItemAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    拆合单结果回传接口
*/
func (ability *Defaultability) TaobaoFulfillmentOrderAssemble(req *request.TaobaoFulfillmentOrderAssembleRequest,session string) (*response.TaobaoFulfillmentOrderAssembleResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fulfillment.order.assemble",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFulfillmentOrderAssembleResponse{}
    if(err != nil){
        log.Println("taobaoFulfillmentOrderAssemble error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    确认收款
*/
func (ability *Defaultability) TaobaoFenxiaoOrderConfirmPaid(req *request.TaobaoFenxiaoOrderConfirmPaidRequest,session string) (*response.TaobaoFenxiaoOrderConfirmPaidResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.order.confirm.paid",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoOrderConfirmPaidResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoOrderConfirmPaid error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询商品下载记录
*/
func (ability *Defaultability) TaobaoFenxiaoDistributorItemsGet(req *request.TaobaoFenxiaoDistributorItemsGetRequest,session string) (*response.TaobaoFenxiaoDistributorItemsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.distributor.items.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoDistributorItemsGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoDistributorItemsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    供应商或分销商获取合作关系信息
*/
func (ability *Defaultability) TaobaoFenxiaoCooperationGet(req *request.TaobaoFenxiaoCooperationGetRequest,session string) (*response.TaobaoFenxiaoCooperationGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.cooperation.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoCooperationGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoCooperationGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取用户宝贝详情页模板名称
*/
func (ability *Defaultability) TaobaoItemTemplatesGet(req *request.TaobaoItemTemplatesGetRequest,session string) (*response.TaobaoItemTemplatesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.templates.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemTemplatesGetResponse{}
    if(err != nil){
        log.Println("taobaoItemTemplatesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取商家可发布商品的市场信息
*/
func (ability *Defaultability) AlibabaItemPublishMarketGet(req *request.AlibabaItemPublishMarketGetRequest,session string) (*response.AlibabaItemPublishMarketGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.publish.market.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemPublishMarketGetResponse{}
    if(err != nil){
        log.Println("alibabaItemPublishMarketGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据收件人信息查询交易单号
*/
func (ability *Defaultability) TaobaoTradesSoldQuery(req *request.TaobaoTradesSoldQueryRequest,session string) (*response.TaobaoTradesSoldQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trades.sold.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradesSoldQueryResponse{}
    if(err != nil){
        log.Println("taobaoTradesSoldQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    发布单条消息
*/
func (ability *Defaultability) TaobaoTmcMessageProduce(req *request.TaobaoTmcMessageProduceRequest) (*response.TaobaoTmcMessageProduceResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.message.produce",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcMessageProduceResponse{}
    if(err != nil){
        log.Println("taobaoTmcMessageProduce error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    取消用户的消息服务
*/
func (ability *Defaultability) TaobaoTmcUserCancel(req *request.TaobaoTmcUserCancelRequest) (*response.TaobaoTmcUserCancelResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.user.cancel",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcUserCancelResponse{}
    if(err != nil){
        log.Println("taobaoTmcUserCancel error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    为已授权的用户开通消息服务
*/
func (ability *Defaultability) TaobaoTmcUserPermit(req *request.TaobaoTmcUserPermitRequest,session string) (*response.TaobaoTmcUserPermitResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.tmc.user.permit",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTmcUserPermitResponse{}
    if(err != nil){
        log.Println("taobaoTmcUserPermit error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫商品/SKU商家编码更新接口
*/
func (ability *Defaultability) TmallItemOuteridUpdate(req *request.TmallItemOuteridUpdateRequest,session string) (*response.TmallItemOuteridUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.outerid.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemOuteridUpdateResponse{}
    if(err != nil){
        log.Println("tmallItemOuteridUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询指定账户的子账号列表
*/
func (ability *Defaultability) TaobaoSellercenterSubusersGet(req *request.TaobaoSellercenterSubusersGetRequest,session string) (*response.TaobaoSellercenterSubusersGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.subusers.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterSubusersGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterSubusersGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定用户的权限集合
*/
func (ability *Defaultability) TaobaoSellercenterUserPermissionsGet(req *request.TaobaoSellercenterUserPermissionsGetRequest,session string) (*response.TaobaoSellercenterUserPermissionsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.user.permissions.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterUserPermissionsGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterUserPermissionsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取商品已生效营销活动更新规则
*/
func (ability *Defaultability) TaobaoItemPromotionRuleGet(req *request.TaobaoItemPromotionRuleGetRequest,session string) (*response.TaobaoItemPromotionRuleGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.promotion.rule.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemPromotionRuleGetResponse{}
    if(err != nil){
        log.Println("taobaoItemPromotionRuleGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取分销用户登录信息
*/
func (ability *Defaultability) TaobaoFenxiaoLoginUserGet(req *request.TaobaoFenxiaoLoginUserGetRequest,session string) (*response.TaobaoFenxiaoLoginUserGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.login.user.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoLoginUserGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoLoginUserGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询指定的子账号的权限和角色信息
*/
func (ability *Defaultability) TaobaoSellercenterSubuserPermissionsRolesGet(req *request.TaobaoSellercenterSubuserPermissionsRolesGetRequest,session string) (*response.TaobaoSellercenterSubuserPermissionsRolesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.subuser.permissions.roles.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterSubuserPermissionsRolesGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterSubuserPermissionsRolesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取指定卖家的角色列表
*/
func (ability *Defaultability) TaobaoSellercenterRolesGet(req *request.TaobaoSellercenterRolesGetRequest,session string) (*response.TaobaoSellercenterRolesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.roles.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterRolesGetResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterRolesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    子账号角色的新增（指定卖家）
*/
func (ability *Defaultability) TaobaoSellercenterRoleAdd(req *request.TaobaoSellercenterRoleAddRequest,session string) (*response.TaobaoSellercenterRoleAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.role.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterRoleAddResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterRoleAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    服务平台评价查询接口
*/
func (ability *Defaultability) TaobaoFuwuScoresGet(req *request.TaobaoFuwuScoresGetRequest) (*response.TaobaoFuwuScoresGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.fuwu.scores.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoFuwuScoresGetResponse{}
    if(err != nil){
        log.Println("taobaoFuwuScoresGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    外部库存变化通知（企业物流用户使用）
*/
func (ability *Defaultability) TaobaoWlbOutInventoryChangeNotify(req *request.TaobaoWlbOutInventoryChangeNotifyRequest,session string) (*response.TaobaoWlbOutInventoryChangeNotifyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.out.inventory.change.notify",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbOutInventoryChangeNotifyResponse{}
    if(err != nil){
        log.Println("taobaoWlbOutInventoryChangeNotify error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    自己联系物流发货
*/
func (ability *Defaultability) AlibabaAscpLogisticsOfflineSend(req *request.AlibabaAscpLogisticsOfflineSendRequest,session string) (*response.AlibabaAscpLogisticsOfflineSendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.ascp.logistics.offline.send",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaAscpLogisticsOfflineSendResponse{}
    if(err != nil){
        log.Println("alibabaAscpLogisticsOfflineSend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改物流公司和运单号
*/
func (ability *Defaultability) AlibabaAscpLogisticsConsignResend(req *request.AlibabaAscpLogisticsConsignResendRequest,session string) (*response.AlibabaAscpLogisticsConsignResendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.ascp.logistics.consign.resend",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaAscpLogisticsConsignResendResponse{}
    if(err != nil){
        log.Println("alibabaAscpLogisticsConsignResend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    更新图片分类
*/
func (ability *Defaultability) TaobaoPictureCategoryUpdate(req *request.TaobaoPictureCategoryUpdateRequest,session string) (*response.TaobaoPictureCategoryUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.category.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureCategoryUpdateResponse{}
    if(err != nil){
        log.Println("taobaoPictureCategoryUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    替换图片
*/
func (ability *Defaultability) TaobaoPictureReplace(req *request.TaobaoPictureReplaceRequest,session string) (*response.TaobaoPictureReplaceResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.replace",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureReplaceResponse{}
    if(err != nil){
        log.Println("taobaoPictureReplace error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    修改图片名字
*/
func (ability *Defaultability) TaobaoPictureUpdate(req *request.TaobaoPictureUpdateRequest,session string) (*response.TaobaoPictureUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureUpdateResponse{}
    if(err != nil){
        log.Println("taobaoPictureUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询图片空间用户的信息
*/
func (ability *Defaultability) TaobaoPictureUserinfoGet(req *request.TaobaoPictureUserinfoGetRequest,session string) (*response.TaobaoPictureUserinfoGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.userinfo.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureUserinfoGetResponse{}
    if(err != nil){
        log.Println("taobaoPictureUserinfoGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    图片是否被引用
*/
func (ability *Defaultability) TaobaoPictureIsreferencedGet(req *request.TaobaoPictureIsreferencedGetRequest,session string) (*response.TaobaoPictureIsreferencedGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.isreferenced.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPictureIsreferencedGetResponse{}
    if(err != nil){
        log.Println("taobaoPictureIsreferencedGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    SKU库存修改
*/
func (ability *Defaultability) TaobaoSkusQuantityUpdate(req *request.TaobaoSkusQuantityUpdateRequest,session string) (*response.TaobaoSkusQuantityUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.skus.quantity.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSkusQuantityUpdateResponse{}
    if(err != nil){
        log.Println("taobaoSkusQuantityUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取商品发布规则信息
*/
func (ability *Defaultability) AlibabaItemPublishSchemaGet(req *request.AlibabaItemPublishSchemaGetRequest,session string) (*response.AlibabaItemPublishSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.publish.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemPublishSchemaGetResponse{}
    if(err != nil){
        log.Println("alibabaItemPublishSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫商品/SKU价格更新接口
*/
func (ability *Defaultability) TmallItemPriceUpdate(req *request.TmallItemPriceUpdateRequest,session string) (*response.TmallItemPriceUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.price.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemPriceUpdateResponse{}
    if(err != nil){
        log.Println("tmallItemPriceUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品SKU删除接口
*/
func (ability *Defaultability) TaobaoFenxiaoProductSkuDelete(req *request.TaobaoFenxiaoProductSkuDeleteRequest,session string) (*response.TaobaoFenxiaoProductSkuDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.sku.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductSkuDeleteResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductSkuDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品sku添加接口
*/
func (ability *Defaultability) TaobaoFenxiaoProductSkuAdd(req *request.TaobaoFenxiaoProductSkuAddRequest,session string) (*response.TaobaoFenxiaoProductSkuAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.sku.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductSkuAddResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductSkuAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品sku编辑接口
*/
func (ability *Defaultability) TaobaoFenxiaoProductSkuUpdate(req *request.TaobaoFenxiaoProductSkuUpdateRequest,session string) (*response.TaobaoFenxiaoProductSkuUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.sku.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductSkuUpdateResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductSkuUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    SKU查询接口
*/
func (ability *Defaultability) TaobaoFenxiaoProductSkusGet(req *request.TaobaoFenxiaoProductSkusGetRequest,session string) (*response.TaobaoFenxiaoProductSkusGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.skus.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductSkusGetResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductSkusGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品图片上传
*/
func (ability *Defaultability) TaobaoFenxiaoProductImageUpload(req *request.TaobaoFenxiaoProductImageUploadRequest,session string) (*response.TaobaoFenxiaoProductImageUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.image.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductImageUploadResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductImageUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品图片删除
*/
func (ability *Defaultability) TaobaoFenxiaoProductImageDelete(req *request.TaobaoFenxiaoProductImageDeleteRequest,session string) (*response.TaobaoFenxiaoProductImageDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.fenxiao.product.image.delete",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoFenxiaoProductImageDeleteResponse{}
    if(err != nil){
        log.Println("taobaoFenxiaoProductImageDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品发布
*/
func (ability *Defaultability) AlibabaItemPublishSubmit(req *request.AlibabaItemPublishSubmitRequest,session string) (*response.AlibabaItemPublishSubmitResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.publish.submit",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemPublishSubmitResponse{}
    if(err != nil){
        log.Println("alibabaItemPublishSubmit error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品编辑获取schema信息
*/
func (ability *Defaultability) AlibabaItemEditSchemaGet(req *request.AlibabaItemEditSchemaGetRequest,session string) (*response.AlibabaItemEditSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.edit.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemEditSchemaGetResponse{}
    if(err != nil){
        log.Println("alibabaItemEditSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品编辑提交schema信息
*/
func (ability *Defaultability) AlibabaItemEditSubmit(req *request.AlibabaItemEditSubmitRequest,session string) (*response.AlibabaItemEditSubmitResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.edit.submit",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemEditSubmitResponse{}
    if(err != nil){
        log.Println("alibabaItemEditSubmit error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品下架
*/
func (ability *Defaultability) AlibabaItemOperateDownshelf(req *request.AlibabaItemOperateDownshelfRequest,session string) (*response.AlibabaItemOperateDownshelfResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.operate.downshelf",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemOperateDownshelfResponse{}
    if(err != nil){
        log.Println("alibabaItemOperateDownshelf error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品上架
*/
func (ability *Defaultability) AlibabaItemOperateUpshelf(req *request.AlibabaItemOperateUpshelfRequest,session string) (*response.AlibabaItemOperateUpshelfResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.operate.upshelf",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemOperateUpshelfResponse{}
    if(err != nil){
        log.Println("alibabaItemOperateUpshelf error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品级联属性信息获取
*/
func (ability *Defaultability) AlibabaItemPublishPropsGet(req *request.AlibabaItemPublishPropsGetRequest,session string) (*response.AlibabaItemPublishPropsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.item.publish.props.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaItemPublishPropsGetResponse{}
    if(err != nil){
        log.Println("alibabaItemPublishPropsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫根据规则发布商品
*/
func (ability *Defaultability) TmallItemSchemaAdd(req *request.TmallItemSchemaAddRequest,session string) (*response.TmallItemSchemaAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.schema.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemSchemaAddResponse{}
    if(err != nil){
        log.Println("tmallItemSchemaAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫发布商品规则获取
*/
func (ability *Defaultability) TmallItemAddSchemaGet(req *request.TmallItemAddSchemaGetRequest,session string) (*response.TmallItemAddSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.add.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemAddSchemaGetResponse{}
    if(err != nil){
        log.Println("tmallItemAddSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品发布规则获取接口
*/
func (ability *Defaultability) TmallProductAddSchemaGet(req *request.TmallProductAddSchemaGetRequest,session string) (*response.TmallProductAddSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.product.add.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallProductAddSchemaGetResponse{}
    if(err != nil){
        log.Println("tmallProductAddSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    批次库存查询接口
*/
func (ability *Defaultability) TaobaoWlbItemBatchQuery(req *request.TaobaoWlbItemBatchQueryRequest,session string) (*response.TaobaoWlbItemBatchQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.wlb.item.batch.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoWlbItemBatchQueryResponse{}
    if(err != nil){
        log.Println("taobaoWlbItemBatchQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取匹配产品规则
*/
func (ability *Defaultability) TmallProductMatchSchemaGet(req *request.TmallProductMatchSchemaGetRequest,session string) (*response.TmallProductMatchSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.product.match.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallProductMatchSchemaGetResponse{}
    if(err != nil){
        log.Println("tmallProductMatchSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    product匹配接口
*/
func (ability *Defaultability) TmallProductSchemaMatch(req *request.TmallProductSchemaMatchRequest,session string) (*response.TmallProductSchemaMatchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.product.schema.match",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallProductSchemaMatchResponse{}
    if(err != nil){
        log.Println("tmallProductSchemaMatch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    使用Schema文件发布一个产品
*/
func (ability *Defaultability) TmallProductSchemaAdd(req *request.TmallProductSchemaAddRequest,session string) (*response.TmallProductSchemaAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.product.schema.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallProductSchemaAddResponse{}
    if(err != nil){
        log.Println("tmallProductSchemaAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    通过主账号登陆态分页查询子账号列表
*/
func (ability *Defaultability) TaobaoSellercenterSubusersPage(req *request.TaobaoSellercenterSubusersPageRequest,session string) (*response.TaobaoSellercenterSubusersPageResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.sellercenter.subusers.page",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSellercenterSubusersPageResponse{}
    if(err != nil){
        log.Println("taobaoSellercenterSubusersPage error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    创建物流订单
*/
func (ability *Defaultability) TaobaoLogisticsOrderCreate(req *request.TaobaoLogisticsOrderCreateRequest,session string) (*response.TaobaoLogisticsOrderCreateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.logistics.order.create",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoLogisticsOrderCreateResponse{}
    if(err != nil){
        log.Println("taobaoLogisticsOrderCreate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    分页获取指定账户的子账号简易信息列表（新isv建议使用taobao.sellercenter.subusers.page接口）
*/
func (ability *Defaultability) TaobaoSubusersPage(req *request.TaobaoSubusersPageRequest,session string) (*response.TaobaoSubusersPageResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subusers.page",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubusersPageResponse{}
    if(err != nil){
        log.Println("taobaoSubusersPage error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫根据规则编辑商品
*/
func (ability *Defaultability) TmallItemSchemaUpdate(req *request.TmallItemSchemaUpdateRequest,session string) (*response.TmallItemSchemaUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.schema.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemSchemaUpdateResponse{}
    if(err != nil){
        log.Println("tmallItemSchemaUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫编辑商品规则获取
*/
func (ability *Defaultability) TmallItemUpdateSchemaGet(req *request.TmallItemUpdateSchemaGetRequest,session string) (*response.TmallItemUpdateSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.update.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemUpdateSchemaGetResponse{}
    if(err != nil){
        log.Println("tmallItemUpdateSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据当前子账号登陆态，获取该子账号基本信息
*/
func (ability *Defaultability) TaobaoSubusersInfoQuery(req *request.TaobaoSubusersInfoQueryRequest,session string) (*response.TaobaoSubusersInfoQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.subusers.info.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoSubusersInfoQueryResponse{}
    if(err != nil){
        log.Println("taobaoSubusersInfoQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取可用宝贝描述规范化模块
*/
func (ability *Defaultability) TaobaoItemAnchorGet(req *request.TaobaoItemAnchorGetRequest,session string) (*response.TaobaoItemAnchorGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.item.anchor.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoItemAnchorGetResponse{}
    if(err != nil){
        log.Println("taobaoItemAnchorGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    同意退款
*/
func (ability *Defaultability) TaobaoRpRefundsAgree(req *request.TaobaoRpRefundsAgreeRequest,session string) (*response.TaobaoRpRefundsAgreeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.rp.refunds.agree",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRpRefundsAgreeResponse{}
    if(err != nil){
        log.Println("taobaoRpRefundsAgree error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    卖家同意退货
*/
func (ability *Defaultability) TaobaoRpReturngoodsAgree(req *request.TaobaoRpReturngoodsAgreeRequest,session string) (*response.TaobaoRpReturngoodsAgreeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.rp.returngoods.agree",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoRpReturngoodsAgreeResponse{}
    if(err != nil){
        log.Println("taobaoRpReturngoodsAgree error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    查询卖家已卖出的增量交易数据（根据入库时间）
*/
func (ability *Defaultability) TaobaoTradesSoldIncrementvGet(req *request.TaobaoTradesSoldIncrementvGetRequest,session string) (*response.TaobaoTradesSoldIncrementvGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.trades.sold.incrementv.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoTradesSoldIncrementvGetResponse{}
    if(err != nil){
        log.Println("taobaoTradesSoldIncrementvGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    图片获取
*/
func (ability *Defaultability) TaobaoPicturePicturesGet(req *request.TaobaoPicturePicturesGetRequest,session string) (*response.TaobaoPicturePicturesGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.pictures.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPicturePicturesGetResponse{}
    if(err != nil){
        log.Println("taobaoPicturePicturesGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    图片总数查询
*/
func (ability *Defaultability) TaobaoPicturePicturesCount(req *request.TaobaoPicturePicturesCountRequest,session string) (*response.TaobaoPicturePicturesCountResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.picture.pictures.count",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoPicturePicturesCountResponse{}
    if(err != nil){
        log.Println("taobaoPicturePicturesCount error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
