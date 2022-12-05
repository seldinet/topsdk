package request

import (
        "topsdk/util"
    )

type TaobaoProductImgUploadRequest struct {
    /*
        产品图片ID.修改图片时需要传入     */
    Id  *int64 `json:"id,omitempty" required:"false" `
    /*
        产品ID.Product的id     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        图片内容.图片最大为500K,只支持JPG,GIF格式.     */
    Image  *[]byte `json:"image" required:"true" `
    /*
        图片序号 defalutValue��0    */
    Position  *int64 `json:"position,omitempty" required:"false" `
    /*
        是否将该图片设为主图.可选值:true,false;默认值:false.     */
    IsMajor  *bool `json:"is_major,omitempty" required:"false" `
}

func (s *TaobaoProductImgUploadRequest) SetId(v int64) *TaobaoProductImgUploadRequest {
    s.Id = &v
    return s
}
func (s *TaobaoProductImgUploadRequest) SetProductId(v int64) *TaobaoProductImgUploadRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductImgUploadRequest) SetImage(v []byte) *TaobaoProductImgUploadRequest {
    s.Image = &v
    return s
}
func (s *TaobaoProductImgUploadRequest) SetPosition(v int64) *TaobaoProductImgUploadRequest {
    s.Position = &v
    return s
}
func (s *TaobaoProductImgUploadRequest) SetIsMajor(v bool) *TaobaoProductImgUploadRequest {
    s.IsMajor = &v
    return s
}

func (req *TaobaoProductImgUploadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.Position != nil) {
        paramMap["position"] = *req.Position
    }
    if(req.IsMajor != nil) {
        paramMap["is_major"] = *req.IsMajor
    }
    return paramMap
}

func (req *TaobaoProductImgUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}