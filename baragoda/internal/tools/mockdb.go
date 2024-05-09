package tools

type mockDB struct {}

var mockBarcodeSequence = map[string]BarcodeSequence{
	"barcode": {
		Barcode: "barcode",
		Sequence: 1,
	},
}

func (d *mockDB) GetBarcodeSequence(barcode string) *BarcodeSequence {
	barcodeSequenceData, ok := mockBarcodeSequence[barcode]
	if !ok {
		return nil
	}

	return &barcodeSequenceData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}