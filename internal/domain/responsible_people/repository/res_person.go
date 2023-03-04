package repository

import (
	"fmt"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/responsible_people/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (repo *repository) GetResponsible(id int) (*entities.Responsible, *utils.Error) {
	query := "SELECT id_person FROM RESPONSIBLE_PERSON WHERE id_responsible = ?"

	rows, err := repo.database.Query(query, id)
	if err != nil {
		return nil, utils.NewError(err, "error executing query")
	}
	defer rows.Close()

	responsible := entities.Responsible{
		IdResponsible: &id,
		People:        []int{},
	}

	for rows.Next() {
		var personID int
		if err := rows.Scan(&personID); err != nil {
			return nil, utils.NewError(err, "error scanning row")
		}

		responsible.People = append(responsible.People, personID)
	}

	if len(responsible.People) == 0 {
		return nil, utils.NewError(fmt.Errorf("responsible %d not found", id), utils.NotFound)
	}

	return &responsible, nil
}

func (repo *repository) GetResponsibles(size, page int) (*entities.ResponsiblesPage, *utils.Error) {
	pages, count, err := repo.getPageData(size)
	if err != nil {
		return nil, err
	}

	responsibles, err := repo.getResponsibles(size, page)
	if err != nil {
		return nil, err
	}

	return &entities.ResponsiblesPage{
		Page: utils.Page{
			Size:  size,
			Page:  page,
			Pages: pages,
			Count: count,
		},
		Data: responsibles,
	}, nil
}

func (repo *repository) CreateResponsible(responsible entities.ResponsiblePayload) (*entities.Responsible, *utils.Error) {
	stmt, err := repo.database.Prepare("INSERT INTO RESPONSIBLE_PERSON(id_responsible,id_person) VALUES(?,?)")
	if err != nil {
		return nil, utils.NewError(err, "error preparing statement")
	}

	result, err := stmt.Exec(responsible.IdResponsible, responsible.IdPerson)
	if err != nil {
		return nil, utils.NewError(err, "error executing statement")
	}

	rawId, err := result.LastInsertId()
	if err != nil {
		return nil, utils.NewError(err, "error getting last insert rawId")
	}

	id := int(rawId)

	return &entities.Responsible{
		IdResponsible: &id,
		IdPerson:      responsible.IdPerson,
	}, nil
}

func (repo *repository) DeleteResponsible(id int) *utils.Error {
	stmt, err := repo.database.Prepare("DELETE FROM RESPONSIBLE_PERSON WHERE id_person = ?")
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
	err := repo.database.QueryRow("SELECT COUNT(*) AS counter FROM RESPONSIBLE_PERSON").Scan(&count)
	if err != nil {
		return 0, 0, utils.NewError(err, "error getting page count")
	}

	pages = count / size
	if count%size != 0 {
		pages++
	}

	return count, pages, nil
}

func (repo *repository) getResponsibles(size, page int) (entities.Responsibles, *utils.Error) {
	query := "SELECT id_responsible, id_person FROM RESPONSIBLE_PERSON LIMIT ? OFFSET ?"
	rows, err := repo.database.Query(query, size, (page-1)*size)
	if err != nil {
		return nil, utils.NewError(err, "error executing query")
	}
	defer rows.Close()

	responsibles := entities.Responsibles{}
	for rows.Next() {
		var responsible entities.Responsible
		err := rows.Scan(&responsible.IdResponsible, &responsible.IdPerson)
		if err != nil {
			return nil, utils.NewError(err, "error scanning row")
		}
		responsibles = append(responsibles, responsible)
	}

	return responsibles, nil
}
