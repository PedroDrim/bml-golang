package main

import (
    "fmt"
    "os"
    "strconv"
    "time"
    "github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Método de inicialização do projeto
// @param args Lista de parametros obtidos via console
func main() {

    var tamanho int = getParam(os.Args)
    
    antes := time.Now().UnixNano()
    lista := make([]*userInfo.UserInfo, tamanho)
    
    for index, _ := range lista {
        id := strconv.Itoa(index)
        user := userInfo.NewUserInfo("user" + id, "password" + id)
        lista[index] = user
    }
    
    depois := time.Now().UnixNano()
    resultado := float64(depois - antes) / float64(time.Millisecond)
    
    fmt.Println("[START] Golang", tamanho)
    fmt.Println("[OK]Tamanho:", tamanho)
    fmt.Println("[OK]Tempo:", resultado, "ms")
    fmt.Println("[END] Golang", tamanho)
}

// Método para captura e tratamento dos parametros obtidos via console
// @param codes Lista de parametros obtidos via console
// @return Tamanho de usuários á serem gerados
func getParam(codes []string) int {
    if len(codes) != 2 {
        fmt.Println("Parametros inválidos.")
        os.Exit(-1)
    }
    
    var line int
    line, _ = strconv.Atoi(codes[1])

    if line <= 0 {
        fmt.Println("Quantidade de linhas menor que 1.")
        os.Exit(-1)
    }

    return line    
}
