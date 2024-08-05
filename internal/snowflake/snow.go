package snowflake

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"

	_ "github.com/snowflakedb/gosnowflake"
)

type SnowConfig struct {
	User      string `json:"user"`
	Password  string `json:"password"`
	Account   string `json:"account"`
	Warehouse string `json:"warehouse"`
	Database  string `json:"database"`
	Schema    string `json:"schema"`
	Role      string `json:"role"`
}

func (sn *SnowConfig) Read() {
	// Read the JSON file
	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read file contents
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Unmarshal JSON data into SnowflakeConfig struct
	err = json.Unmarshal(data, &sn)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
}

func createConnectionString(cfg SnowConfig) string {
	//dsn := "user:password@account/warehouse/db/schema?role=role"
	//"user:password@account/db/schema?warehouse=?&role=?"
	//dsn := "%s:%s@%s/%s/%s/%s?role=%s"
	dsn := "%s:%s@%s/%s/%s?warehouse=%s&role=%s"
	final := fmt.Sprintf(dsn, cfg.User, url.QueryEscape(cfg.Password), cfg.Account, cfg.Database, cfg.Schema, cfg.Warehouse, cfg.Role)
	return final
}

// Create a snowflake connection
//
// make sure to close it.
func SnowflakeConn() (*sql.DB, error) {
	var cfg SnowConfig
	cfg.Read()

	if cfg.User == "" || cfg.Account == "" || cfg.Password == "" || cfg.Database == "" || cfg.Schema == "" || cfg.Warehouse == "" || cfg.Role == "" {
		return nil, fmt.Errorf("missing required configuration fields")
	}
	dsn := createConnectionString(cfg)
	return sql.Open("snowflake", dsn)
}

// run a sql statement without returning anything
func RunSQL(db *sql.DB, sql string, args ...any) error {
	_, err := db.Exec(sql, args...)
	return err
}
