package domain

import (
	"github.com/ednailson/hash-challenge/discount-calculator/time_now"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"testing"
	"time"
)

func TestProduct(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("creating a new product", func(t *testing.T) {
		sut := CreateProduct(1799, "Notebook Gamer", "A great notebook")

		g.Expect(sut).Should(MatchAllFields(Fields{
			"Id":           BeEquivalentTo(""),
			"Title":        BeEquivalentTo("Notebook Gamer"),
			"Description":  BeEquivalentTo("A great notebook"),
			"PriceInCents": BeEquivalentTo(1799),
		}))
	})
	t.Run("calculate discount on black friday", func(t *testing.T) {
		date := time.Date(2020, 11, 25, 10, 47, 30, 0, time.UTC)
		time_now.ReplaceFunctionTime(func() time.Time {
			return date
		})
		sut := CreateProduct(2000, "Notebook Gamer", "A great notebook")

		discount, value := sut.CalculateDiscount(fakeUser(time.Now().AddDate(-18, -1, 0)))

		g.Expect(discount).Should(BeEquivalentTo(10))
		g.Expect(value).Should(BeEquivalentTo(200))
	})
	t.Run("calculate discount on its birthday", func(t *testing.T) {
		date := time.Date(2020, 11, 28, 10, 47, 30, 0, time.UTC)
		time_now.ReplaceFunctionTime(func() time.Time {
			return date
		})
		sut := CreateProduct(2000, "Notebook Gamer", "A great notebook")

		discount, value := sut.CalculateDiscount(fakeUser(date.AddDate(-18, 0, 0)))

		g.Expect(discount).Should(BeEquivalentTo(5))
		g.Expect(value).Should(BeEquivalentTo(100))
	})
	t.Run("calculate none discount", func(t *testing.T) {
		date := time.Date(2020, 11, 28, 10, 47, 30, 0, time.UTC)
		time_now.ReplaceFunctionTime(func() time.Time {
			return date
		})
		sut := CreateProduct(2000, "Notebook Gamer", "A great notebook")

		discount, value := sut.CalculateDiscount(fakeUser(date.AddDate(-21, -1, 0)))

		g.Expect(discount).Should(BeEquivalentTo(0))
		g.Expect(value).Should(BeEquivalentTo(0))
	})
}

func fakeUser(date time.Time) *User {
	user := CreateUser("Albert", "Einstein", date)
	return &user
}
