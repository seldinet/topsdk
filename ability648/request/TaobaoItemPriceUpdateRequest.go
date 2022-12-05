package request

import (
        "topsdk/util"
    )

type TaobaoItemPriceUpdateRequest struct {
    /*
        商品数字ID，该参数必须     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        商品上传后的状态。可选值:onsale（出售中）,instock（库中），如果同时更新商品状态为出售中及list_time为将来的时间，则商品还是处于定时上架的状态, 此时item.is_timing为true     */
    ApproveStatus  *string `json:"approve_status,omitempty" required:"false" `
    /*
        自动重发。可选值:true,false;     */
    AutoRepost  *bool `json:"auto_repost,omitempty" required:"false" `
    /*
        叶子类目id     */
    Cid  *int64 `json:"cid,omitempty" required:"false" `
    /*
        商品描述. 字数要大于5个字符，小于25000个字符 ，受违禁词控制     */
    Desc  *string `json:"desc,omitempty" required:"false" `
    /*
        有效期。可选值:7,14;单位:天;     */
    ValidThru  *int64 `json:"valid_thru,omitempty" required:"false" `
    /*
        平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分, 注:post_fee,express_fee,ems_fee需一起填写     */
    PostFee  *string `json:"post_fee,omitempty" required:"false" `
    /*
        快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分     */
    ExpressFee  *string `json:"express_fee,omitempty" required:"false" `
    /*
        ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分     */
    EmsFee  *string `json:"ems_fee,omitempty" required:"false" `
    /*
        是否有保修。可选值:true,false;     */
    HasWarranty  *bool `json:"has_warranty,omitempty" required:"false" `
    /*
        是否有发票。可选值:true,false (商城卖家此字段必须为true)     */
    HasInvoice  *bool `json:"has_invoice,omitempty" required:"false" `
    /*
        加价幅度 如果为0，代表系统代理幅度     */
    Increment  *string `json:"increment,omitempty" required:"false" `
    /*
        商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。     */
    Props  *string `json:"props,omitempty" required:"false" `
    /*
        商品数量，取值范围:0-999999的整数。且需要等于Sku所有数量的和     */
    Num  *int64 `json:"num,omitempty" required:"false" `
    /*
        运费承担方式。运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);     */
    FreightPayer  *string `json:"freight_payer,omitempty" required:"false" `
    /*
        重新关联商品与店铺类目，结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。     */
    SellerCids  *string `json:"seller_cids,omitempty" required:"false" `
    /*
        橱窗推荐。可选值:true,false;     */
    HasShowcase  *bool `json:"has_showcase,omitempty" required:"false" `
    /*
        上架时间。不论是更新架下的商品还是出售中的商品，如果这个字段小于当前时间则直接上架商品，并且上架的时间为更新商品的时间，此时item.is_timing为false，如果大于当前时间则宝贝会下架进入定时上架的宝贝中。     */
    ListTime  *util.LocalTime `json:"list_time,omitempty" required:"false" `
    /*
        商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。     */
    StuffStatus  *string `json:"stuff_status,omitempty" required:"false" `
    /*
        宝贝标题. 不能超过60字符,受违禁词控制     */
    Title  *string `json:"title,omitempty" required:"false" `
    /*
        商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。     */
    Price  *string `json:"price,omitempty" required:"false" `
    /*
        支持会员打折。可选值:true,false;     */
    HasDiscount  *bool `json:"has_discount,omitempty" required:"false" `
    /*
        商家编码     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.postages.get获得当前会话用户的所有邮费模板）     */
    PostageId  *int64 `json:"postage_id,omitempty" required:"false" `
    /*
        商品所属的产品ID(B商家发布商品需要用)     */
    ProductId  *int64 `json:"product_id,omitempty" required:"false" `
    /*
        是否在淘宝上显示     */
    IsTaobao  *bool `json:"is_taobao,omitempty" required:"false" `
    /*
        是否在外店显示     */
    IsEx  *bool `json:"is_ex,omitempty" required:"false" `
    /*
        是否是3D     */
    Is3D  *bool `json:"is_3D,omitempty" required:"false" `
    /*
        商品的积分返点比例。如：5 表示返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50%     */
    AuctionPoint  *int64 `json:"auction_point,omitempty" required:"false" `
    /*
        属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过512字节     */
    PropertyAlias  *string `json:"property_alias,omitempty" required:"false" `
    /*
        商品文字的版本，繁体传入”zh_HK”，简体传入”zh_CN”     */
    Lang  *string `json:"lang,omitempty" required:"false" `
    /*
        商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path     */
    PicPath  *string `json:"pic_path,omitempty" required:"false" `
    /*
        是否替换sku     */
    IsReplaceSku  *bool `json:"is_replace_sku,omitempty" required:"false" `
    /*
        商品图片。类型:JPG,GIF;最大长度:500k     */
    Image  *[]byte `json:"image,omitempty" required:"false" `
    /*
        代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： 
no_mark(不做类型标记) 
time_card(点卡软件代充) 
fee_card(话费软件代充)     */
    AutoFill  *string `json:"auto_fill,omitempty" required:"false" `
    /*
        是否承诺退换货服务!虚拟商品无须设置此项!     */
    SellPromise  *bool `json:"sell_promise,omitempty" required:"false" `
    /*
        货到付款运费模板ID     */
    CodPostageId  *int64 `json:"cod_postage_id,omitempty" required:"false" `
    /*
        实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记     */
    IsLightningConsignment  *bool `json:"is_lightning_consignment,omitempty" required:"false" `
    /*
        商品的重量(商超卖家专用字段)     */
    Weight  *int64 `json:"weight,omitempty" required:"false" `
    /*
        宝贝形态:
1代表电子券;0或其他代表实物     */
    Shape  *string `json:"shape,omitempty" required:"false" `
    /*
        商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。     */
    IsXinpin  *bool `json:"is_xinpin,omitempty" required:"false" `
    /*
        商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存 defalutValue��0    */
    SubStock  *int64 `json:"sub_stock,omitempty" required:"false" `
    /*
        忽略警告提示.     */
    Ignorewarning  *string `json:"ignorewarning,omitempty" required:"false" `
    /*
        用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。     */
    InputPids  *[]string `json:"input_pids,omitempty" required:"false" `
    /*
        用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。     */
    InputStr  *[]string `json:"input_str,omitempty" required:"false" `
    /*
        更新的Sku的属性串，调用taobao.itemprops.get获取类目属性，如果属性是销售属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid;pid:vid。该字段内的销售属性也需要在props字段填写 。如果更新时有对Sku进行操作，则Sku的properties一定要传入。     */
    SkuProperties  *string `json:"sku_properties,omitempty" required:"false" `
    /*
        更新的Sku的数量串，结构如：num1,num2,num3 如:2,3,4     */
    SkuQuantities  *string `json:"sku_quantities,omitempty" required:"false" `
    /*
        更新的Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分     */
    SkuPrices  *string `json:"sku_prices,omitempty" required:"false" `
    /*
        Sku的外部id串，结构如：1234,1342,… sku_properties, sku_quantities, sku_prices, sku_outer_ids在输入数据时要一一对应，如果没有sku_outer_ids也要写上这个参数，入参是","(这个是两个sku的示列，逗号数应该是sku个数减1)；该参数最大长度是512个字节     */
    SkuOuterIds  *string `json:"sku_outer_ids,omitempty" required:"false" `
    /*
        所在地省份。如浙江 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到     */
    LocationState  *string `json:"location.state,omitempty" required:"false" `
    /*
        所在地城市。如杭州 具体可以下载http://dl.open.taobao.com/sdk/商品城市列表.rar 取到     */
    LocationCity  *string `json:"location.city,omitempty" required:"false" `
}

