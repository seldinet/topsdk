package request


type TaobaoScitemMapDeleteRequest struct {
    /*
        后台商品ID     */
    ScItemId  *int64 `json:"sc_item_id" required:"true" `
    /*
        店铺用户nick。 如果该参数为空则删除后端商品与当前调用人的商品映射关系;如果不为空则删除指定用户与后端商品的映射关系     */
    UserNick  *string `json:"user_nick,omitempty" required:"false" `
}

func (s *TaobaoScitemMapDeleteRequest) SetScItemId(v int64) *TaobaoScitemMapDeleteRequest {
    s.ScItemId = &v
    return s
}
func (s *TaobaoScitemMapDeleteRequest) SetUserNick(v string) *TaobaoScitemMapDeleteRequest {
    s.UserNick = &v
    return s
}

func (req *TaobaoScitemMapDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ScItemId != nil) {
        paramMap["sc_item_id"] = *req.ScItemId
    }
    if(req.UserNick != nil) {
        paramMap["user_nick"] = *req.UserNick
    }
    return paramMap
}

func (req *TaobaoScitemMapDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}