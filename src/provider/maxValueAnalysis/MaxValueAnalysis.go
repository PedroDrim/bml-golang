package maxvalueanalysis

import(
	"math"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Classe de analise que implementa a interface "SimpleTableAnalysis"
type MaxValueAnalysis struct { }

// Construtor publico da classe
func NewMaxValueAnalysis() *MaxValueAnalysis {  
    return &MaxValueAnalysis{}
}

// Método responsável por obter o maior valor de credit na lista de usuarios
// @list Lista de usuarios
// @return Valor maximo da lista
func (obj *MaxValueAnalysis) Analysis(userInfoList []userinfo.UserInfo) float64 {
	maxValue := math.SmallestNonzeroFloat64

	for i := range userInfoList {
		user := userInfoList[i]
		if user.GetCredit() > maxValue {
			maxValue = user.GetCredit()
		}
	}
	return maxValue
}
