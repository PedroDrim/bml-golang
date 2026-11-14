package languagesortanalysis

import (
	"sort"

	"github.com/pedrodrim/bml-golang/src/model/exception/invalidParameterException"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Struct responsavel por analisar os dados ordenando por credito decrescente
// @see TableAnalysis
type LanguageSortAnalysis struct{}

// NewLanguageSortAnalysis cria uma nova instância de LanguageSortAnalysis
// @return Retorna um ponteiro para LanguageSortAnalysis
func NewLanguageSortAnalysis() *LanguageSortAnalysis {
	return &LanguageSortAnalysis{}
}

// Analysis realiza a ordenação dos dados por crédito decrescente
// @param userInfoList Lista de dados a ser analisada
// @return Lista ordenada por crédito decrescente e o erro ocorrido (nil em caso de sucesso)
// @see TableAnalysis
func (l *LanguageSortAnalysis) Analysis(userInfoList []userinfo.UserInfo) ([]userinfo.UserInfo, error) {
	if len(userInfoList) == 0 {
		return nil, invalidparameterexception.NewInvalidParameterException(
			"'userInfoList' é nil ou vazio", nil,
		)
	}

	sortedList := make([]userinfo.UserInfo, len(userInfoList))
	copy(sortedList, userInfoList)

	sort.SliceStable(sortedList, func(i, j int) bool {
		return sortedList[i].GetCredit() > sortedList[j].GetCredit()
	})

	return sortedList, nil
}