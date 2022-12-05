package request

import (
        "topsdk/util"
    )

type TaobaoItemUpdateRequest struct {
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
        加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。     */
    Increment  *string `json:"increment,omitempty" required:"false" `
    /*
        商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值。     */
    Props  *string `json:"props,omitempty" required:"false" `
    /*
        商品数量，取值范围:0-900000000的整数。且需要等于Sku所有数量的和 拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。     */
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
        上架时间。大于当前时间则宝贝会下架进入定时上架的宝贝中。     */
    ListTime  *util.LocalTime `json:"list_time,omitempty" required:"false" `
    /*
        商品新旧程度。可选值:new（全新）,unused（闲置）,second（二手）。     */
    StuffStatus  *string `json:"stuff_status,omitempty" required:"false" `
    /*
        宝贝标题. 不能超过30字符,受违禁词控制     */
    Title  *string `json:"title,omitempty" required:"false" `
    /*
        商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。 拍卖商品对应的起拍价。     */
    Price  *string `json:"price,omitempty" required:"false" `
    /*
        支持会员打折。可选值:true,false;     */
    HasDiscount  *bool `json:"has_discount,omitempty" required:"false" `
    /*
        商家编码     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.templates.get获得当前会话用户的所有邮费模板）     */
    PostageId  *int64 `json:"postage_id,omitempty" required:"false" `
    /*
        商品所属的产品ID(B商家发布商品需要用)     */
    ProductId  *int64 `json:"product_id,omitempty" required:"false" `
    /*
        是否在淘宝上显示（如果传FALSE，则在淘宝主站无法显示该商品）     */
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
        属性值别名。如pid:vid:别名;pid1:vid1:别名1， pid:属性id vid:值id。总长度不超过800个字符，如"123:333:你好"，引号内的是10个字符。     */
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
        商品图片。类型:JPG,GIF;最大长度:3M     */
    Image  *[]byte `json:"image,omitempty" required:"false" `
    /*
        代充商品类型。只有少数类目下的商品可以标记上此字段，具体哪些类目可以上传可以通过taobao.itemcat.features.get获得。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充)     */
    AutoFill  *string `json:"auto_fill,omitempty" required:"false" `
    /*
        是否承诺退换货服务!虚拟商品无须设置此项!     */
    SellPromise  *bool `json:"sell_promise,omitempty" required:"false" `
    /*
        货到付款运费模板ID该字段已经废弃，货到付款模板已经集成到运费模板中。     */
    CodPostageId  *int64 `json:"cod_postage_id,omitempty" required:"false" `
    /*
        实物闪电发货。注意：在售的闪电发货产品不允许取消闪电发货，需要先下架商品才能取消闪电发货标记     */
    IsLightningConsignment  *bool `json:"is_lightning_consignment,omitempty" required:"false" `
    /*
        商品的重量(商超卖家专用字段)     */
    Weight  *int64 `json:"weight,omitempty" required:"false" `
    /*
        宝贝形态:1代表电子券;0或其他代表实物     */
    Shape  *string `json:"shape,omitempty" required:"false" `
    /*
        商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置参数就保持原有值。     */
    IsXinpin  *bool `json:"is_xinpin,omitempty" required:"false" `
    /*
        商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改 集市卖家默认拍下减库存; 商城卖家默认付款减库存 defalutValue��0    */
    SubStock  *int64 `json:"sub_stock,omitempty" required:"false" `
    /*
        宝贝特征值，格式为：【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp,是指尺码库对应的key     */
    Features  *string `json:"features,omitempty" required:"false" `
    /*
        景区门票类宝贝编辑时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付     */
    ScenicTicketPayWay  *int64 `json:"scenic_ticket_pay_way,omitempty" required:"false" `
    /*
        景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100     */
    ScenicTicketBookCost  *string `json:"scenic_ticket_book_cost,omitempty" required:"false" `
    /*
        支持宝贝信息的删除,如需删除对应的食品安全信息中的储藏方法、保质期， 则应该设置此参数的值为：food_security.plan_storage,food_security.period; 各个参数的名称之间用【,】分割, 如果对应的参数有设置过值，即使在这个列表中，也不会被删除; 目前支持此功能的宝贝信息如下：食品安全信息所有字段、电子交易凭证字段（locality_life，locality_life.verification，locality_life.refund_ratio，locality_life.network_id ，locality_life.onsale_auto_refund_ratio）。支持对全球购宝贝信息的清除（字符串中包含global_stock）     */
    EmptyFields  *string `json:"empty_fields,omitempty" required:"false" `
    /*
        全球购商品采购地（库存类型）全球购商品有两种库存类型：现货和代购 参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用     */
    GlobalStockType  *string `json:"global_stock_type,omitempty" required:"false" `
    /*
        全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值为（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 澳洲, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾, 其他）     */
    GlobalStockCountry  *string `json:"global_stock_country,omitempty" required:"false" `
    /*
        全球购商品发货地，发货地现在有两种类型：“国内”和“海外及港澳台”，参数值为1时代表“国内”，值为2时代表“海外及港澳台”     */
    GlobalStockDeliveryPlace  *string `json:"global_stock_delivery_place,omitempty" required:"false" `
    /*
        全球购商品卖家包税承诺，当值为true时，代表卖家承诺包税。     */
    GlobalStockTaxFreePromise  *bool `json:"global_stock_tax_free_promise,omitempty" required:"false" `
    /*
        商品的重量，用于按重量计费的运费模板。注意：单位为kg。 只能传入数值类型（包含小数），不能带单位，单位默认为kg。 在编辑时候，如果需要在商品里删除重量的信息，就需要将值设置为0     */
    ItemWeight  *string `json:"item_weight,omitempty" required:"false" `
    /*
        表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)在编辑的时候，如果需要删除体积属性，请设置该值为0，如bulk:0     */
    ItemSize  *string `json:"item_size,omitempty" required:"false" `
    /*
        针对当前商品的自定义属性值     */
    InputCustomCpv  *string `json:"input_custom_cpv,omitempty" required:"false" `
    /*
        针对当前商品的标准属性值的补充说明，让买家更加了解商品信息减少交易纠纷     */
    CpvMemo  *string `json:"cpv_memo,omitempty" required:"false" `
    /*
        该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；<br>注意：使用该API修改商品其它属性如标题title时，如需保持商品不支持7天无理由退货状态，该字段需传入0 。     */
    Newprepay  *string `json:"newprepay,omitempty" required:"false" `
    /*
        商品条形码     */
    Barcode  *string `json:"barcode,omitempty" required:"false" `
    /*
        商品卖点信息，最长150个字符。天猫和集市都可用     */
    SellPoint  *string `json:"sell_point,omitempty" required:"false" `
    /*
        商品资质信息     */
    Qualification  *string `json:"qualification,omitempty" required:"false" `
    /*
        汽车O2O绑定线下服务标记，如不为空，表示关联服务，否则，不关联服务。     */
    O2oBindService  *bool `json:"o2o_bind_service,omitempty" required:"false" `
    /*
        忽略警告提示.     */
    Ignorewarning  *string `json:"ignorewarning,omitempty" required:"false" `
    /*
        售后说明模板id     */
    AfterSaleId  *int64 `json:"after_sale_id,omitempty" required:"false" `
    /*
        基础色数据，淘宝不使用     */
    ChangeProp  *string `json:"change_prop,omitempty" required:"false" `
    /*
        已废弃     */
    DescModules  *string `json:"desc_modules,omitempty" required:"false" `
    /*
        是否是线下商品。1：线上商品（默认值）；2：线上或线下商品；3：线下商品。     */
    IsOffline  *string `json:"is_offline,omitempty" required:"false" `
    /*
        无线的宝贝描述     */
    WirelessDesc  *string `json:"wireless_desc,omitempty" required:"false" `
    /*
        天猫超市扩展字段，天猫超市专用。     */
    ChaoshiExtendsInfo  *string `json:"chaoshi_extends_info,omitempty" required:"false" `
    /*
        手机类目spu 产品信息确认声明     */
    SpuConfirm  *bool `json:"spu_confirm,omitempty" required:"false" `
    /*
        主图视频id     */
    VideoId  *int64 `json:"video_id,omitempty" required:"false" `
    /*
        主图视频互动信息id，必须有主图视频id才能传互动信息id     */
    InteractiveId  *int64 `json:"interactive_id,omitempty" required:"false" `
    /*
        淘宝租赁扩展信息     */
    LeaseExtendsInfo  *string `json:"lease_extends_info,omitempty" required:"false" `
    /*
        用户自行输入的类目属性ID串，结构："pid1,pid2,pid3"，如："20000"（表示品牌） 注：通常一个类目下用户可输入的关键属性不超过1个。     */
    InputPids  *[]string `json:"input_pids,omitempty" required:"false" `
    /*
        用户自行输入的子属性名和属性值，结构:"父属性值;一级子属性名;一级子属性值;二级子属性名;自定义输入值,....",如：“耐克;耐克系列;科比系列;科比系列;2K5,Nike乔丹鞋;乔丹系列;乔丹鞋系列;乔丹鞋系列;json5”，多个自定义属性用','分割，input_str需要与input_pids一一对应，注：通常一个类目下用户可输入的关键属性不超过1个。所有属性别名加起来不能超过3999字节。此处不可以使用“其他”、“其它”和“其她”这三个词。     */
    InputStr  *[]string `json:"input_str,omitempty" required:"false" `
    /*
        更新的sku的属性串，调用taobao.itemprops.get获取。格式:pid1:vid;pid2:vid,多个sku属性之间用逗号分隔。该字段内的属性需要在props字段同时包含。如果新增商品包含了sku，则此字段一定要传入,字段长度要控制在512个字节以内。     */
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
        sku层面的条形码，多个SKU情况，与SKU价格库存格式类似，用逗号分隔     */
    SkuBarcode  *string `json:"sku_barcode,omitempty" required:"false" `
    /*
        此参数暂时不起作用     */
    SkuSpecIds  *string `json:"sku_spec_ids,omitempty" required:"false" `
    /*
        此参数暂时不起作用     */
    SkuDeliveryTimes  *string `json:"sku_delivery_times,omitempty" required:"false" `
    /*
        家装建材类目，商品SKU的长度，正整数，单位为cm，部分类目必选。 数据和SKU一一对应，用,分隔，如：20,30,30     */
    SkuHdLength  *string `json:"sku_hd_length,omitempty" required:"false" `
    /*
        家装建材类目，商品SKU的高度，单位为cm，部分类目必选。 天猫和淘宝格式不同。天猫：可选值为："0-15", "15-25", "25-50", "50-60", "60-80", "80-120", "120-160", "160-200"。 数据和SKU一一对应，用,分隔，格式如：15-25,25-50,25-50。 淘宝：正整数，单位为cm,格式如：20,30,30     */
    SkuHdHeight  *string `json:"sku_hd_height,omitempty" required:"false" `
    /*
        家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。 数据和SKU一一对应，用,分隔，如：3,5,7     */
    SkuHdLampQuantity  *string `json:"sku_hd_lamp_quantity,omitempty" required:"false" `
    /*
        所在地省份。如浙江     */
    LocationState  *string `json:"location.state,omitempty" required:"false" `
    /*
        所在地城市。如杭州     */
    LocationCity  *string `json:"location.city,omitempty" required:"false" `
    /*
        生产许可证号     */
    FoodSecurityPrdLicenseNo  *string `json:"food_security.prd_license_no,omitempty" required:"false" `
    /*
        产品标准号     */
    FoodSecurityDesignCode  *string `json:"food_security.design_code,omitempty" required:"false" `
    /*
        厂名     */
    FoodSecurityFactory  *string `json:"food_security.factory,omitempty" required:"false" `
    /*
        厂址     */
    FoodSecurityFactorySite  *string `json:"food_security.factory_site,omitempty" required:"false" `
    /*
        厂家联系方式     */
    FoodSecurityContact  *string `json:"food_security.contact,omitempty" required:"false" `
    /*
        配料表     */
    FoodSecurityMix  *string `json:"food_security.mix,omitempty" required:"false" `
    /*
        储藏方法     */
    FoodSecurityPlanStorage  *string `json:"food_security.plan_storage,omitempty" required:"false" `
    /*
        保质期，默认有单位，传入数字     */
    FoodSecurityPeriod  *string `json:"food_security.period,omitempty" required:"false" `
    /*
        食品添加剂     */
    FoodSecurityFoodAdditive  *string `json:"food_security.food_additive,omitempty" required:"false" `
    /*
        供货商     */
    FoodSecuritySupplier  *string `json:"food_security.supplier,omitempty" required:"false" `
    /*
        生产开始日期，格式必须为yyyy-MM-dd     */
    FoodSecurityProductDateStart  *string `json:"food_security.product_date_start,omitempty" required:"false" `
    /*
        生产结束日期,格式必须为yyyy-MM-dd     */
    FoodSecurityProductDateEnd  *string `json:"food_security.product_date_end,omitempty" required:"false" `
    /*
        进货开始日期，要在生产日期之后，格式必须为yyyy-MM-dd     */
    FoodSecurityStockDateStart  *string `json:"food_security.stock_date_start,omitempty" required:"false" `
    /*
        进货结束日期，要在生产日期之后，格式必须为yyyy-MM-dd     */
    FoodSecurityStockDateEnd  *string `json:"food_security.stock_date_end,omitempty" required:"false" `
    /*
        健字号，保健品/膳食营养补充剂 这个类目下特有的信息，此类目下无需填写生产许可证编号（QS），如果填写了生产许可证编号（QS）将被忽略不保存；保存宝贝时，标题前会自动加上健字号产品名称一起作为宝贝标题；     */
    FoodSecurityHealthProductNo  *string `json:"food_security.health_product_no,omitempty" required:"false" `
    /*
        本地生活电子交易凭证业务，目前此字段只涉及到的信息为有效期;如果有效期为起止日期类型，此值为2012-08-06,2012-08-16如果有效期为【购买成功日 至】类型则格式为2012-08-16如果有效期为天数类型则格式为15     */
    LocalityLifeExpirydate  *string `json:"locality_life.expirydate,omitempty" required:"false" `
    /*
        网点ID,在参数empty_fields里设置locality_life.network_id可删除网点ID     */
    LocalityLifeNetworkId  *string `json:"locality_life.network_id,omitempty" required:"false" `
    /*
        码商信息，格式为 码商id:nick     */
    LocalityLifeMerchant  *string `json:"locality_life.merchant,omitempty" required:"false" `
    /*
        核销打款,1代表核销打款 0代表非核销打款; 在参数empty_fields里设置locality_life.verification可删除核销打款     */
    LocalityLifeVerification  *string `json:"locality_life.verification,omitempty" required:"false" `
    /*
        退款比例，百分比%前的数字,1-100的正整数值; 在参数empty_fields里设置locality_life.refund_ratio可删除退款比例     */
    LocalityLifeRefundRatio  *int64 `json:"locality_life.refund_ratio,omitempty" required:"false" `
    /*
        电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数     */
    LocalityLifeOnsaleAutoRefundRatio  *int64 `json:"locality_life.onsale_auto_refund_ratio,omitempty" required:"false" `
    /*
        编辑电子凭证宝贝时候表示是否使用邮寄0: 代表不使用邮寄；1：代表使用邮寄；如果不设置这个值，代表不使用邮寄     */
    LocalityLifeChooseLogis  *string `json:"locality_life.choose_logis,omitempty" required:"false" `
    /*
        退款码费承担方。发布电子凭证宝贝的时候会增加“退款码费承担方”配置项，可选填：(1)s（卖家承担） (2)b(买家承担)     */
    LocalityLifeRefundmafee  *string `json:"locality_life.refundmafee,omitempty" required:"false" `
    /*
        预约门店是否支持门店自提,1:是     */
    LocalityLifeObs  *string `json:"locality_life.obs,omitempty" required:"false" `
    /*
        电子凭证业务属性，数据字典是: 1、is_card:1 (暂时不用) 2、consume_way:4 （1 串码 ，4 身份证）3、consume_midmnick ：(核销放行账号:用户id-用户名，支持多个，用逗号分隔,例如 1234-测试账号35,1345-测试账号56）4、market:eticket (电子凭证商品标记) 5、has_pos:1 (1 表示商品配置线下门店，在detail上进行展示 ，没有或者其他值只不展示)格式是: k1:v2;k2:v2;........ 如：has_pos:1;market:eticket;consume_midmnick:901409638-OPPO;consume_way:4     */
    LocalityLifeEticket  *string `json:"locality_life.eticket,omitempty" required:"false" `
    /*
        电子凭证版本 新版电子凭证值:1     */
    LocalityLifeVersion  *string `json:"locality_life.version,omitempty" required:"false" `
    /*
        新版电子凭证包id     */
    LocalityLifePackageid  *string `json:"locality_life.packageid,omitempty" required:"false" `
    /*
        拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。     */
    PaimaiInfoMode  *int64 `json:"paimai_info.mode,omitempty" required:"false" `
    /*
        拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数。如果该参数不传或传入0则代表使用默认。     */
    PaimaiInfoDeposit  *int64 `json:"paimai_info.deposit,omitempty" required:"false" `
    /*
        降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。     */
    PaimaiInfoInterval  *int64 `json:"paimai_info.interval,omitempty" required:"false" `
    /*
        降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数     */
    PaimaiInfoReserve  *string `json:"paimai_info.reserve,omitempty" required:"false" `
    /*
        自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。     */
    PaimaiInfoValidHour  *int64 `json:"paimai_info.valid_hour,omitempty" required:"false" `
    /*
        自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。     */
    PaimaiInfoValidMinute  *int64 `json:"paimai_info.valid_minute,omitempty" required:"false" `
    /*
        订金     */
    MsPaymentPrice  *string `json:"ms_payment.price,omitempty" required:"false" `
    /*
        尾款可抵扣金额     */
    MsPaymentVoucherPrice  *string `json:"ms_payment.voucher_price,omitempty" required:"false" `
    /*
        参考价     */
    MsPaymentReferencePrice  *string `json:"ms_payment.reference_price,omitempty" required:"false" `
    /*
        商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11     */
    DeliveryTimeDeliveryTime  *string `json:"delivery_time.delivery_time,omitempty" required:"false" `
    /*
        发货时间类型：绝对发货时间或者相对发货时间     */
    DeliveryTimeDeliveryTimeType  *string `json:"delivery_time.delivery_time_type,omitempty" required:"false" `
    /*
        设置是否使用发货时间，商品级别，sku级别     */
    DeliveryTimeNeedDeliveryTime  *string `json:"delivery_time.need_delivery_time,omitempty" required:"false" `
}

