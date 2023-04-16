package fuzzMySqlOrm

import "strconv"
import "github.com/folospace/go-mysql-orm/orm"
import fuzz "github.com/AdaLogics/go-fuzz-headers"

type User struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

func (User) TableName() string {
    return "user"
}
func (User) DatabaseName() string {
    return "mydb"
}

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 2 {
        num, _ = strconv.Atoi(string(bytes[0]))
        bytes = bytes[1:]

        switch num {

        case 0:
            fuzzConsumer := fuzz.NewConsumer(bytes)
            var query string
            err := fuzzConsumer.CreateSlice(&query)
            if err != nil {
                return 0
            }

            orm.OpenMysql(query)

        default:
            fuzzConsumer := fuzz.NewConsumer(bytes)
            var testName string
            err := fuzzConsumer.CreateSlice(&testName)
            if err != nil {
                return 0
            }

            var db, _ = orm.Open("sqlite", "test")
            var UserTable = orm.NewQuery(User{}, db)
            UserTable.Where(&UserTable.T.Name, testName)

            return 0
        }

        return 0
    }

    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}