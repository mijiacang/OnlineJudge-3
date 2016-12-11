package models

import (
	"github.com/jmoiron/sqlx"
)

type TestCase struct {
	CaseId    int64 `db:"case_id"`
	Input     []byte
	InputMD5  []byte `db:"input_md5"`
	Output    []byte
	OutputMD5 []byte `db:"output_md5"`

	LocalPidFK int64 `db:"local_pid_fk"`
}

type TestCaseModel struct {
	Model
}

func NewTestCaseModel() *TestCaseModel {
	return &TestCaseModel{Model{Table: "TestCases"}}
}

func (this *TestCaseModel) Insert(tx *sqlx.Tx, tc *TestCase) (int64, error) {
	// hash input and output
	last_insert_id, err := this.InlineInsert(tx, tc, nil, []string{"case_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}
