package domain

import (
        "topsdk/util"
    )

type TaobaoProductPropimgUploadProductPropImg struct {
    /*
        图片地址.(绝对地址,格式:http://host/image_path)     */
    Url  *string `json:"url,omitempty" `

    /*
        产品属性图片ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        添加时间.格式:yyyy-mm-dd hh:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        修改时间.格式:yyyy-mm-dd hh:mm:ss     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoProductPropimgUploadProductPropImg) SetUrl(v string) *TaobaoProductPropimgUploadProductPropImg {
    s.Url = &v
    return s
}
func (s *TaobaoProductPropimgUploadProductPropImg) SetId(v int64) *TaobaoProductPropimgUploadProductPropImg {
    s.Id = &v
    return s
}
func (s *TaobaoProductPropimgUploadProductPropImg) SetCreated(v util.LocalTime) *TaobaoProductPropimgUploadProductPropImg {
    s.Created = &v
    return s
}
func (s *TaobaoProductPropimgUploadProductPropImg) SetModified(v util.LocalTime) *TaobaoProductPropimgUploadProductPropImg {
    s.Modified = &v
    return s
}
