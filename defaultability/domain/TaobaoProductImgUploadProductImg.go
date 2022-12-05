package domain

import (
        "topsdk/util"
    )

type TaobaoProductImgUploadProductImg struct {
    /*
        产品图片ID     */
    Id  *int64 `json:"id,omitempty" `

    /*
        图片地址.(绝对地址,格式:http://host/image_path)     */
    Url  *string `json:"url,omitempty" `

    /*
        添加时间.格式:yyyy-mm-dd hh:mm:ss     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        修改时间.格式:yyyy-mm-dd hh:mm:ss     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoProductImgUploadProductImg) SetId(v int64) *TaobaoProductImgUploadProductImg {
    s.Id = &v
    return s
}
func (s *TaobaoProductImgUploadProductImg) SetUrl(v string) *TaobaoProductImgUploadProductImg {
    s.Url = &v
    return s
}
func (s *TaobaoProductImgUploadProductImg) SetCreated(v util.LocalTime) *TaobaoProductImgUploadProductImg {
    s.Created = &v
    return s
}
func (s *TaobaoProductImgUploadProductImg) SetModified(v util.LocalTime) *TaobaoProductImgUploadProductImg {
    s.Modified = &v
    return s
}
