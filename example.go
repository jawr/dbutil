package main

type Foo {
    ID uint64
    Data string
}

const (
    INSERT_STMT string = "INSERT INTO foo (data) VALUES ($1) RETURNING id"
    UPDATE_STMT string = "UPDATE foo SET data = $1 WHERE id = $2"
    SELECT_STMT string = "SELECT * FROM foo"
    BY_ID_STMT string = SELECT_STMT + " WHERE id = $1"
)

func Create(data string) (f *Foo, err error) {
    id, err := Insert(INSERT_STMT, data)
    if err != nil { return f, err }
    f = &Foo{
        ID: id,
        Data: data,
    }
    return f, nil
}

func (f *Foo) Save() error {
    dbconn := Get()
    _, err := dbconn.Exec(
        UPDATE_STMT,
        f.Data,
    )
    return err
}

/* this is helpful when using trickier objects */
func ParseRow(row Row) (Object, error) {
    var id uint64
    var data string
    err := row.Scan(&id, &data)
    f := &Foo{
        ID: id,
        Data: data,
    }
    return f, err
}

func GetByID(id uint64) (*Foo, error) {
    item, err := Select(
        ParseRowFunc(ParseRow),
        BY_ID_STMT,
        id,
    )
    return item.(*Foo), err
}

func GetAll() ([]*Foo, error) {
    var items []*Foo
    objects, err := SelectList(
        ParseRowFunc(ParseRow),
        SELECT_STMT,
    )
    for _, i := range objects {
        items = append(items, i.(*Foo))
    }
    return items, err
}
