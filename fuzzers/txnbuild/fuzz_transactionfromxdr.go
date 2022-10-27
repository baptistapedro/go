package stellargofuzz
import "github.com/stellar/go/txnbuild"

func Fuzz(data []byte) int {
	txnbuild.TransactionFromXDR(string(data))
	return 0	
}
