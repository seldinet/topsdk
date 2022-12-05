package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoDistributorProductsGetRequest struct {
    /*
        产品ID列表（最大限制30），用逗号分割，例如：“1001,1002,1003,1004,1005”     */
    Pids  *[]int64 `json:"pids,omitempty" required:"false" `
    /*
        根据商品ID列表查询，优先级次于产品ID列表，高于其他分页查询条件。如果商品不是分销商品，自动过滤。最大限制20，用逗号分割，例如：“1001,1002,1003,1004,1005”     */
    ItemIds  *[]int64 `json:"item_ids,omitempty" required:"false" `
    /*
        供应商nick，分页查询时，必传     */
    SupplierNick  *string `json:"supplier_nick,omitempty" required:"false" `
    /*
        分销方式；可选择：AGENT ： 代销；DEALER：经销；DIRECT：直营     */
    TradeType  *string `json:"trade_type,omitempty" required:"false" `
    /*
        下载状态，默认是未下载；可选值：UNDOWNLOAD：未下载 ；DOWNLOADED：已下载；下载：指将供应商授权的产品发布为店铺新宝贝的过程，下载成功后，分销商需要将新生成的宝贝重新编辑并上架后售卖。     */
    DownloadStatus  *string `json:"download_status,omitempty" required:"false" `
    /*
        产品线ID     */
    ProductcatId  *int64 `json:"productcat_id,omitempty" required:"false" `
    /*
        指定查询额外的信息，可选值：skus（sku数据）、images（多图），多个可选值用逗号分割。     */
    Fields  *[]string `json:"fields,omitempty" required:"false" `
    /*
        开始修改时间     */
    StartTime  *util.LocalTime `json:"start_time,omitempty" required:"false" `
    /*
        结束时间     */
    EndTime  *util.LocalTime `json:"end_time,omitempty" required:"false" `
    /*
        time_type     */
    TimeType  *string `json:"time_type,omitempty" required:"false" `
    /*
        页码（大于0的整数，默认1）     */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页记录数（默认20，最大50）     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
    /*
        order_by     */
    OrderBy  *string `json:"order_by,omitempty" required:"false" `
}

func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetPids(v []int64) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.Pids = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetItemIds(v []int64) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.ItemIds = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetSupplierNick(v string) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.SupplierNick = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetTradeType(v string) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.TradeType = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetDownloadStatus(v string) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.DownloadStatus = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetProductcatId(v int64) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.ProductcatId = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetFields(v []string) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetStartTime(v util.LocalTime) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.StartTime = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetEndTime(v util.LocalTime) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.EndTime = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetTimeType(v string) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.TimeType = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetPageNo(v int64) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetPageSize(v int64) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.PageSize = &v
    return s
}
func (s *TaobaoFenxiaoDistributorProductsGetRequest) SetOrderBy(v string) *TaobaoFenxiaoDistributorProductsGetRequest {
    s.OrderBy = &v
    return s
}

func (req *TaobaoFenxiaoDistributorProductsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Pids != nil) {
        paramMap["pids"] = util.ConvertBasicList(*req.Pids)
    }
    if(req.ItemIds != nil) {
        paramMap["item_ids"] = util.ConvertBasicList(*req.ItemIds)
    }
    if(req.SupplierNick != nil) {
        paramMap["supplier_nick"] = *req.SupplierNick
    }
    if(req.TradeType != nil) {
        paramMap["trade_type"] = *req.TradeType
    }
    if(req.DownloadStatus != nil) {
        paramMap["download_status"] = *req.DownloadStatus
    }
    if(req.ProductcatId != nil) {
        paramMap["productcat_id"] = *req.ProductcatId
    }
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.StartTime != nil) {
        paramMap["start_time"] = *req.StartTime
    }
    if(req.EndTime != nil) {
        paramMap["end_time"] = *req.EndTime
    }
    if(req.TimeType != nil) {
        paramMap["time_type"] = *req.TimeType
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    if(req.OrderBy != nil) {
        paramMap["order_by"] = *req.OrderBy
    }
    return paramMap
}

func (req *TaobaoFenxiaoDistributorProductsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}