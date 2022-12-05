package request

import (
        "topsdk/util"
    )

type TaobaoItemImgUploadRequest struct {
    /*
        商品数字ID，该参数必须     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        商品图片id(如果是更新图片，则需要传该参数)     */
    Id  *int64 `json:"id,omitempty" required:"false" `
    /*
        图片序号     */
    Position  *int64 `json:"position,omitempty" required:"false" `
    /*
        商品图片内容类型:JPG;最大:3M 。支持的文件类型：jpg,jpeg,png     */
    Image  *[]byte `json:"image,omitempty" required:"false" `
    /*
        是否将该图片设为主图,可选值:true,false;默认值:false(非主图) defalutValue��false    */
    IsMajor  *bool `json:"is_major,omitempty" required:"false" `
    /*
        是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图 defalutValue��false    */
    IsRectangle  *bool `json:"is_rectangle,omitempty" required:"false" `
}

func (s *TaobaoItemImgUploadRequest) SetNumIid(v int64) *TaobaoItemImgUploadRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemImgUploadRequest) SetId(v int64) *TaobaoItemImgUploadRequest {
    s.Id = &v
    return s
}
func (s *TaobaoItemImgUploadRequest) SetPosition(v int64) *TaobaoItemImgUploadRequest {
    s.Position = &v
    return s
}
func (s *TaobaoItemImgUploadRequest) SetImage(v []byte) *TaobaoItemImgUploadRequest {
    s.Image = &v
    return s
}
func (s *TaobaoItemImgUploadRequest) SetIsMajor(v bool) *TaobaoItemImgUploadRequest {
    s.IsMajor = &v
    return s
}
func (s *TaobaoItemImgUploadRequest) SetIsRectangle(v bool) *TaobaoItemImgUploadRequest {
    s.IsRectangle = &v
    return s
}

func (req *TaobaoItemImgUploadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    if(req.Position != nil) {
        paramMap["position"] = *req.Position
    }
    if(req.IsMajor != nil) {
        paramMap["is_major"] = *req.IsMajor
    }
    if(req.IsRectangle != nil) {
        paramMap["is_rectangle"] = *req.IsRectangle
    }
    return paramMap
}

func (req *TaobaoItemImgUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}