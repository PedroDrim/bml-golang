package simpletableanalysis

import(
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Interface para analise de tabela simples
type SimpleTableAnalysis interface {
	Analysis(userInfoList []userinfo.UserInfo) float64
}