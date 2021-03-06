package stellar

import (
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
)

func CreateAccount(address string) error{
	// pair is the pair that was generated from previous example, or create a pair based on
	// existing keys.

	//address := "GBDG2ITTCWWNALG2OG5IYEBV3LFMF74WLSXL2XQGZXBM5I7W3OIMELNG"
	resp, err := http.Get("https://horizon-testnet.stellar.org/friendbot?addr=" + address)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	return err
}
