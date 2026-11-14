package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pedrodrim/bml-golang/src/model/timeFormat"
	"github.com/pedrodrim/bml-golang/src/provider/benchmarkMeasure"
	"github.com/pedrodrim/bml-golang/src/provider/languageSortAnalysis"
	"github.com/pedrodrim/bml-golang/src/provider/mergeSortAnalysis"
	"github.com/pedrodrim/bml-golang/src/provider/quickSortAnalysis"
	"github.com/pedrodrim/bml-golang/src/provider/summaryAnalysis"
	"github.com/pedrodrim/bml-golang/src/provider/tableReader"
)

// Config representa o arquivo JSON de configuracao
type Config struct {
	InputFileNameList []string `json:"INPUT_FILENAME_LIST"`
	OutputFileName    string   `json:"OUTPUT_FILENAME"`
}

// Metodo de inicializacao do projeto
// @param args Lista de parametros obtidos via console
func main() {
	configFile := getParam(os.Args)
	properties, err := getConfig(configFile)
	if err != nil {
		fmt.Println("[ERROR] Falha ao ler configuracao:", err)
		os.Exit(1)
	}

	input := properties.InputFileNameList
	output := properties.OutputFileName

	benchmark := benchmarkmeasure.NewBenchmarkMeasure()

	summaryAnalysis := summaryanalysis.NewSummaryAnalysis()
	mergeSortAnalysis := mergesortanalysis.NewMergeSortAnalysis()
	quickSortAnalysis := quicksortanalysis.NewQuickSortAnalysis()
	languageSortAnalysis := languagesortanalysis.NewLanguageSortAnalysis()

	for index := 0; index < len(input); index++ {
		fmt.Printf("[START] Arquivo: %d\n", index)
		fileName := input[index]

		// ==================================================
		// Leitura dos dados
		fmt.Println("\t[LOG] Read")
		startBenchmark(benchmark, fmt.Sprintf("Read@%d", index))
		reader, err := tablereader.NewTableReader(fileName)
		if err != nil {
			fmt.Println("[ERROR] Falha ao ler arquivo:", err)
			os.Exit(1)
		}
		list := reader.ReadAll()
		endBenchmark(benchmark, fmt.Sprintf("Read@%d", index))

		// ==================================================
		// Analise dos dados (Summary)
		fmt.Println("\t[LOG] Summary")
		startBenchmark(benchmark, fmt.Sprintf("SummaryAnalysis@%d", index))
		if _, err := summaryAnalysis.Analysis(list); err != nil {
			fmt.Println("[ERROR] Falha no SummaryAnalysis:", err)
			os.Exit(1)
		}
		endBenchmark(benchmark, fmt.Sprintf("SummaryAnalysis@%d", index))

		// ==================================================
		// Analise dos dados (Merge)
		fmt.Println("\t[LOG] Merge")
		startBenchmark(benchmark, fmt.Sprintf("MergeAnalysis@%d", index))
		if _, err := mergeSortAnalysis.Analysis(list); err != nil {
			fmt.Println("[ERROR] Falha no MergeSortAnalysis:", err)
			os.Exit(1)
		}
		endBenchmark(benchmark, fmt.Sprintf("MergeAnalysis@%d", index))

		// ==================================================
		// Analise dos dados (Quick)
		fmt.Println("\t[LOG] Quick")
		startBenchmark(benchmark, fmt.Sprintf("QuickAnalysis@%d", index))
		if _, err := quickSortAnalysis.Analysis(list); err != nil {
			fmt.Println("[ERROR] Falha no QuickSortAnalysis:", err)
			os.Exit(1)
		}
		endBenchmark(benchmark, fmt.Sprintf("QuickAnalysis@%d", index))

		// ==================================================
		// Analise dos dados (Language)
		fmt.Println("\t[LOG] Language")
		startBenchmark(benchmark, fmt.Sprintf("LanguageAnalysis@%d", index))
		if _, err := languageSortAnalysis.Analysis(list); err != nil {
			fmt.Println("[ERROR] Falha no LanguageSortAnalysis:", err)
			os.Exit(1)
		}
		endBenchmark(benchmark, fmt.Sprintf("LanguageAnalysis@%d", index))

		// ==================================================
		fmt.Printf("[END] Arquivo: %d\n", index)
	}

	if err := benchmark.Export(output, timeformat.Milissegundos); err != nil {
		fmt.Println("[ERROR] Falha ao exportar benchmark:", err)
		os.Exit(1)
	}
}

// Metodo para captura e tratamento dos parametros obtidos via console
// @param codes Lista de parametros obtidos via console
// @return Caminho do arquivo de configuracao
func getParam(codes []string) string {
	if len(codes) != 2 {
		fmt.Println("Parametros inválidos.")
		os.Exit(-1)
	}

	return codes[1]
}

// getConfig le e desserializa o arquivo JSON de configuracao
// @param fileName Caminho do arquivo de configuracao
// @return Configuracao carregada e o erro ocorrido (nil em caso de sucesso)
func getConfig(fileName string) (*Config, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// startBenchmark marca o inicio de uma captura, registrando erro caso ocorra
// @param benchmark Instancia de BenchmarkMeasure
// @param tag Nome da captura
func startBenchmark(benchmark *benchmarkmeasure.BenchmarkMeasure, tag string) {
	if err := benchmark.Start(tag); err != nil {
		fmt.Println("[ERROR] Falha ao iniciar benchmark:", err)
		os.Exit(1)
	}
}

// endBenchmark marca o fim de uma captura, registrando erro caso ocorra
// @param benchmark Instancia de BenchmarkMeasure
// @param tag Nome da captura
func endBenchmark(benchmark *benchmarkmeasure.BenchmarkMeasure, tag string) {
	if err := benchmark.End(tag); err != nil {
		fmt.Println("[ERROR] Falha ao encerrar benchmark:", err)
		os.Exit(1)
	}
}