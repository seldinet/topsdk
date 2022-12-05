package request

import (
        "topsdk/util"
    )

type TaobaoProductAddRequest struct {
    /*
        商品类目ID.调用taobao.itemcats.get获取;注意:必须是叶子类目 id.     */
    Cid  *int64 `json:"cid" required:"true" `
    /*
        产品名称,最大30个字符.     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        产品市场价.精确到2位小数;单位为元.如：200.07     */
    Price  *string `json:"price,omitempty" required:"false" `
    /*
        产品主图片.最大1M,目前仅支持GIF,JPG.     */
    Image  *[]byte `json:"image" required:"true" `
    /*
        外部产品ID     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        关键属性 结构:pid:vid;pid:vid.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;如果碰到用户自定义属性,请用customer_props.     */
    Props  *string `json:"props,omitempty" required:"false" `
    /*
        非关键属性结构:pid:vid;pid:vid.<br>
非关键属性<font color=red>不包含</font>关键属性、销售属性、用户自定义属性、商品属性;
<br>调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid.<br><font color=red>注:支持最大长度为512字节</font>     */
    Binds  *string `json:"binds,omitempty" required:"false" `
    /*
        销售属性结构:pid:vid;pid:vid.调用taobao.itemprops.get获取is_sale_prop＝true的pid,调用taobao.itempropvalues.get获取vid.     */
    SaleProps  *string `json:"sale_props,omitempty" required:"false" `
    /*
        用户自定义属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”
<br><font color=red>注：包含所有自定义属性的传入</font>     */
    CustomerProps  *string `json:"customer_props,omitempty" required:"false" `
    /*
        产品描述.最大不超过25000个字符     */
    Desc  *string `json:"desc,omitempty" required:"false" `
    /*
        是不是主图 defalutValue��true    */
    Major  *bool `json:"major,omitempty" required:"false" `
    /*
        native_unkeyprops     */
    NativeUnkeyprops  *string `json:"native_unkeyprops,omitempty" required:"false" `
    /*
        加入垂直市场，目前只支持以鞋城卖家身份加入名鞋馆(暂时此字段还不起作用，不对外开放)     */
    VerticalMarket  *int64 `json:"vertical_market,omitempty" required:"false" `
    /*
        上市时间。目前只支持鞋城类目传入此参数     */
    MarketTime  *util.LocalTime `json:"market_time,omitempty" required:"false" `
    /*
        销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。只有少数销售属性值支持传入别名，比如颜色和尺寸     */
    PropertyAlias  *string `json:"property_alias,omitempty" required:"false" `
}

func (s *TaobaoProductAddRequest) SetCid(v int64) *TaobaoProductAddRequest {
    s.Cid = &v
    return s
}
func (s *TaobaoProductAddRequest) SetName(v string) *TaobaoProductAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoProductAddRequest) SetPrice(v string) *TaobaoProductAddRequest {
    s.Price = &v
    return s
}
func (s *TaobaoProductAddRequest) SetImage(v []byte) *TaobaoProductAddRequest {
    s.Image = &v
    return s
}
func (s *TaobaoProductAddRequest) SetOuterId(v string) *TaobaoProductAddRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoProductAddRequest) SetProps(v string) *TaobaoProductAddRequest {
    s.Props = &v
    return s
}
func (s *TaobaoProductAddRequest) SetBinds(v string) *TaobaoProductAddRequest {
    s.Binds = &v
    return s
}
func (s *TaobaoProductAddRequest) SetSaleProps(v string) *TaobaoProductAddRequest {
    s.SaleProps = &v
    return s
}
func (s *TaobaoProductAddRequest) SetCustomerProps(v string) *TaobaoProductAddRequest {
    s.CustomerProps = &v
    return s
}
func (s *TaobaoProductAddRequest) SetDesc(v string) *TaobaoProductAddRequest {
    s.Desc = &v
    return s
}
func (s *TaobaoProductAddRequest) SetMajor(v bool) *TaobaoProductAddRequest {
    s.Major = &v
    return s
}
func (s *TaobaoProductAddRequest) SetNativeUnkeyprops(v string) *TaobaoProductAddRequest {
    s.NativeUnkeyprops = &v
    return s
}
func (s *TaobaoProductAddRequest) SetVerticalMarket(v int64) *TaobaoProductAddRequest {
    s.VerticalMarket = &v
    return s
}
func (s *TaobaoProductAddRequest) SetMarketTime(v util.LocalTime) *TaobaoProductAddRequest {
    s.MarketTime = &v
    return s
}
func (s *TaobaoProductAddRequest) SetPropertyAlias(v string) *TaobaoProductAddRequest {
    s.PropertyAlias = &v
    return s
}

func (req *TaobaoProductAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Cid != nil) {
        paramMap["cid"] = *req.Cid
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.Price != nil) {
        paramMap["price"] = *req.Price
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Props != nil) {
        paramMap["props"] = *req.Props
    }
    if(req.Binds != nil) {
        paramMap["binds"] = *req.Binds
    }
    if(req.SaleProps != nil) {
        paramMap["sale_props"] = *req.SaleProps
    }
    if(req.CustomerProps != nil) {
        paramMap["customer_props"] = *req.CustomerProps
    }
    if(req.Desc != nil) {
        paramMap["desc"] = *req.Desc
    }
    if(req.Major != nil) {
        paramMap["major"] = *req.Major
    }
    if(req.NativeUnkeyprops != nil) {
        paramMap["native_unkeyprops"] = *req.NativeUnkeyprops
    }
    if(req.VerticalMarket != nil) {
        paramMap["vertical_market"] = *req.VerticalMarket
    }
    if(req.MarketTime != nil) {
        paramMap["market_time"] = *req.MarketTime
    }
    if(req.PropertyAlias != nil) {
        paramMap["property_alias"] = *req.PropertyAlias
    }
    return paramMap
}

func (req *TaobaoProductAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}