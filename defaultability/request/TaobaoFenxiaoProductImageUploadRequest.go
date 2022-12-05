package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoProductImageUploadRequest struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        产品主图图片空间相对路径或绝对路径     */
    PicPath  *string `json:"pic_path,omitempty" required:"false" `
    /*
        产品图片     */
    Image  *[]byte `json:"image,omitempty" required:"false" `
    /*
        图片位置，0-14之间。0：操作sku属性图片，1：主图，2-5：细节图，6-14：额外主图     */
    Position  *int64 `json:"position" required:"true" `
    /*
        properties表示sku图片的属性。key:value形式，key是pid，value是vid。如果position是0的话，则properties需要是必传项     */
    Properties  *string `json:"properties,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoProductImageUploadRequest) SetProductId(v int64) *TaobaoFenxiaoProductImageUploadRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoProductImageUploadRequest) SetPicPath(v string) *TaobaoFenxiaoProductImageUploadRequest {
    s.PicPath = &v
    return s
}
func (s *TaobaoFenxiaoProductImageUploadRequest) SetImage(v []byte) *TaobaoFenxiaoProductImageUploadRequest {
    s.Image = &v
    return s
}
func (s *TaobaoFenxiaoProductImageUploadRequest) SetPosition(v int64) *TaobaoFenxiaoProductImageUploadRequest {
    s.Position = &v
    return s
}
func (s *TaobaoFenxiaoProductImageUploadRequest) SetProperties(v string) *TaobaoFenxiaoProductImageUploadRequest {
    s.Properties = &v
    return s
}

func (req *TaobaoFenxiaoProductImageUploadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.PicPath != nil) {
        paramMap["pic_path"] = *req.PicPath
    }
    if(req.Position != nil) {
        paramMap["position"] = *req.Position
    }
    if(req.Properties != nil) {
        paramMap["properties"] = *req.Properties
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductImageUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}