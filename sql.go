package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func newSQLClient(configs sqlConfigs) (*sqlClient, error) {
	sqlConn, err := sql.Open("mysql", configs.String())
	if err != nil {
		return nil, fmt.Errorf("unable to open sql connection due: %v", err)
	}
	return &sqlClient{sqlConn: sqlConn}, nil
}

// sqlClient is used for interacting with MySQL
type sqlClient struct {
	sqlConn *sql.DB
}

func (c *sqlClient) sendUserID(data postUserID) (*postUserID, error) {
	query := "INSERT INTO usersInfo (user_id) VALUES (?)"
	if data.UserID <= 0 {
		return nil, fmt.Errorf("You must enter integer which is greater than 0")
	}
	_, err := c.sqlConn.Exec(query, data.UserID)
	if err != nil {
		return nil, fmt.Errorf("unable to execute query due : %v", err)
	}
	postUserID := postUserID{
		UserID: data.UserID,
	}
	return &postUserID, nil
}

func (c *sqlClient) removeUserID(data postUserID) (*postUserID, error) {
	query := "DELETE FROM usersInfo WHERE user_id = ?"
	if data.UserID <= 0 {
		return nil, fmt.Errorf("You must enter integer which is greater than 0")
	}
	_, err := c.sqlConn.Exec(query, data.UserID)
	if err != nil {
		return nil, fmt.Errorf("unable to execute query due : %v", err)
	}
	postUserID := postUserID{
		UserID: data.UserID,
	}
	return &postUserID, nil
}

func (c *sqlClient) getUserIDIfExists(data postUserID) (int, error) {
	query := "SELECT EXISTS (SELECT user_id FROM usersInfo WHERE user_id = ?)"
	row := c.sqlConn.QueryRow(query, data.UserID)
	var getUserID int
	err := row.Scan(&getUserID)
	if err != nil {
		return 0, fmt.Errorf("couldn't scan the rows due : %v", err)
	}
	return getUserID, nil
}

func (c *sqlClient) getAllUsersID(limit int) ([]postUserID, error) {
	query := "SELECT user_id FROM usersInfo LIMIT ?"
	rows, err := c.sqlConn.Query(query, limit)
	if err != nil {
		return nil, fmt.Errorf("unable to execute query due : %v", err)
	}
	defer rows.Close()
	var ids []postUserID
	for rows.Next() {
		var id postUserID
		err := rows.Scan(&id.UserID)
		if err != nil {
			continue
		}
		ids = append(ids, id)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("unable to iterate through rows due % v", err)
	}
	return ids, nil
}

func (c *sqlClient) getDocCount() (int, error) {
	query := "SELECT COUNT(*) FROM usersInfo"
	row := c.sqlConn.QueryRow(query)
	var getCount int
	err := row.Scan(&getCount)
	if err != nil {
		return 0, fmt.Errorf("couldn't scan the rows due : %v", err)
	}
	return getCount, nil
}

// sqlConfigs represent configs for sql connection
type sqlConfigs struct {
	Host                 string `json:"host"`
	Port                 int    `json:"port"`
	User                 string `json:"user"`
	Password             string `json:"password"`
	DBName               string `json:"db_name"`
	DialTimeoutInSeconds int    `json:"dial_timeout"`
}

func (c *sqlConfigs) String() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%d)/%v?timeout=%vs", c.User, c.Password, c.Host, c.Port, c.DBName, c.DialTimeoutInSeconds)
}
