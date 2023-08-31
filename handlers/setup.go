package handlers

import (
	"log"
	"net/http"

	"github.com/guilhermewolke/fts-go-api/config"
)

func Setup(w http.ResponseWriter, r *http.Request) {
	db, err := config.DBConnect()

	if err != nil {
		panic(err)
	}

	defer db.Close()

	//Criação e preenchimento das tabelas
	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS ` + "`fts_bookshelf`.`author`" + ` (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(40) NOT NULL,
			nationality VARCHAR(40) NOT NULL
		)ENGINE=InnoDB;`)

	if err != nil {
		panic(err)
	}

	log.Println("Tabela author criada com sucesso!")

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS ` + "`fts_bookshelf`.`book`" + ` (
			id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
			author_id INT NOT NULL,
			title VARCHAR(100) NOT NULL
		)ENGINE=InnoDB;`)

	if err != nil {
		panic(err)
	}

	log.Println("Tabela book criada com sucesso!")

	_, err = db.Exec(
		`INSERT INTO ` + "`fts_bookshelf`.`author`" + ` (id, name, nationality) VALUES
			(1, "Fiodor Dostoievski", "Rússia"),
			(2, "Liev Tolstoi", "Rússia"),
			(3, "Homero", "Grécia"),
			(4, "Machado de Assis", "Brasil"),
			(5, "Eça de Queirós", "Portugal"),
			(6, "Rosana Rios", "Brasil");`)

	if err != nil {
		panic(err)
	}

	log.Println("Autores inseridos com sucesso!")

	_, err = db.Exec(
		`INSERT INTO ` + "`fts_bookshelf`.`book`" + ` (author_id, title) VALUES
			(1, "Humilhados e Ofendidos"),
			(1, "Noites Brancas"),
			(1, "Crime e Castigo"),
			(2, "A morte de Ivan Ilítch"),
			(2, "Guerra e Paz"),
			(2, "Anna Karenina"),
			(3, "Ilíadas"),
			(3, "Odisséia"),
			(4, "Helena"),
			(4, "Memórias Póstumas de Brás Cubas"),
			(4, "Dom Casmurro"),
			(4, "Quincas Borba"),
			(4, "Relíquias de Casa Velha"),
			(5, "A cidade e as serras"),
			(5, "A Relíquia"),
			(5, "A ilustre casa de Ramires"),
			(5, "O crime do Padre Amaro"),
			(5, "O primo Basílio"),
			(6, "Timóteo, o tatu poeta");`)

	if err != nil {
		panic(err)
	}

	log.Println("Livros inseridos com sucesso!")

	w.Write([]byte("Tabelas criadas e preenchidas com sucesso!"))
}
