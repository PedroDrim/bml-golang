package datareaderexception

// Struct responsavel por abstrair o erro de leitura de dados
type DataReaderException struct {
    Message string
    Cause error
}

// Construtor publico da classe
// @param message Mensagem de erro
// @param cause Causa do erro
// @return Retorna um ponteiro para o exception
func NewDataReaderException(message string, cause error) *DataReaderException {
    return &DataReaderException{
        Message: message,
        Cause: cause,
    }
}

// Error implementa a interface error, retornando a mensagem do exception
// @return Retorna a mensagem de erro
func (e *DataReaderException) Error() string {
    return e.Message
}

// Unwrap expoe a causa original para integracao com errors.Is e errors.As
// @return Retorna a causa do erro, ou nil caso nao exista
func (e *DataReaderException) Unwrap() error {
    return e.Cause
}