package manipulador

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ccasado/cursodego/bancodedados/model"
	"github.com/ccasado/cursodego/bancodedados/repo"
)

func Local(w http.ResponseWriter, r *http.Request) {
	local := model.Local{}
	codigoTelefone, err := strconv.Atoi(r.URL.Path[7:])
	if err != nil {
		http.Error(w, "Não foi enviado um número válido.", http.StatusBadRequest)
		fmt.Println("[local] erro ao converter o numero enviado: ", err.Error())
	}
	sql := "select country, city, telcode from place where telcode = ?"
	linha, err := repo.Db.Queryx(sql, codigoTelefone)
	if err != nil {
		http.Error(w, "Não foi possível executar a query: ", http.StatusInternalServerError)
		fmt.Println("[local] erro ao converter o numero enviado: ", err.Error())
		return
	}
	for linha.Next() {
		err = linha.StructScan(&local)
		if err != nil {
			http.Error(w, "Não foi possível pesquisar esse número.", http.StatusInternalServerError)
			fmt.Println("[local] não foi possível fazer o binding dos dados do banco na struct", err.Error())
			return
		}

	}

	if err := ModeloLocal.ExecuteTemplate(w, "local.html", local); err != nil {
		http.Error(w, "Houve um erro na render da página.", http.StatusInternalServerError)
		fmt.Println("[local] Erro na execução do request")
	}
	sql = "insert into logquery (datarequest) values (?)"
	resultado, err := repo.Db.Exec(sql, time.Now().Format("02/01/2006 15:04:05"))
	if err != nil {
		fmt.Println("[local] Erro na inclusao do log: ", sql, " - ", err.Error())
	}
	linhasAfetadas, err := resultado.RowsAffected()
	if err != nil {
		fmt.Println("[local] Erro ao pegar o numero de linhas afetadas pelo comando: ", sql, " - ", err.Error())
	}
	fmt.Println("Sucesso! ", linhasAfetadas, " linha(s) afetada(s)")

}
