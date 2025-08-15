package agreement

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_agreement.sql
var saveAgreementQuery string

//go:embed sql/find_agreement_by_id.sql
var findAgreementByIDQuery string

//go:embed sql/find_categories_by_agreement_id.sql
var findCategoriesByAgreementIDQuery string

//go:embed sql/find_salary_complements_by_id.sql
var findSalaryComplementsByIDQuery string

//go:embed sql/update_agreement.sql
var updateAgreementQuery string

//go:embed sql/save_category.sql
var saveCategoryQuery string

//go:embed sql/save_salary_complements.sql
var saveSalaryComplementsQuery string

//go:embed sql/delete_categories.sql
var deleteCategoriesQuery string

//go:embed sql/delete_salary_complements.sql
var deleteSalaryComplementsQuery string

//go:embed sql/find_agreement.sql
var findAgreementQuery string

//go:embed sql/find_category_by_id.sql
var findCategoryByIDQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveAgreementStmt, err := db.Prepare(saveAgreementQuery)
	if err != nil {
		return nil, err
	}
	findAgreementByIDStmt, err := db.Prepare(findAgreementByIDQuery)
	if err != nil {
		return nil, err
	}
	findCategoriesByAgreementIDStmt, err := db.Prepare(findCategoriesByAgreementIDQuery)
	if err != nil {
		return nil, err
	}
	findSalaryComplementByIDStmt, err := db.Prepare(findSalaryComplementsByIDQuery)
	if err != nil {
		return nil, err
	}
	updateAgreementStmt, err := db.Prepare(updateAgreementQuery)
	if err != nil {
		return nil, err
	}
	saveCategoryStmt, err := db.Prepare(saveCategoryQuery)
	if err != nil {
		return nil, err
	}
	saveSalaryComplementStmt, err := db.Prepare(saveSalaryComplementsQuery)
	if err != nil {
		return nil, err
	}
	deleteCategoriesStmt, err := db.Prepare(deleteCategoriesQuery)
	if err != nil {
		return nil, err
	}
	deleteSalaryComplementsStmt, err := db.Prepare(deleteSalaryComplementsQuery)
	if err != nil {
		return nil, err
	}
	findAgreementsStmt, err := db.Prepare(findAgreementQuery)
	if err != nil {
		return nil, err
	}
	findCategoryByIDStmt, err := db.Prepare(findCategoryByIDQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                          db,
		saveAgreement:               saveAgreementStmt,
		findAgreementByID:           findAgreementByIDStmt,
		findCategoriesByAgreementID: findCategoriesByAgreementIDStmt,
		findSalaryComplementByID:    findSalaryComplementByIDStmt,
		updateAgreement:             updateAgreementStmt,
		saveCategory:                saveCategoryStmt,
		saveSalaryComplements:       saveSalaryComplementStmt,
		deleteCategories:            deleteCategoriesStmt,
		deleteSalaryComplements:     deleteSalaryComplementsStmt,
		findAgreement:               findAgreementsStmt,
		findCategoryByID:            findCategoryByIDStmt,
	}, nil
}

type Repository struct {
	db                          *sql.DB
	saveAgreement               *sql.Stmt
	findAgreementByID           *sql.Stmt
	findCategoriesByAgreementID *sql.Stmt
	findSalaryComplementByID    *sql.Stmt
	updateAgreement             *sql.Stmt
	saveCategory                *sql.Stmt
	saveSalaryComplements       *sql.Stmt
	deleteCategories            *sql.Stmt
	deleteSalaryComplements     *sql.Stmt
	findAgreement               *sql.Stmt
	findCategoryByID            *sql.Stmt
}
