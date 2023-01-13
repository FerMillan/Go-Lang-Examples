package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"awesomeProject/database/models"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	_ "github.com/go-sql-driver/mysql"
)

func getConnection() (context.Context, *sql.DB) {
	ctx := context.Background()
	db, err := sql.Open("mysql", user_db+":"+password_db+"@tcp("+host_db+")/"+scheme_db+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	boil.SetDB(db) // Sets the current db object globally
	return ctx, db
}

func ExampleSelectCountToUsers() {
	ctx, db := getConnection()
	// SELECT COUNT(*) FROM tbl_users
	fmt.Println("-----------------------------------------------")
	fmt.Println("SELECT COUNT(*) FROM tbl_users")
	// Simplest possible query. models.TBLUsers is a "starter method"
	// and specifies the table. Count is a "finisher" and specifies
	// the operation
	numUsers, err := models.TBLUsers().Count(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numUsers)
}

func ExampleSelectLimitORM() {
	ctx, db := getConnection()
	//##########################################
	// SELECT * FROM tbl_users LIMIT 1
	fmt.Println("-----------------------------------------------")
	fmt.Println("SELECT * FROM tbl_users LIMIT 1 (ORM)")
	// .One is equivalent to SELECT ... LIMIT 1
	// it returns a pointer to a tbl_users model object.
	// qm.Select is a "query mod" and is used to modify the query
	userP, err := models.TBLUsers(
		qm.Select(models.TBLUserColumns.ID, // Maps to "user_id"
			models.TBLUserColumns.Email), // The args do jack shit, the model brings every column anyways
	).One(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	user := *userP
	fmt.Printf("%+v\n", user)
	// The column values are stored in the user object
	// as attributes whose names start with a capital letter
	//fmt.Printf("%s, %s, %d\n",user.Name, user.Email, user.Mobile)
	fmt.Println("-----------------------------------------------")
	fmt.Printf("%s, %s, %d\n", user.Name, user.Email, user.Mobile.Int64) // the mobile is of type null.Int64
}

func ExampleNativeQuery() {
	ctx, db := getConnection()
	var err error
	//##########################################
	// Native query equivalent to the above.
	fmt.Println("-----------------------------------------------")
	fmt.Println("SELECT * FROM tbl_users LIMIT 1 (native)")
	// Here's how it works: the following struct's attributes must be
	// sqlboiler-generated types, one for each table (model) that will
	// participate in the query. The struct tag states:
	// - boil: this attribute will be used by the sqlboiler package
	// - tbl_users: any value belonging to the tbl_users table will be
	//              applied to this model
	// - bind: The query's results, subject to the above, will be
	//         added to this model as an attribute
	type TblUsersWrapper struct {
		TblUser models.TBLUser `boil:"tbl_users,bind"` // Remember the capital letter in the name
	}

	tblUsersWrapper := TblUsersWrapper{}
	// The bind method will execute the query defined in the Raw method
	// and assign the results to the attributes of its third parameter
	err = queries.Raw("SELECT * FROM tbl_users LIMIT 1").Bind(ctx, db, &tblUsersWrapper)
	if err != nil {
		log.Fatal(err)
	}
	// The TblUser attribute contains the query's results
	fmt.Println(tblUsersWrapper.TblUser)
	// An example of a join: https://pkg.go.dev/github.com/volatiletech/sqlboiler/v4/queries#Bind
}

func ExampleMoreHardQueryUseORM() {
	ctx, db := getConnection()
	//##########################################
	fmt.Println("-----------------------------------------------")
	fmt.Println("SELECT * FROM tbl_users WHERE email = 'vikram@undostres.com.mx' LIMIT 1")
	// To build more complicated queries, use the "query mod system" (functions in the qm package).
	// These functions are passed into instances of models, and there can be an arbitrary
	// number of them (since the model's argument is a slice).

	// ONE
	tblUser, err := models.TBLUsers(
		qm.Where("email = ?", "vikram@undostres.com.mx"),
	).One(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	// ALL
	tblUser2, err := models.TBLUsers(
		qm.Limit(10),
	).All(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	// Print values
	fmt.Println(tblUser)
	// Print slice of pointers
	fmt.Println(tblUser2)
}

// To Excecute correctly this funcction need have to configure globally the database if not its correctly this
// throw a panic message
func ExampleStillMoreHardQueryUseORM() {
	ctx, _ := getConnection()
	var err error
	//##########################################
	// More complex query
	fmt.Println("-----------------------------------------------")
	fmt.Println("SELECT * " +
		"FROM tbl_referralcodes tr " +
		"JOIN tbl_users tu " +
		"ON tr.user_id = tu.id " +
		"WHERE tu.mobile IS NULL " +
		"AND tu.facebookid = '' " +
		"LIMIT 1")

	// Once again, use a bound struct
	type WrapperForJoinQuery struct {
		TblUser          models.TBLUser         `boil:"tbl_users,bind"`
		TblReferralcodes models.TBLReferralcode `boil:"tbl_referralcodes,bind"`
	}

	wrapperForJoinQuery := WrapperForJoinQuery{}

	err = models.NewQuery(
		qm.Select("tr.*, tu.*"),
		qm.From("tbl_referralcodes tr"),
		qm.InnerJoin("tbl_users tu ON tr.user_id = tu.id"),
		qm.Where("tu.mobile IS NULL"),
		qm.Where("tu.facebookid = ?", ""),
		qm.Limit(1),
	).BindG(ctx, &wrapperForJoinQuery) // G for global (it uses the db executor set at the beginning)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wrapperForJoinQuery.TblReferralcodes, wrapperForJoinQuery.TblUser)
	print("-----------------------Other form to print this-----------------------\n")
	fmt.Println(wrapperForJoinQuery)
}

func ExampleEveryRowInATable() {
	_, db := getConnection()
	//##########################################
	// Return every row in a table
	fmt.Println("-----------------------------------------------")
	fmt.Println("SELECT * FROM referralcode_rejection_reason (All rows)")

	// .QueryP is the same as .Query but it panics on error instead of returning an error var
	rows := models.ReferralcodeRejectionReasons().QueryP(db) // Returns an *sql.Rows type.
	// Another possibility is to use the .All() finisher, which .Query() calls anyways,
	// the difference being that the former returns a *[]ReferralcodeRejectionReasons
	// rather than a *sql.Rows
	var (
		rejectionCode         uint8
		rejectionReason       null.String
		createdAt, modifiedAt null.Time
	)
	for rows.Next() {
		rows.Scan(&rejectionCode, &rejectionReason, &createdAt, &modifiedAt)
		fmt.Println(rejectionCode, rejectionReason, createdAt, modifiedAt)
	}
}

func ExampleQueryToUseSafeConstants() {
	_, db := getConnection()
	//##########################################
	// Query with "type-safe" constants
	fmt.Println("-----------------------------------------------")
	fmt.Println("SELECT * FROM tbl_referralcodes WHERE user_id = 2647712 (Type safe version)")
	// Each of these models.* structs' attributes maps to a string containing its name
	// in the DB. These are safe in the sense that if the table or column
	// names in the DB are changed, it's only necessary to rebuild the models
	// instead of changing the names in a bunch of places in the code.

	// For the Select, there's also models.TBLReferralcodeTableColumns,
	// which appends the tbl_referralcodes scope to the columns name
	// (e.g. tbl_referralcodes.user_id)
	row := models.NewQuery(
		qm.Select(
			models.TBLReferralcodeColumns.UserID,       // Maps to "user_id"
			models.TBLReferralcodeColumns.Referralcode, // "referralcode"
		),
		qm.From(models.TableNames.TBLReferralcodes),    // "tbl_referralcodes"
		models.TBLReferralcodeWhere.UserID.EQ(2647712), // Seems it only generated WHEREs for the primary key
	).QueryRow(db) // Returns *sql.Row (note the singular)

	// What do you know, this actually returns only the two columns in the Select
	var userId int
	var referralcode string

	row.Scan(&userId, &referralcode) // Obviously lacks a .Next() method
	fmt.Println(userId, referralcode)
}
