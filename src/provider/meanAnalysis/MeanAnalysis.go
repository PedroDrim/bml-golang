package meananalysis

import(
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Classe de analise que implementa a interface "SimpleTableAnalysis"
type MeanAnalysis struct { }

// Construtor publico da classe
func NewMeanAnalysis() *MeanAnalysis {  
    return &MeanAnalysis{}
}

// Método responsável por obter a media de valores de credit na lista de usuarios
// @list Lista de usuarios
// @return Media de valores de credit da lista
func (obj *MeanAnalysis) Analysis(userInfoList []userInfo.UserInfo) float64 {
	var sum float64

	for i := range userInfoList {
		user := userInfoList[i]
		sum += user.GetCredit()
	}
	return sum/float64(len(userInfoList))
}
