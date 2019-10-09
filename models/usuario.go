package models

//Usuario - estrutura para armazenar as informações do usuário
type Usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
