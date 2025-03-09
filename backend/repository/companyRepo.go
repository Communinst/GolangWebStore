package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	customErrors "github.com/Communinst/GolangWebStore/backend/errors"
	"github.com/jmoiron/sqlx"
)

type companyRepo struct {
	db *sqlx.DB
}

func NewCompanyRepo(db *sqlx.DB) *companyRepo {
	return &companyRepo{
		db: db,
	}
}

func (repo *companyRepo) PostCompany(ctx context.Context, company *entities.Company) (int, error) {
	var resultId int

	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return -1, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := fmt.Sprintf(`INSERT INTO %s (name)
		VALUES ($1) RETURNING company_id`, companiesTable)

	err = tx.QueryRowContext(ctx, query,
		company.Name).Scan(&resultId)

	if err != nil {
		tx.Rollback()
		return -1, err
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return -1, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	log.Printf("Company: %s posted successfully", company.Name)
	return resultId, err
}

func (repo *companyRepo) GetCompany(ctx context.Context, companyId int) (*entities.Company, error) {
	var resultCompany entities.Company

	query := fmt.Sprintf(`SELECT * FROM %s WHERE company_id = $1`, companiesTable)

	err := repo.db.GetContext(ctx, &resultCompany, query, companyId)
	if err == nil {
		log.Printf("Company by %d id was obtained", companyId)
		return &resultCompany, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return &resultCompany, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        fmt.Sprintf("Company with %d id wasn't found", companyId),
		}
	}

	slog.Error("unknown error obtaining company by id")
	return &resultCompany, &customErrors.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "unknown interanal server error occured",
	}
}

func (repo *companyRepo) GetAllCompanies(ctx context.Context) ([]entities.Company, error) {
	var resultCompanys []entities.Company

	query := fmt.Sprintf(`SELECT * FROM %s`, companiesTable)

	err := repo.db.SelectContext(ctx, &resultCompanys, query)
	if err == nil {
		log.Print("Companies were obtained")
		return resultCompanys, err
	}

	if errors.Is(err, sql.ErrNoRows) {
		return resultCompanys, &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "no company found",
		}
	}

	slog.Error("unknown error obtaining company by id")
	return resultCompanys, &customErrors.ErrorWithStatusCode{
		HTTPStatus: http.StatusInternalServerError,
		Msg:        "unknown interanal server error occured",
	}
}

func (repo *companyRepo) DeleteCompany(ctx context.Context, companyId int) error {
	tx, err := repo.db.BeginTx(ctx, nil)
	if err != nil {
		slog.Error("transaction initiation error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction initiation failed",
		}
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE company_id = $1`, companiesTable)

	result, err := tx.ExecContext(ctx, query, companyId)
	if err != nil {
		tx.Rollback()
		slog.Error(fmt.Sprintf("error deleting company by %d id", companyId), "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete company",
		}
	}

	affectedAmount, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		slog.Error("error getting amount of affected rows", "err", err.Error())
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "failed to delete company",
		}
	}

	if err = tx.Commit(); err != nil {
		slog.Error("transaction fulfillment error")
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        "transaction fulfillment failed",
		}
	}

	if affectedAmount == 0 {
		return &customErrors.ErrorWithStatusCode{
			HTTPStatus: http.StatusInternalServerError,
			Msg:        fmt.Sprintf("Company by %d id wasn't found", companyId),
		}
	}

	log.Printf("Company by %d id was deleted", companyId)
	return nil
}
