package models

import (
	"errors"
	"time"
)

type Submission struct {
	RunId           int `db:"run_id"`
	Status          string
	StatusCode      string    `db:"status_code"`
	TestCasesPassed int       `db:"testcases_passed"`
	SubmitTime      time.Time `db:"submit_time"`
	TimeUsed        int       `db:"time_used"`
	MemoryUsed      int       `db:"memory_used"`
	Code            string
	LangIdFK        int    `db:"lang_id_fk"`
	CEInfo          string `db:"ce_info"`
	IPAddr          string `db:"ip_addr"`
	IsShared        bool   `db:"is_shared"`

	IsContest bool `db:"is_contest"`
	CPIdFK    int  `db:"cp_id_fk"`
	CUIdFK    int  `db:"cu_id_fk"`
	MetaPidFK int  `db:"meta_pid_fk"`
	UserIdFK  int  `db:"user_id_fk"`
}

type SubmissionModel struct {
	Model
}

func NewSubmissionModel() *SubmissionModel {
	return &SubmissionModel{Model{Table: "Submissions"}}
}

func (this *SubmissionModel) Validate(sub *Submission) error {
	if sub.IsContest == true {
		if sub.CPIdFK == 0 {
			return errors.New("Lack of foreign key references to ContestsProblems.")
		}
		if sub.CUIdFK == 0 {
			return errors.New("Lack of foreign key references to ContestsUsers.")
		}
	} else {
		if sub.MetaPidFK == 0 {
			return errors.New("Lack of foreign key references to MetaProblems.")
		}
		if sub.UserIdFK == 0 {
			return errors.New("Lack of foreign key references to Users.")
		}
	}
	return nil
}

func (this *SubmissionModel) Insert(sub *Submission) (int64, error) {
	if err := this.OpenDB(); err != nil {
		return 0, err
	}
	defer this.CloseDB()

	if err := this.Validate(sub); err != nil {
		return 0, err
	}
	last_insert_id, err := this.InlineInsert(sub, nil, []string{"run_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil

}

func (this *SubmissionModel) Update(sub *Submission, required []string, excepts []string) error {
	if err := this.OpenDB(); err != nil {
		return err
	}
	defer this.CloseDB()
	if err := this.InlineUpdate(sub, "run_id", required, excepts); err != nil {
		return err
	}
	return nil
}

func (this *SubmissionModel) UpdateStatus(sub *Submission) error {
	required := []string{"status", "status_code", "testcases_passed"}
	if err := this.Update(sub, required, nil); err != nil {
		return err
	}
	return nil
}
