package repository

import (
	"fmt"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/municipalities/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (repo *repository) GetMunicipality(id int) (*entities.Municipality, *utils.Error) {
	var entity entities.Municipality
	err := repo.database.QueryRow("SELECT * FROM MUNICIPALITY WHERE id_municipality = ?", id).
		Scan(&entity.Id, &entity.Name)
	if err != nil {
		return nil, utils.NewError(err, "error executing query")
	}

	if entity.IsEmpty() {
		return nil, utils.NewError(fmt.Errorf("municipality %d not found", id), utils.NotFound)
	}

	return &entity, nil
}

func (repo *repository) GetMunicipalities(size, page int) (*entities.MunicipalitiesPage, *utils.Error) {
	pages, count, err := repo.getPageData(size)
	if err != nil {
		return nil, err
	}

	municipalities, err := repo.getMunicipalities(size, page)
	if err != nil {
		return nil, err
	}

	return &entities.MunicipalitiesPage{
		Page: utils.Page{
			Size:  size,
			Page:  page,
			Pages: pages,
			Count: count,
		},
		Data: municipalities,
	}, nil
}

func (repo *repository) CreateMunicipality(municipality entities.MunicipalityPayload) (*entities.Municipality, *utils.Error) {
	stmt, err := repo.database.Prepare("INSERT INTO MUNICIPALITY(municipality_name) VALUES(?)")
	if err != nil {
		return nil, utils.NewError(err, "error preparing statement")
	}

	result, err := stmt.Exec(municipality.Name)
	if err != nil {
		return nil, utils.NewError(err, "error executing statement")
	}

	rawId, err := result.LastInsertId()
	if err != nil {
		return nil, utils.NewError(err, "error getting last insert rawId")
	}

	id := int(rawId)

	return &entities.Municipality{
		Id:   &id,
		Name: municipality.Name,
	}, nil
}

func (repo *repository) UpdateMunicipality(municipality entities.MunicipalityPayload) (*entities.Municipality, *utils.Error) {
	query := "UPDATE MUNICIPALITY SET municipality_name = ? WHERE id_municipality = ?"
	_, err := repo.database.Exec(query, municipality.Name, municipality.Id)
	if err != nil {
		return nil, utils.NewError(err, "error updating municipality")
	}

	return &entities.Municipality{
		Id:   municipality.Id,
		Name: municipality.Name,
	}, nil
}

func (repo *repository) DeleteMunicipality(id int) *utils.Error {
	stmt, err := repo.database.Prepare("DELETE FROM MUNICIPALITY WHERE id_municipality = ?")
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
	err := repo.database.QueryRow("SELECT COUNT(*) AS counter FROM MUNICIPALITY").Scan(&count)
	if err != nil {
		return 0, 0, utils.NewError(err, "error getting page count")
	}

	pages = count / size
	if count%size != 0 {
		pages++
	}

	return count, pages, nil
}

func (repo *repository) getMunicipalities(size, page int) (entities.Municipalities, *utils.Error) {
	query := "SELECT id_municipality, municipality_name FROM MUNICIPALITY LIMIT ? OFFSET ?"
	rows, err := repo.database.Query(query, size, (page-1)*size)
	if err != nil {
		return nil, utils.NewError(err, "error executing query")
	}
	defer rows.Close()

	municipalities := entities.Municipalities{}
	for rows.Next() {
		var municipality entities.Municipality
		err := rows.Scan(&municipality.Id, &municipality.Name)
		if err != nil {
			return nil, utils.NewError(err, "error scanning row")
		}
		municipalities = append(municipalities, municipality)
	}

	return municipalities, nil
}
