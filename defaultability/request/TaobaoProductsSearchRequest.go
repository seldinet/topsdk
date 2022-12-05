package request

import (
        "topsdk/util"
    )

type TaobaoProductsSearchRequest struct {
    /*
        需返回的字段列表.可选值:Product数据结构中的以下字段:product_id,name,pic_url,cid,props,price,tsc;多个字段之间用","分隔.新增字段status(product的当前状态)     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        搜索的关键词是用来搜索产品的title.　注:q,cid和props至少传入一个     */
    Q  *string `json:"q,omitempty" required:"false" `
    /*
        商品类目ID.调用taobao.itemcats.get获取.     */
    Cid  *int64 `json:"cid,omitempty" required:"false" `
    /*
        属性,属性值的组合.格式:pid:vid;pid:vid;调用taobao.itemprops.get获取类目属性pid 
,再用taobao.itempropvalues.get取得vid.     */
    Props  *string `json:"props,omitempty" required:"false" `
    /*
        想要获取的产品的状态列表，支持多个状态并列获取，多个状态之间用","分隔，最多同时指定5种状态。例如，只获取小二确认的spu传入"3",只要商家确认的传入"0"，既要小二确认又要商家确认的传入"0,3"。目前只支持者两种类型的状态搜索，输入其他状态无效。     */
    Status  *string `json:"status,omitempty" required:"false" `
    /*
        页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始.     */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页条数.每页返回最多返回100条,默认值为40.     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        传入值为：3表示3C表示3C垂直市场产品，4表示鞋城垂直市场产品，8表示网游垂直市场产品。一次只能指定一种垂直市场类型     */
    VerticalMarket  *int64 `json:"vertical_market,omitempty" required:"false" `
    /*
        用户自定义关键属性,结构：pid1:value1;pid2:value2，如果有型号，系列等子属性用: 隔开 例如：“20000:优衣库:型号:001;632501:1234”，表示“品牌:优衣库:型号:001;货号:1234”     */
    CustomerProps  *string `json:"customer_props,omitempty" required:"false" `
    /*
        按关联产品规格specs搜索套装产品     */
    SuiteItemsStr  *string `json:"suite_items_str,omitempty" required:"false" `
    /*
        按条码搜索产品信息,多个逗号隔开，不支持条码为全零的方式     */
    BarcodeStr  *string `json:"barcode_str,omitempty" required:"false" `
    /*
        市场ID，1为取C2C市场的产品信息， 2为取B2C市场的产品信息。  不填写此值则默认取C2C的产品信息。     */
    MarketId  *string `json:"market_id,omitempty" required:"false" `
}

func (s *TaobaoProductsSearchRequest) SetFields(v []string) *TaobaoProductsSearchRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoProductsSearchRequest) SetQ(v string) *TaobaoProductsSearchRequest {
    s.Q = &v
    return s
}
func (s *TaobaoProductsSearchRequest) SetCid(v int64) *TaobaoProductsSearchRequest {
    s.Cid = &v
    return s
}
func (s *TaobaoProductsSearchRequest) SetProps(v string) *TaobaoProductsSearchRequest {
    s.Props = &v
    return s
}
func (s *TaobaoProductsSearchRequest) SetStatus(v string) *TaobaoProductsSearchRequest {
    s.Status = &v
    return s
}
func (s *TaobaoProductsSearchRequest) SetPageNo(v int64) *TaobaoProductsSearchRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoProductsSearchRequest) SetPageSize(v int64) *TaobaoProductsSearchRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoProductsSearchRequest) SetVerticalMarket(v int64) *TaobaoProductsSearchRequest {
    s.VerticalMarket = &v
    return s
}
func (s *TaobaoProductsSearchRequest) SetCustomerProps(v string) *TaobaoProductsSearchRequest {
    s.CustomerProps = &v
    return s
}
func (s *TaobaoProductsSearchRequest) SetSuiteItemsStr(v string) *TaobaoProductsSearchRequest {
    s.SuiteItemsStr = &v
    return s
}
func (s *TaobaoProductsSearchRequest) SetBarcodeStr(v string) *TaobaoProductsSearchRequest {
    s.BarcodeStr = &v
    return s
}
func (s *TaobaoProductsSearchRequest) SetMarketId(v string) *TaobaoProductsSearchRequest {
    s.MarketId = &v
    return s
}

func (req *TaobaoProductsSearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.Q != nil) {
        paramMap["q"] = *req.Q
    }
    if(req.Cid != nil) {
        paramMap["cid"] = *req.Cid
    }
    if(req.Props != nil) {
        paramMap["props"] = *req.Props
    }
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.VerticalMarket != nil) {
        paramMap["vertical_market"] = *req.VerticalMarket
    }
    if(req.CustomerProps != nil) {
        paramMap["customer_props"] = *req.CustomerProps
    }
    if(req.SuiteItemsStr != nil) {
        paramMap["suite_items_str"] = *req.SuiteItemsStr
    }
    if(req.BarcodeStr != nil) {
        paramMap["barcode_str"] = *req.BarcodeStr
    }
    if(req.MarketId != nil) {
        paramMap["market_id"] = *req.MarketId
    }
    return paramMap
}

func (req *TaobaoProductsSearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}