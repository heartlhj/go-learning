package db

func CreateByteArry(name string, bytes string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO bytearry (name , bytes) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(name, bytes)
	if err != nil {
		return err
	}
	stmtIns.Close()
	return nil
}
