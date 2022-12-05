package request


type TaobaoItemJointPropimgRequest struct {
    /*
        商品数字ID，必选     */
    NumIid  *int64 `json:"num_iid" required:"true" `
    /*
        属性列表。调用taobao.itemprops.get获取，属性必须是颜色属性，格式:pid:vid。     */
    Properties  *string `json:"properties" required:"true" `
    /*
        属性图片ID。如果是新增不需要填写 defalutValue��0    */
    Id  *int64 `json:"id,omitempty" required:"false" `
    /*
        图片序号 defalutValue��0    */
    Position  *int64 `json:"position,omitempty" required:"false" `
    /*
        图片地址(传入图片相对地址即可,即不需包含 http://img02.taobao.net/bao/uploaded )     */
    PicPath  *string `json:"pic_path" required:"true" `
}

func (s *TaobaoItemJointPropimgRequest) SetNumIid(v int64) *TaobaoItemJointPropimgRequest {
    s.NumIid = &v
    return s
}
func (s *TaobaoItemJointPropimgRequest) SetProperties(v string) *TaobaoItemJointPropimgRequest {
    s.Properties = &v
    return s
}
func (s *TaobaoItemJointPropimgRequest) SetId(v int64) *TaobaoItemJointPropimgRequest {
    s.Id = &v
    return s
}
func (s *TaobaoItemJointPropimgRequest) SetPosition(v int64) *TaobaoItemJointPropimgRequest {
    s.Position = &v
    return s
}
func (s *TaobaoItemJointPropimgRequest) SetPicPath(v string) *TaobaoItemJointPropimgRequest {
    s.PicPath = &v
    return s
}

func (req *TaobaoItemJointPropimgRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NumIid != nil) {
        paramMap["num_iid"] = *req.NumIid
    }
    if(req.Properties != nil) {
        paramMap["properties"] = *req.Properties
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
    return paramMap
}

func (req *TaobaoItemJointPropimgRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}