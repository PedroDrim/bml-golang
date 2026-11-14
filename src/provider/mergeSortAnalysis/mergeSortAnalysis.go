package mergesortanalysis

import (
	"github.com/pedrodrim/bml-golang/src/model/exception/invalidParameterException"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Struct responsavel por analisar os dados ordenando por credito decrescente usando MergeSort
// @see TableAnalysis
type MergeSortAnalysis struct{}

// NewMergeSortAnalysis cria uma nova instância de MergeSortAnalysis
// @return Retorna um ponteiro para MergeSortAnalysis
func NewMergeSortAnalysis() *MergeSortAnalysis {
	return &MergeSortAnalysis{}
}

// Analysis realiza a ordenação MergeSort por crédito decrescente
// @param userInfoList Lista de dados a ser analisada
// @return Lista ordenada por crédito decrescente e o erro ocorrido (nil em caso de sucesso)
// @see TableAnalysis
func (m *MergeSortAnalysis) Analysis(userInfoList []userinfo.UserInfo) ([]userinfo.UserInfo, error) {
	if len(userInfoList) == 0 {
		return nil, invalidparameterexception.NewInvalidParameterException(
			"'userInfoList' é nil ou vazio", nil,
		)
	}

	return m.mergeSort(userInfoList), nil
}

// mergeSort inicia o MergeSort recursivo
// @param array Lista a ser ordenada
// @return Lista ordenada
func (m *MergeSortAnalysis) mergeSort(array []userinfo.UserInfo) []userinfo.UserInfo {
	if len(array) <= 1 {
		return array
	}

	mid := len(array) / 2
	left := m.mergeSort(array[:mid])
	right := m.mergeSort(array[mid:])

	return m.merge(left, right)
}

// merge unifica os vetores ordenados
// @param left Vetor da esquerda
// @param right Vetor da direita
// @return Lista unificada em ordem decrescente
func (m *MergeSortAnalysis) merge(left, right []userinfo.UserInfo) []userinfo.UserInfo {
	response := make([]userinfo.UserInfo, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i].GetCredit() > right[j].GetCredit() {
			response = append(response, left[i])
			i++
		} else {
			response = append(response, right[j])
			j++
		}
	}

	response = append(response, left[i:]...)
	response = append(response, right[j:]...)

	return response
}