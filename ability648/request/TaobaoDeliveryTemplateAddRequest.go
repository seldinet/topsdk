package request


type TaobaoDeliveryTemplateAddRequest struct {
    /*
        运费模板的名称，长度不能超过50个字节     */
    Name  *string `json:"name" required:"true" `
    /*
        可选值：0、1 ，说明如下<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费     */
    Assumer  *int64 `json:"assumer" required:"true" `
    /*
        可选值：0、1、3，说明如下。<br>0:表示按宝贝件数计算运费 <br>1:表示按宝贝重量计算运费
 <br>3:表示按宝贝体积计算运费     */
    Valuation  *int64 `json:"valuation" required:"true" `
    /*
        卖家发货地址区域ID
<br/><br/><font color=blue>可以不填，如果没有填写取卖家默认发货地址的区域ID，如果需要输入必须用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm 
</font>

<br/><br/><font color=red>注意：填入该值时必须取您的发货地址最小区域级别ID，比如您的发货地址是：某省XX市xx区（县）时需要填入区(县)的ID，当然有些地方没有区或县可以直接填市级别的ID</font>     */
    ConsignAreaId  *int64 `json:"consign_area_id,omitempty" required:"false" `
    /*
        运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod),物流宝保障速递(wlb),家装物流(furniture)结构:value1;value2;value3;value4 
如: post;express;ems;cod 
<br/><font color=red>
注意:在添加多个运费方式时,字符串中使用 ";" 分号区分
。template_dests(指定地区)
template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. </font>
<br>
<font color=blue>
注意：<br/>
1、post,ems,express三种运费方式必须填写一个<br/>
2、只有订购了货到付款用户可以填cod;只有家装物流用户可以填写furniture
只有订购了保障速递的用户可以填写bzsd,只有物流宝用户可以填写wlb<br/>
3、如果是货到付款用户当没有填写cod运送方式时，运费模板会默认继承express的费用为cod的费用<br/>
4、如果是保障速递户当没有填写bzsd运送方式时，运费模板会默认继承express的费用为bzsd的费用<br/>
5、由于3和4的原因所以当是货到付款用户或保障速递用户时如果没填写对应的运送方式express是必须填写的
</font>     */
    TemplateTypes  *string `json:"template_types" required:"true" `
    /*
        邮费子项涉及的地区.结构: value1;value2;value3,value4
<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)
如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应
<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm<br/>
<br/><font color=red>每个运费方式设置涉及的地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font>
<br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，
template_start_standards(首费标准)、template_start_fees(首费)、
template_add_standards(增费标准)、
template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font>
<font color=red>如果卖家没有传入发货地址则默认地址库的发货地址</font>     */
    TemplateDests  *string `json:"template_dests" required:"true" `
    /*
        首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）
<br/>
<font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>     */
    TemplateStartStandards  *string `json:"template_start_standards" required:"true" `
    /*
        首费：输入0.00-999.99（最多包含两位小数）
<br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>     */
    TemplateStartFees  *string `json:"template_start_fees" required:"true" `
    /*
        增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br/><font color=red>当valuation(记价方式)为1时输入0.1-9999.9范围内的小数只能包含以为小数（单位为千克）<br/><font color=blue>当valuation(记价方式)为3时输入0.1-999.9范围内的数值，数值只能包含一位小数（单位为 立方米）
<br/>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>     */
    TemplateAddStandards  *string `json:"template_add_standards" required:"true" `
    /*
        增费：输入0.00-999.99（最多包含两位小数）

<br/><br/><font color=blue>增费必须小于等于首费，但是当首费为0时增费可以大于首费</font>


<br/><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>     */
    TemplateAddFees  *string `json:"template_add_fees" required:"true" `
}

func (s *TaobaoDeliveryTemplateAddRequest) SetName(v string) *TaobaoDeliveryTemplateAddRequest {
    s.Name = &v
    return s
}
func (s *TaobaoDeliveryTemplateAddRequest) SetAssumer(v int64) *TaobaoDeliveryTemplateAddRequest {
    s.Assumer = &v
    return s
}
func (s *TaobaoDeliveryTemplateAddRequest) SetValuation(v int64) *TaobaoDeliveryTemplateAddRequest {
    s.Valuation = &v
    return s
}
func (s *TaobaoDeliveryTemplateAddRequest) SetConsignAreaId(v int64) *TaobaoDeliveryTemplateAddRequest {
    s.ConsignAreaId = &v
    return s
}
func (s *TaobaoDeliveryTemplateAddRequest) SetTemplateTypes(v string) *TaobaoDeliveryTemplateAddRequest {
    s.TemplateTypes = &v
    return s
}
func (s *TaobaoDeliveryTemplateAddRequest) SetTemplateDests(v string) *TaobaoDeliveryTemplateAddRequest {
    s.TemplateDests = &v
    return s
}
func (s *TaobaoDeliveryTemplateAddRequest) SetTemplateStartStandards(v string) *TaobaoDeliveryTemplateAddRequest {
    s.TemplateStartStandards = &v
    return s
}
func (s *TaobaoDeliveryTemplateAddRequest) SetTemplateStartFees(v string) *TaobaoDeliveryTemplateAddRequest {
    s.TemplateStartFees = &v
    return s
}
func (s *TaobaoDeliveryTemplateAddRequest) SetTemplateAddStandards(v string) *TaobaoDeliveryTemplateAddRequest {
    s.TemplateAddStandards = &v
    return s
}
func (s *TaobaoDeliveryTemplateAddRequest) SetTemplateAddFees(v string) *TaobaoDeliveryTemplateAddRequest {
    s.TemplateAddFees = &v
    return s
}

func (req *TaobaoDeliveryTemplateAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.Assumer != nil) {
        paramMap["assumer"] = *req.Assumer
    }
    if(req.Valuation != nil) {
        paramMap["valuation"] = *req.Valuation
    }
    if(req.ConsignAreaId != nil) {
        paramMap["consign_area_id"] = *req.ConsignAreaId
    }
    if(req.TemplateTypes != nil) {
        paramMap["template_types"] = *req.TemplateTypes
    }
    if(req.TemplateDests != nil) {
        paramMap["template_dests"] = *req.TemplateDests
    }
    if(req.TemplateStartStandards != nil) {
        paramMap["template_start_standards"] = *req.TemplateStartStandards
    }
    if(req.TemplateStartFees != nil) {
        paramMap["template_start_fees"] = *req.TemplateStartFees
    }
    if(req.TemplateAddStandards != nil) {
        paramMap["template_add_standards"] = *req.TemplateAddStandards
    }
    if(req.TemplateAddFees != nil) {
        paramMap["template_add_fees"] = *req.TemplateAddFees
    }
    return paramMap
}

func (req *TaobaoDeliveryTemplateAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}