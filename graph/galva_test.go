package graph

import (
	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/store"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGalva(t *testing.T) {
	orm, _ := store.NewORM()

	c := setUp(orm)
	tearDown(orm)

	t.Run("hello world", func(t *testing.T) {
		var resp map[string]string

		c.MustPost(`query { hello }`, &resp)

		require.Equal(t, "Hello", resp["hello"])
	})

	t.Run("sign up user", func(t *testing.T) {
		var resp struct {
			AddUser struct {
				Address   string
				Signature string
			}
		}

		c.MustPost(
			`mutation($address: String!, $signature: String!) { addUser(input: {address: $address, signature: $signature}) { address signature } }`,
			&resp,
			client.Var("address", "0x40D054170DB5417369D170D1343063EeE55fb0cC"),
			client.Var("signature", "unique signature"),
		)

		require.Equal(t, "0x40D054170DB5417369D170D1343063EeE55fb0cC", resp.AddUser.Address)
		require.Equal(t, "unique signature", resp.AddUser.Signature)
	})

	t.Run("should panic for duplicate user sign up", func(t *testing.T) {
		var resp struct {
			AddUser struct {
				Address   string
				Signature string
			}
		}

		err := c.Post(
			`mutation($address: String!, $signature: String!) { addUser(input: {address: $address, signature: $signature}) { address } }`,
			&resp,
			client.Var("address", "0x40D054170DB5417369D170D1343063EeE55fb0cC"),
			client.Var("signature", "unique signature"),
		)

		require.EqualError(t, err, "[{\"message\":\"user already exists\",\"path\":[\"addUser\"]}]")
	})

	t.Run("should panic listing not tokenized property to market", func(t *testing.T) {
		var resp struct {
			AddListing struct {
				ID         int64
				PostalCode string
			}
		}

		err := c.Post(
			`mutation($id: ID!, $postalCode: String!, $location: String!, $sateliteImage: String!, $userAddress: String!) { addListing(input: {id: $id, postalCode: $postalCode, sateliteImage: $sateliteImage, location: $location, userAddress: $userAddress}) { id postalCode } }`,
			&resp,
			client.Var("id", 4325),
			client.Var("postalCode", "50300"),
			client.Var("location", "Mbale, Kenya"),
			client.Var("sateliteImage", "image"),
			client.Var("userAddress", "0x40D054170DB5417369D170D1343063EeE55fb0cC"),
		)

		require.EqualError(t, err, "[{\"message\":\"Error 'VM execution error.' querying token owner\",\"path\":[\"addListing\"]}]")
	})
}

func setUp(orm *store.ORM) *client.Client {
	db, _ := store.NewORM()
	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: New(db)})))
	return c
}

func tearDown(orm *store.ORM) {
	orm.Store.Exec("DELETE FROM users")
	orm.Store.Exec("DELETE FROM offers")
	orm.Store.Exec("DELETE FROM properties")
}
