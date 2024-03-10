package GoPolyglot

// ColumnConfig defines the configuration for database columns related to translations.
type ColumnConfig struct {
	TableName         string // TableName is the name of the table in the database.
	RCColumn          string // RCColumn is the name of the column storing Response Code Identifiers.
	TitleColumn       string // TitleColumn is the name of the column storing translation titles.
	DescriptionColumn string // DescriptionColumn is the name of the column storing translation descriptions.
}

// QueryConfig defines the configuration for SQL queries used to fetch translations.
type QueryConfig struct {
	SelectClause string // SelectClause is the SQL SELECT clause for fetching translations.
	WhereClause  string // WhereClause is the SQL WHERE clause for filtering translations.
}

// Config holds the overall configuration for translation and dynamic translation handling.
type Config struct {
	ColumnConfig ColumnConfig // ColumnConfig defines the database column configuration.
	QueryConfig  QueryConfig  // QueryConfig defines the SQL query configuration for translations.
}
