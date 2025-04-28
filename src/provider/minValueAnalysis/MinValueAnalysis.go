package minvalueanalysis

import(
	"math"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Classe de analise que implementa a interface "SimpleTableAnalysis"
type MinValueAnalysis struct { }

// Construtor publico da classe
func NewMinValueAnalysis() *MinValueAnalysis {  
    return &MinValueAnalysis{}
}

// Método responsável por obter o menor valor de credit na lista de usuarios
// @list Lista de usuarios
// @return Valor minimo da lista
func (obj *MinValueAnalysis) Analysis(userInfoList []userinfo.UserInfo) float64 {
	minValue := math.MaxFloat64

	for i := range userInfoList {
		user := userInfoList[i]
		if user.GetCredit() < minValue {
			minValue = user.GetCredit()
		}
	}
	return minValue
}