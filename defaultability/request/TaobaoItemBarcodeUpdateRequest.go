package request


type TaobaoItemBarcodeUpdateRequest struct {
    /*
        被更新商品的ID     */
    ItemId  *int64 `json:"item_id" required:"true" `
    /*
        商品条形码，如果不用更新，可选择不填     */
    ItemBarcode  *string `json:"item_barcode,omitempty" required:"false" `
    /*
        被更新SKU的ID列表，中间以英文逗号进行分隔。如果没有SKU或者不需要更新SKU的条形码，不需要设置     */
    SkuIds  *string `json:"sku_ids,omitempty" required:"false" `
    /*
        SKU维度的条形码，和sku_ids字段一一对应，中间以英文逗号分隔     */
    SkuBarcodes  *string `json:"sku_barcodes,omitempty" required:"false" `
    /*
        是否强制保存商品条码。true：强制保存false ：需要执行条码库校验 defalutValue��false    */
    Isforce  *bool `json:"isforce,omitempty" required:"false" `
    /*
        访问来源，这字段提供给千牛扫码枪用，其他调用方，不需要填写     */
    Src  *string `json:"src,omitempty" required:"false" `
}

func (s *TaobaoItemBarcodeUpdateRequest) SetItemId(v int64) *TaobaoItemBarcodeUpdateRequest {
    s.ItemId = &v
    return s
}
func (s *TaobaoItemBarcodeUpdateRequest) SetItemBarcode(v string) *TaobaoItemBarcodeUpdateRequest {
    s.ItemBarcode = &v
    return s
}
func (s *TaobaoItemBarcodeUpdateRequest) SetSkuIds(v string) *TaobaoItemBarcodeUpdateRequest {
    s.SkuIds = &v
    return s
}
func (s *TaobaoItemBarcodeUpdateRequest) SetSkuBarcodes(v string) *TaobaoItemBarcodeUpdateRequest {
    s.SkuBarcodes = &v
    return s
}
func (s *TaobaoItemBarcodeUpdateRequest) SetIsforce(v bool) *TaobaoItemBarcodeUpdateRequest {
    s.Isforce = &v
    return s
}
func (s *TaobaoItemBarcodeUpdateRequest) SetSrc(v string) *TaobaoItemBarcodeUpdateRequest {
    s.Src = &v
    return s
}

func (req *TaobaoItemBarcodeUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemId != nil) {
        paramMap["item_id"] = *req.ItemId
    }
    if(req.ItemBarcode != nil) {
        paramMap["item_barcode"] = *req.ItemBarcode
    }
    if(req.SkuIds != nil) {
        paramMap["sku_ids"] = *req.SkuIds
    }
    if(req.SkuBarcodes != nil) {
        paramMap["sku_barcodes"] = *req.SkuBarcodes
    }
    if(req.Isforce != nil) {
        paramMap["isforce"] = *req.Isforce
    }
    if(req.Src != nil) {
        paramMap["src"] = *req.Src
    }
    return paramMap
}

func (req *TaobaoItemBarcodeUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}