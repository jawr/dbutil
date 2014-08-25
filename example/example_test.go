package example

import (
    "github.com/jawr/dbutil"
    "testing"
)

func TestExample(t *testing.T) {
    err := dbutil.Setup("dbutil", "test", "dbutil")
    if err != nil {
        t.Errorf("Unable to setup database connection: %q", err)
    }
    dbconn := dbutil.Get()
    _, err = dbconn.Exec(`
        DROP TABLE IF EXISTS foo;
        DROP SEQUENCE IF EXISTS foo_id;
        CREATE TABLE foo (
            id SERIAL,
            data TEXT,
            PRIMARY KEY (id)
        )`,
    )
    if err != nil {
        t.Errorf("Unable to create test table: %q", err)
    }

    f, err := Create("milk chocolate peanuts")
    if err != nil {
        t.Errorf("Unable to create Foo object: %q", err)
    }

    f.Data = "stick of truth"
    err = f.Save()
    if err != nil {
        t.Errorf("Unable to save Foo object: %q", err)
    }

    f2, err := GetByID(f.ID)
    if err != nil {
        t.Errorf("Unable to GetByID: %q", err)
    }
    if f.ID != f2.ID || f.Data != f2.Data {
        t.Errorf("GetByID does not match: %+v vs %+v", f, f2)
    }

    fList, err := GetAll()
    if err != nil {
        t.Errorf("Unable to GetAll: %q", err)
    }
    if len(fList) != 1 {
        t.Errorf("GetByAll return mismatch, expected 1, got %d", len(fList))
    }

    _, err = dbconn.Exec(`
        DROP TABLE foo;
        DROP SEQUENCE IF EXISTS foo_id`,
    )
    if err != nil {
        t.Errorf("Unable to destroy test table: %q", err)
    }
}