func (s *TaobaoItemPriceUpdateRequest) SetNumIid(v int64) *TaobaoItemPriceUpdateRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetApproveStatus(v string) *TaobaoItemPriceUpdateRequest {
    s.ApproveStatus = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetAutoRepost(v bool) *TaobaoItemPriceUpdateRequest {
    s.AutoRepost = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetCid(v int64) *TaobaoItemPriceUpdateRequest {
    s.Cid = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetDesc(v string) *TaobaoItemPriceUpdateRequest {
    s.Desc = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetValidThru(v int64) *TaobaoItemPriceUpdateRequest {
    s.ValidThru = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetPostFee(v string) *TaobaoItemPriceUpdateRequest {
    s.PostFee = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetExpressFee(v string) *TaobaoItemPriceUpdateRequest {
    s.ExpressFee = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetEmsFee(v string) *TaobaoItemPriceUpdateRequest {
    s.EmsFee = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetHasWarranty(v bool) *TaobaoItemPriceUpdateRequest {
    s.HasWarranty = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetHasInvoice(v bool) *TaobaoItemPriceUpdateRequest {
    s.HasInvoice = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetIncrement(v string) *TaobaoItemPriceUpdateRequest {
    s.Increment = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetProps(v string) *TaobaoItemPriceUpdateRequest {
    s.Props = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetNum(v int64) *TaobaoItemPriceUpdateRequest {
    s.Num = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetFreightPayer(v string) *TaobaoItemPriceUpdateRequest {
    s.FreightPayer = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetSellerCids(v string) *TaobaoItemPriceUpdateRequest {
    s.SellerCids = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetHasShowcase(v bool) *TaobaoItemPriceUpdateRequest {
    s.HasShowcase = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetListTime(v util.LocalTime) *TaobaoItemPriceUpdateRequest {
    s.ListTime = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetStuffStatus(v string) *TaobaoItemPriceUpdateRequest {
    s.StuffStatus = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetTitle(v string) *TaobaoItemPriceUpdateRequest {
    s.Title = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetPrice(v string) *TaobaoItemPriceUpdateRequest {
    s.Price = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetHasDiscount(v bool) *TaobaoItemPriceUpdateRequest {
    s.HasDiscount = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetOuterId(v string) *TaobaoItemPriceUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetPostageId(v int64) *TaobaoItemPriceUpdateRequest {
    s.PostageId = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetProductId(v int64) *TaobaoItemPriceUpdateRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetIsTaobao(v bool) *TaobaoItemPriceUpdateRequest {
    s.IsTaobao = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetIsEx(v bool) *TaobaoItemPriceUpdateRequest {
    s.IsEx = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetIs3D(v bool) *TaobaoItemPriceUpdateRequest {
    s.Is3D = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetAuctionPoint(v int64) *TaobaoItemPriceUpdateRequest {
    s.AuctionPoint = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetPropertyAlias(v string) *TaobaoItemPriceUpdateRequest {
    s.PropertyAlias = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetLang(v string) *TaobaoItemPriceUpdateRequest {
    s.Lang = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetPicPath(v string) *TaobaoItemPriceUpdateRequest {
    s.PicPath = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetIsReplaceSku(v bool) *TaobaoItemPriceUpdateRequest {
    s.IsReplaceSku = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetImage(v []byte) *TaobaoItemPriceUpdateRequest {
    s.Image = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetAutoFill(v string) *TaobaoItemPriceUpdateRequest {
    s.AutoFill = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetSellPromise(v bool) *TaobaoItemPriceUpdateRequest {
    s.SellPromise = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetCodPostageId(v int64) *TaobaoItemPriceUpdateRequest {
    s.CodPostageId = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetIsLightningConsignment(v bool) *TaobaoItemPriceUpdateRequest {
    s.IsLightningConsignment = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetWeight(v int64) *TaobaoItemPriceUpdateRequest {
    s.Weight = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetShape(v string) *TaobaoItemPriceUpdateRequest {
    s.Shape = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetIsXinpin(v bool) *TaobaoItemPriceUpdateRequest {
    s.IsXinpin = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetSubStock(v int64) *TaobaoItemPriceUpdateRequest {
    s.SubStock = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetIgnorewarning(v string) *TaobaoItemPriceUpdateRequest {
    s.Ignorewarning = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetInputPids(v []string) *TaobaoItemPriceUpdateRequest {
    s.InputPids = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetInputStr(v []string) *TaobaoItemPriceUpdateRequest {
    s.InputStr = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetSkuProperties(v string) *TaobaoItemPriceUpdateRequest {
    s.SkuProperties = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetSkuQuantities(v string) *TaobaoItemPriceUpdateRequest {
    s.SkuQuantities = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetSkuPrices(v string) *TaobaoItemPriceUpdateRequest {
    s.SkuPrices = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetSkuOuterIds(v string) *TaobaoItemPriceUpdateRequest {
    s.SkuOuterIds = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetLocationState(v string) *TaobaoItemPriceUpdateRequest {
    s.LocationState = &v
    return s
}
func (s *TaobaoItemPriceUpdateRequest) SetLocationCity(v string) *TaobaoItemPriceUpdateRequest {
    s.LocationCity = &v
    return s
}

func (req *TaobaoItemPriceUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    if(req.ApproveStatus != nil) {
        paramMap["approve_status"] = *req.ApproveStatus
    }
    if(req.AutoRepost != nil) {
        paramMap["auto_repost"] = *req.AutoRepost
    }
    if(req.Cid != nil) {
        paramMap["cid"] = *req.Cid
    }
    if(req.Desc != nil) {
        paramMap["desc"] = *req.Desc
    }
    if(req.ValidThru != nil) {
        paramMap["valid_thru"] = *req.ValidThru
    }
    if(req.PostFee != nil) {
        paramMap["post_fee"] = *req.PostFee
    }
    if(req.ExpressFee != nil) {
        paramMap["express_fee"] = *req.ExpressFee
    }
    if(req.EmsFee != nil) {
        paramMap["ems_fee"] = *req.EmsFee
    }
    if(req.HasWarranty != nil) {
        paramMap["has_warranty"] = *req.HasWarranty
    }
    if(req.HasInvoice != nil) {
        paramMap["has_invoice"] = *req.HasInvoice
    }
    if(req.Increment != nil) {
        paramMap["increment"] = *req.Increment
    }
    if(req.Props != nil) {
        paramMap["props"] = *req.Props
    }
    if(req.Num != nil) {
        paramMap["num"] = *req.Num
    }
    if(req.FreightPayer != nil) {
        paramMap["freight_payer"] = *req.FreightPayer
    }
    if(req.SellerCids != nil) {
        paramMap["seller_cids"] = *req.SellerCids
    }
    if(req.HasShowcase != nil) {
        paramMap["has_showcase"] = *req.HasShowcase
    }
    if(req.ListTime != nil) {
        paramMap["list_time"] = *req.ListTime
    }
    if(req.StuffStatus != nil) {
        paramMap["stuff_status"] = *req.StuffStatus
    }
    if(req.Title != nil) {
        paramMap["title"] = *req.Title
    }
    if(req.Price != nil) {
        paramMap["price"] = *req.Price
    }
    if(req.HasDiscount != nil) {
        paramMap["has_discount"] = *req.HasDiscount
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.PostageId != nil) {
        paramMap["postage_id"] = *req.PostageId
    }
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.IsTaobao != nil) {
        paramMap["is_taobao"] = *req.IsTaobao
    }
    if(req.IsEx != nil) {
        paramMap["is_ex"] = *req.IsEx
    }
    if(req.Is3D != nil) {
        paramMap["is_3D"] = *req.Is3D
    }
    if(req.AuctionPoint != nil) {
        paramMap["auction_point"] = *req.AuctionPoint
    }
    if(req.PropertyAlias != nil) {
        paramMap["property_alias"] = *req.PropertyAlias
    }
    if(req.Lang != nil) {
        paramMap["lang"] = *req.Lang
    }
    if(req.PicPath != nil) {
        paramMap["pic_path"] = *req.PicPath
    }
    if(req.IsReplaceSku != nil) {
        paramMap["is_replace_sku"] = *req.IsReplaceSku
    }
    if(req.AutoFill != nil) {
        paramMap["auto_fill"] = *req.AutoFill
    }
    if(req.SellPromise != nil) {
        paramMap["sell_promise"] = *req.SellPromise
    }
    if(req.CodPostageId != nil) {
        paramMap["cod_postage_id"] = *req.CodPostageId
    }
    if(req.IsLightningConsignment != nil) {
        paramMap["is_lightning_consignment"] = *req.IsLightningConsignment
    }
    if(req.Weight != nil) {
        paramMap["weight"] = *req.Weight
    }
    if(req.Shape != nil) {
        paramMap["shape"] = *req.Shape
    }
    if(req.IsXinpin != nil) {
        paramMap["is_xinpin"] = *req.IsXinpin
    }
    if(req.SubStock != nil) {
        paramMap["sub_stock"] = *req.SubStock
    }
    if(req.Ignorewarning != nil) {
        paramMap["ignorewarning"] = *req.Ignorewarning
    }
    if(req.InputPids != nil) {
        paramMap["input_pids"] = util.ConvertBasicList(*req.InputPids)
    }
    if(req.InputStr != nil) {
        paramMap["input_str"] = util.ConvertBasicList(*req.InputStr)
    }
    if(req.SkuProperties != nil) {
        paramMap["sku_properties"] = *req.SkuProperties
    }
    if(req.SkuQuantities != nil) {
        paramMap["sku_quantities"] = *req.SkuQuantities
    }
    if(req.SkuPrices != nil) {
        paramMap["sku_prices"] = *req.SkuPrices
    }
    if(req.SkuOuterIds != nil) {
        paramMap["sku_outer_ids"] = *req.SkuOuterIds
    }
    if(req.LocationState != nil) {
        paramMap["location.state"] = *req.LocationState
    }
    if(req.LocationCity != nil) {
        paramMap["location.city"] = *req.LocationCity
    }
    return paramMap
}

func (req *TaobaoItemPriceUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}