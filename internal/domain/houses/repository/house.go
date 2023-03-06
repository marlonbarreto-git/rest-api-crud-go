package repository

import (
	"fmt"
	"github.com/marlonbarreto-git/rest-api-crud-go/internal/domain/houses/entities"
	"github.com/marlonbarreto-git/rest-api-crud-go/utils"
)

func (repo *repository) GetHouse(id string) (*entities.House, *utils.Error) {
	var entity entities.House
	err := repo.database.QueryRow("SELECT * FROM HOUSE WHERE id_cadastral = ?", id).
		Scan(&entity.Id, &entity.Address, &entity.IdOwner, &entity.IdMunicipality)
	if err != nil {
		return nil, utils.NewError(err, "error executing query")
	}

	if entity.IsEmpty() {
		return nil, utils.NewError(fmt.Errorf("House %d not found", id), utils.NotFound)
	}

	return &entity, nil
}

func (repo *repository) GetHouses(size, page int) (*entities.HousesPage, *utils.Error) {
	pages, count, err := repo.getPageData(size)
	if err != nil {
		return nil, err
	}

	houses, err := repo.getHouses(size, page)
	if err != nil {
		return nil, err
	}

	return &entities.HousesPage{
		Page: utils.Page{
			Size:  size,
			Page:  page,
			Pages: pages,
			Count: count,
		},
		Data: houses,
	}, nil
}

func (repo *repository) CreateHouse(house entities.HousePayload) (*entities.House, *utils.Error) {
	stmt, err := repo.database.Prepare("INSERT INTO HOUSE(id_cadastral,address,id_owner,id_municipality) VALUES(?,?,?,?)")
	if err != nil {
		return nil, utils.NewError(err, "error preparing statement")
	}

	result, err := stmt.Exec(house.Id, house.Address, house.IdOwner, house.IdMunicipality)
	if err != nil {
		return nil, utils.NewError(err, "error executing statement")
	}

	if rows, err := result.RowsAffected(); err != nil || rows < 1 {
		return nil, utils.NewError(err, "error validating affected rows")
	}

	return &entities.House{
		Id:             house.Id,
		Address:        house.Address,
		IdOwner:        house.IdOwner,
		IdMunicipality: house.IdMunicipality,
	}, nil
}

func (repo *repository) UpdateHouse(house entities.HousePayload) (*entities.House, *utils.Error) {
	query := "UPDATE HOUSE SET id_owner = ? WHERE id_cadastral = ?"
	_, err := repo.database.Exec(query, house.IdOwner, house.Id)
	if err != nil {
		return nil, utils.NewError(err, "error updating house")
	}

	return &entities.House{
		Id:             house.Id,
		Address:        house.Address,
		IdOwner:        house.IdOwner,
		IdMunicipality: house.IdMunicipality,
	}, nil
}

func (repo *repository) DeleteHouse(id string) *utils.Error {
	stmt, err := repo.database.Prepare("DELETE FROM HOUSE WHERE id_cadastral = ?")
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
	err := repo.database.QueryRow("SELECT COUNT(*) AS counter FROM HOUSE").Scan(&count)
	if err != nil {
		return 0, 0, utils.NewError(err, "error getting page count")
	}

	pages = count / size
	if count%size != 0 {
		pages++
	}

	return count, pages, nil
}

func (repo *repository) getHouses(size, page int) (entities.Houses, *utils.Error) {
	query := "SELECT id_cadastral, address, id_owner, id_municipality FROM HOUSE LIMIT ? OFFSET ?"
	rows, err := repo.database.Query(query, size, (page-1)*size)
	if err != nil {
		return nil, utils.NewError(err, "error executing query")
	}
	defer rows.Close()

	houses := entities.Houses{}
	for rows.Next() {
		var house entities.House
		err := rows.Scan(&house.Id, &house.Address, &house.IdOwner, &house.IdMunicipality)
		if err != nil {
			return nil, utils.NewError(err, "error scanning row")
		}
		houses = append(houses, house)
	}

	return houses, nil
}
