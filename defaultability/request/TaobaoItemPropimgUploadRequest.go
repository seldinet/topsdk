package request

import (
        "topsdk/util"
    )

type TaobaoItemPropimgUploadRequest struct {
    /*
        商品数字ID，必选     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        属性图片ID。如果是新增不需要填写     */
    Id  *int64 `json:"id,omitempty" required:"false" `
    /*
        属性列表。调用taobao.itemprops.get获取类目属性，属性必须是颜色属性，再用taobao.itempropvalues.get取得vid。格式:pid:vid。     */
    Properties  *string `json:"properties" required:"true" `
    /*
        图片位置 defalutValue��0    */
    Position  *int64 `json:"position,omitempty" required:"false" `
    /*
        属性图片内容。类型:JPG,GIF;图片大小不超过:3M     */
    Image  *[]byte `json:"image,omitempty" required:"false" `
}

func (s *TaobaoItemPropimgUploadRequest) SetNumIid(v int64) *TaobaoItemPropimgUploadRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemPropimgUploadRequest) SetId(v int64) *TaobaoItemPropimgUploadRequest {
    s.Id = &v
    return s
}
func (s *TaobaoItemPropimgUploadRequest) SetProperties(v string) *TaobaoItemPropimgUploadRequest {
    s.Properties = &v
    return s
}
func (s *TaobaoItemPropimgUploadRequest) SetPosition(v int64) *TaobaoItemPropimgUploadRequest {
    s.Position = &v
    return s
}
func (s *TaobaoItemPropimgUploadRequest) SetImage(v []byte) *TaobaoItemPropimgUploadRequest {
    s.Image = &v
    return s
}

func (req *TaobaoItemPropimgUploadRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    if(req.Id != nil) {
        paramMap["id"] = *req.Id
    }
    if(req.Properties != nil) {
        paramMap["properties"] = *req.Properties
    }
    if(req.Position != nil) {
        paramMap["position"] = *req.Position
    }
    return paramMap
}

func (req *TaobaoItemPropimgUploadRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}