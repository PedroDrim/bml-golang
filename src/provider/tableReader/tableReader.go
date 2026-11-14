package tablereader

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/pedrodrim/bml-golang/src/model/exception/dataReaderException"
	"github.com/pedrodrim/bml-golang/src/model/exception/invalidParameterException"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Struct responsavel por ler e disponibilizar os dados em um formato desejado
// @see DataReader
type TableReader struct {
	fileName     string
	userInfoList []userinfo.UserInfo
}

// NewTableReader cria uma nova instância de TableReader lendo o arquivo informado
// @param fileName Nome do arquivo de dados a ser lido
// @return Retorna um ponteiro para TableReader e o erro ocorrido (nil em caso de sucesso)
// @throws DataReaderException Caso nao seja possivel ler os dados corretamente
func NewTableReader(fileName string) (*TableReader, error) {
	if fileName == "" {
		return nil, invalidparameterexception.NewInvalidParameterException("'fileName' é nil", nil)
	}

	list, err := deserializeFile(fileName)
	if err != nil {
		return nil, err
	}

	return &TableReader{
		fileName:     fileName,
		userInfoList: list,
	}, nil
}

// GetFileName obtem o nome do arquivo de dados a ser lido
// @return Nome do arquivo de dados a ser lido
func (t *TableReader) GetFileName() string {
	return t.fileName
}

// ReadAll obtem todos os dados disponiveis
// @return Lista contendo todos os dados disponiveis
func (t *TableReader) ReadAll() []userinfo.UserInfo {
	return t.userInfoList
}

// Read obtem os dados disponiveis dentro de um intervalo (inclusivo em ambos os extremos)
// @param startIndex Inicio do intervalo
// @param endIndex Fim do intervalo (inclusivo)
// @return Lista contendo os dados do intervalo e o erro ocorrido (nil em caso de sucesso)
func (t *TableReader) Read(startIndex int, endIndex int) ([]userinfo.UserInfo, error) {
	if startIndex < 0 {
		return nil, invalidparameterexception.NewInvalidParameterException("'startIndex' é menor que 0", nil)
	}
	if endIndex < 0 {
		return nil, invalidparameterexception.NewInvalidParameterException("'endIndex' é menor que 0", nil)
	}
	if startIndex >= endIndex {
		return nil, invalidparameterexception.NewInvalidParameterException("'startIndex' é maior ou igual á 'endIndex'", nil)
	}
	if endIndex >= len(t.userInfoList) {
		return nil, invalidparameterexception.NewInvalidParameterException("'endIndex' é maior que o tamanho da lista", nil)
	}

	// Ruby [start..end] é inclusivo; Go [start:end] é exclusivo → endIndex+1
	return t.userInfoList[startIndex : endIndex+1], nil
}

// deserializeFile desserializa o arquivo de dados em uma lista de UserInfo
// @param fileName Nome do arquivo de dados
// @return Lista contendo os dados desserializados e o erro ocorrido
// @throws DataReaderException Caso nao seja possivel ler os dados corretamente
func deserializeFile(fileName string) ([]userinfo.UserInfo, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, datareaderexception.NewDataReaderException("'fileName' é invalido", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Pula o cabeçalho (equivalente ao .drop(1) do Ruby)
	if !scanner.Scan() {
		return []userinfo.UserInfo{}, nil
	}

	list := make([]userinfo.UserInfo, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		values := strings.Split(line, ",")
		if len(values) < 3 {
			return nil, datareaderexception.NewDataReaderException("'fileName' é invalido", nil)
		}

		user := strings.TrimSpace(values[0])
		password := strings.TrimSpace(values[1])
		credit, err := strconv.ParseFloat(strings.TrimSpace(values[2]), 64)
		if err != nil {
			return nil, datareaderexception.NewDataReaderException("'fileName' é invalido", err)
		}

		list = append(list, *userinfo.NewUserInfo(user, password, credit))
	}

	if err := scanner.Err(); err != nil {
		return nil, datareaderexception.NewDataReaderException("'fileName' é invalido", err)
	}

	return list, nil
}