package stellargofuzz
import "github.com/stellar/go/xdr"

//type String64 = string

func Fuzz(data []byte) int {
	// Fuzzing xdr.LedgerEntryTypeData.DataName
	// by triggering switch case LedgerEntryTypeData
	// in ledgerKeyCompressEncodeTo()
	
	ledgerKey := xdr.LedgerKey{
		Type: xdr.LedgerEntryTypeData,
		Data: &xdr.LedgerKeyData{
			AccountId: xdr.MustAddress("GAOQJGUAB7NI7K7I62ORBXMN3J4SSWQUQ7FOEPSDJ322W2HMCNWPHXFB"),
			DataName:  xdr.String64(string(data)),
		},
	
	}
	e := xdr.NewEncodingBuffer()
	_, err := e.LedgerKeyUnsafeMarshalBinaryCompress(ledgerKey)
	if err != nil {
		return 1
	}
	return 0
}
