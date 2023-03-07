package database

import (
	"database/sql"
	"github/pgabrieldeveloper/intensivo_go/internal/entity"
	"testing"

	"github.com/stretchr/testify/suite"
	_ "modernc.org/sqlite"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	db *sql.DB
}

func (suite *OrderRepositoryTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite", ":memory:")
	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varcher(255) not null, price float not null, tax float not null, final_price float not null, primary key (id))")
	suite.db = db
}

func (suite *OrderRepositoryTestSuite) TearDownSuite() {

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestSavingOrder() {
	order, err := entity.NewOrder("123", 10, 2)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())
	repo := NewOrderRepository(suite.db)
	err = repo.Save(order)
	suite.NoError(err)
	total, err := repo.GetTotal()
	suite.NoError(err)
	suite.Equal(1, total)
}
