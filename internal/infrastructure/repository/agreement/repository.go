package agreement

import (
	"database/sql"

	_ "embed"
)

//go:embed sql/save_agreement.sql
var saveAgreementQuery string

//go:embed sql/find_agreement_by_id.sql
var findAgreementByIDQuery string

//go:embed sql/find_categories_by_id.sql
var findCategoriesByIDQuery string

//go:embed sql/find_complements_by_id.sql
var findComplementsByIDQuery string

//go:embed sql/update_agreement.sql
var updateAgreementQuery string

//go:embed sql/save_categories.sql
var saveCategoriesQuery string

//go:embed sql/save_complements.sql
var saveComplementsQuery string

//go:embed sql/delete_categories.sql
var deleteCategoriesQuery string

//go:embed sql/delete_complements.sql
var deleteComplementsQuery string

func NewRepository(db *sql.DB) (*Repository, error) {
	saveAgreementStmt, err := db.Prepare(saveAgreementQuery)
	if err != nil {
		return nil, err
	}
	findAgreementByIDStmt, err := db.Prepare(findAgreementByIDQuery)
	if err != nil {
		return nil, err
	}
	findCategoriesByIDStmt, err := db.Prepare(findCategoriesByIDQuery)
	if err != nil {
		return nil, err
	}
	findComplementByIDStmt, err := db.Prepare(findComplementsByIDQuery)
	if err != nil {
		return nil, err
	}
	updateAgreementStmt, err := db.Prepare(updateAgreementQuery)
	if err != nil {
		return nil, err
	}
	saveCategoriesStmt, err := db.Prepare(saveCategoriesQuery)
	if err != nil {
		return nil, err
	}
	saveComplementStmt, err := db.Prepare(saveComplementsQuery)
	if err != nil {
		return nil, err
	}
	deleteCategoriesStmt, err := db.Prepare(deleteCategoriesQuery)
	if err != nil {
		return nil, err
	}
	deleteComplementsStmt, err := db.Prepare(deleteComplementsQuery)
	if err != nil {
		return nil, err
	}

	return &Repository{
		db:                 db,
		saveAgreement:      saveAgreementStmt,
		findAgreementByID:  findAgreementByIDStmt,
		findCategoriesByID: findCategoriesByIDStmt,
		findComplementByID: findComplementByIDStmt,
		updateAgreement:    updateAgreementStmt,
		saveCategories:     saveCategoriesStmt,
		saveComplements:    saveComplementStmt,
		deleteCategories:   deleteCategoriesStmt,
		deleteComplements:  deleteComplementsStmt,
	}, nil
}

type Repository struct {
	db                 *sql.DB
	saveAgreement      *sql.Stmt
	findAgreementByID  *sql.Stmt
	findCategoriesByID *sql.Stmt
	findComplementByID *sql.Stmt
	updateAgreement    *sql.Stmt
	saveCategories     *sql.Stmt
	saveComplements    *sql.Stmt
	deleteCategories   *sql.Stmt
	deleteComplements  *sql.Stmt
}
