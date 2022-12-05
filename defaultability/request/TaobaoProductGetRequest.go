package request

import (
        "topsdk/util"
    )

type TaobaoProductGetRequest struct {
    /*
        需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔.     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        Product的id.两种方式来查看一个产品:1.传入product_id来查询 2.传入cid和props来查询     */
    ProductId  *int64 `json:"product_id,omitempty" required:"false" `
    /*
        商品类目id.调用taobao.itemcats.get获取;必须是叶子类目id,如果没有传product_id,那么cid和props必须要传.     */
    Cid  *int64 `json:"cid,omitempty" required:"false" `
    /*
        比如:诺基亚N73这个产品的关键属性列表就是:品牌:诺基亚;型号:N73,对应的PV值就是10005:10027;10006:29729.     */
    Props  *string `json:"props,omitempty" required:"false" `
}

func (s *TaobaoProductGetRequest) SetFields(v []string) *TaobaoProductGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoProductGetRequest) SetProductId(v int64) *TaobaoProductGetRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductGetRequest) SetCid(v int64) *TaobaoProductGetRequest {
    s.Cid = &v
    return s
}
func (s *TaobaoProductGetRequest) SetProps(v string) *TaobaoProductGetRequest {
    s.Props = &v
    return s
}

func (req *TaobaoProductGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.Cid != nil) {
        paramMap["cid"] = *req.Cid
    }
    if(req.Props != nil) {
        paramMap["props"] = *req.Props
    }
    return paramMap
}

func (req *TaobaoProductGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}