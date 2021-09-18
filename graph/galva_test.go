package graph

import (
	"testing"

	"github.com/3dw1nM0535/galva/graph/generated"
	"github.com/3dw1nM0535/galva/store"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"

	"github.com/stretchr/testify/require"
)

var id string

func TestGalva(t *testing.T) {
	orm, _ := store.NewORM()

	// setup store before running tests
	c := setUp(orm)
	// cleanup store after running tests
	defer tearDown(orm)

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

	t.Run("sign up user 2", func(t *testing.T) {
		var resp struct {
			AddUser struct {
				Address   string
				Signature string
			}
		}

		c.MustPost(
			`mutation($address: String!, $signature: String!) { addUser(input: {address: $address, signature: $signature}) { address signature } }`,
			&resp,
			client.Var("address", "0x7f42226E0aB236Ebc578C642AdF2D0C7CE0A7FbB"),
			client.Var("signature", "unique signature 2"),
		)

		require.Equal(t, "0x7f42226E0aB236Ebc578C642AdF2D0C7CE0A7FbB", resp.AddUser.Address)
		require.Equal(t, "unique signature 2", resp.AddUser.Signature)
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

	// t.Run("should panic listing not tokenized property to market", func(t *testing.T) {
	// 	var resp struct {
	// 		AddListing struct {
	// 			ID         int64
	// 			PostalCode string
	// 		}
	// 	}

	// 	err := c.Post(
	// 		`mutation($id: ID!, $postalCode: String!, $location: String!, $sateliteImage: String!) { addListing(input: {id: $id, postalCode: $postalCode, sateliteImage: $sateliteImage, location: $location}) { id postalCode } }`,
	// 		&resp,
	// 		client.Var("id", 4325),
	// 		client.Var("postalCode", "50300"),
	// 		client.Var("location", "Mbale, Kenya"),
	// 		client.Var("sateliteImage", "image"),
	// 	)

	// 	require.EqualError(t, err, "[{\"message\":\"Error 'VM execution error.' querying token owner\",\"path\":[\"addListing\"]}]")
	// })

	// t.Run("should panic making offer to a not tokenized property", func(t *testing.T) {
	// 	var resp struct {
	// 		MakeOffer struct {
	// 			Purpose  string
	// 			Size     string
	// 			Duration time.Time
	// 		}
	// 	}

	// 	err := c.Post(
	// 		`mutation($purpose: String!, $duration: Time!, $cost: String!, $size: String!, $userAddress: String!, $propertyId: Int!) { makeOffer(input: {purpose: $purpose, size: $size, duration: $duration, cost: $cost, userAddress: $userAddress, propertyId: $propertyId}) { purpose duration size cost } }`,
	// 		&resp,
	// 		client.Var("purpose", "Apple plantation"),
	// 		client.Var("duration", time.Now().Add(time.Hour*24*10)),
	// 		client.Var("size", "3.4"),
	// 		client.Var("cost", "32"),
	// 		client.Var("userAddress", "0x40D054170DB5417369D170D1343063EeE55fb0cC"),
	// 		client.Var("propertyId", 8583),
	// 	)

	// 	require.EqualError(t, err, "[{\"message\":\"Error 'VM execution error.' querying token owner\",\"path\":[\"makeOffer\"]}]")
	// })

	// t.Run("should panic accepting nonexistent offer", func(t *testing.T) {
	// 	var resp struct {
	// 		AcceptOffer struct {
	// 			ID       string
	// 			Accepted bool
	// 		}
	// 	}

	// 	err := c.Post(
	// 		`mutation($id: ID!, $userAddress: String!) { acceptOffer(input: {id: $id, userAddress: $userAddress}) { id accepted } }`,
	// 		&resp,
	// 		client.Var("id", "id"),
	// 		client.Var("userAddress", "0x40D054170DB5417369D170D1343063EeE55fb0cC"),
	// 	)

	// 	require.EqualError(t, err, "[{\"message\":\"cannot find offer with id id\",\"path\":[\"acceptOffer\"]}]")
	// })

	// t.Run("should panic querying user profile for nonexistent user", func(t *testing.T) {
	// 	var resp struct {
	// 		GetUser struct {
	// 			Address   string
	// 			Signature string
	// 		}
	// 	}

	// 	c.MustPost(
	// 		`query($address: String!) { getUser(address: $address) { address signature } }`,
	// 		&resp,
	// 		client.Var("address", "0x40D054170DB5417369D170D1343063EeE55fb0cC"),
	// 	)

	// 	require.Equal(t, "unique signature", resp.GetUser.Signature)
	// })

	// t.Run("should panic querying user profile for nonexistent user", func(t *testing.T) {
	// 	var resp struct {
	// 		GetUser struct {
	// 			Address   string
	// 			Signature string
	// 		}
	// 	}

	// 	err := c.Post(
	// 		`query($address: String!) { getUser(address: $address) { address signature } }`,
	// 		&resp,
	// 		client.Var("address", "0x40D054170DB5417369D170D1343063EeE45fb0cC"),
	// 	)

	// 	require.EqualError(t, err, "[{\"message\":\"cannot find user\",\"path\":[\"getUser\"]}]")
	// })

	// t.Run("should panic querying info for nonlisted property", func(t *testing.T) {
	// 	var resp struct {
	// 		GetProperty struct {
	// 			PostalCode string
	// 			Location   string
	// 		}
	// 	}

	// 	err := c.Post(
	// 		`query($id: ID!) { getProperty(id: $id) { postalCode location } }`,
	// 		&resp,
	// 		client.Var("id", "id"),
	// 	)

	// 	require.EqualError(t, err, "[{\"message\":\"cannot find property\",\"path\":[\"getProperty\"]}]")
	// })

	// t.Run("should list tokenized property to market successfully", func(t *testing.T) {
	// 	var resp struct {
	// 		AddListing struct {
	// 			UserAddress string
	// 		}
	// 	}

	// 	c.MustPost(
	// 		`mutation($id: ID!, $postalCode: String!, $location: String!, $sateliteImage: String!) {
	// 			addListing(input: {
	// 				id: $id
	// 				postalCode: $postalCode
	// 				location: $location
	// 				sateliteImage: $sateliteImage
	// 			}) {
	// 				userAddress
	// 			}
	// 		}`,
	// 		&resp,
	// 		client.Var("id", 9432),
	// 		client.Var("postalCode", "50300"),
	// 		client.Var("location", "Mbale, Kenya"),
	// 		client.Var("sateliteImage", "image/jpg"),
	// 	)

	// 	require.Equal(t, "0x40D054170DB5417369D170D1343063EeE55fb0cC", resp.AddListing.UserAddress)
	// })

	// t.Run("should querying info for listed property successfully", func(t *testing.T) {
	// 	var resp struct {
	// 		GetProperty struct {
	// 			PostalCode string
	// 			Location   string
	// 		}
	// 	}

	// 	c.MustPost(
	// 		`query($id: ID!) { getProperty(id: $id) { postalCode location } }`,
	// 		&resp,
	// 		client.Var("id", 9432),
	// 	)

	// 	require.Equal(t, "Mbale, Kenya", resp.GetProperty.Location)
	// 	require.Equal(t, "50300", resp.GetProperty.PostalCode)
	// })

	// t.Run("should panic making offer to ownself tokenized property", func(t *testing.T) {
	// 	var resp struct {
	// 		MakeOffer struct {
	// 			Owner      string
	// 			FullFilled bool
	// 		}
	// 	}

	// 	err := c.Post(
	// 		`mutation($purpose: String!, $duration: Time!, $cost: String!, $size: String!, $userAddress: String!, $propertyId: Int!) {
	// 			makeOffer(input: {
	// 				purpose: $purpose,
	// 				size: $size,
	// 				duration: $duration,
	// 				cost: $cost,
	// 				userAddress: $userAddress,
	// 				propertyId: $propertyId
	// 			}) { owner fullFilled } }`,
	// 		&resp,
	// 		client.Var("purpose", "Apple plantation"),
	// 		client.Var("duration", time.Now().Add(time.Hour*24*10)),
	// 		client.Var("size", "3.4"),
	// 		client.Var("cost", "32"),
	// 		client.Var("userAddress", "0x40D054170DB5417369D170D1343063EeE55fb0cC"),
	// 		client.Var("propertyId", 9432),
	// 	)

	// 	require.EqualError(t, err, "[{\"message\":\"cannot make offer to ownself\",\"path\":[\"makeOffer\"]}]")
	// })

	// t.Run("should make offer to tokenized property successfully", func(t *testing.T) {
	// 	var resp struct {
	// 		MakeOffer struct {
	// 			ID          string
	// 			Owner       string
	// 			FullFilled  bool
	// 			UserAddress string
	// 		}
	// 	}

	// 	c.MustPost(
	// 		`mutation($purpose: String!, $duration: Time!, $cost: String!, $size: String!, $userAddress: String!, $propertyId: Int!) {
	// 			makeOffer(input: {
	// 				purpose: $purpose,
	// 				size: $size,
	// 				duration: $duration,
	// 				cost: $cost,
	// 				userAddress: $userAddress,
	// 				propertyId: $propertyId
	// 			}) { id owner fullFilled userAddress } }`,
	// 		&resp,
	// 		client.Var("purpose", "Apple plantation"),
	// 		client.Var("duration", time.Now().Add(time.Hour*24*10)),
	// 		client.Var("size", "3.4"),
	// 		client.Var("cost", "32"),
	// 		client.Var("userAddress", "0x7f42226E0aB236Ebc578C642AdF2D0C7CE0A7FbB"),
	// 		client.Var("propertyId", 9432),
	// 	)

	// 	id = resp.MakeOffer.ID
	// 	require.Equal(t, "0x40D054170DB5417369D170D1343063EeE55fb0cC", resp.MakeOffer.Owner)
	// 	require.Equal(t, "0x7f42226E0aB236Ebc578C642AdF2D0C7CE0A7FbB", resp.MakeOffer.UserAddress)
	// 	require.False(t, resp.MakeOffer.FullFilled)
	// })

	// t.Run("should panic accepting offer for nonexistent offer", func(t *testing.T) {
	// 	var resp struct {
	// 		AcceptOffer struct {
	// 			Accepted bool
	// 		}
	// 	}

	// 	err := c.Post(
	// 		`mutation($id: ID!, $userAddress: String!) { acceptOffer(input: {id: $id, userAddress: $userAddress}) { accepted } }`,
	// 		&resp,
	// 		client.Var("id", "id"),
	// 		client.Var("userAddress", "0x7f42226E0aB236Ebc578C642AdF2D0C7CE0A7FbB"),
	// 	)

	// 	require.EqualError(t, err, "[{\"message\":\"cannot find offer with id id\",\"path\":[\"acceptOffer\"]}]")
	// })

	// t.Run("should panic accepting offer for unauthorized offer author", func(t *testing.T) {
	// 	var resp struct {
	// 		AcceptOffer struct {
	// 			Accepted bool
	// 		}
	// 	}

	// 	err := c.Post(
	// 		`mutation($id: ID!, $userAddress: String!) { acceptOffer(input: {id: $id, userAddress: $userAddress}) { accepted } }`,
	// 		&resp,
	// 		client.Var("id", id),
	// 		client.Var("userAddress", "0x7f42226E0aB236Ebc578C642AdF2D0C7CE0A7FbB"),
	// 	)

	// 	require.EqualError(t, err, "[{\"message\":\"forbidden only to offer author\",\"path\":[\"acceptOffer\"]}]")
	// })

	// t.Run("should accept a not expired offer successfully", func(t *testing.T) {
	// 	var resp struct {
	// 		AcceptOffer struct {
	// 			Accepted bool
	// 		}
	// 	}

	// 	c.MustPost(
	// 		`mutation($id: ID!, $userAddress: String!) { acceptOffer(input: {id: $id, userAddress: $userAddress}) { accepted } }`,
	// 		&resp,
	// 		client.Var("id", id),
	// 		client.Var("userAddress", "0x40D054170DB5417369D170D1343063EeE55fb0cC"),
	// 	)

	// 	require.True(t, resp.AcceptOffer.Accepted)
	// })

	// t.Run("should query user and user market offers", func(t *testing.T) {
	// 	var resp struct {
	// 		GetUser struct {
	// 			Address string
	// 			Offers  []struct {
	// 				UserAddress string
	// 				Accepted    bool
	// 			}
	// 		}
	// 	}

	// 	c.MustPost(
	// 		`query($address: String!) { getUser(address: $address) { address offers { userAddress accepted } } }`,
	// 		&resp,
	// 		client.Var("address", "0x7f42226E0aB236Ebc578C642AdF2D0C7CE0A7FbB"),
	// 	)

	// 	require.Equal(t, 1, len(resp.GetUser.Offers))
	// 	require.Equal(t, "0x7f42226E0aB236Ebc578C642AdF2D0C7CE0A7FbB", resp.GetUser.Offers[0].UserAddress)
	// 	require.True(t, resp.GetUser.Offers[0].Accepted)
	// })

	// t.Run("should query user and user market listings", func(t *testing.T) {
	// 	var resp struct {
	// 		GetUser struct {
	// 			Address    string
	// 			Properties []struct {
	// 				ID int64
	// 			}
	// 		}
	// 	}

	// 	c.MustPost(
	// 		`query($address: String!) { getUser(address: $address) { address properties { id } } }`,
	// 		&resp,
	// 		client.Var("address", "0x40D054170DB5417369D170D1343063EeE55fb0cC"),
	// 	)

	// 	require.Equal(t, 1, len(resp.GetUser.Properties))
	// 	require.Equal(t, int64(9432), resp.GetUser.Properties[0].ID)
	// })

	// t.Run("should query property and property market offers", func(t *testing.T) {
	// 	var resp struct {
	// 		GetProperty struct {
	// 			ID     int64
	// 			Offers []struct {
	// 				PropertyID int64
	// 			}
	// 		}
	// 	}

	// 	c.MustPost(
	// 		`query($id: ID!) { getProperty(id: $id) { id offers{ propertyId } } }`,
	// 		&resp,
	// 		client.Var("id", 9432),
	// 	)

	// 	require.Equal(t, 1, len(resp.GetProperty.Offers))
	// 	require.Equal(t, int64(9432), resp.GetProperty.Offers[0].PropertyID)
	// })

	// t.Run("should panic accepting an expired offer", func(t *testing.T) {
	// 	var resp struct {
	// 		AcceptOffer struct {
	// 			Accepted bool
	// 		}
	// 	}

	// 	// Sleep for a minute
	// 	fmt.Println("waiting for offer expiration time ....")
	// 	time.Sleep(3 * time.Minute)

	// 	err := c.Post(
	// 		`mutation($id: ID!, $userAddress: String!) { acceptOffer(input: {id: $id, userAddress: $userAddress}) { accepted } }`,
	// 		&resp,
	// 		client.Var("id", id),
	// 		client.Var("userAddress", "0x40D054170DB5417369D170D1343063EeE55fb0cC"),
	// 	)

	// 	require.EqualError(t, err, "[{\"message\":\"offer already expired\",\"path\":[\"acceptOffer\"]}]")
	// })

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
