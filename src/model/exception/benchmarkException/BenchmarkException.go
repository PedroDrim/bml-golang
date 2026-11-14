package benchmarkexception

// Struct responsavel por abstrair o erro de benchmark
type BenchmarkException struct {
    Message string
    Cause error
}

// Construtor publico da classe
// @param message Mensagem de erro
// @param cause Causa do erro
// @return Retorna um ponteiro para o exception
func NewBenchmarkException(message string, cause error) *BenchmarkException {
    return &BenchmarkException{
        Message: message,
        Cause: cause,
    }
}

// Error implementa a interface error, retornando a mensagem do exception
// @return Retorna a mensagem de erro
func (e *BenchmarkException) Error() string {
    return e.Message
}

// Unwrap expoe a causa original para integracao com errors.Is e errors.As
// @return Retorna a causa do erro, ou nil caso nao exista
func (e *BenchmarkException) Unwrap() error {
    return e.Cause
}