package GoPolyglot

import (
	"database/sql"
	"fmt"
	"github.com/xans-me/GoPolyglot/dao"
)

// Translator provides functionalities for message translations.
type Translator struct {
	dao    *dao.TranslationDAO
	config Config
}

// NewTranslator creates a new Translator with a database connection and configuration.
func NewTranslator(db *sql.DB, config Config) *Translator {
	translationDAO := dao.NewTranslationDAO(db)
	return &Translator{dao: translationDAO, config: config}
}

// TranslateWithParams retrieves a translation based on the provided parameters and language.
func (t *Translator) TranslateWithParams(rc, trxType, trxFeature, language string) (*dao.Translation, error) {
	colConfig := t.config.ColumnConfig
	queryConfig := t.config.QueryConfig

	titleColumn := colConfig.TitleColumn + "_" + language
	descriptionColumn := colConfig.DescriptionColumn + "_" + language

	fullQuery := fmt.Sprintf("%s %s",
		fmt.Sprintf(queryConfig.SelectClause, colConfig.RCColumn, titleColumn, descriptionColumn, colConfig.TableName), queryConfig.WhereClause)

	translation, err := t.dao.GetTranslationCustomQuery(fullQuery, rc, trxType, trxFeature)
	if err != nil {
		return nil, err
	}

	return translation, nil
}
