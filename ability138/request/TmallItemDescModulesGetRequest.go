package request


type TmallItemDescModulesGetRequest struct {
    /*
        淘宝后台发布商品的叶子类目id，可通过taobao.itemcats.get查到。api 访问地址http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.CFhhk4&path=cid:3-apiId:122     */
    CatId  *int64 `json:"cat_id" required:"true" `
    /*
        商家主帐号id     */
    UsrId  *string `json:"usr_id" required:"true" `
}

func (s *TmallItemDescModulesGetRequest) SetCatId(v int64) *TmallItemDescModulesGetRequest {
    s.CatId = &v
    return s
}
func (s *TmallItemDescModulesGetRequest) SetUsrId(v string) *TmallItemDescModulesGetRequest {
    s.UsrId = &v
    return s
}

func (req *TmallItemDescModulesGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CatId != nil) {
        paramMap["cat_id"] = *req.CatId
    }
    if(req.UsrId != nil) {
        paramMap["usr_id"] = *req.UsrId
    }
    return paramMap
}

func (req *TmallItemDescModulesGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}