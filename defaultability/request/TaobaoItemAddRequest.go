package request

import (
        "topsdk/util"
    )

type TaobaoItemAddRequest struct {
    /*
        商品上传后的状态。可选值:onsale(出售中),instock(仓库中);默认值:onsale defalutValue��onsale    */
    ApproveStatus  *string `json:"approve_status,omitempty" required:"false" `
    /*
        发布类型。可选值:fixed(一口价),auction(拍卖)。B商家不能发布拍卖商品，而且拍卖商品是没有SKU的。拍卖商品发布时需要附加拍卖商品信息：拍卖类型(paimai_info.mode，拍卖类型包括三种：增价拍[1]，荷兰拍[2]以及降价拍[3])，商品数量(num)，起拍价(price)，价格幅度(increament)，保证金(paimai_info.deposit)。另外拍卖商品支持自定义销售周期，通过paimai_info.valid_hour和paimai_info.valid_minute来指定。对于降价拍来说需要设置降价周期(paimai_info.interval)和拍卖保留价(paimai_info.reserve)。注意：通过taobao.item.get接口获取拍卖信息时，会返回除了valid_hour和valid_minute之外的所有拍卖信息。     */
    Type  *string `json:"type" required:"true" `
    /*
        自动重发。可选值:true,false;默认值:false(不重发) defalutValue��false    */
    AutoRepost  *bool `json:"auto_repost,omitempty" required:"false" `
    /*
        叶子类目id     */
    Cid  *int64 `json:"cid" required:"true" `
    /*
        宝贝描述。字数要大于5个字符，小于25000个字符，受违禁词控制     */
    Desc  *string `json:"desc" required:"true" `
    /*
        有效期。可选值:7,14;单位:天;默认值:14 defalutValue��14    */
    ValidThru  *int64 `json:"valid_thru,omitempty" required:"false" `
    /*
        平邮费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:5.07，表示:5元7分. 注:post_fee,express_fee,ems_fee需要一起填写     */
    PostFee  *string `json:"post_fee,omitempty" required:"false" `
    /*
        快递费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:15.07，表示:15元7分     */
    ExpressFee  *string `json:"express_fee,omitempty" required:"false" `
    /*
        ems费用。取值范围:0.01-999.00;精确到2位小数;单位:元。如:25.07，表示:25元7分     */
    EmsFee  *string `json:"ems_fee,omitempty" required:"false" `
    /*
        是否有保修。可选值:true,false;默认值:false(不保修) defalutValue��false    */
    HasWarranty  *bool `json:"has_warranty,omitempty" required:"false" `
    /*
        是否有发票。可选值:true,false (商城卖家此字段必须为true);默认值:false(无发票) defalutValue��false    */
    HasInvoice  *bool `json:"has_invoice,omitempty" required:"false" `
    /*
        加价(降价)幅度。如果为0，代表系统代理幅度。对于增价拍和荷兰拍来说是加价幅度，对于降价拍来说是降价幅度。     */
    Increment  *string `json:"increment,omitempty" required:"false" `
    /*
        商品属性列表。格式:pid:vid;pid:vid。属性的pid调用taobao.itemprops.get取得，属性值的vid用taobao.itempropvalues.get取得vid。 如果该类目下面没有属性，可以不用填写。如果有属性，必选属性必填，其他非必选属性可以选择不填写.属性不能超过35对。所有属性加起来包括分割符不能超过549字节，单个属性没有限制。 如果有属性是可输入的话，则用字段input_str填入属性的值     */
    Props  *string `json:"props,omitempty" required:"false" `
    /*
        商品数量。取值范围:0-900000000的整数。且需要等于Sku所有数量的和。拍卖商品中增加拍只能为1，荷兰拍要在[2,500)范围内。     */
    Num  *int64 `json:"num" required:"true" `
    /*
        运费承担方式。可选值:seller（卖家承担）,buyer(买家承担);默认值:seller。卖家承担不用设置邮费和postage_id.买家承担的时候，必填邮费和postage_id 如果用户设置了运费模板会优先使用运费模板，否则要同步设置邮费（post_fee,express_fee,ems_fee） defalutValue��seller    */
    FreightPayer  *string `json:"freight_payer,omitempty" required:"false" `
    /*
        商品所属的店铺类目列表。按逗号分隔。结构:",cid1,cid2,...,"，如果店铺类目存在二级类目，必须传入子类目cids。     */
    SellerCids  *string `json:"seller_cids,omitempty" required:"false" `
    /*
        橱窗推荐。可选值:true,false;默认值:false(不推荐) defalutValue��false    */
    HasShowcase  *bool `json:"has_showcase,omitempty" required:"false" `
    /*
        定时上架时间。(时间格式：yyyy-MM-dd HH:mm:ss)     */
    ListTime  *util.LocalTime `json:"list_time,omitempty" required:"false" `
    /*
        新旧程度。可选值：new(新)，second(二手)。B商家不能发布二手商品。如果是二手商品，特定类目下属性里面必填新旧成色属性     */
    StuffStatus  *string `json:"stuff_status" required:"true" `
    /*
        宝贝标题。不能超过30字符，受违禁词控制。天猫图书管控类目最大允许120字符；     */
    Title  *string `json:"title" required:"true" `
    /*
        商品价格。取值范围:0-100000000;精确到2位小数;单位:元。如:200.07，表示:200元7分。需要在正确的价格区间内。拍卖商品对应的起拍价。     */
    Price  *string `json:"price,omitempty" required:"false" `
    /*
        支持会员打折。可选值:true,false;默认值:false(不打折) defalutValue��false    */
    HasDiscount  *bool `json:"has_discount,omitempty" required:"false" `
    /*
        商品外部编码，该字段的最大长度是64个字节     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        宝贝所属的运费模板ID。取值范围：整数且必须是该卖家的运费模板的ID（可通过taobao.delivery.template.get获得当前会话用户的所有邮费模板）     */
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
        商品的积分返点比例。如:5,表示:返点比例0.5%. 注意：返点比例必须是>0的整数，而且最大是90,即为9%.B商家在发布非虚拟商品时，返点必须是 5的倍数，即0.5%的倍数。其它是1的倍数，即0.1%的倍数。无名良品商家发布商品时，复用该字段记录积分宝返点比例，返点必须是对应类目的返点步长的整数倍，默认是5，即0.5%。注意此时该字段值依旧必须是>0的整数，最高值不超过500，即50%     */
    AuctionPoint  *int64 `json:"auction_point,omitempty" required:"false" `
    /*
        属性值别名。如pid:vid:别名;pid1:vid1:别名1 ，其中：pid是属性id vid是属性值id。总长度不超过800个字符，如"123:333:你好"，引号内的是10个字符。     */
    PropertyAlias  *string `json:"property_alias,omitempty" required:"false" `
    /*
        商品文字的字符集。繁体传入"zh_HK"，简体传入"zh_CN"，不传默认为简体 defalutValue��zh_CN    */
    Lang  *string `json:"lang,omitempty" required:"false" `
    /*
        （推荐）商品主图需要关联的图片空间的相对url。这个url所对应的图片必须要属于当前用户。pic_path和image只需要传入一个,如果两个都传，默认选择pic_path     */
    PicPath  *string `json:"pic_path,omitempty" required:"false" `
    /*
        商品主图片。类型:JPG,GIF;最大长度:3M。（推荐使用pic_path字段，先把图片上传到卖家图片空间）     */
    Image  *[]byte `json:"image,omitempty" required:"false" `
    /*
        代充商品类型。在代充商品的类目下，不传表示不标记商品类型（交易搜索中就不能通过标记搜到相关的交易了）。可选类型： no_mark(不做类型标记) time_card(点卡软件代充) fee_card(话费软件代充)     */
    AutoFill  *string `json:"auto_fill,omitempty" required:"false" `
    /*
        是否承诺退换货服务!虚拟商品无须设置此项!     */
    SellPromise  *bool `json:"sell_promise,omitempty" required:"false" `
    /*
        此为货到付款运费模板的ID，对应的JAVA类型是long,如果COD卖家应用了货到付款运费模板，此值要进行设置。该字段已经废弃     */
    CodPostageId  *int64 `json:"cod_postage_id,omitempty" required:"false" `
    /*
        实物闪电发货     */
    IsLightningConsignment  *bool `json:"is_lightning_consignment,omitempty" required:"false" `
    /*
        商品的重量(商超卖家专用字段)     */
    Weight  *int64 `json:"weight,omitempty" required:"false" `
    /*
        宝贝形态:1代表电子券;0或其他代表实物     */
    Shape  *string `json:"shape,omitempty" required:"false" `
    /*
        商品是否为新品。只有在当前类目开通新品,并且当前用户拥有该类目下发布新品权限时才能设置is_xinpin为true，否则设置true后会返回错误码:isv.invalid-permission:add-xinpin。同时只有一口价全新的宝贝才能设置为新品，否则会返回错误码：isv.invalid-parameter:xinpin。不设置该参数值或设置为false效果一致。 defalutValue��false    */
    IsXinpin  *bool `json:"is_xinpin,omitempty" required:"false" `
    /*
        商品是否支持拍下减库存:1支持;2取消支持(付款减库存);0(默认)不更改集市卖家默认拍下减库存;商城卖家默认付款减库存 defalutValue��0    */
    SubStock  *int64 `json:"sub_stock,omitempty" required:"false" `
    /*
        宝贝特征值，格式为：【key1:value1;key2:value2;key3:value3;】，key和value用【:】分隔，key&value之间用【;】分隔，只有在Top支持的特征值才能保存到宝贝上，目前支持的Key列表为：mysize_tp     */
    Features  *string `json:"features,omitempty" required:"false" `
    /*
        景区门票类宝贝发布时候，当卖家签订了支付宝代扣协议时候，需要选择支付方式：全额支付和订金支付。当scenic_ticket_pay_way为1时表示全额支付，为2时表示订金支付     */
    ScenicTicketPayWay  *int64 `json:"scenic_ticket_pay_way,omitempty" required:"false" `
    /*
        景区门票在选择订金支付时候，需要交的预订费。传入的值是1到20之间的数值，小数点后最多可以保留两位（多余的部分将做四舍五入的处理）。这个数值表示的是预订费的比例，最终的预订费为 scenic_ticket_book_cost乘一口价除以100     */
    ScenicTicketBookCost  *string `json:"scenic_ticket_book_cost,omitempty" required:"false" `
    /*
        全球购商品采购地（库存类型），有两种库存类型：现货和代购参数值为1时代表现货，值为2时代表代购。注意：使用时请与 全球购商品采购地（地区/国家）配合使用     */
    GlobalStockType  *string `json:"global_stock_type,omitempty" required:"false" `
    /*
        全球购商品采购地（地区/国家）,默认值只在全球购商品采购地（库存类型选择情况生效），地区国家值请填写法定的国家名称，类如（美国, 香港, 日本, 英国, 新西兰, 德国, 韩国, 荷兰, 法国, 意大利, 台湾, 澳门, 加拿大, 瑞士, 西班牙, 泰国, 新加坡, 马来西亚, 菲律宾），不要使用其他     */
    GlobalStockCountry  *string `json:"global_stock_country,omitempty" required:"false" `
    /*
        全球购商品发货地，发货地现在有两种类型：“国内”和“海外及港澳台”，参数值为1时代表“国内”，值为2时代表“海外及港澳台”，默认为国内。注意：卖家必须已经签署并启用“海外直邮”合约，才能选择发货地为“海外及港澳台”     */
    GlobalStockDeliveryPlace  *string `json:"global_stock_delivery_place,omitempty" required:"false" `
    /*
        全球购商品卖家包税承诺，当值为true时，代表卖家承诺包税。注意：请与“全球购商品发货地”配合使用，包税承诺必须满足：1、发货地为海外及港澳台 2、卖家已经签署并启用“包税合约”合约     */
    GlobalStockTaxFreePromise  *bool `json:"global_stock_tax_free_promise,omitempty" required:"false" `
    /*
        商品的重量，用于按重量计费的运费模板。注意：单位为kg。只能传入数值类型（包含小数），不能带单位，单位默认为kg。     */
    ItemWeight  *string `json:"item_weight,omitempty" required:"false" `
    /*
        表示商品的体积，如果需要使用按体积计费的运费模板，一定要设置这个值。该值的单位为立方米（m3），如果是其他单位，请转换成成立方米。该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）。体积和长宽高都支持小数类型。在传入体积或长宽高时候，不能带单位。体积的单位默认为立方米（m3），长宽高的单位默认为米(m)该值支持两种格式的设置：格式1：bulk:3,单位为立方米(m3),表示直接设置为商品的体积。格式2：length:10;breadth:10;height:10，单位为米（m）     */
    ItemSize  *string `json:"item_size,omitempty" required:"false" `
    /*
        针对当前商品的自定义属性值，目前是针对销售属性值自定义的，所以调用方需要把自定义属性值对应的虚拟属性值ID（负整数，例如例子中的 -1和-2）像标准属性值值的id一样设置到SKU上，如果自定义属性值有属性值图片，也要设置到属性图片上     */
    InputCustomCpv  *string `json:"input_custom_cpv,omitempty" required:"false" `
    /*
        针对当前商品的标准属性值的补充说明，让买家更加了解商品信息减少交易纠纷     */
    CpvMemo  *string `json:"cpv_memo,omitempty" required:"false" `
    /*
        是否支持定制市场 true代表支持，false代表支持,如果为空代表与之前保持不变不会修改     */
    SupportCustomMade  *bool `json:"support_custom_made,omitempty" required:"false" `
    /*
        定制工具Id如果支持定制市场，这个值不填写，就用之前的定制工具Id，之前的定制工具Id没有值就默认为-1     */
    CustomMadeTypeId  *string `json:"custom_made_type_id,omitempty" required:"false" `
    /*
        该宝贝是否支持【7天无理由退货】，卖家选择的值只是一个因素，最终以类目和选择的属性条件来确定是否支持7天。填入字符0，表示不支持；未填写或填人字符1，表示支持7天无理由退货；     */
    Newprepay  *string `json:"newprepay,omitempty" required:"false" `
    /*
        商品条形码     */
    Barcode  *string `json:"barcode,omitempty" required:"false" `
    /*
        商品卖点信息，最长150个字符。天猫商家和集市卖家都可用。     */
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
        手机类目spu 优化，信息确认字段     */
    SpuConfirm  *bool `json:"spu_confirm,omitempty" required:"false" `
    /*
        主图视频id     */
    VideoId  *int64 `json:"video_id,omitempty" required:"false" `
    /*
        主图视频互动信息id，必须填写主图视频id才能有互动信息id     */
    InteractiveId  *int64 `json:"interactive_id,omitempty" required:"false" `
    /*
        租赁扩展信息     */
    LeaseExtendsInfo  *string `json:"lease_extends_info,omitempty" required:"false" `
    /*
        仅淘小铺卖家需要。佣金比例(15.3对应的佣金比例为15.3%).只支持小数点后1位。多余的位数四舍五入(15.32会保存为15.3%     */
    Brokerage  *string `json:"brokerage,omitempty" required:"false" `
    /*
        业务身份编码。淘小铺编码为"taobao-taoxiaopu"     */
    BizCode  *string `json:"biz_code,omitempty" required:"false" `
    /*
        此字段已经废弃，不再使用     */
    ImageUrls  *[]string `json:"image_urls,omitempty" required:"false" `
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
        Sku的数量串，结构如：num1,num2,num3 如：2,3     */
    SkuQuantities  *string `json:"sku_quantities,omitempty" required:"false" `
    /*
        Sku的价格串，结构如：10.00,5.00,… 精确到2位小数;单位:元。如:200.07，表示:200元7分     */
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
        家装建材类目，商品SKU的灯头数量，正整数，大于等于3，部分类目必选。天猫商家专用。 数据和SKU一一对应，用,分隔，如：3,5,7     */
    SkuHdLampQuantity  *string `json:"sku_hd_lamp_quantity,omitempty" required:"false" `
    /*
        所在地省份。如浙江     */
    LocationState  *string `json:"location.state" required:"true" `
    /*
        所在地城市。如杭州 。     */
    LocationCity  *string `json:"location.city" required:"true" `
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
        网点ID     */
    LocalityLifeNetworkId  *string `json:"locality_life.network_id,omitempty" required:"false" `
    /*
        码商信息，格式为 码商id:nick     */
    LocalityLifeMerchant  *string `json:"locality_life.merchant,omitempty" required:"false" `
    /*
        核销打款 1代表核销打款 0代表非核销打款     */
    LocalityLifeVerification  *string `json:"locality_life.verification,omitempty" required:"false" `
    /*
        退款比例，百分比%前的数字,1-100的正整数值     */
    LocalityLifeRefundRatio  *int64 `json:"locality_life.refund_ratio,omitempty" required:"false" `
    /*
        电子凭证售中自动退款比例，百分比%前的数字，介于1-100之间的整数     */
    LocalityLifeOnsaleAutoRefundRatio  *int64 `json:"locality_life.onsale_auto_refund_ratio,omitempty" required:"false" `
    /*
        发布电子凭证宝贝时候表示是否使用邮寄 0: 代表不使用邮寄； 1：代表使用邮寄；如果不设置这个值，代表不使用邮寄     */
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
        新版电子凭证字段     */
    LocalityLifeVersion  *string `json:"locality_life.version,omitempty" required:"false" `
    /*
        新版电子凭证包 id     */
    LocalityLifePackageid  *string `json:"locality_life.packageid,omitempty" required:"false" `
    /*
        拍卖商品选择的拍卖类型，拍卖类型包括三种：增价拍(1)，荷兰拍(2)和降价拍(3)。     */
    PaimaiInfoMode  *int64 `json:"paimai_info.mode,omitempty" required:"false" `
    /*
        拍卖宝贝的保证金。对于增价拍和荷兰拍来说保证金有两种模式：淘宝默认模式（首次出价金额的10%），自定义固定保证金（固定冻结金额只能输入不超过30万的正整数），并且保证金只冻结1次。对于降价拍来说保证金只有淘宝默认的（竞拍价格的10% * 竞拍数量），并且每次出价都需要冻结保证金。对于拍卖宝贝来说，保证金是必须的，但是默认使用淘宝默认保证金模式，只有用户需要使用自定义固定保证金的时候才需要使用到这个参数，如果该参数不传或传入0则代表使用默认。     */
    PaimaiInfoDeposit  *int64 `json:"paimai_info.deposit,omitempty" required:"false" `
    /*
        降价拍宝贝的降价周期(分钟)。降价拍宝贝的价格每隔paimai_info.interval时间会下降一次increment。     */
    PaimaiInfoInterval  *int64 `json:"paimai_info.interval,omitempty" required:"false" `
    /*
        降价拍宝贝的保留价。对于降价拍来说，paimai_info.reserve必须大于0，且小于price-increment，而且（price-paimai_info.reserve）/increment的计算结果必须为整数     */
    PaimaiInfoReserve  *string `json:"paimai_info.reserve,omitempty" required:"false" `
    /*
        自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。     */
    PaimaiInfoValidHour  *int64 `json:"paimai_info.valid_hour,omitempty" required:"false" `
    /*
        自定义销售周期的分钟数。拍卖宝贝可以自定义销售周期，这里是指定销售周期的分钟数。自定义销售周期的小时数。拍卖宝贝可以自定义销售周期，这里指定销售周期的小时数。注意，该参数只作为输入参数，不能通过taobao.item.get接口获取。     */
    PaimaiInfoValidMinute  *int64 `json:"paimai_info.valid_minute,omitempty" required:"false" `
    /*
        参考价。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm     */
    MsPaymentReferencePrice  *string `json:"ms_payment.reference_price,omitempty" required:"false" `
    /*
        尾款可抵扣金额。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm     */
    MsPaymentVoucherPrice  *string `json:"ms_payment.voucher_price,omitempty" required:"false" `
    /*
        订金。在“线上付订金线下付尾款”模式中，有订金、尾款可抵扣金额和参考价，三者需要同时填写。该商品订单首次支付价格为 订金 价格，用户可根据 参考价 估算全款。该模式有别于“一口价”付款方式，针对一个商品，只能选择两种付款方式中的一种，其适用于家装、二手车等场景。详见说明：http://bangpai.taobao.com/group/thread/15031186-303287205.htm     */
    MsPaymentPrice  *string `json:"ms_payment.price,omitempty" required:"false" `
    /*
        设置是否使用发货时间，商品级别，sku级别     */
    DeliveryTimeNeedDeliveryTime  *string `json:"delivery_time.need_delivery_time,omitempty" required:"false" `
    /*
        发货时间类型：绝对发货时间或者相对发货时间     */
    DeliveryTimeDeliveryTimeType  *string `json:"delivery_time.delivery_time_type,omitempty" required:"false" `
    /*
        商品级别设置的发货时间。设置了商品级别的发货时间，相对发货时间，则填写相对发货时间的天数（大于3）；绝对发货时间，则填写yyyy-mm-dd格式，如2013-11-11     */
    DeliveryTimeDeliveryTime  *string `json:"delivery_time.delivery_time,omitempty" required:"false" `
}

