package main

import (
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/gocraft/dbr"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	db *sql.DB //database
)

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	fmt.Println(dbURI)

	conn, err := sql.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
}

//returns a handle to the DB object
func GetDB() *sql.DB {
	return db
}

func gravarExame(post Exames) int {
	var CdIntegracao int

	temp := strings.Split(post.DataExame, "/")
	post.DataExame = temp[2] + temp[1] + temp[0]
	temp = strings.Split(post.DataNascimento, "/")
	post.DataNascimento = temp[2] + temp[1] + temp[0]
	fmt.Println("Atendimento, Empresa", post.AccessionNumber, post.Exame)
	err := GetDB().QueryRow(`insert into integracao_promedico(
                         accession_number,
                         cd_paciente,
                         ds_paciente,
                         dt_nascimento,
						 dt_exame,
                         ds_solicitante,
                         ds_crm_solicitante,
						 cd_oficial,
						 ds_exame,
						 ds_modalidade,
						 ds_sexo,
						 cd_medico,
						 cd_procedimento,
						 codigoexame,
						 created_at
						 ) 
                         values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,now()) returning cd_integracao`,

		post.AccessionNumber,
		post.PatientID,
		post.NomePaciente,
		post.DataNascimento,
		post.DataExame,
		post.Solicitante,
		post.CrmSolicitante,
		post.CodigoOficial,
		post.Exame,
		post.ModalidadeDicom,
		post.Sexo,
		post.CodigoMedico,
		post.CodigoProcedimento).Scan(&CdIntegracao)

	fmt.Println(err)
	if err != nil {
		return 0
		//panic(err)

	}
	return CdIntegracao

}

func laudosExame(busca int) (Laudos, int, error) {
	var post Laudos

	i := 0
	rows, err := GetDB().Query(`select 
                              ip.accession_number as accession_number, 
							  case when le.ds_crm_nr is null and le.ds_crm_uf is null then me.ds_crm_nr || me.ds_crm_uf
							  else le.ds_crm_nr || le.ds_crm_uf end as crm_medico,
                              encode(la.bb_laudo, 'base64') as Laudo
							  from integracao_promedico ip
							  join laudos la on (ip.accession_number = la.cd_laudo)
							  left join laudos_externo le on ip.accession_number = le.nr_controle
							  left join atendimentos ae on ip.accession_number = ae.cd_atendimento
							  left join medicos me on ae.cd_medico = me.cd_medico
							  where ip.accession_number = $1 limit 1`, busca)

	defer rows.Close()

	if err != nil {
		fmt.Println("Erro SQL")
		return post, i, err
	}

	i, err = dbr.Load(rows, &post)
	return post, i, err
}

func assinaturaMedicos(busca string) (Assinatura, int, error) {
	var post Assinatura

	crm := busca[:len(busca)-2]
	uf := busca[len(busca)-2:]

	i := 0
	rows, err := GetDB().Query(`select ds_medico                       as Nome,
								ds_crm_nr || ds_crm_uf                 as CRM,
	                            md5(bb_assinatura_imagem)::text        as Md5,
	                            encode(bb_assinatura_imagem, 'base64') as Assinatura
								from medicos where sn_ativo is true
								and bb_assinatura_imagem is not null
								and ds_crm_nr = $1 and ds_crm_uf = $2 limit 1`, crm, uf)

	defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return post, i, err
	}

	i, err = dbr.Load(rows, &post)
	return post, i, err
}

func verificaAssinado(busca int) (Laudo, int, error) {
	var post Laudo

	i := 0
	rows, err := GetDB().Query(`select cd_laudo
                              from laudos
							  where cd_laudo = $1 limit 1`, busca)

	defer rows.Close()

	if err != nil {
		fmt.Println("Erro SQL")
		return post, i, err
	}

	i, err = dbr.Load(rows, &post)
	return post, i, err
}

func deletaExames(busca int) (Laudo, error) {
	var post Laudo

	_, err := GetDB().Exec(`delete from exames where cd_exame = $1`, busca)

	if err == nil {
		_, err := GetDB().Exec(`delete from atendimentos where cd_atendimento = $1`, busca)
		if err == nil {
			_, err := GetDB().Exec(`update integracao_promedico set sn_cancelado = true where accession_number = $1`, busca)
			return post, err
		}
	}

	return post, err
}
