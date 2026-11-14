package summaryanalysis

import (
	"github.com/pedrodrim/bml-golang/src/model/exception/invalidParameterException"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// SummaryResult representa os valores maximo, minimo e media dos creditos
type SummaryResult struct {
	Min  float64
	Max  float64
	Mean float64
}

// Struct responsavel por analisar os dados obtendo valores maximo, minimo e media dos creditos
// @see TableAnalysis
type SummaryAnalysis struct{}

// NewSummaryAnalysis cria uma nova instância de SummaryAnalysis
// @return Retorna um ponteiro para SummaryAnalysis
func NewSummaryAnalysis() *SummaryAnalysis {
	return &SummaryAnalysis{}
}

// Analysis realiza a analise maxima, minima e media dos dados
// @param userInfoList Lista de dados a ser analisada
// @return Valores maximo, minimo e media dos creditos e o erro ocorrido (nil em caso de sucesso)
// @see TableAnalysis
func (s *SummaryAnalysis) Analysis(userInfoList []userinfo.UserInfo) (*SummaryResult, error) {
	if len(userInfoList) == 0 {
		return nil, invalidparameterexception.NewInvalidParameterException(
			"'userInfoList' é nil ou vazio", nil,
		)
	}

	max := -1_000_000_000.0
	min := 1_000_000_000.0
	sum := 0.0

	for _, userInfo := range userInfoList {
		credit := userInfo.GetCredit()
		sum += credit
		if credit > max {
			max = credit
		}
		if credit < min {
			min = credit
		}
	}

	mean := sum / float64(len(userInfoList))

	return &SummaryResult{
		Min:  min,
		Max:  max,
		Mean: mean,
	}, nil
}