package include

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/hashicorp/serf/serf"
	"go.uber.org/zap"
)

type ConnectionVerifier func(string) bool

func TestMemberLoggingErrors(member serf.Member, errs []error) {
	zap.Errors(member.Name, errs)
	zap.Errors(member.Tags["rpc_addr"], errs)
}

func TestServe(w http.ResponseWriter, r *http.Request, t *testing.T) {
	fmt.Fprint(w, "", r)

	t.Run("Return Proper HTTP Response Code\n\t", func(t *testing.T) {
		got := r.Response.StatusCode
		want := 200

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestDataConnection(verify ConnectionVerifier, paths []string) map[string]bool {
	resN := make(map[string]bool)

	for _, path := range paths {
		resN[path] = verify(path)
	}

	return resN
}

func CheckConnection(path string) bool {
	localhost := "localhost"
	localHttps := ":8080"
	postgre := "postgre://"
	postgrePort := ":5432"

	conns := map[string]bool{
		localhost + localHttps: true,
		postgre + postgrePort:  true,
	}

	vals := reflect.ValueOf(conns).MapKeys()
	slices := make([]string, len(vals))
	for idx := 0; idx < len(conns); idx++ {
		slices[idx] = vals[idx].String()
	}

	return path == strings.Join(slices, "")
} // Available [online]: https://stackoverflow.com/a/41690880

func TestPathing(t *testing.T) {
	domains := []string{
		"postgres://localhost",
		"postgres://localhost:5432",
	}

	want := map[string]bool{
		"postgres://localhost":      true,
		"postgres://localhost:5432": true,
	}

	got := TestDataConnection(CheckConnection, domains)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("Wanted %v, got %v", want, got)
	}
} // Available [online]: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency
