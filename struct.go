package main

type Exames struct {
	AccessionNumber    int    `json:"AccessionNumber"`
	PatientID          int    `json:"PatientId"`
	NomePaciente       string `json:"NomePaciente"`
	DataNascimento     string `json:"DataNascimento"`
	DataExame          string `json:"DataExame"`
	Exame              string `json:"Exame"`
	Solicitante        string `json:"Solicitante"`
	CrmSolicitante     string `json:"CrmSolicitante"`
	CodigoOficial      string `json:"CodigoOficial"`
	ModalidadeDicom    string `json:"ModalidadeDicom"`
	Sexo               string `json:"Sexo"`
	CodigoProcedimento string `json:"CodigoProcedimento"`
	CodigoMedico       int    `json:"CodigoMedico"`
	CodigoExame		   int    `json:"CodigoExame"`
}

type Laudos struct {
	AccessionNumber string `json:"AccessionNumber"`
	CrmMedico       *string `json:"crm_medico"`
	Laudo           string `json:"Laudo"`
}

type Laudo struct {
	CodigoLaudo int `json:"CodigoLaudo"`
}

/*type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}*/

type RetornoLaudo struct {
	Status    string `json:"Status"`
	Descricao string `json:"Descricao"`
}

type Retorno struct {
	AccessionNumber int    `json:"AccessionNumber"`
	Status          string `json:"Status"`
	Descricao       string `json:"Descricao"`
}

type Assinatura struct {
	Nome       string `json:"Nome"`
	CRM        string `json:"CRM"`
	Md5        string `json:"Md5Assinatura"`
	Assinatura string `json:"Assinatura"`
}
