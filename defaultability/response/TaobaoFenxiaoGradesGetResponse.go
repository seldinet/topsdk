package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoFenxiaoGradesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        分销商等级信息
    */
    FenxiaoGrades  []domain.TaobaoFenxiaoGradesGetFenxiaoGrade `json:"fenxiao_grades,omitempty" `
}
