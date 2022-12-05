package domain

import (
        "topsdk/util"
    )

type TaobaoProductsGetProduct struct {
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
        产品的子图片.目前最多支持4张。fields中设置为product_imgs.id、product_imgs.url、product_imgs.position 等形式就会返回相应的字段     */
    ProductImgs  *[]TaobaoProductsGetProductImg `json:"product_imgs,omitempty" `

    /*
        产品的属性图片.比如说黄色对应的产品图片,绿色对应的产品图片。fields中设置为product_prop_imgs.id、 
product_prop_imgs.props、product_prop_imgs.url、product_prop_imgs.position等形式就会返回相应的字段     */
    ProductPropImgs  *[]TaobaoProductsGetProductPropImg `json:"product_prop_imgs,omitempty" `

}

func (s *TaobaoProductsGetProduct) SetProductId(v int64) *TaobaoProductsGetProduct {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetOuterId(v string) *TaobaoProductsGetProduct {
    s.OuterId = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetTsc(v string) *TaobaoProductsGetProduct {
    s.Tsc = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetCid(v int64) *TaobaoProductsGetProduct {
    s.Cid = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetCatName(v string) *TaobaoProductsGetProduct {
    s.CatName = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetProps(v string) *TaobaoProductsGetProduct {
    s.Props = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetPropsStr(v string) *TaobaoProductsGetProduct {
    s.PropsStr = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetName(v string) *TaobaoProductsGetProduct {
    s.Name = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetBinds(v string) *TaobaoProductsGetProduct {
    s.Binds = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetBindsStr(v string) *TaobaoProductsGetProduct {
    s.BindsStr = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetSaleProps(v string) *TaobaoProductsGetProduct {
    s.SaleProps = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetSalePropsStr(v string) *TaobaoProductsGetProduct {
    s.SalePropsStr = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetPrice(v string) *TaobaoProductsGetProduct {
    s.Price = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetDesc(v string) *TaobaoProductsGetProduct {
    s.Desc = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetPicUrl(v string) *TaobaoProductsGetProduct {
    s.PicUrl = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetCreated(v util.LocalTime) *TaobaoProductsGetProduct {
    s.Created = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetModified(v util.LocalTime) *TaobaoProductsGetProduct {
    s.Modified = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetStatus(v int64) *TaobaoProductsGetProduct {
    s.Status = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetProductImgs(v []TaobaoProductsGetProductImg) *TaobaoProductsGetProduct {
    s.ProductImgs = &v
    return s
}
func (s *TaobaoProductsGetProduct) SetProductPropImgs(v []TaobaoProductsGetProductPropImg) *TaobaoProductsGetProduct {
    s.ProductPropImgs = &v
    return s
}
