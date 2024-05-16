package tools

type mockDB struct {}

var mockBarcodeGroup= map[string]BarcodeGroup{
	"barcode": {
		Barcode: "barcode",
		Sequence: 1,
	},
	"sqpd": {
		Barcode: "sqpd",
		Sequence: 10,
	},
	"test": {
		Barcode: "test",
		Sequence: 100,
	},
}

func (d *mockDB) CreateBarcode(barcodeGroup string) *BarcodeGroup {
	barcodeGroupData, ok := mockBarcodeGroup[barcodeGroup]
	if !ok {
		return nil
	}

	barcodeGroupData.Sequence++
	mockBarcodeGroup[barcodeGroup] = barcodeGroupData

	return &barcodeGroupData
}

func (d *mockDB) GetBarcodeGroup(barcodeGroup string) *BarcodeGroup {
	barcodeGroupData, ok := mockBarcodeGroup[barcodeGroup]
	if !ok {
		return nil
	}

	return &barcodeGroupData
}

func (d *mockDB) GetBarcodeGroups() *[]BarcodeGroup {
	var barcodeGroups []BarcodeGroup
	for _, barcodeGroup := range mockBarcodeGroup {
		barcodeGroups = append(barcodeGroups, barcodeGroup)
	}

	return &barcodeGroups
}

func (d *mockDB) SetupDatabase() error {
	return nil
}