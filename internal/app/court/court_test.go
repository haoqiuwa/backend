package court

import (
	"os"
	"testing"
	"wxcloudrun-golang/internal/pkg/db"
)

func TestMain(m *testing.M) {
	db.Init()
	os.Exit(m.Run())
}

func TestService_GetCourts(t *testing.T) {

}