func (s *TaobaoItemUpdateRequest) SetNumIid(v int64) *TaobaoItemUpdateRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetApproveStatus(v string) *TaobaoItemUpdateRequest {
    s.ApproveStatus = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetAutoRepost(v bool) *TaobaoItemUpdateRequest {
    s.AutoRepost = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetCid(v int64) *TaobaoItemUpdateRequest {
    s.Cid = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetDesc(v string) *TaobaoItemUpdateRequest {
    s.Desc = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetValidThru(v int64) *TaobaoItemUpdateRequest {
    s.ValidThru = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetPostFee(v string) *TaobaoItemUpdateRequest {
    s.PostFee = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetExpressFee(v string) *TaobaoItemUpdateRequest {
    s.ExpressFee = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetEmsFee(v string) *TaobaoItemUpdateRequest {
    s.EmsFee = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetHasWarranty(v bool) *TaobaoItemUpdateRequest {
    s.HasWarranty = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetHasInvoice(v bool) *TaobaoItemUpdateRequest {
    s.HasInvoice = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetIncrement(v string) *TaobaoItemUpdateRequest {
    s.Increment = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetProps(v string) *TaobaoItemUpdateRequest {
    s.Props = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetNum(v int64) *TaobaoItemUpdateRequest {
    s.Num = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFreightPayer(v string) *TaobaoItemUpdateRequest {
    s.FreightPayer = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSellerCids(v string) *TaobaoItemUpdateRequest {
    s.SellerCids = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetHasShowcase(v bool) *TaobaoItemUpdateRequest {
    s.HasShowcase = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetListTime(v util.LocalTime) *TaobaoItemUpdateRequest {
    s.ListTime = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetStuffStatus(v string) *TaobaoItemUpdateRequest {
    s.StuffStatus = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetTitle(v string) *TaobaoItemUpdateRequest {
    s.Title = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetPrice(v string) *TaobaoItemUpdateRequest {
    s.Price = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetHasDiscount(v bool) *TaobaoItemUpdateRequest {
    s.HasDiscount = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetOuterId(v string) *TaobaoItemUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetPostageId(v int64) *TaobaoItemUpdateRequest {
    s.PostageId = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetProductId(v int64) *TaobaoItemUpdateRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetIsTaobao(v bool) *TaobaoItemUpdateRequest {
    s.IsTaobao = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetIsEx(v bool) *TaobaoItemUpdateRequest {
    s.IsEx = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetIs3D(v bool) *TaobaoItemUpdateRequest {
    s.Is3D = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetAuctionPoint(v int64) *TaobaoItemUpdateRequest {
    s.AuctionPoint = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetPropertyAlias(v string) *TaobaoItemUpdateRequest {
    s.PropertyAlias = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLang(v string) *TaobaoItemUpdateRequest {
    s.Lang = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetPicPath(v string) *TaobaoItemUpdateRequest {
    s.PicPath = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetIsReplaceSku(v bool) *TaobaoItemUpdateRequest {
    s.IsReplaceSku = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetImage(v []byte) *TaobaoItemUpdateRequest {
    s.Image = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetAutoFill(v string) *TaobaoItemUpdateRequest {
    s.AutoFill = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSellPromise(v bool) *TaobaoItemUpdateRequest {
    s.SellPromise = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetCodPostageId(v int64) *TaobaoItemUpdateRequest {
    s.CodPostageId = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetIsLightningConsignment(v bool) *TaobaoItemUpdateRequest {
    s.IsLightningConsignment = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetWeight(v int64) *TaobaoItemUpdateRequest {
    s.Weight = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetShape(v string) *TaobaoItemUpdateRequest {
    s.Shape = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetIsXinpin(v bool) *TaobaoItemUpdateRequest {
    s.IsXinpin = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSubStock(v int64) *TaobaoItemUpdateRequest {
    s.SubStock = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFeatures(v string) *TaobaoItemUpdateRequest {
    s.Features = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetScenicTicketPayWay(v int64) *TaobaoItemUpdateRequest {
    s.ScenicTicketPayWay = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetScenicTicketBookCost(v string) *TaobaoItemUpdateRequest {
    s.ScenicTicketBookCost = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetEmptyFields(v string) *TaobaoItemUpdateRequest {
    s.EmptyFields = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetGlobalStockType(v string) *TaobaoItemUpdateRequest {
    s.GlobalStockType = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetGlobalStockCountry(v string) *TaobaoItemUpdateRequest {
    s.GlobalStockCountry = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetGlobalStockDeliveryPlace(v string) *TaobaoItemUpdateRequest {
    s.GlobalStockDeliveryPlace = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetGlobalStockTaxFreePromise(v bool) *TaobaoItemUpdateRequest {
    s.GlobalStockTaxFreePromise = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetItemWeight(v string) *TaobaoItemUpdateRequest {
    s.ItemWeight = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetItemSize(v string) *TaobaoItemUpdateRequest {
    s.ItemSize = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetInputCustomCpv(v string) *TaobaoItemUpdateRequest {
    s.InputCustomCpv = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetCpvMemo(v string) *TaobaoItemUpdateRequest {
    s.CpvMemo = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetNewprepay(v string) *TaobaoItemUpdateRequest {
    s.Newprepay = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetBarcode(v string) *TaobaoItemUpdateRequest {
    s.Barcode = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSellPoint(v string) *TaobaoItemUpdateRequest {
    s.SellPoint = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetQualification(v string) *TaobaoItemUpdateRequest {
    s.Qualification = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetO2oBindService(v bool) *TaobaoItemUpdateRequest {
    s.O2oBindService = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetIgnorewarning(v string) *TaobaoItemUpdateRequest {
    s.Ignorewarning = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetAfterSaleId(v int64) *TaobaoItemUpdateRequest {
    s.AfterSaleId = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetChangeProp(v string) *TaobaoItemUpdateRequest {
    s.ChangeProp = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetDescModules(v string) *TaobaoItemUpdateRequest {
    s.DescModules = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetIsOffline(v string) *TaobaoItemUpdateRequest {
    s.IsOffline = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetWirelessDesc(v string) *TaobaoItemUpdateRequest {
    s.WirelessDesc = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetChaoshiExtendsInfo(v string) *TaobaoItemUpdateRequest {
    s.ChaoshiExtendsInfo = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSpuConfirm(v bool) *TaobaoItemUpdateRequest {
    s.SpuConfirm = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetVideoId(v int64) *TaobaoItemUpdateRequest {
    s.VideoId = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetInteractiveId(v int64) *TaobaoItemUpdateRequest {
    s.InteractiveId = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLeaseExtendsInfo(v string) *TaobaoItemUpdateRequest {
    s.LeaseExtendsInfo = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetInputPids(v []string) *TaobaoItemUpdateRequest {
    s.InputPids = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetInputStr(v []string) *TaobaoItemUpdateRequest {
    s.InputStr = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSkuProperties(v string) *TaobaoItemUpdateRequest {
    s.SkuProperties = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSkuQuantities(v string) *TaobaoItemUpdateRequest {
    s.SkuQuantities = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSkuPrices(v string) *TaobaoItemUpdateRequest {
    s.SkuPrices = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSkuOuterIds(v string) *TaobaoItemUpdateRequest {
    s.SkuOuterIds = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSkuBarcode(v string) *TaobaoItemUpdateRequest {
    s.SkuBarcode = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSkuSpecIds(v string) *TaobaoItemUpdateRequest {
    s.SkuSpecIds = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSkuDeliveryTimes(v string) *TaobaoItemUpdateRequest {
    s.SkuDeliveryTimes = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSkuHdLength(v string) *TaobaoItemUpdateRequest {
    s.SkuHdLength = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSkuHdHeight(v string) *TaobaoItemUpdateRequest {
    s.SkuHdHeight = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetSkuHdLampQuantity(v string) *TaobaoItemUpdateRequest {
    s.SkuHdLampQuantity = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocationState(v string) *TaobaoItemUpdateRequest {
    s.LocationState = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocationCity(v string) *TaobaoItemUpdateRequest {
    s.LocationCity = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityPrdLicenseNo(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityPrdLicenseNo = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityDesignCode(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityDesignCode = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityFactory(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityFactory = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityFactorySite(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityFactorySite = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityContact(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityContact = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityMix(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityMix = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityPlanStorage(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityPlanStorage = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityPeriod(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityPeriod = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityFoodAdditive(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityFoodAdditive = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecuritySupplier(v string) *TaobaoItemUpdateRequest {
    s.FoodSecuritySupplier = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityProductDateStart(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityProductDateStart = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityProductDateEnd(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityProductDateEnd = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityStockDateStart(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityStockDateStart = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityStockDateEnd(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityStockDateEnd = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetFoodSecurityHealthProductNo(v string) *TaobaoItemUpdateRequest {
    s.FoodSecurityHealthProductNo = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifeExpirydate(v string) *TaobaoItemUpdateRequest {
    s.LocalityLifeExpirydate = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifeNetworkId(v string) *TaobaoItemUpdateRequest {
    s.LocalityLifeNetworkId = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifeMerchant(v string) *TaobaoItemUpdateRequest {
    s.LocalityLifeMerchant = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifeVerification(v string) *TaobaoItemUpdateRequest {
    s.LocalityLifeVerification = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifeRefundRatio(v int64) *TaobaoItemUpdateRequest {
    s.LocalityLifeRefundRatio = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifeOnsaleAutoRefundRatio(v int64) *TaobaoItemUpdateRequest {
    s.LocalityLifeOnsaleAutoRefundRatio = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifeChooseLogis(v string) *TaobaoItemUpdateRequest {
    s.LocalityLifeChooseLogis = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifeRefundmafee(v string) *TaobaoItemUpdateRequest {
    s.LocalityLifeRefundmafee = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifeObs(v string) *TaobaoItemUpdateRequest {
    s.LocalityLifeObs = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifeEticket(v string) *TaobaoItemUpdateRequest {
    s.LocalityLifeEticket = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifeVersion(v string) *TaobaoItemUpdateRequest {
    s.LocalityLifeVersion = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetLocalityLifePackageid(v string) *TaobaoItemUpdateRequest {
    s.LocalityLifePackageid = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetPaimaiInfoMode(v int64) *TaobaoItemUpdateRequest {
    s.PaimaiInfoMode = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetPaimaiInfoDeposit(v int64) *TaobaoItemUpdateRequest {
    s.PaimaiInfoDeposit = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetPaimaiInfoInterval(v int64) *TaobaoItemUpdateRequest {
    s.PaimaiInfoInterval = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetPaimaiInfoReserve(v string) *TaobaoItemUpdateRequest {
    s.PaimaiInfoReserve = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetPaimaiInfoValidHour(v int64) *TaobaoItemUpdateRequest {
    s.PaimaiInfoValidHour = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetPaimaiInfoValidMinute(v int64) *TaobaoItemUpdateRequest {
    s.PaimaiInfoValidMinute = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetMsPaymentPrice(v string) *TaobaoItemUpdateRequest {
    s.MsPaymentPrice = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetMsPaymentVoucherPrice(v string) *TaobaoItemUpdateRequest {
    s.MsPaymentVoucherPrice = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetMsPaymentReferencePrice(v string) *TaobaoItemUpdateRequest {
    s.MsPaymentReferencePrice = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetDeliveryTimeDeliveryTime(v string) *TaobaoItemUpdateRequest {
    s.DeliveryTimeDeliveryTime = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetDeliveryTimeDeliveryTimeType(v string) *TaobaoItemUpdateRequest {
    s.DeliveryTimeDeliveryTimeType = &v
    return s
}
func (s *TaobaoItemUpdateRequest) SetDeliveryTimeNeedDeliveryTime(v string) *TaobaoItemUpdateRequest {
    s.DeliveryTimeNeedDeliveryTime = &v
    return s
}

func (req *TaobaoItemUpdateRequest) ToMap() map[string]interface{} {
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
    if(req.Features != nil) {
        paramMap["features"] = *req.Features
    }
    if(req.ScenicTicketPayWay != nil) {
        paramMap["scenic_ticket_pay_way"] = *req.ScenicTicketPayWay
    }
    if(req.ScenicTicketBookCost != nil) {
        paramMap["scenic_ticket_book_cost"] = *req.ScenicTicketBookCost
    }
    if(req.EmptyFields != nil) {
        paramMap["empty_fields"] = *req.EmptyFields
    }
    if(req.GlobalStockType != nil) {
        paramMap["global_stock_type"] = *req.GlobalStockType
    }
    if(req.GlobalStockCountry != nil) {
        paramMap["global_stock_country"] = *req.GlobalStockCountry
    }
    if(req.GlobalStockDeliveryPlace != nil) {
        paramMap["global_stock_delivery_place"] = *req.GlobalStockDeliveryPlace
    }
    if(req.GlobalStockTaxFreePromise != nil) {
        paramMap["global_stock_tax_free_promise"] = *req.GlobalStockTaxFreePromise
    }
    if(req.ItemWeight != nil) {
        paramMap["item_weight"] = *req.ItemWeight
    }
    if(req.ItemSize != nil) {
        paramMap["item_size"] = *req.ItemSize
    }
    if(req.InputCustomCpv != nil) {
        paramMap["input_custom_cpv"] = *req.InputCustomCpv
    }
    if(req.CpvMemo != nil) {
        paramMap["cpv_memo"] = *req.CpvMemo
    }
    if(req.Newprepay != nil) {
        paramMap["newprepay"] = *req.Newprepay
    }
    if(req.Barcode != nil) {
        paramMap["barcode"] = *req.Barcode
    }
    if(req.SellPoint != nil) {
        paramMap["sell_point"] = *req.SellPoint
    }
    if(req.Qualification != nil) {
        paramMap["qualification"] = *req.Qualification
    }
    if(req.O2oBindService != nil) {
        paramMap["o2o_bind_service"] = *req.O2oBindService
    }
    if(req.Ignorewarning != nil) {
        paramMap["ignorewarning"] = *req.Ignorewarning
    }
    if(req.AfterSaleId != nil) {
        paramMap["after_sale_id"] = *req.AfterSaleId
    }
    if(req.ChangeProp != nil) {
        paramMap["change_prop"] = *req.ChangeProp
    }
    if(req.DescModules != nil) {
        paramMap["desc_modules"] = *req.DescModules
    }
    if(req.IsOffline != nil) {
        paramMap["is_offline"] = *req.IsOffline
    }
    if(req.WirelessDesc != nil) {
        paramMap["wireless_desc"] = *req.WirelessDesc
    }
    if(req.ChaoshiExtendsInfo != nil) {
        paramMap["chaoshi_extends_info"] = *req.ChaoshiExtendsInfo
    }
    if(req.SpuConfirm != nil) {
        paramMap["spu_confirm"] = *req.SpuConfirm
    }
    if(req.VideoId != nil) {
        paramMap["video_id"] = *req.VideoId
    }
    if(req.InteractiveId != nil) {
        paramMap["interactive_id"] = *req.InteractiveId
    }
    if(req.LeaseExtendsInfo != nil) {
        paramMap["lease_extends_info"] = *req.LeaseExtendsInfo
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
    if(req.SkuBarcode != nil) {
        paramMap["sku_barcode"] = *req.SkuBarcode
    }
    if(req.SkuSpecIds != nil) {
        paramMap["sku_spec_ids"] = *req.SkuSpecIds
    }
    if(req.SkuDeliveryTimes != nil) {
        paramMap["sku_delivery_times"] = *req.SkuDeliveryTimes
    }
    if(req.SkuHdLength != nil) {
        paramMap["sku_hd_length"] = *req.SkuHdLength
    }
    if(req.SkuHdHeight != nil) {
        paramMap["sku_hd_height"] = *req.SkuHdHeight
    }
    if(req.SkuHdLampQuantity != nil) {
        paramMap["sku_hd_lamp_quantity"] = *req.SkuHdLampQuantity
    }
    if(req.LocationState != nil) {
        paramMap["location.state"] = *req.LocationState
    }
    if(req.LocationCity != nil) {
        paramMap["location.city"] = *req.LocationCity
    }
    if(req.FoodSecurityPrdLicenseNo != nil) {
        paramMap["food_security.prd_license_no"] = *req.FoodSecurityPrdLicenseNo
    }
    if(req.FoodSecurityDesignCode != nil) {
        paramMap["food_security.design_code"] = *req.FoodSecurityDesignCode
    }
    if(req.FoodSecurityFactory != nil) {
        paramMap["food_security.factory"] = *req.FoodSecurityFactory
    }
    if(req.FoodSecurityFactorySite != nil) {
        paramMap["food_security.factory_site"] = *req.FoodSecurityFactorySite
    }
    if(req.FoodSecurityContact != nil) {
        paramMap["food_security.contact"] = *req.FoodSecurityContact
    }
    if(req.FoodSecurityMix != nil) {
        paramMap["food_security.mix"] = *req.FoodSecurityMix
    }
    if(req.FoodSecurityPlanStorage != nil) {
        paramMap["food_security.plan_storage"] = *req.FoodSecurityPlanStorage
    }
    if(req.FoodSecurityPeriod != nil) {
        paramMap["food_security.period"] = *req.FoodSecurityPeriod
    }
    if(req.FoodSecurityFoodAdditive != nil) {
        paramMap["food_security.food_additive"] = *req.FoodSecurityFoodAdditive
    }
    if(req.FoodSecuritySupplier != nil) {
        paramMap["food_security.supplier"] = *req.FoodSecuritySupplier
    }
    if(req.FoodSecurityProductDateStart != nil) {
        paramMap["food_security.product_date_start"] = *req.FoodSecurityProductDateStart
    }
    if(req.FoodSecurityProductDateEnd != nil) {
        paramMap["food_security.product_date_end"] = *req.FoodSecurityProductDateEnd
    }
    if(req.FoodSecurityStockDateStart != nil) {
        paramMap["food_security.stock_date_start"] = *req.FoodSecurityStockDateStart
    }
    if(req.FoodSecurityStockDateEnd != nil) {
        paramMap["food_security.stock_date_end"] = *req.FoodSecurityStockDateEnd
    }
    if(req.FoodSecurityHealthProductNo != nil) {
        paramMap["food_security.health_product_no"] = *req.FoodSecurityHealthProductNo
    }
    if(req.LocalityLifeExpirydate != nil) {
        paramMap["locality_life.expirydate"] = *req.LocalityLifeExpirydate
    }
    if(req.LocalityLifeNetworkId != nil) {
        paramMap["locality_life.network_id"] = *req.LocalityLifeNetworkId
    }
    if(req.LocalityLifeMerchant != nil) {
        paramMap["locality_life.merchant"] = *req.LocalityLifeMerchant
    }
    if(req.LocalityLifeVerification != nil) {
        paramMap["locality_life.verification"] = *req.LocalityLifeVerification
    }
    if(req.LocalityLifeRefundRatio != nil) {
        paramMap["locality_life.refund_ratio"] = *req.LocalityLifeRefundRatio
    }
    if(req.LocalityLifeOnsaleAutoRefundRatio != nil) {
        paramMap["locality_life.onsale_auto_refund_ratio"] = *req.LocalityLifeOnsaleAutoRefundRatio
    }
    if(req.LocalityLifeChooseLogis != nil) {
        paramMap["locality_life.choose_logis"] = *req.LocalityLifeChooseLogis
    }
    if(req.LocalityLifeRefundmafee != nil) {
        paramMap["locality_life.refundmafee"] = *req.LocalityLifeRefundmafee
    }
    if(req.LocalityLifeObs != nil) {
        paramMap["locality_life.obs"] = *req.LocalityLifeObs
    }
    if(req.LocalityLifeEticket != nil) {
        paramMap["locality_life.eticket"] = *req.LocalityLifeEticket
    }
    if(req.LocalityLifeVersion != nil) {
        paramMap["locality_life.version"] = *req.LocalityLifeVersion
    }
    if(req.LocalityLifePackageid != nil) {
        paramMap["locality_life.packageid"] = *req.LocalityLifePackageid
    }
    if(req.PaimaiInfoMode != nil) {
        paramMap["paimai_info.mode"] = *req.PaimaiInfoMode
    }
    if(req.PaimaiInfoDeposit != nil) {
        paramMap["paimai_info.deposit"] = *req.PaimaiInfoDeposit
    }
    if(req.PaimaiInfoInterval != nil) {
        paramMap["paimai_info.interval"] = *req.PaimaiInfoInterval
    }
    if(req.PaimaiInfoReserve != nil) {
        paramMap["paimai_info.reserve"] = *req.PaimaiInfoReserve
    }
    if(req.PaimaiInfoValidHour != nil) {
        paramMap["paimai_info.valid_hour"] = *req.PaimaiInfoValidHour
    }
    if(req.PaimaiInfoValidMinute != nil) {
        paramMap["paimai_info.valid_minute"] = *req.PaimaiInfoValidMinute
    }
    if(req.MsPaymentPrice != nil) {
        paramMap["ms_payment.price"] = *req.MsPaymentPrice
    }
    if(req.MsPaymentVoucherPrice != nil) {
        paramMap["ms_payment.voucher_price"] = *req.MsPaymentVoucherPrice
    }
    if(req.MsPaymentReferencePrice != nil) {
        paramMap["ms_payment.reference_price"] = *req.MsPaymentReferencePrice
    }
    if(req.DeliveryTimeDeliveryTime != nil) {
        paramMap["delivery_time.delivery_time"] = *req.DeliveryTimeDeliveryTime
    }
    if(req.DeliveryTimeDeliveryTimeType != nil) {
        paramMap["delivery_time.delivery_time_type"] = *req.DeliveryTimeDeliveryTimeType
    }
    if(req.DeliveryTimeNeedDeliveryTime != nil) {
        paramMap["delivery_time.need_delivery_time"] = *req.DeliveryTimeNeedDeliveryTime
    }
    return paramMap
}

func (req *TaobaoItemUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}