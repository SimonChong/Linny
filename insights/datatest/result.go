package datatest

type MockDBResult struct {
	ExpectedInsertID     int64
	ExpectedRowsAffected int64
}

func (m *MockDBResult) LastInsertId() (int64, error) { return m.ExpectedInsertID, nil }
func (m *MockDBResult) RowsAffected() (int64, error) { return m.ExpectedRowsAffected, nil }
