package tools

type mockDB struct {}

var mockBarcodeGroup= map[string]BarcodeGroup{
	"barcode": {
		Prefix: "barcode",
		Sequence: 1,
	},
	"sqpd": {
		Prefix: "sqpd",
		Sequence: 10,
	},
	"test": {
		Prefix: "test",
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


func (d *mockDB) CreateBarcodeGroup(prefix string, sequence int) *BarcodeGroup {
	barcodeGroupData := BarcodeGroup{
		Prefix: prefix,
		Sequence: int64(sequence),
	}
	mockBarcodeGroup[prefix] = barcodeGroupData

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