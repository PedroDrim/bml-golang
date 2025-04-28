package main

import (
    "fmt"
    "os"
    "time"
	"github.com/pedrodrim/bml-golang/src/model/simpleTableAnalysis"
    "github.com/pedrodrim/bml-golang/src/model/table"
    "github.com/pedrodrim/bml-golang/src/provider/maxValueAnalysis"
    "github.com/pedrodrim/bml-golang/src/provider/minValueAnalysis"
    "github.com/pedrodrim/bml-golang/src/provider/meanAnalysis"
)

// Método de inicialização do projeto
// @param args Lista de parametros obtidos via console
func main() {

    fileName := getParam(os.Args)
   
    // Obtendo o tempo inicial de leitura em milissegundos
    leitura_inicio := time.Now().UnixNano()

    // Convertendo arquivo em lista de "UserInfo"
    userTable := table.NewTable(fileName)

    // Obtendo o tempo final de leitura em milissegundos
    leitura_fim := time.Now().UnixNano()

    userInfoList := userTable.GetUserInfoList()

    var maxValue simpletableanalysis.SimpleTableAnalysis = maxvalueanalysis.NewMaxValueAnalysis()
    var minValue simpletableanalysis.SimpleTableAnalysis = minvalueanalysis.NewMinValueAnalysis()
    var meanValue simpletableanalysis.SimpleTableAnalysis = meananalysis.NewMeanAnalysis()

    // Obtendo o tempo inicial de analise em milissegundos
    analise_inicio := time.Now().UnixNano()

    // Realizando analises
    vmax := maxValue.Analysis(userInfoList)
    vmin := minValue.Analysis(userInfoList)
    vmean := meanValue.Analysis(userInfoList)

    // Obtendo o tempo final de analise em milissegundos
    analise_fim := time.Now().UnixNano()

    resultado_leitura := float64(leitura_fim - leitura_inicio) / float64(time.Millisecond)
    resultado_analise := float64(analise_fim - analise_inicio) / float64(time.Millisecond)

    // Dados de saida
    fmt.Println("[START] Golang_", fileName)
    fmt.Println("[OK]Arquivo:", fileName)
    fmt.Println("[OK]TempoLeitura:", resultado_leitura, "ms")
    fmt.Println("[OK]TempoAnalise:", resultado_analise, "ms")
    fmt.Println("[OK]Max:", vmax)
    fmt.Println("[OK]Min:", vmin)
    fmt.Println("[OK]Mean:", vmean)
    fmt.Println("[END] Golang_", fileName)
}

// Método para captura e tratamento dos parametros obtidos via console
// @param codes Lista de parametros obtidos via console
// @return Tamanho de usuários á serem gerados
func getParam(codes []string) string {
    if len(codes) != 2 {
        fmt.Println("Parametros inválidos.")
        os.Exit(-1)
    }
    
    var line = codes[1]

    return line    
}
