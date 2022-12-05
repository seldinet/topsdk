package request

import (
        "topsdk/util"
    )

type TaobaoItempropsGetRequest struct {
    /*
        需要返回的字段列表，见：ItemProp，默认返回：pid, name, must, multi, prop_values defalutValue��pid,name,must,multi,prop_values    */
    Fields  *[]string `json:"fields,omitempty" required:"false" `
    /*
        叶子类目ID，如果只传cid，则只返回一级属性,通过taobao.itemcats.get获得叶子类目ID     */
    Cid  *int64 `json:"cid" required:"true" `
    /*
        属性id (取类目属性时，传pid，不用同时传PID和parent_pid)     */
    Pid  *int64 `json:"pid,omitempty" required:"false" `
    /*
        父属性ID     */
    ParentPid  *int64 `json:"parent_pid,omitempty" required:"false" `
    /*
        是否关键属性。可选值:true(是),false(否)     */
    IsKeyProp  *bool `json:"is_key_prop,omitempty" required:"false" `
    /*
        是否销售属性。可选值:true(是),false(否)     */
    IsSaleProp  *bool `json:"is_sale_prop,omitempty" required:"false" `
    /*
        是否颜色属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)     */
    IsColorProp  *bool `json:"is_color_prop,omitempty" required:"false" `
    /*
        是否枚举属性。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)。如果返回true，属性值是下拉框选择输入，如果返回false，属性值是用户自行手工输入。     */
    IsEnumProp  *bool `json:"is_enum_prop,omitempty" required:"false" `
    /*
        在is_enum_prop是true的前提下，是否是卖家可以自行输入的属性（注：如果is_enum_prop返回false，该参数统一返回false）。可选值:true(是),false(否) (删除的属性不会匹配和返回这个条件)     */
    IsInputProp  *bool `json:"is_input_prop,omitempty" required:"false" `
    /*
        是否商品属性，这个属性只能放于发布商品时使用。可选值:true(是),false(否)     */
    IsItemProp  *bool `json:"is_item_prop,omitempty" required:"false" `
    /*
        增量时间戳。格式:yyyy-MM-dd HH:mm:ss假如传2005-01-01 00:00:00，则取所有的属性和子属性ID(如果传了pid会忽略datetime)     */
    Datetime  *util.LocalTime `json:"datetime,omitempty" required:"false" `
    /*
        类目子属性路径,由该子属性上层的类目属性和类目属性值组成,格式pid:vid;pid:vid.取类目子属性需要传child_path,cid     */
    ChildPath  *string `json:"child_path,omitempty" required:"false" `
    /*
        获取类目的类型：1代表集市、2代表天猫 defalutValue��1    */
    Type  *int64 `json:"type,omitempty" required:"false" `
    /*
        属性的Key，支持多条，以“,”分隔     */
    AttrKeys  *[]string `json:"attr_keys,omitempty" required:"false" `
}

func (s *TaobaoItempropsGetRequest) SetFields(v []string) *TaobaoItempropsGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetCid(v int64) *TaobaoItempropsGetRequest {
    s.Cid = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetPid(v int64) *TaobaoItempropsGetRequest {
    s.Pid = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetParentPid(v int64) *TaobaoItempropsGetRequest {
    s.ParentPid = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetIsKeyProp(v bool) *TaobaoItempropsGetRequest {
    s.IsKeyProp = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetIsSaleProp(v bool) *TaobaoItempropsGetRequest {
    s.IsSaleProp = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetIsColorProp(v bool) *TaobaoItempropsGetRequest {
    s.IsColorProp = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetIsEnumProp(v bool) *TaobaoItempropsGetRequest {
    s.IsEnumProp = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetIsInputProp(v bool) *TaobaoItempropsGetRequest {
    s.IsInputProp = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetIsItemProp(v bool) *TaobaoItempropsGetRequest {
    s.IsItemProp = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetDatetime(v util.LocalTime) *TaobaoItempropsGetRequest {
    s.Datetime = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetChildPath(v string) *TaobaoItempropsGetRequest {
    s.ChildPath = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetType(v int64) *TaobaoItempropsGetRequest {
    s.Type = &v
    return s
}
func (s *TaobaoItempropsGetRequest) SetAttrKeys(v []string) *TaobaoItempropsGetRequest {
    s.AttrKeys = &v
    return s
}

func (req *TaobaoItempropsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.Cid != nil) {
        paramMap["cid"] = *req.Cid
    }
    if(req.Pid != nil) {
        paramMap["pid"] = *req.Pid
    }
    if(req.ParentPid != nil) {
        paramMap["parent_pid"] = *req.ParentPid
    }
    if(req.IsKeyProp != nil) {
        paramMap["is_key_prop"] = *req.IsKeyProp
    }
    if(req.IsSaleProp != nil) {
        paramMap["is_sale_prop"] = *req.IsSaleProp
    }
    if(req.IsColorProp != nil) {
        paramMap["is_color_prop"] = *req.IsColorProp
    }
    if(req.IsEnumProp != nil) {
        paramMap["is_enum_prop"] = *req.IsEnumProp
    }
    if(req.IsInputProp != nil) {
        paramMap["is_input_prop"] = *req.IsInputProp
    }
    if(req.IsItemProp != nil) {
        paramMap["is_item_prop"] = *req.IsItemProp
    }
    if(req.Datetime != nil) {
        paramMap["datetime"] = *req.Datetime
    }
    if(req.ChildPath != nil) {
        paramMap["child_path"] = *req.ChildPath
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.AttrKeys != nil) {
        paramMap["attr_keys"] = util.ConvertBasicList(*req.AttrKeys)
    }
    return paramMap
}

func (req *TaobaoItempropsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}