package datareader

import (
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Interface para analise de tabela simples
type DataReader interface {

	// Abre o arquivo de dados
	// @return Retorna nil em caso de sucesso ou o erro ocorrido
	Open() error

	// Obtem todos os dados disponiveis
	// @return Lista contendo todos os dados disponiveis e o erro ocorrido (nil em caso de sucesso)
	ReadAll() ([]userinfo.UserInfo, error)

	// Obtem os dados disponiveis dentro de um intervalo
	// @param startIndex Inicio do intervalo (inclusivo)
	// @param endIndex Fim do intervalo (exclusivo)
	// @return Lista contendo os dados do intervalo e o erro ocorrido (nil em caso de sucesso)
	Read(startIndex int, endIndex int) ([]userinfo.UserInfo, error)
}