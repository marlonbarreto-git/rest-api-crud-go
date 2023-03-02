package repository

import (
	"fmt"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (repo *repository) GetPerson(id int) (*entities.Person, *utils.Error) {
	var entity entities.Person
	err := repo.database.QueryRow("SELECT * FROM PERSON WHERE id_number = ?", id).
		Scan(&entity.Id, &entity.Forename, &entity.Surname, &entity.BirthDate, &entity.Sex, &entity.IdHome)
	if err != nil {
		return nil, utils.NewError(err, "error executing query")
	}

	if entity.IsEmpty() {
		return nil, utils.NewError(fmt.Errorf("PERSON %d not found", id), utils.NotFound)
	}

	return &entity, nil
}

func (repo *repository) GetPeople(size, page int) (*entities.PeoplePage, *utils.Error) {
	pages, count, err := repo.getPageData(size)
	if err != nil {
		return nil, err
	}

	people, err := repo.getPeople(size, page)
	if err != nil {
		return nil, err
	}

	return &entities.PeoplePage{
		Page: utils.Page{
			Size:  size,
			Page:  page,
			Pages: pages,
			Count: count,
		},
		Data: people,
	}, nil
}

func (repo *repository) CreatePerson(person entities.PersonPayload) (*entities.Person, *utils.Error) {
	stmt, err := repo.database.Prepare("INSERT INTO PERSON(id_number,forename,surname,birth_date,sex,id_home) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return nil, utils.NewError(err, "error preparing statement")
	}

	result, err := stmt.Exec(person.Id, person.Forename, person.Surname, person.BirthDate, person.Sex, person.IdHome)
	if err != nil {
		return nil, utils.NewError(err, "error executing statement")
	}

	rawId, err := result.LastInsertId()
	if err != nil {
		return nil, utils.NewError(err, "error getting last insert rawId")
	}

	id := int(rawId)

	return &entities.Person{
		Id:        &id,
		Forename:  person.Forename,
		Surname:   person.Surname,
		BirthDate: person.BirthDate,
		Sex:       person.Sex,
		IdHome:    person.IdHome,
	}, nil
}

func (repo *repository) UpdatePerson(person entities.PersonPayload) (*entities.Person, *utils.Error) {
	query := "UPDATE PERSON SET forename=?, surname=?, id_home=? WHERE id_number = ?"
	_, err := repo.database.Exec(query, person.Forename, person.Surname, person.IdHome, person.Id)
	if err != nil {
		return nil, utils.NewError(err, "error updating person")
	}

	return &entities.Person{
		Id:        person.Id,
		Forename:  person.Forename,
		Surname:   person.Surname,
		BirthDate: person.BirthDate,
		Sex:       person.Sex,
		IdHome:    person.IdHome,
	}, nil
}

func (repo *repository) DeletePerson(id int) *utils.Error {
	stmt, err := repo.database.Prepare("DELETE FROM PERSON WHERE id_number = ?")
	if err != nil {
		return utils.NewError(err, "error preparing delete statement")
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return utils.NewError(err, "error executing delete statement")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return utils.NewError(err, "error getting rows affected")
	}

	if rowsAffected == 0 {
		return utils.NewNotFoundError("entity not found")
	}

	return nil
}

func (repo *repository) getPageData(size int) (count int, pages int, error *utils.Error) {
	err := repo.database.QueryRow("SELECT COUNT(*) AS counter FROM PERSON").Scan(&count)
	if err != nil {
		return 0, 0, utils.NewError(err, "error getting page count")
	}

	pages = count / size
	if count%size != 0 {
		pages++
	}

	return count, pages, nil
}

func (repo *repository) getPeople(size, page int) (entities.People, *utils.Error) {
	query := "SELECT id_number,forename,surname,birth_date,sex,id_home FROM PERSON LIMIT ? OFFSET ?"
	rows, err := repo.database.Query(query, size, (page-1)*size)
	if err != nil {
		return nil, utils.NewError(err, "error executing query")
	}
	defer rows.Close()

	people := entities.People{}
	for rows.Next() {
		var person entities.Person
		err := rows.Scan(&person.Id, &person.Forename, &person.Surname, &person.BirthDate, &person.Sex, &person.IdHome)
		if err != nil {
			return nil, utils.NewError(err, "error scanning row")
		}
		people = append(people, person)
	}

	return people, nil
}
