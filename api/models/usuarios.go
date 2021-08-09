package models

import "api/db"

type Usuario struct {
	Id     int
	Nome   string
	Nivel  string
	Idade  int
	Status string
}

func BuscaTodosUsuarios() []Usuario {

	db := db.ConectaComBancoDeDados()

	selectDeTodosOsUsuarios, err := db.Query("select * from usuarios order by id asc")
	if err != nil {
		panic(err.Error())
	}

	u := Usuario{}
	usuarios := []Usuario{}

	for selectDeTodosOsUsuarios.Next() {
		var id, idade int
		var nome, nivel, status string

		err = selectDeTodosOsUsuarios.Scan(&id, &nome, &nivel, &idade, &status)
		if err != nil {
			panic(err.Error())
		}

		u.Id = id
		u.Nome = nome
		u.Nivel = nivel
		u.Idade = idade
		u.Status = status

		usuarios = append(usuarios, u)
	}
	print(usuarios)
	defer db.Close()
	return usuarios
}

func CriaNovoUsuario(nome, nivel, status string, idadeConvertida int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into usuarios(nome, nivel, idade, status) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, nivel, idadeConvertida, status)
	defer db.Close()
}

func DeletaUsuario(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOUsuario, err := db.Prepare("delete from usuarios where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOUsuario.Exec(id)
	defer db.Close()
}

func EditaUsuario(id string) Usuario {
	db := db.ConectaComBancoDeDados()

	usuarioDoBanco, err := db.Query("select * from usuarios where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	usuarioParaAtualizar := Usuario{}

	for usuarioDoBanco.Next() {
		var id, idade int
		var nome, nivel, status string

		err = usuarioDoBanco.Scan(&id, &nome, &nivel, &idade, &status)
		if err != nil {
			panic(err.Error())
		}

		usuarioParaAtualizar.Id = id
		usuarioParaAtualizar.Nome = nome
		usuarioParaAtualizar.Nivel = nivel
		usuarioParaAtualizar.Idade = idade
		usuarioParaAtualizar.Status = status
	}
	defer db.Close()
	return usuarioParaAtualizar
}

func AtualizaUsuario(id int, nome string, nivel string, idade int, status string) {
	db := db.ConectaComBancoDeDados()

	AtualizaUsuario, err := db.Prepare("update usuarios set nome=$1, nivel=$2, idade=$3, status=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaUsuario.Exec(nome, nivel, idade, status, id)
	defer db.Close()
}
