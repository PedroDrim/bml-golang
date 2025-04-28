package table

import(
	"os"
	"bufio"
	"strings"
	"strconv"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Estrutura para gerenciar uma tabela de usuarios
type Table struct {
	UserInfoList []userinfo.UserInfo
}

// cria uma nova instância da tabela
// @param fileName Nome do arquivo .csv
// @return Retorna um ponteiro da tabela
func NewTable(fileName string) *Table {  
    return &Table{deserializeFile(fileName)}
}

// GetUserInfoList retorna a lista de usuarios
// @return lista de usuarios
func (u *Table) GetUserInfoList() []userinfo.UserInfo {
	return u.UserInfoList
}

// Método privado para conversão do arquivo .csv em uma lista de usuarios
// @param fileName Nome do arquivo
// @return Lista convertida de usuarios    
func deserializeFile(fileName string) []userinfo.UserInfo {
	file, _ := os.Open(fileName)
	scanner := bufio.NewScanner(file)


	defer file.Close()

	userInfoList := make([]userinfo.UserInfo, 0)

	// Eliminando primeira linha (cabecalho)
	var first bool = true

    for scanner.Scan() {
		if(!first) {
			line := scanner.Text()
			values := strings.Split(line, ",")

			credit, _ := strconv.ParseFloat(values[2], 64)
			user := userinfo.NewUserInfo(values[0], values[1], credit)
			userInfoList = append(userInfoList, *user)	
		}
		
		first = false
	}
	
	return userInfoList
}