package benchmarkmeasure

import (
	"encoding/json"
	"os"
	"time"

	"github.com/pedrodrim/bml-golang/src/model/exception/benchmarkException"
	"github.com/pedrodrim/bml-golang/src/model/exception/invalidParameterException"
	"github.com/pedrodrim/bml-golang/src/model/timeFormat"
)

// Classe para captura de estados de tempo
type BenchmarkMeasure struct {
	startMark string
	endMark   string
	benchMap  map[string]int64
}

// Funcao de criacao da classe
// @return Retorna um ponteiro para BenchmarkMeasure
func NewBenchmarkMeasure() *BenchmarkMeasure {
	return &BenchmarkMeasure{
		startMark: "_S",
		endMark:   "_E",
		benchMap:  make(map[string]int64),
	}
}

// Marca o início de uma captura
// @param tag tag de inicio
// @return Retorna InvalidParameterException caso 'tag' seja vazia
func (b *BenchmarkMeasure) Start(tag string) error {
	if tag == "" {
		return invalidparameterexception.NewInvalidParameterException("'tag' é nil", nil)
	}
	b.benchMap[tag + b.startMark] = time.Now().UnixMilli()
	return nil
}

// Marca o fim de uma captura
// @param tag tag de encerramento
// @return Retorna InvalidParameterException caso 'tag' seja vazia
func (b *BenchmarkMeasure) End(tag string) error {
	if tag == "" {
		return invalidparameterexception.NewInvalidParameterException("'tag' é nil", nil)
	}
	b.benchMap[tag + b.endMark] = time.Now().UnixMilli()
	return nil
}

// Resultado da captura de estado especifica
// @param tag Nome da captura referente
// @param format Formato do resultado
// @return Tempo decorrido entre o inicio e o fim da captura, ou erro
func (b *BenchmarkMeasure) ResultByTag(tag string, format timeformat.TimeFormat) (float64, error) {
	if tag == "" {
		return 0, invalidparameterexception.NewInvalidParameterException("'tag' é nil", nil)
	}
	if format == 0 {
		return 0, invalidparameterexception.NewInvalidParameterException("'format' é nil", nil)
	}

	startKey := tag + b.startMark
	endKey := tag + b.endMark

	startTime, startOk := b.benchMap[startKey]
	endTime, endOk := b.benchMap[endKey]

	if !startOk || !endOk {
		return 0, benchmarkexception.NewBenchmarkException(
			"Não encontrado par 'inicio-fim' de:"+tag, nil,
		)
	}

	return float64(endTime-startTime) * float64(format), nil
}

// Resultado de todas as capturas de estado
// @param format Formato do resultado
// @return Mapa contendo o tempo decorrido por tag, ou erro
func (b *BenchmarkMeasure) Result(format timeformat.TimeFormat) (map[string]float64, error) {
	if format == 0 {
		return nil, invalidparameterexception.NewInvalidParameterException("'format' é nil", nil)
	}

	result := make(map[string]float64)

	// Itera sobre as chaves de início (evita processar cada tag duas vezes)
	for key := range b.benchMap {
		if len(key) <= len(b.startMark) || key[len(key)-len(b.startMark):] != b.startMark {
			continue
		}
		tag := key[:len(key)-len(b.startMark)]

		time, err := b.ResultByTag(tag, format)
		if err != nil {
			return nil, err
		}
		result[tag] = time
	}

	return result, nil
}

// Exporta o resultado no formato de um arquivo
// @param fileName Nome do arquivo de saida
// @param format Formato do resultado
// @return Retorna erro caso a escrita falhe
func (b *BenchmarkMeasure) Export(fileName string, format timeformat.TimeFormat) error {
	if fileName == "" {
		return invalidparameterexception.NewInvalidParameterException("'fileName' é nil", nil)
	}
	if format == 0 {
		return invalidparameterexception.NewInvalidParameterException("'format' é nil", nil)
	}

	result, err := b.Result(format)
	if err != nil {
		return err
	}

	serialized, err := json.Marshal(result)
	if err != nil {
		return benchmarkexception.NewBenchmarkException("Falha ao serializar resultado", err)
	}

	if err := os.WriteFile(fileName, serialized, 0644); err != nil {
		return benchmarkexception.NewBenchmarkException("Falha ao escrever arquivo:"+fileName, err)
	}

	return nil
}