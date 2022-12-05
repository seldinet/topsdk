package request

import (
        "topsdk/util"
    )

type TaobaoProductUpdateRequest struct {
    /*
        产品ID     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        外部产品ID     */
    OuterId  *string `json:"outer_id,omitempty" required:"false" `
    /*
        非关键属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid     */
    Binds  *string `json:"binds,omitempty" required:"false" `
    /*
        销售属性.调用taobao.itemprops.get获取pid,调用taobao.itempropvalues.get获取vid;格式:pid:vid;pid:vid     */
    SaleProps  *string `json:"sale_props,omitempty" required:"false" `
    /*
        产品名称.最大不超过30个字符     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        产品市场价.精确到2位小数;单位为元.如:200.07     */
    Price  *string `json:"price,omitempty" required:"false" `
    /*
        产品描述.最大不超过25000个字符     */
    Desc  *string `json:"desc,omitempty" required:"false" `
    /*
        是否是主图 defalutValue��true    */
    Major  *bool `json:"major,omitempty" required:"false" `
    /*
        自定义非关键属性 defalutValue��native_unkeyprops    */
    NativeUnkeyprops  *string `json:"native_unkeyprops,omitempty" required:"false" `
    /*
        产品主图.最大500K,目前仅支持GIF,JPG     */
    Image  *[]byte `json:"image,omitempty" required:"false" `
}

func (s *TaobaoProductUpdateRequest) SetProductId(v int64) *TaobaoProductUpdateRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoProductUpdateRequest) SetOuterId(v string) *TaobaoProductUpdateRequest {
    s.OuterId = &v
    return s
}
func (s *TaobaoProductUpdateRequest) SetBinds(v string) *TaobaoProductUpdateRequest {
    s.Binds = &v
    return s
}
func (s *TaobaoProductUpdateRequest) SetSaleProps(v string) *TaobaoProductUpdateRequest {
    s.SaleProps = &v
    return s
}
func (s *TaobaoProductUpdateRequest) SetName(v string) *TaobaoProductUpdateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoProductUpdateRequest) SetPrice(v string) *TaobaoProductUpdateRequest {
    s.Price = &v
    return s
}
func (s *TaobaoProductUpdateRequest) SetDesc(v string) *TaobaoProductUpdateRequest {
    s.Desc = &v
    return s
}
func (s *TaobaoProductUpdateRequest) SetMajor(v bool) *TaobaoProductUpdateRequest {
    s.Major = &v
    return s
}
func (s *TaobaoProductUpdateRequest) SetNativeUnkeyprops(v string) *TaobaoProductUpdateRequest {
    s.NativeUnkeyprops = &v
    return s
}
func (s *TaobaoProductUpdateRequest) SetImage(v []byte) *TaobaoProductUpdateRequest {
    s.Image = &v
    return s
}

func (req *TaobaoProductUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.OuterId != nil) {
        paramMap["outer_id"] = *req.OuterId
    }
    if(req.Binds != nil) {
        paramMap["binds"] = *req.Binds
    }
    if(req.SaleProps != nil) {
        paramMap["sale_props"] = *req.SaleProps
    }
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.Price != nil) {
        paramMap["price"] = *req.Price
    }
    if(req.Desc != nil) {
        paramMap["desc"] = *req.Desc
    }
    if(req.Major != nil) {
        paramMap["major"] = *req.Major
    }
    if(req.NativeUnkeyprops != nil) {
        paramMap["native_unkeyprops"] = *req.NativeUnkeyprops
    }
    return paramMap
}

func (req *TaobaoProductUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    if req.Image != nil {
        fileMap["image"] = *req.Image
    }
    return fileMap
}