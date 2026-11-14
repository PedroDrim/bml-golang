package quicksortanalysis

import (
	"github.com/pedrodrim/bml-golang/src/model/exception/invalidParameterException"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Struct responsavel por analisar os dados ordenando por credito decrescente usando QuickSort
// @see TableAnalysis
type QuickSortAnalysis struct{}

// NewQuickSortAnalysis cria uma nova instância de QuickSortAnalysis
// @return Retorna um ponteiro para QuickSortAnalysis
func NewQuickSortAnalysis() *QuickSortAnalysis {
	return &QuickSortAnalysis{}
}

// Analysis realiza a ordenação QuickSort por crédito decrescente
// @param userInfoList Lista de dados a ser analisada
// @return Lista ordenada por crédito decrescente e o erro ocorrido (nil em caso de sucesso)
// @see TableAnalysis
func (q *QuickSortAnalysis) Analysis(userInfoList []userinfo.UserInfo) ([]userinfo.UserInfo, error) {
	if len(userInfoList) == 0 {
		return nil, invalidparameterexception.NewInvalidParameterException(
			"'userInfoList' é nil ou vazio", nil,
		)
	}

	return q.quickSort(userInfoList), nil
}

// quickSort inicia o QuickSort recursivo com particionamento em três vias
// @param array Lista a ser ordenada
// @return Lista ordenada em ordem decrescente
func (q *QuickSortAnalysis) quickSort(array []userinfo.UserInfo) []userinfo.UserInfo {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2
	pivot := array[mid]

	menores := make([]userinfo.UserInfo, 0)
	iguais := make([]userinfo.UserInfo, 0)
	maiores := make([]userinfo.UserInfo, 0)

	for _, value := range array {
		switch {
		case value.GetCredit() < pivot.GetCredit():
			menores = append(menores, value)
		case value.GetCredit() == pivot.GetCredit():
			iguais = append(iguais, value)
		default:
			maiores = append(maiores, value)
		}
	}

	// Ordem decrescente: maiores primeiro, depois iguais, depois menores
	result := make([]userinfo.UserInfo, 0, len(array))
	result = append(result, q.quickSort(maiores)...)
	result = append(result, iguais...)
	result = append(result, q.quickSort(menores)...)

	return result
}