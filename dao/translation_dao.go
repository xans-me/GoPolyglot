package dao

import (
	"database/sql"
)

// TranslationDAO handles database operations related to translations.
type TranslationDAO struct {
	db *sql.DB
}

// Translation represents the structure for a translation entry.
type Translation struct {
	RC          string // RC should be the Response Code Identifier of the Message
	Title       string
	Description string
}

// NewTranslationDAO creates a new instance of TranslationDAO.
func NewTranslationDAO(db *sql.DB) *TranslationDAO {
	return &TranslationDAO{db: db}
}

// GetTranslationCustomQuery retrieves a custom translation from the database based on the provided query and arguments.
func (dao *TranslationDAO) GetTranslationCustomQuery(query string, args ...interface{}) (*Translation, error) {
	var translation Translation
	err := dao.db.QueryRow(query, args...).Scan(&translation.RC, &translation.Title, &translation.Description)
	if err != nil {
		return nil, err
	}
	return &translation, nil
}
