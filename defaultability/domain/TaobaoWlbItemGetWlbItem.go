package domain

import (
        "topsdk/util"
    )

type TaobaoWlbItemGetWlbItem struct {
    /*
        商品id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        商品所有人淘宝ID     */
    UserId  *int64 `json:"user_id,omitempty" `

    /*
        商品所有人淘宝nick     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        商品的名称     */
    Name  *string `json:"name,omitempty" `

    /*
        前台商品名称     */
    Title  *string `json:"title,omitempty" `

    /*
        商家编码     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        是不是sku商品
值为true或false     */
    IsSku  *bool `json:"is_sku,omitempty" `

    /*
        标记，用逗号隔开的字符串。
BIT_HAS_AUTHORIZE 第1位 是否有授权规则;
BATCH  第2位 是否有批次规则；
SYNCHRONIZATION 第3位 是否有同步规则。     */
    Flag  *string `json:"flag,omitempty" `

    /*
        商品类型：
NORMAL--普通类型;
COMBINE--组合商品;
DISTRIBUTION--分销商品;
默认为NORMAL     */
    Type  *string `json:"type,omitempty" `

    /*
        商品备注     */
    Remark  *string `json:"remark,omitempty" `

    /*
        父item的id，当item为物流宝子商品时，parent_id必填,否则不必填
可通过父ID来得知商品的关系。     */
    ParentId  *int64 `json:"parent_id,omitempty" `

    /*
        状态，item_status_valid -- 1 表示 有效 item_status_lock -- 2 表示锁住     */
    Status  *string `json:"status,omitempty" `

    /*
        发布版本号，用来同步商     */
    PublishVersion  *int64 `json:"publish_version,omitempty" `

    /*
        创建人     */
    Creator  *string `json:"creator,omitempty" `

    /*
        品牌ID     */
    BrandId  *int64 `json:"brand_id,omitempty" `

    /*
        创建日期     */
    GmtCreate  *util.LocalTime `json:"gmt_create,omitempty" `

    /*
        最后修改人     */
    LastModifier  *string `json:"last_modifier,omitempty" `

    /*
        修改日期     */
    GmtModified  *util.LocalTime `json:"gmt_modified,omitempty" `

    /*
        属性key:value; key:value     */
    Properties  *string `json:"properties,omitempty" `

    /*
        是否易碎     */
    IsFriable  *bool `json:"is_friable,omitempty" `

    /*
        是否危险品     */
    IsDangerous  *bool `json:"is_dangerous,omitempty" `

    /*
        颜色     */
    Color  *string `json:"color,omitempty" `

    /*
        重量     */
    Weight  *int64 `json:"weight,omitempty" `

    /*
        mm     */
    Length  *int64 `json:"length,omitempty" `

    /*
        宽     */
    Width  *int64 `json:"width,omitempty" `

    /*
        高     */
    Height  *int64 `json:"height,omitempty" `

    /*
        立方mm     */
    Volume  *int64 `json:"volume,omitempty" `

    /*
        货类     */
    GoodsCat  *string `json:"goods_cat,omitempty" `

    /*
        计价货类     */
    PricingCat  *string `json:"pricing_cat,omitempty" `

    /*
        包装材料     */
    PackageMaterial  *string `json:"package_material,omitempty" `

    /*
        价格     */
    Price  *int64 `json:"price,omitempty" `

}

func (s *TaobaoWlbItemGetWlbItem) SetId(v int64) *TaobaoWlbItemGetWlbItem {
    s.Id = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetUserId(v int64) *TaobaoWlbItemGetWlbItem {
    s.UserId = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetUserNick(v string) *TaobaoWlbItemGetWlbItem {
    s.UserNick = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetName(v string) *TaobaoWlbItemGetWlbItem {
    s.Name = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetTitle(v string) *TaobaoWlbItemGetWlbItem {
    s.Title = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetItemCode(v string) *TaobaoWlbItemGetWlbItem {
    s.ItemCode = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetIsSku(v bool) *TaobaoWlbItemGetWlbItem {
    s.IsSku = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetFlag(v string) *TaobaoWlbItemGetWlbItem {
    s.Flag = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetType(v string) *TaobaoWlbItemGetWlbItem {
    s.Type = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetRemark(v string) *TaobaoWlbItemGetWlbItem {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetParentId(v int64) *TaobaoWlbItemGetWlbItem {
    s.ParentId = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetStatus(v string) *TaobaoWlbItemGetWlbItem {
    s.Status = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetPublishVersion(v int64) *TaobaoWlbItemGetWlbItem {
    s.PublishVersion = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetCreator(v string) *TaobaoWlbItemGetWlbItem {
    s.Creator = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetBrandId(v int64) *TaobaoWlbItemGetWlbItem {
    s.BrandId = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetGmtCreate(v util.LocalTime) *TaobaoWlbItemGetWlbItem {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetLastModifier(v string) *TaobaoWlbItemGetWlbItem {
    s.LastModifier = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetGmtModified(v util.LocalTime) *TaobaoWlbItemGetWlbItem {
    s.GmtModified = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetProperties(v string) *TaobaoWlbItemGetWlbItem {
    s.Properties = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetIsFriable(v bool) *TaobaoWlbItemGetWlbItem {
    s.IsFriable = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetIsDangerous(v bool) *TaobaoWlbItemGetWlbItem {
    s.IsDangerous = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetColor(v string) *TaobaoWlbItemGetWlbItem {
    s.Color = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetWeight(v int64) *TaobaoWlbItemGetWlbItem {
    s.Weight = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetLength(v int64) *TaobaoWlbItemGetWlbItem {
    s.Length = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetWidth(v int64) *TaobaoWlbItemGetWlbItem {
    s.Width = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetHeight(v int64) *TaobaoWlbItemGetWlbItem {
    s.Height = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetVolume(v int64) *TaobaoWlbItemGetWlbItem {
    s.Volume = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetGoodsCat(v string) *TaobaoWlbItemGetWlbItem {
    s.GoodsCat = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetPricingCat(v string) *TaobaoWlbItemGetWlbItem {
    s.PricingCat = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetPackageMaterial(v string) *TaobaoWlbItemGetWlbItem {
    s.PackageMaterial = &v
    return s
}
func (s *TaobaoWlbItemGetWlbItem) SetPrice(v int64) *TaobaoWlbItemGetWlbItem {
    s.Price = &v
    return s
}
