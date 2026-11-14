package benchmarkoutput

import (
	"github.com/pedrodrim/bml-golang/src/model/timeFormat"
)

// BenchmarkOutput é a interface para captura de estados de tempo (benchmark)
type BenchmarkOutput interface {

	// Start marca o início da captura de estado
	// @param tag Nome da captura referente
	// @return Retorna erro caso a tag seja inválida
	Start(tag string) error

	// End marca o fim da captura de estado
	// @param tag Nome da captura referente
	// @return Retorna erro caso a tag seja inválida
	End(tag string) error

	// ResultByTag retorna o resultado da captura de estado especifica
	// @param tag Nome da captura referente
	// @param format Formato do resultado
	// @return Tempo decorrido entre o inicio e o fim da captura e o erro ocorrido
	ResultByTag(tag string, format timeformat.TimeFormat) (float64, error)

	// Result retorna o resultado de todas as capturas de estado
	// @param format Formato do resultado
	// @return Mapa contendo o tempo decorrido por tag e o erro ocorrido
	Result(format timeformat.TimeFormat) (map[string]float64, error)

	// Export exporta o resultado para um arquivo
	// @param fileName Nome do arquivo de saida
	// @param format Formato do resultado
	// @return Retorna erro caso a escrita falhe
	Export(fileName string, format timeformat.TimeFormat) error
}