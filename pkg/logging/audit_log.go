package logging

import (
	"fmt"
	"time"

	"github.com/satriaprayoga/cukurin-user/pkg/utils"
)

type auditLog struct {
	ID        int64     `bson:"id"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Level     string    `json:"level"`
	UUID      string    `json:"uuid"`
	FuncName  string    `json:"func_name"`
	FileName  string    `json:"file_name"`
	Line      int       `json:"line"`
	Time      string    `json:"time"`
	Message   string    `json:"message"`
}

func (a *auditLog) saveAudit() {

	a.ID = utils.GetTimeNow().Unix()
	a.Message = "API User : " + a.Message
	fmt.Printf("Inserted a single document: %v", a.Message)

}
