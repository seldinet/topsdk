package request


type TaobaoItemJointImgRequest struct {
    /*
        商品数字ID，必选     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        商品图片id(如果是更新图片，则需要传该参数) defalutValue��0    */
    Id  *int64 `json:"id,omitempty" required:"false" `
    /*
        图片序号 defalutValue��0    */
    Position  *int64 `json:"position,omitempty" required:"false" `
    /*
        图片URL,图片空间图片的相对地址，支持的文件类型：jpg,jpeg,png     */
    PicPath  *string `json:"pic_path" required:"true" `
    /*
        上传的图片是否关联为商品主图 defalutValue��false    */
    IsMajor  *bool `json:"is_major,omitempty" required:"false" `
    /*
        是否3:4长方形图片，绑定3:4主图视频时用于上传3:4商品主图 defalutValue��false    */
    IsRectangle  *bool `json:"is_rectangle,omitempty" required:"false" `
}

func (s *TaobaoItemJointImgRequest) SetNumIid(v int64) *TaobaoItemJointImgRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemJointImgRequest) SetId(v int64) *TaobaoItemJointImgRequest {
    s.Id = &v
    return s
}
func (s *TaobaoItemJointImgRequest) SetPosition(v int64) *TaobaoItemJointImgRequest {
    s.Position = &v
    return s
}
func (s *TaobaoItemJointImgRequest) SetPicPath(v string) *TaobaoItemJointImgRequest {
    s.PicPath = &v
    return s
}
func (s *TaobaoItemJointImgRequest) SetIsMajor(v bool) *TaobaoItemJointImgRequest {
    s.IsMajor = &v
    return s
}
func (s *TaobaoItemJointImgRequest) SetIsRectangle(v bool) *TaobaoItemJointImgRequest {
    s.IsRectangle = &v
    return s
}

func (req *TaobaoItemJointImgRequest) ToMap() map[string]interface{} {
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
    if(req.PicPath != nil) {
        paramMap["pic_path"] = *req.PicPath
    }
    if(req.IsMajor != nil) {
        paramMap["is_major"] = *req.IsMajor
    }
    if(req.IsRectangle != nil) {
        paramMap["is_rectangle"] = *req.IsRectangle
    }
    return paramMap
}

func (req *TaobaoItemJointImgRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}