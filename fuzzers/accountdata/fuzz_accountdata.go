package stellargofuzz
import (
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/support/http/httptest"
)

func Fuzz(data []byte) int {
	// fuzz DataKey 
	hmock := httptest.NewClient()
	client := &horizonclient.Client{
		HorizonURL: "https://localhost/",
		HTTP: hmock,
	}
	accountRequest := horizonclient.AccountRequest{
		AccountID: "GCLWGQPMKXQSPF776IU33AH4PZNOOWNAWGGKVTBQMIC5IMKUNP3E6NVU",
		DataKey:   string(data),
	}
	_, err := client.AccountData(accountRequest)
	if err != nil { return 1 }
	return 0
}
