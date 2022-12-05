package domain

import (
        "topsdk/util"
    )

type TaobaoProductGetProduct struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id,omitempty" `

    /*
        外部产品ID     */
    OuterId  *string `json:"outer_id,omitempty" `

    /*
        淘宝标准产品编码     */
    Tsc  *string `json:"tsc,omitempty" `

    /*
        商品类目ID.必须是叶子类目ID     */
    Cid  *int64 `json:"cid,omitempty" `

    /*
        商品类目名称     */
    CatName  *string `json:"cat_name,omitempty" `

    /*
        产品的关键属性列表.格式：pid:vid;pid:vid     */
    Props  *string `json:"props,omitempty" `

    /*
        产品的关键属性字符串列表.比如:品牌:诺基亚;型号:N73(<strong>注：</strong><font color="red">属性名称中的冒号":"被转换为："#cln#";  
分号";"被转换为："#scln#"
</font>)     */
    PropsStr  *string `json:"props_str,omitempty" `

    /*
        产品名称     */
    Name  *string `json:"name,omitempty" `

    /*
        产品的非关键属性列表.格式:pid:vid;pid:vid.     */
    Binds  *string `json:"binds,omitempty" `

    /*
        产品的非关键属性字符串列表.格式同props_str(<strong>注：</strong><font color="red">属性名称中的冒号":"被转换为："#cln#";  
分号";"被转换为："#scln#"
</font>)     */
    BindsStr  *string `json:"binds_str,omitempty" `

    /*
        产品的销售属性列表.格式:pid:vid;pid:vid     */
    SaleProps  *string `json:"sale_props,omitempty" `

    /*
        产品的销售属性字符串列表.格式同props_str(<strong>注：</strong><font color="red">属性名称中的冒号":"被转换为："#cln#";  
分号";"被转换为："#scln#"
</font>)     */
    SalePropsStr  *string `json:"sale_props_str,omitempty" `

    /*
        产品的市场价.单位为元.精确到2位小数;如:200.07     */
    Price  *string `json:"price,omitempty" `

    /*
        产品的描述.最大25000个字节     */
    Desc  *string `json:"desc,omitempty" `

    /*
        产品的主图片地址.(绝对地址,格式:http://host/image_path)     */
    PicUrl  *string `json:"pic_url,omitempty" `

    /*
        创建时间.格式:yyyy-mm-dd hh:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        修改时间.格式:yyyy-mm-dd hh:mm:ss     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

    /*
        当前状态(0 商家确认 1 屏蔽 3 小二确认 2 未确认 -1 删除)     */
    Status  *int64 `json:"status,omitempty" `

    /*
        垂直市场,如：3（3C），4（鞋城）     */
    VerticalMarket  *int64 `json:"vertical_market,omitempty" `

    /*
        销售属性值别名。格式为pid1:vid1:alias1;pid1:vid2:alia2。     */
    PropertyAlias  *string `json:"property_alias,omitempty" `

    /*
        用户自定义属性,结构：pid1:value1;pid2:value2 例如：“20000:优衣库”，表示“品牌:优衣库”     */
    CustomerProps  *string `json:"customer_props,omitempty" `

    /*
        产品卖点描述，长度限制20个汉字     */
    SellPt  *string `json:"sell_pt,omitempty" `

    /*
        产品的子图片.目前最多支持4张。fields中设置为product_imgs.id、product_imgs.url、product_imgs.position 等形式就会返回相应的字段     */
    ProductImgs  *[]TaobaoProductGetProductImg `json:"product_imgs,omitempty" `

    /*
        产品的属性图片.比如说黄色对应的产品图片,绿色对应的产品图片。fields中设置为product_prop_imgs.id、 
product_prop_imgs.props、product_prop_imgs.url、product_prop_imgs.position等形式就会返回相应的字段     */
    ProductPropImgs  *[]TaobaoProductGetProductPropImg `json:"product_prop_imgs,omitempty" `

}

func (s *TaobaoProductGetProduct) SetProductId(v int64) *TaobaoProductGetProduct {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductGetProduct) SetOuterId(v string) *TaobaoProductGetProduct {
    s.OuterId = &v
    return s
}
func (s *TaobaoProductGetProduct) SetTsc(v string) *TaobaoProductGetProduct {
    s.Tsc = &v
    return s
}
func (s *TaobaoProductGetProduct) SetCid(v int64) *TaobaoProductGetProduct {
    s.Cid = &v
    return s
}
func (s *TaobaoProductGetProduct) SetCatName(v string) *TaobaoProductGetProduct {
    s.CatName = &v
    return s
}
func (s *TaobaoProductGetProduct) SetProps(v string) *TaobaoProductGetProduct {
    s.Props = &v
    return s
}
func (s *TaobaoProductGetProduct) SetPropsStr(v string) *TaobaoProductGetProduct {
    s.PropsStr = &v
    return s
}
func (s *TaobaoProductGetProduct) SetName(v string) *TaobaoProductGetProduct {
    s.Name = &v
    return s
}
func (s *TaobaoProductGetProduct) SetBinds(v string) *TaobaoProductGetProduct {
    s.Binds = &v
    return s
}
func (s *TaobaoProductGetProduct) SetBindsStr(v string) *TaobaoProductGetProduct {
    s.BindsStr = &v
    return s
}
func (s *TaobaoProductGetProduct) SetSaleProps(v string) *TaobaoProductGetProduct {
    s.SaleProps = &v
    return s
}
func (s *TaobaoProductGetProduct) SetSalePropsStr(v string) *TaobaoProductGetProduct {
    s.SalePropsStr = &v
    return s
}
func (s *TaobaoProductGetProduct) SetPrice(v string) *TaobaoProductGetProduct {
    s.Price = &v
    return s
}
func (s *TaobaoProductGetProduct) SetDesc(v string) *TaobaoProductGetProduct {
    s.Desc = &v
    return s
}
func (s *TaobaoProductGetProduct) SetPicUrl(v string) *TaobaoProductGetProduct {
    s.PicUrl = &v
    return s
}
func (s *TaobaoProductGetProduct) SetCreated(v util.LocalTime) *TaobaoProductGetProduct {
    s.Created = &v
    return s
}
func (s *TaobaoProductGetProduct) SetModified(v util.LocalTime) *TaobaoProductGetProduct {
    s.Modified = &v
    return s
}
func (s *TaobaoProductGetProduct) SetStatus(v int64) *TaobaoProductGetProduct {
    s.Status = &v
    return s
}
func (s *TaobaoProductGetProduct) SetVerticalMarket(v int64) *TaobaoProductGetProduct {
    s.VerticalMarket = &v
    return s
}
func (s *TaobaoProductGetProduct) SetPropertyAlias(v string) *TaobaoProductGetProduct {
    s.PropertyAlias = &v
    return s
}
func (s *TaobaoProductGetProduct) SetCustomerProps(v string) *TaobaoProductGetProduct {
    s.CustomerProps = &v
    return s
}
func (s *TaobaoProductGetProduct) SetSellPt(v string) *TaobaoProductGetProduct {
    s.SellPt = &v
    return s
}
func (s *TaobaoProductGetProduct) SetProductImgs(v []TaobaoProductGetProductImg) *TaobaoProductGetProduct {
    s.ProductImgs = &v
    return s
}
func (s *TaobaoProductGetProduct) SetProductPropImgs(v []TaobaoProductGetProductPropImg) *TaobaoProductGetProduct {
    s.ProductPropImgs = &v
    return s
}
