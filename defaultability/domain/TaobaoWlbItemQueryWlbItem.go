package domain

import (
        "topsdk/util"
    )

type TaobaoWlbItemQueryWlbItem struct {
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

func (s *TaobaoWlbItemQueryWlbItem) SetId(v int64) *TaobaoWlbItemQueryWlbItem {
    s.Id = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetUserId(v int64) *TaobaoWlbItemQueryWlbItem {
    s.UserId = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetUserNick(v string) *TaobaoWlbItemQueryWlbItem {
    s.UserNick = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetName(v string) *TaobaoWlbItemQueryWlbItem {
    s.Name = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetTitle(v string) *TaobaoWlbItemQueryWlbItem {
    s.Title = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetItemCode(v string) *TaobaoWlbItemQueryWlbItem {
    s.ItemCode = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetIsSku(v bool) *TaobaoWlbItemQueryWlbItem {
    s.IsSku = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetFlag(v string) *TaobaoWlbItemQueryWlbItem {
    s.Flag = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetType(v string) *TaobaoWlbItemQueryWlbItem {
    s.Type = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetRemark(v string) *TaobaoWlbItemQueryWlbItem {
    s.Remark = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetParentId(v int64) *TaobaoWlbItemQueryWlbItem {
    s.ParentId = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetStatus(v string) *TaobaoWlbItemQueryWlbItem {
    s.Status = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetPublishVersion(v int64) *TaobaoWlbItemQueryWlbItem {
    s.PublishVersion = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetCreator(v string) *TaobaoWlbItemQueryWlbItem {
    s.Creator = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetBrandId(v int64) *TaobaoWlbItemQueryWlbItem {
    s.BrandId = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetGmtCreate(v util.LocalTime) *TaobaoWlbItemQueryWlbItem {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetLastModifier(v string) *TaobaoWlbItemQueryWlbItem {
    s.LastModifier = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetGmtModified(v util.LocalTime) *TaobaoWlbItemQueryWlbItem {
    s.GmtModified = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetProperties(v string) *TaobaoWlbItemQueryWlbItem {
    s.Properties = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetIsFriable(v bool) *TaobaoWlbItemQueryWlbItem {
    s.IsFriable = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetIsDangerous(v bool) *TaobaoWlbItemQueryWlbItem {
    s.IsDangerous = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetColor(v string) *TaobaoWlbItemQueryWlbItem {
    s.Color = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetWeight(v int64) *TaobaoWlbItemQueryWlbItem {
    s.Weight = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetLength(v int64) *TaobaoWlbItemQueryWlbItem {
    s.Length = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetWidth(v int64) *TaobaoWlbItemQueryWlbItem {
    s.Width = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetHeight(v int64) *TaobaoWlbItemQueryWlbItem {
    s.Height = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetVolume(v int64) *TaobaoWlbItemQueryWlbItem {
    s.Volume = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetGoodsCat(v string) *TaobaoWlbItemQueryWlbItem {
    s.GoodsCat = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetPricingCat(v string) *TaobaoWlbItemQueryWlbItem {
    s.PricingCat = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetPackageMaterial(v string) *TaobaoWlbItemQueryWlbItem {
    s.PackageMaterial = &v
    return s
}
func (s *TaobaoWlbItemQueryWlbItem) SetPrice(v int64) *TaobaoWlbItemQueryWlbItem {
    s.Price = &v
    return s
}
