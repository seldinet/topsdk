package request

import (
        "topsdk/util"
    )

type TaobaoItempropvaluesGetRequest struct {
    /*
        需要返回的字段。目前支持有：cid,pid,prop_name,vid,name,name_alias,status,sort_order     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        属性和属性值 id串，格式例如(pid1;pid2)或(pid1:vid1;pid2:vid2)或(pid1;pid2:vid2)     */
    Pvs  *string `json:"pvs,omitempty" required:"false" `
    /*
        叶子类目ID ,通过taobao.itemcats.get获得叶子类目ID     */
    Cid  *int64 `json:"cid" required:"true" `
    /*
        假如传2005-01-01 00:00:00，则取所有的属性和子属性(状态为删除的属性值不返回prop_name)     */
    Datetime  *util.LocalTime `json:"datetime,omitempty" required:"false" `
    /*
        获取类目的类型：1代表集市、2代表天猫 defalutValue��1    */
    Type  *int64 `json:"type,omitempty" required:"false" `
    /*
        属性的Key，支持多条，以“,”分隔     */
    AttrKeys  *[]string `json:"attr_keys,omitempty" required:"false" `
}

func (s *TaobaoItempropvaluesGetRequest) SetFields(v []string) *TaobaoItempropvaluesGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoItempropvaluesGetRequest) SetPvs(v string) *TaobaoItempropvaluesGetRequest {
    s.Pvs = &v
    return s
}
func (s *TaobaoItempropvaluesGetRequest) SetCid(v int64) *TaobaoItempropvaluesGetRequest {
    s.Cid = &v
    return s
}
func (s *TaobaoItempropvaluesGetRequest) SetDatetime(v util.LocalTime) *TaobaoItempropvaluesGetRequest {
    s.Datetime = &v
    return s
}
func (s *TaobaoItempropvaluesGetRequest) SetType(v int64) *TaobaoItempropvaluesGetRequest {
    s.Type = &v
    return s
}
func (s *TaobaoItempropvaluesGetRequest) SetAttrKeys(v []string) *TaobaoItempropvaluesGetRequest {
    s.AttrKeys = &v
    return s
}

func (req *TaobaoItempropvaluesGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.Pvs != nil) {
        paramMap["pvs"] = *req.Pvs
    }
    if(req.Cid != nil) {
        paramMap["cid"] = *req.Cid
    }
    if(req.Datetime != nil) {
        paramMap["datetime"] = *req.Datetime
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    if(req.AttrKeys != nil) {
        paramMap["attr_keys"] = util.ConvertBasicList(*req.AttrKeys)
    }
    return paramMap
}

func (req *TaobaoItempropvaluesGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}