func (s *TaobaoItemAddRequest) SetApproveStatus(v string) *TaobaoItemAddRequest {
    s.ApproveStatus = &v
    return s
}
func (s *TaobaoItemAddRequest) SetType(v string) *TaobaoItemAddRequest {
    s.Type = &v
    return s
}
func (s *TaobaoItemAddRequest) SetAutoRepost(v bool) *TaobaoItemAddRequest {
    s.AutoRepost = &v
    return s
}
func (s *TaobaoItemAddRequest) SetCid(v int64) *TaobaoItemAddRequest {
    s.Cid = &v
    return s
}
func (s *TaobaoItemAddRequest) SetDesc(v string) *TaobaoItemAddRequest {
    s.Desc = &v
    return s
}
func (s *TaobaoItemAddRequest) SetValidThru(v int64) *TaobaoItemAddRequest {
    s.ValidThru = &v
    return s
}
func (s *TaobaoItemAddRequest) SetPostFee(v string) *TaobaoItemAddRequest {
    s.PostFee = &v
    return s
}
func (s *TaobaoItemAddRequest) SetExpressFee(v string) *TaobaoItemAddRequest {
    s.ExpressFee = &v
    return s
}
func (s *TaobaoItemAddRequest) SetEmsFee(v string) *TaobaoItemAddRequest {
    s.EmsFee = &v
    return s
}
func (s *TaobaoItemAddRequest) SetHasWarranty(v bool) *TaobaoItemAddRequest {
    s.HasWarranty = &v
    return s
}
func (s *TaobaoItemAddRequest) SetHasInvoice(v bool) *TaobaoItemAddRequest {
    s.HasInvoice = &v
    return s
}
func (s *TaobaoItemAddRequest) SetIncrement(v string) *TaobaoItemAddRequest {
    s.Increment = &v
    return s
}
func (s *TaobaoItemAddRequest) SetProps(v string) *TaobaoItemAddRequest {
    s.Props = &v
    return s
}
func (s *TaobaoItemAddRequest) SetNum(v int64) *TaobaoItemAddRequest {
    s.Num = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFreightPayer(v string) *TaobaoItemAddRequest {
    s.FreightPayer = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSellerCids(v string) *TaobaoItemAddRequest {
    s.SellerCids = &v
    return s
}
func (s *TaobaoItemAddRequest) SetHasShowcase(v bool) *TaobaoItemAddRequest {
    s.HasShowcase = &v
    return s
}
func (s *TaobaoItemAddRequest) SetListTime(v util.LocalTime) *TaobaoItemAddRequest {
    s.ListTime = &v
    return s
}
func (s *TaobaoItemAddRequest) SetStuffStatus(v string) *TaobaoItemAddRequest {
    s.StuffStatus = &v
    return s
}
func (s *TaobaoItemAddRequest) SetTitle(v string) *TaobaoItemAddRequest {
    s.Title = &v
    return s
}
func (s *TaobaoItemAddRequest) SetPrice(v string) *TaobaoItemAddRequest {
    s.Price = &v
    return s
}
func (s *TaobaoItemAddRequest) SetHasDiscount(v bool) *TaobaoItemAddRequest {
    s.HasDiscount = &v
    return s
}
func (s *TaobaoItemAddRequest) SetOuterId(v string) *TaobaoItemAddRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoItemAddRequest) SetPostageId(v int64) *TaobaoItemAddRequest {
    s.PostageId = &v
    return s
}
func (s *TaobaoItemAddRequest) SetProductId(v int64) *TaobaoItemAddRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoItemAddRequest) SetIsTaobao(v bool) *TaobaoItemAddRequest {
    s.IsTaobao = &v
    return s
}
func (s *TaobaoItemAddRequest) SetIsEx(v bool) *TaobaoItemAddRequest {
    s.IsEx = &v
    return s
}
func (s *TaobaoItemAddRequest) SetIs3D(v bool) *TaobaoItemAddRequest {
    s.Is3D = &v
    return s
}
func (s *TaobaoItemAddRequest) SetAuctionPoint(v int64) *TaobaoItemAddRequest {
    s.AuctionPoint = &v
    return s
}
func (s *TaobaoItemAddRequest) SetPropertyAlias(v string) *TaobaoItemAddRequest {
    s.PropertyAlias = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLang(v string) *TaobaoItemAddRequest {
    s.Lang = &v
    return s
}
func (s *TaobaoItemAddRequest) SetPicPath(v string) *TaobaoItemAddRequest {
    s.PicPath = &v
    return s
}
func (s *TaobaoItemAddRequest) SetImage(v []byte) *TaobaoItemAddRequest {
    s.Image = &v
    return s
}
func (s *TaobaoItemAddRequest) SetAutoFill(v string) *TaobaoItemAddRequest {
    s.AutoFill = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSellPromise(v bool) *TaobaoItemAddRequest {
    s.SellPromise = &v
    return s
}
func (s *TaobaoItemAddRequest) SetCodPostageId(v int64) *TaobaoItemAddRequest {
    s.CodPostageId = &v
    return s
}
func (s *TaobaoItemAddRequest) SetIsLightningConsignment(v bool) *TaobaoItemAddRequest {
    s.IsLightningConsignment = &v
    return s
}
func (s *TaobaoItemAddRequest) SetWeight(v int64) *TaobaoItemAddRequest {
    s.Weight = &v
    return s
}
func (s *TaobaoItemAddRequest) SetShape(v string) *TaobaoItemAddRequest {
    s.Shape = &v
    return s
}
func (s *TaobaoItemAddRequest) SetIsXinpin(v bool) *TaobaoItemAddRequest {
    s.IsXinpin = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSubStock(v int64) *TaobaoItemAddRequest {
    s.SubStock = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFeatures(v string) *TaobaoItemAddRequest {
    s.Features = &v
    return s
}
func (s *TaobaoItemAddRequest) SetScenicTicketPayWay(v int64) *TaobaoItemAddRequest {
    s.ScenicTicketPayWay = &v
    return s
}
func (s *TaobaoItemAddRequest) SetScenicTicketBookCost(v string) *TaobaoItemAddRequest {
    s.ScenicTicketBookCost = &v
    return s
}
func (s *TaobaoItemAddRequest) SetGlobalStockType(v string) *TaobaoItemAddRequest {
    s.GlobalStockType = &v
    return s
}
func (s *TaobaoItemAddRequest) SetGlobalStockCountry(v string) *TaobaoItemAddRequest {
    s.GlobalStockCountry = &v
    return s
}
func (s *TaobaoItemAddRequest) SetGlobalStockDeliveryPlace(v string) *TaobaoItemAddRequest {
    s.GlobalStockDeliveryPlace = &v
    return s
}
func (s *TaobaoItemAddRequest) SetGlobalStockTaxFreePromise(v bool) *TaobaoItemAddRequest {
    s.GlobalStockTaxFreePromise = &v
    return s
}
func (s *TaobaoItemAddRequest) SetItemWeight(v string) *TaobaoItemAddRequest {
    s.ItemWeight = &v
    return s
}
func (s *TaobaoItemAddRequest) SetItemSize(v string) *TaobaoItemAddRequest {
    s.ItemSize = &v
    return s
}
func (s *TaobaoItemAddRequest) SetInputCustomCpv(v string) *TaobaoItemAddRequest {
    s.InputCustomCpv = &v
    return s
}
func (s *TaobaoItemAddRequest) SetCpvMemo(v string) *TaobaoItemAddRequest {
    s.CpvMemo = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSupportCustomMade(v bool) *TaobaoItemAddRequest {
    s.SupportCustomMade = &v
    return s
}
func (s *TaobaoItemAddRequest) SetCustomMadeTypeId(v string) *TaobaoItemAddRequest {
    s.CustomMadeTypeId = &v
    return s
}
func (s *TaobaoItemAddRequest) SetNewprepay(v string) *TaobaoItemAddRequest {
    s.Newprepay = &v
    return s
}
func (s *TaobaoItemAddRequest) SetBarcode(v string) *TaobaoItemAddRequest {
    s.Barcode = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSellPoint(v string) *TaobaoItemAddRequest {
    s.SellPoint = &v
    return s
}
func (s *TaobaoItemAddRequest) SetQualification(v string) *TaobaoItemAddRequest {
    s.Qualification = &v
    return s
}
func (s *TaobaoItemAddRequest) SetO2oBindService(v bool) *TaobaoItemAddRequest {
    s.O2oBindService = &v
    return s
}
func (s *TaobaoItemAddRequest) SetIgnorewarning(v string) *TaobaoItemAddRequest {
    s.Ignorewarning = &v
    return s
}
func (s *TaobaoItemAddRequest) SetAfterSaleId(v int64) *TaobaoItemAddRequest {
    s.AfterSaleId = &v
    return s
}
func (s *TaobaoItemAddRequest) SetChangeProp(v string) *TaobaoItemAddRequest {
    s.ChangeProp = &v
    return s
}
func (s *TaobaoItemAddRequest) SetDescModules(v string) *TaobaoItemAddRequest {
    s.DescModules = &v
    return s
}
func (s *TaobaoItemAddRequest) SetIsOffline(v string) *TaobaoItemAddRequest {
    s.IsOffline = &v
    return s
}
func (s *TaobaoItemAddRequest) SetWirelessDesc(v string) *TaobaoItemAddRequest {
    s.WirelessDesc = &v
    return s
}
func (s *TaobaoItemAddRequest) SetChaoshiExtendsInfo(v string) *TaobaoItemAddRequest {
    s.ChaoshiExtendsInfo = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSpuConfirm(v bool) *TaobaoItemAddRequest {
    s.SpuConfirm = &v
    return s
}
func (s *TaobaoItemAddRequest) SetVideoId(v int64) *TaobaoItemAddRequest {
    s.VideoId = &v
    return s
}
func (s *TaobaoItemAddRequest) SetInteractiveId(v int64) *TaobaoItemAddRequest {
    s.InteractiveId = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLeaseExtendsInfo(v string) *TaobaoItemAddRequest {
    s.LeaseExtendsInfo = &v
    return s
}
func (s *TaobaoItemAddRequest) SetBrokerage(v string) *TaobaoItemAddRequest {
    s.Brokerage = &v
    return s
}
func (s *TaobaoItemAddRequest) SetBizCode(v string) *TaobaoItemAddRequest {
    s.BizCode = &v
    return s
}
func (s *TaobaoItemAddRequest) SetImageUrls(v []string) *TaobaoItemAddRequest {
    s.ImageUrls = &v
    return s
}
func (s *TaobaoItemAddRequest) SetInputPids(v []string) *TaobaoItemAddRequest {
    s.InputPids = &v
    return s
}
func (s *TaobaoItemAddRequest) SetInputStr(v []string) *TaobaoItemAddRequest {
    s.InputStr = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSkuProperties(v string) *TaobaoItemAddRequest {
    s.SkuProperties = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSkuQuantities(v string) *TaobaoItemAddRequest {
    s.SkuQuantities = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSkuPrices(v string) *TaobaoItemAddRequest {
    s.SkuPrices = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSkuOuterIds(v string) *TaobaoItemAddRequest {
    s.SkuOuterIds = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSkuBarcode(v string) *TaobaoItemAddRequest {
    s.SkuBarcode = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSkuSpecIds(v string) *TaobaoItemAddRequest {
    s.SkuSpecIds = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSkuDeliveryTimes(v string) *TaobaoItemAddRequest {
    s.SkuDeliveryTimes = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSkuHdLength(v string) *TaobaoItemAddRequest {
    s.SkuHdLength = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSkuHdHeight(v string) *TaobaoItemAddRequest {
    s.SkuHdHeight = &v
    return s
}
func (s *TaobaoItemAddRequest) SetSkuHdLampQuantity(v string) *TaobaoItemAddRequest {
    s.SkuHdLampQuantity = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocationState(v string) *TaobaoItemAddRequest {
    s.LocationState = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocationCity(v string) *TaobaoItemAddRequest {
    s.LocationCity = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityPrdLicenseNo(v string) *TaobaoItemAddRequest {
    s.FoodSecurityPrdLicenseNo = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityDesignCode(v string) *TaobaoItemAddRequest {
    s.FoodSecurityDesignCode = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityFactory(v string) *TaobaoItemAddRequest {
    s.FoodSecurityFactory = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityFactorySite(v string) *TaobaoItemAddRequest {
    s.FoodSecurityFactorySite = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityContact(v string) *TaobaoItemAddRequest {
    s.FoodSecurityContact = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityMix(v string) *TaobaoItemAddRequest {
    s.FoodSecurityMix = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityPlanStorage(v string) *TaobaoItemAddRequest {
    s.FoodSecurityPlanStorage = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityPeriod(v string) *TaobaoItemAddRequest {
    s.FoodSecurityPeriod = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityFoodAdditive(v string) *TaobaoItemAddRequest {
    s.FoodSecurityFoodAdditive = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecuritySupplier(v string) *TaobaoItemAddRequest {
    s.FoodSecuritySupplier = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityProductDateStart(v string) *TaobaoItemAddRequest {
    s.FoodSecurityProductDateStart = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityProductDateEnd(v string) *TaobaoItemAddRequest {
    s.FoodSecurityProductDateEnd = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityStockDateStart(v string) *TaobaoItemAddRequest {
    s.FoodSecurityStockDateStart = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityStockDateEnd(v string) *TaobaoItemAddRequest {
    s.FoodSecurityStockDateEnd = &v
    return s
}
func (s *TaobaoItemAddRequest) SetFoodSecurityHealthProductNo(v string) *TaobaoItemAddRequest {
    s.FoodSecurityHealthProductNo = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifeExpirydate(v string) *TaobaoItemAddRequest {
    s.LocalityLifeExpirydate = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifeNetworkId(v string) *TaobaoItemAddRequest {
    s.LocalityLifeNetworkId = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifeMerchant(v string) *TaobaoItemAddRequest {
    s.LocalityLifeMerchant = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifeVerification(v string) *TaobaoItemAddRequest {
    s.LocalityLifeVerification = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifeRefundRatio(v int64) *TaobaoItemAddRequest {
    s.LocalityLifeRefundRatio = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifeOnsaleAutoRefundRatio(v int64) *TaobaoItemAddRequest {
    s.LocalityLifeOnsaleAutoRefundRatio = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifeChooseLogis(v string) *TaobaoItemAddRequest {
    s.LocalityLifeChooseLogis = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifeRefundmafee(v string) *TaobaoItemAddRequest {
    s.LocalityLifeRefundmafee = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifeObs(v string) *TaobaoItemAddRequest {
    s.LocalityLifeObs = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifeEticket(v string) *TaobaoItemAddRequest {
    s.LocalityLifeEticket = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifeVersion(v string) *TaobaoItemAddRequest {
    s.LocalityLifeVersion = &v
    return s
}
func (s *TaobaoItemAddRequest) SetLocalityLifePackageid(v string) *TaobaoItemAddRequest {
    s.LocalityLifePackageid = &v
    return s
}
func (s *TaobaoItemAddRequest) SetPaimaiInfoMode(v int64) *TaobaoItemAddRequest {
    s.PaimaiInfoMode = &v
    return s
}
func (s *TaobaoItemAddRequest) SetPaimaiInfoDeposit(v int64) *TaobaoItemAddRequest {
    s.PaimaiInfoDeposit = &v
    return s
}
func (s *TaobaoItemAddRequest) SetPaimaiInfoInterval(v int64) *TaobaoItemAddRequest {
    s.PaimaiInfoInterval = &v
    return s
}
func (s *TaobaoItemAddRequest) SetPaimaiInfoReserve(v string) *TaobaoItemAddRequest {
    s.PaimaiInfoReserve = &v
    return s
}
func (s *TaobaoItemAddRequest) SetPaimaiInfoValidHour(v int64) *TaobaoItemAddRequest {
    s.PaimaiInfoValidHour = &v
    return s
}
func (s *TaobaoItemAddRequest) SetPaimaiInfoValidMinute(v int64) *TaobaoItemAddRequest {
    s.PaimaiInfoValidMinute = &v
    return s
}
func (s *TaobaoItemAddRequest) SetMsPaymentReferencePrice(v string) *TaobaoItemAddRequest {
    s.MsPaymentReferencePrice = &v
    return s
}
func (s *TaobaoItemAddRequest) SetMsPaymentVoucherPrice(v string) *TaobaoItemAddRequest {
    s.MsPaymentVoucherPrice = &v
    return s
}
func (s *TaobaoItemAddRequest) SetMsPaymentPrice(v string) *TaobaoItemAddRequest {
    s.MsPaymentPrice = &v
    return s
}
func (s *TaobaoItemAddRequest) SetDeliveryTimeNeedDeliveryTime(v string) *TaobaoItemAddRequest {
    s.DeliveryTimeNeedDeliveryTime = &v
    return s
}
func (s *TaobaoItemAddRequest) SetDeliveryTimeDeliveryTimeType(v string) *TaobaoItemAddRequest {
    s.DeliveryTimeDeliveryTimeType = &v
    return s
}
func (s *TaobaoItemAddRequest) SetDeliveryTimeDeliveryTime(v string) *TaobaoItemAddRequest {
    s.DeliveryTimeDeliveryTime = &v
    return s
}

func (req *TaobaoItemAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ApproveStatus != nil) {
        paramMap["approve_status"] = *req.ApproveStatus
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
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
    if(req.SupportCustomMade != nil) {
        paramMap["support_custom_made"] = *req.SupportCustomMade
    }
    if(req.CustomMadeTypeId != nil) {
        paramMap["custom_made_type_id"] = *req.CustomMadeTypeId
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
    if(req.Brokerage != nil) {
        paramMap["brokerage"] = *req.Brokerage
    }
    if(req.BizCode != nil) {
        paramMap["biz_code"] = *req.BizCode
    }
    if(req.ImageUrls != nil) {
        paramMap["image_urls"] = util.ConvertBasicList(*req.ImageUrls)
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
    if(req.MsPaymentReferencePrice != nil) {
        paramMap["ms_payment.reference_price"] = *req.MsPaymentReferencePrice
    }
    if(req.MsPaymentVoucherPrice != nil) {
        paramMap["ms_payment.voucher_price"] = *req.MsPaymentVoucherPrice
    }
    if(req.MsPaymentPrice != nil) {
        paramMap["ms_payment.price"] = *req.MsPaymentPrice
    }
    if(req.DeliveryTimeNeedDeliveryTime != nil) {
        paramMap["delivery_time.need_delivery_time"] = *req.DeliveryTimeNeedDeliveryTime
    }
    if(req.DeliveryTimeDeliveryTimeType != nil) {
        paramMap["delivery_time.delivery_time_type"] = *req.DeliveryTimeDeliveryTimeType
    }
    if(req.DeliveryTimeDeliveryTime != nil) {
        paramMap["delivery_time.delivery_time"] = *req.DeliveryTimeDeliveryTime
    }
    return paramMap
}

func (req *TaobaoItemAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}