package apb

import (
	"testing"

	ft "github.com/fusor/ansible-service-broker/pkg/fusortest"
)

func TestBind(t *testing.T) {
	t.Skip("skipping bind until we can pass in a mock client")
	output := []byte(`
Login failed (401 Unauthorized)

PLAY [all] *********************************************************************

TASK [setup] *******************************************************************
ok: [localhost]

TASK [Bind] ********************************************************************
changed: [localhost]

TASK [debug] *******************************************************************
ok: [localhost] => {
    "msg": "<BIND_CREDENTIALS>eyJkYiI6ICJmdXNvcl9ndWVzdGJvb2tfZGIiLCAidXNlciI6ICJkdWRlcl90d28iLCAicGFzcyI6ICJkb2c4dHdvIn0=</BIND_CREDENTIALS>"
}

PLAY RECAP *********************************************************************
localhost                  : ok=3    changed=1    unreachable=0    failed=0
`)
	result, err := decodeOutput(output)
	if err != nil {
		t.Fatal(err)
	}

	ft.AssertNotNil(t, result, "result")
	ft.AssertEqual(t, result["db"], "fusor_guestbook_db", "db is not fusor_guestbook_db")
	ft.AssertEqual(t, result["user"], "duder_two", "user is not duder_two")
	ft.AssertEqual(t, result["pass"], "dog8two", "password is not dog8two")
}
