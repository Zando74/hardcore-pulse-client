package value_object

import "slices"

type MapID int 

type InvalidMapIDError struct {}

func (e *InvalidMapIDError) Error() string {
	return "Invalid Map ID"
}

var ValidMapIds = []int{
	33, 34, 36, 43, 47, 48, 70, 90 ,109,
	129, 209, 229, 230, 329, 349, 389, 429,
	1001, 1004, 1007, 409, 469, 509, 531, 249,
	3459, 189, 533,
	1411, 1412, 1413, 1414, 1415, 1416, 
	1417, 1418, 1419, 1420, 1421, 1422, 
	1423, 1424, 1425, 1426, 1427, 1428, 
	1429, 1430, 1431, 1432, 1433, 1434,
	1435, 1436, 1437, 1438, 1439, 1440, 
	1441, 1442, 1443, 1444, 1445, 1446, 
	1447, 1448, 1449, 1450, 1451, 1452, 
	1453, 1454, 1455, 1456, 1457, 1458,
}

func NewMapID(value int) (*MapID, error) {
	if !slices.Contains(ValidMapIds, value) {
		return nil, &InvalidMapIDError{}
	}
	mapID := MapID(value)
	return &mapID, nil
}