package tableanalysis

import(
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Interface para analise dos dados
// @param <T> Tipo de classe do resultado
type TableAnalysis[T any] interface {

    // Realiza a analise dos dados
    // @param userInfoList Lista de dados a ser analisada
    // @return Resultado da analise
    Analysis(userInfoList []userinfo.UserInfo) (T, error)
}
