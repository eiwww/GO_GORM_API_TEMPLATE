package seeder

import (
	model "rest/model"
)

var provinces = []model.Province{
	{Name: "ນະຄອນຫຼວງວຽງຈັນ", Section: "center"}, // 1
	{Name: "ຜົ້ງສາລີ", Section: "north"},         // 2
	{Name: "ຫຼວງນ້ຳທາ", Section: "north"},        // 3
	{Name: "ອຸດົມໄຊ", Section: "north"},          // 4
	{Name: "ບໍ່ແກ້ວ", Section: "north"},          // 5
	{Name: "ຫຼວງພະບາງ", Section: "north"},        // 6
	{Name: "ຫົວພັນ", Section: "north"},           // 7
	{Name: "ໄຊຍະບູລີ", Section: "north"},         // 8
	{Name: "ຊຽງຂວາງ", Section: "north"},          // 9
	{Name: "ວຽງຈັນ", Section: "center"},          // 10
	{Name: "ບໍລິຄຳໄຊ", Section: "center"},        // 11
	{Name: "ຄຳມ່ວນ", Section: "center"},          // 12
	{Name: "ສະຫວັນນະເຂດ", Section: "center"},     // 13
	{Name: "ສາລະວັນ", Section: "south"},          // 14
	{Name: "ເຊກອງ", Section: "south"},            // 15
	{Name: "ຈຳປາສັກ", Section: "south"},          // 16
	{Name: "ອັດຕະປື", Section: "south"},          // 17
	{Name: "ໄຊສົມບູນ", Section: "center"},        // 18
}
