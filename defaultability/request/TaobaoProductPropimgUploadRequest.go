package request

import (
        "topsdk/util"
    )

type TaobaoProductPropimgUploadRequest struct {
    /*
        产品属性图片ID     */
    Id  *int64 `json:"id,omitempty" required:"false" `
    /*
        产品ID.Product的id     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        属性串.目前仅支持颜色属性.调用taobao.itemprops.get获取类目属性,取得颜色属性pid,再用taobao.itempropvalues.get取得vid;格式:pid:vid,只能传入一个颜色pid:vid串;     */
    Props  *string `json:"props" required:"true" `
    /*
        图片内容.图片最大为2M,只支持JPG,GIF.     */
    Image  *[]byte `json:"image" required:"true" `
    /*
        图片序号 defalutValue��0    */
    Position  *int64 `json:"position,omitempty" required:"false" `
}

func (s *TaobaoProductPropimgUploadRequest) SetId(v int64) *TaobaoProductPropimgUploadRequest {
    s.Id = &v
    return s
}
func (s *TaobaoProductPropimgUploadRequest) SetProductId(v int64) *TaobaoProductPropimgUploadRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductPropimgUploadRequest) SetProps(v string) *TaobaoProductPropimgUploadRequest {
    s.Props = &v
    return s
}
func (s *TaobaoProductPropimgUploadRequest) SetImage(v []byte) *TaobaoProductPropimgUploadRequest {
    s.Image = &v
    return s
}
func (s *TaobaoProductPropimgUploadRequest) SetPosition(v int64) *TaobaoProductPropimgUploadRequest {
    s.Position = &v
    return s
}

func (req *TaobaoProductPropimgUploadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.Props != nil) {
        paramMap["props"] = *req.Props
    }
    if(req.Position != nil) {
        paramMap["position"] = *req.Position
    }
    return paramMap
}

func (req *TaobaoProductPropimgUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}