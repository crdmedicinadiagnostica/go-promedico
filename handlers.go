package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func createExames(w http.ResponseWriter, r *http.Request) {

	var post []Exames
	var retorno []Retorno
	var retu Retorno

	reqBody := json.NewDecoder(r.Body)

	err := reqBody.Decode(&post)
	fmt.Println(post)

	recebimento, err := json.Marshal(&post)
	if err != nil {
		fmt.Println("Error marshalling to Json:", err)
	}

	err = ioutil.WriteFile("recebimento.json", recebimento, 0644)
	if err != nil {
		fmt.Println("Error writing json to file:", err)
	}

	for _, pedidos := range post {
		cdIntegracao := gravarExame(pedidos, recebimento)

		retu.AccessionNumber = pedidos.AccessionNumber
		if cdIntegracao == 0 {
			retu.Descricao = "Erro ao Gravar"
			retu.Status = "Erro"
		} else {
			retu.Status = "Sucesso"
		}
		retorno = append(retorno, retu)
	}

	output, err := json.Marshal(&retorno)
	if err != nil {
		fmt.Println("Error marshalling to Json:", err)
	}

	err = ioutil.WriteFile("post-retorno.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing json to file:", err)

	}
	fmt.Fprint(w, string(output))
}

func getLaudos(w http.ResponseWriter, r *http.Request) {

	var laudos Laudos
	var i int
	var err error
	var output []byte

	params := mux.Vars(r)

	/** converting the params variable into an int using Atoi method */
	i1, err := strconv.Atoi(params["accessionNumber"])

	laudos, i, err = laudosExame(i1)

	fmt.Println(laudos)
	output, err = json.Marshal(&laudos)

	if err != nil {
		fmt.Println("Error marshalling to Json:", err)
	}

	err = ioutil.WriteFile("laudos-enviado.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing json to file:", err)

	}

	var retorno RetornoLaudo
	if err != nil || i == 0 {
		fmt.Println("Laudo não encontrado")
		retorno.Status = "Pendente"
		retorno.Descricao = "Laudo em Processo de Elaboracao"
		output, err = json.Marshal(&retorno)
	}

	fmt.Fprint(w, string(output))

}

func deleteExame(w http.ResponseWriter, r *http.Request) {
	var laudo Laudo
	var i int
	var err error
	var output []byte
	var retorno RetornoLaudo

	params := mux.Vars(r)

	/** converting the params variable into an int using Atoi method */
	i1, err := strconv.Atoi(params["accessionNumber"])

	laudo, i, err = verificaAssinado(i1)

	fmt.Println(laudo)
	output, err = json.Marshal(&laudo)

	if err != nil || i == 1 {
		fmt.Println("Exame laudado")
		retorno.Status = "Negado"
		retorno.Descricao = "Exame de codigo " + strconv.Itoa(i1) + " ja esta em processo de elaboracao de laudo"
		output, err = json.Marshal(&retorno)
	} else {

		laudo, err = deletaExames(i1)

		if err == nil {
			retorno.Status = "Cancelado"
			retorno.Descricao = "Exame de codigo " + strconv.Itoa(i1) + " excluido com sucesso"
			output, err = json.Marshal(&retorno)
		}

	}

	fmt.Fprint(w, string(output))

}

func getAssinaturas(w http.ResponseWriter, r *http.Request) {

	var assinatura Assinatura
	var i int
	var err error
	var output []byte

	params := mux.Vars(r)

	//fmt.Println(params)

	/** converting the params variable into an int using Atoi method */
	//i1, err := strconv.Atoi(params["crm"])

	assinatura, i, err = assinaturaMedicos(params["crm"])

	fmt.Println(assinatura)
	output, err = json.Marshal(&assinatura)

	if err != nil {
		fmt.Println("Error marshalling to Json:", err)
	}

	err = ioutil.WriteFile("assinatura-enviada.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing json to file:", err)

	}

	var retorno RetornoLaudo
	if err != nil || i == 0 {
		fmt.Println("Medico nao encontrado")
		retorno.Descricao = "Nenhum médico encontrado para o CRM informado"
		output, err = json.Marshal(&retorno)
	}

	fmt.Fprint(w, string(output))

}
