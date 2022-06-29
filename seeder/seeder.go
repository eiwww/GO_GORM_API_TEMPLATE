package seeder

import (
	"fmt"
	"rest/database"
	model "rest/model"
)

type Province = model.Province

type District = model.District

func SeedDB() {

	db, conErr := database.GetDatabaseConnection()
	if conErr != nil {
		fmt.Println(conErr)
		return
	}

	province := []model.Province{}
	proData := db.Find(&province)
	if proData.RowsAffected <= 0 {
		var provinces = []Province{
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
		db.Create(&provinces)
	}

	district := []model.District{}
	disData := db.Find(&district)
	if disData.RowsAffected <= 0 {
		var districts = []District{
			// 1. ນະຄອນຫຼວງວຽງຈັນ
			{ProvinceId: 1, Name: "ຈັນທະບູລີ"},  // 1
			{ProvinceId: 1, Name: "ສີໂຄດຕະບອງ"}, // 2
			{ProvinceId: 1, Name: "ໄຊເສດຖາ"},    // 3
			{ProvinceId: 1, Name: "ສີສັດຕະນາກ"}, // 4
			{ProvinceId: 1, Name: "ນາຊາຍທອງ"},   // 5
			{ProvinceId: 1, Name: "ໄຊທານີ"},     // 6
			{ProvinceId: 1, Name: "ຫາດຊາຍຟອງ"},  // 7
			{ProvinceId: 1, Name: "ສັງທອງ"},     // 8
			{ProvinceId: 1, Name: "ປາກງື່ມ"},    // 9
			// 2. ຜົ້ງສາລີ
			{ProvinceId: 2, Name: "ຜົ້ງສາລີ"}, // 1
			{ProvinceId: 2, Name: "ໃໝ່"},      // 2
			{ProvinceId: 2, Name: "ຂວາ"},      // 3
			{ProvinceId: 2, Name: "ສໍາພັນ"},   // 4
			{ProvinceId: 2, Name: "ບຸນເໜືອ"},  // 5
			{ProvinceId: 2, Name: "ຍອດອູ"},    // 6
			{ProvinceId: 2, Name: "ບຸນໃຕ້"},   // 7
			// 3. ຫຼວງນໍ້າທາ
			{ProvinceId: 3, Name: "ຫຼວງນໍ້າທາ"}, // 1
			{ProvinceId: 3, Name: "ສິງ"},        // 2
			{ProvinceId: 3, Name: "ລອງ"},        // 3
			{ProvinceId: 3, Name: "ວຽງພູຄໍາ"},   // 4
			{ProvinceId: 3, Name: "ນາແລ"},       // 5
			// 4. ອຸດົມໄຊ
			{ProvinceId: 4, Name: "ໄຊ"},      // 1
			{ProvinceId: 4, Name: "ຫຼາ"},     // 2
			{ProvinceId: 4, Name: "ນາໝໍ້"},   //3
			{ProvinceId: 4, Name: "ງາ"},      // 4
			{ProvinceId: 4, Name: "ແບງ"},     // 5
			{ProvinceId: 4, Name: "ຮຸນ"},     // 6
			{ProvinceId: 4, Name: "ປາກແບ່ງ"}, // 7
			// 5. ບໍ່ແກ້ວ
			{ProvinceId: 5, Name: "ຫ້ວຍຊາຍ"},   // 1
			{ProvinceId: 5, Name: "ຕົ້ນເຜິ້ງ"}, // 2
			{ProvinceId: 5, Name: "ເມິງ"},      // 3
			{ProvinceId: 5, Name: "ຜາອຸດົມ"},   // 4
			{ProvinceId: 5, Name: "ປາກທາ"},     // 5
			// 6. ຫຼວງພະບາງ
			{ProvinceId: 6, Name: "ນະຄອນຫຼວງພະບາງ"}, // 1
			{ProvinceId: 6, Name: "ຊຽງເງິນ"},        // 2
			{ProvinceId: 6, Name: "ນານ"},            // 3
			{ProvinceId: 6, Name: "ປາກອູ"},          // 4
			{ProvinceId: 6, Name: "ນໍ້າບາກ"},        // 5
			{ProvinceId: 6, Name: "ງອຍ"},            // 6
			{ProvinceId: 6, Name: "ປາກແຊງ"},         // 7
			{ProvinceId: 6, Name: "ໂພນໄຊ"},          // 8
			{ProvinceId: 6, Name: "ຈອມເພັດ"},        // 9
			{ProvinceId: 6, Name: "ວຽງຄໍາ"},         // 10
			{ProvinceId: 6, Name: "ພູຄູນ"},          // 11
			{ProvinceId: 6, Name: "ໂພນທອງ"},         // 12
			// 7. ຫົວພັນ
			{ProvinceId: 7, Name: "ຊໍາເໜືອ"},  // 1
			{ProvinceId: 7, Name: "ຊຽງຄໍ້"},   // 2
			{ProvinceId: 7, Name: "ຮ້ຽມ"},     // 3
			{ProvinceId: 7, Name: "ວຽງໄຊ"},    // 4
			{ProvinceId: 7, Name: "ຫົວເມືອງ"}, // 5
			{ProvinceId: 7, Name: "ຊໍາໃຕ້"},   // 6
			{ProvinceId: 7, Name: "ສົບເບົາ"},  // 7
			{ProvinceId: 7, Name: "ແອດ"},      // 8
			{ProvinceId: 7, Name: "ກວັນ"},     // 9
			{ProvinceId: 7, Name: "ຊ່ອນ"},     // 10
			// 8. ໄຊຍະບູລີ
			{ProvinceId: 8, Name: "ໄຊຍະບູລີ"}, // 1
			{ProvinceId: 8, Name: "ຄອບ"},      // 2
			{ProvinceId: 8, Name: "ຫົງສາ"},    // 3
			{ProvinceId: 8, Name: "ເງິນ"},     // 4
			{ProvinceId: 8, Name: "ຊຽງຮ່ອນ"},  // 5
			{ProvinceId: 8, Name: "ພຽງ"},      // 6
			{ProvinceId: 8, Name: "ປາກອອກ"},   // 7
			{ProvinceId: 8, Name: "ແກ່ນທ້າວ"}, // 8
			{ProvinceId: 8, Name: "ບໍ່ແຕນ"},   // 9
			{ProvinceId: 8, Name: "ທົ່ງມີໄຊ"}, // 10
			{ProvinceId: 8, Name: "ໄຊສະຖານ"},  // 11
			// 9. ຊຽງຂວາງ
			{ProvinceId: 9, Name: "ແປກ (ໂພນສະຫວັນ)"}, // 1
			{ProvinceId: 9, Name: "ຄໍາ"},             // 2
			{ProvinceId: 9, Name: "ໜອງແຮດ"},          // 3
			{ProvinceId: 9, Name: "ຄູນ"},             // 4
			{ProvinceId: 9, Name: "ໝອກໃໝ່"},          // 5
			{ProvinceId: 9, Name: "ພູກູດ"},           // 6
			{ProvinceId: 9, Name: "ຜາໄຊ"},            // 7
			// 10. ວຽງຈັນ
			{ProvinceId: 10, Name: "ໂພນໂຮງ"},    // 1
			{ProvinceId: 10, Name: "ທຸລະຄົມ"},   // 2
			{ProvinceId: 10, Name: "ແກ້ວອຸດົມ"}, // 3
			{ProvinceId: 10, Name: "ກາສີ"},      // 4
			{ProvinceId: 10, Name: "ວັງວຽງ"},    // 5
			{ProvinceId: 10, Name: "ເຟືອງ"},     // 6
			{ProvinceId: 10, Name: "ຊະນະຄາມ"},   // 7
			{ProvinceId: 10, Name: "ແມດ"},       // 8
			{ProvinceId: 10, Name: "ຫີນເຫີບ"},   // 9
			{ProvinceId: 10, Name: "ວຽງຄໍາ"},    // 10
			{ProvinceId: 10, Name: "ໝື່ນ"},      // 11
			// 11. ບໍລິຄໍາໄຊ
			{ProvinceId: 11, Name: "ປາກຊັນ"},            // 1
			{ProvinceId: 11, Name: "ທ່າພະບາດ"},          // 2
			{ProvinceId: 11, Name: "ປາກກະດິງ"},          // 3
			{ProvinceId: 11, Name: "ບໍລິຄັນ"},           // 4
			{ProvinceId: 11, Name: "ຄໍາເກີດ (ຫຼັກ 20)"}, // 5
			{ProvinceId: 11, Name: "ວຽງທອງ"},            // 6
			{ProvinceId: 11, Name: "ໄຊຈໍາພອນ"},          // 7
			// 12. ຄໍາມ່ວນ
			{ProvinceId: 12, Name: "ທ່າແຂກ"},   // 1
			{ProvinceId: 12, Name: "ມະຫາໄຊ"},   // 2
			{ProvinceId: 12, Name: "ໜອງບົກ"},   // 3
			{ProvinceId: 12, Name: "ຫີນບູນ"},   // 4
			{ProvinceId: 12, Name: "ຍົມມະລາດ"}, // 5
			{ProvinceId: 12, Name: "ບົວລະພາ"},  // 6
			{ProvinceId: 12, Name: "ນາກາຍ"},    // 7
			{ProvinceId: 12, Name: "ເຊບັ້ງໄຟ"}, // 8
			{ProvinceId: 12, Name: "ໄຊບົວທອງ"}, // 9
			{ProvinceId: 12, Name: "ຄູນຄໍາ"},   // 10
			// 13. ສະຫວັນນະເຂດ
			{ProvinceId: 13, Name: "ນະຄອນໄກສອນ ພົມວິຫານ"}, // 1
			{ProvinceId: 13, Name: "ອຸທົມພອນ"},            // 2
			{ProvinceId: 13, Name: "ອາດສະພັງທອງ"},         // 3
			{ProvinceId: 13, Name: "ພິນ"},                 // 4
			{ProvinceId: 13, Name: "ເຊໂປນ"},               // 5
			{ProvinceId: 13, Name: "ນອງ"},                 // 6
			{ProvinceId: 13, Name: "ທ່າປາງທອງ"},           // 7
			{ProvinceId: 13, Name: "ສອງຄອນ"},              // 8
			{ProvinceId: 13, Name: "ຈໍາພອນ"},              // 9
			{ProvinceId: 13, Name: "ຊົນນະບູລີ"},           // 10
			{ProvinceId: 13, Name: "ໄຊບູລີ"},              // 11
			{ProvinceId: 13, Name: "ວິລະບູລີ"},            // 12
			{ProvinceId: 13, Name: "ອາດສະພອນ"},            // 13
			{ProvinceId: 13, Name: "ໄຊພູທອງ"},             // 14
			{ProvinceId: 13, Name: "ພະລານໄຊ"},             // 15
			// 14. ສາລະວັນ
			{ProvinceId: 14, Name: "ສາລະວັນ"},   // 1
			{ProvinceId: 14, Name: "ຕະໂອ້ຍ"},    // 2
			{ProvinceId: 14, Name: "ຕຸ້ມລານ"},   // 3
			{ProvinceId: 14, Name: "ລະຄອນເພັງ"}, // 4
			{ProvinceId: 14, Name: "ວາປີ"},      // 5
			{ProvinceId: 14, Name: "ຄົງເຊໂດນ"},  // 6
			{ProvinceId: 14, Name: "ເລົ່າງາມ"},  // 7
			{ProvinceId: 14, Name: "ສະໝ້ວຍ"},    // 8
			// 15. ເຊກອງ
			{ProvinceId: 15, Name: "ລະມາມ"},  // 1
			{ProvinceId: 15, Name: "ກະລຶມ"},  // 2
			{ProvinceId: 15, Name: "ດາກຈຶງ"}, // 3
			{ProvinceId: 15, Name: "ທ່າແຕງ"}, // 4
			// 16. ຈໍາປາສັກ
			{ProvinceId: 16, Name: "ນະຄອນປາກເຊ"},     // 1
			{ProvinceId: 16, Name: "ຊະນະສົມບູນ"},     // 2
			{ProvinceId: 16, Name: "ບາຈຽງຈະເລີນສຸກ"}, // 3
			{ProvinceId: 16, Name: "ປາກຊ່ອງ"},        // 4
			{ProvinceId: 16, Name: "ປະທຸມພອນ"},       // 5
			{ProvinceId: 16, Name: "ໂພນທອງ"},         // 6
			{ProvinceId: 16, Name: "ຈໍາປາສັກ"},       // 7
			{ProvinceId: 16, Name: "ສຸຂຸມາ"},         // 8
			{ProvinceId: 16, Name: "ມູນລະປະໂມກ"},     // 9
			{ProvinceId: 16, Name: "ໂຂງ"},            // 10
			// 17. ອັດຕະປື
			{ProvinceId: 17, Name: "ໄຊເສດຖາ"},   // 1
			{ProvinceId: 17, Name: "ສາມັກຄີໄຊ"}, // 2
			{ProvinceId: 17, Name: "ສະໜາມໄຊ"},   // 3
			{ProvinceId: 17, Name: "ຊານໄຊ"},     // 4
			{ProvinceId: 17, Name: "ພູວົງ"},     // 5
			// 18. ໄຊສົມບູນ
			{ProvinceId: 18, Name: "ອະນຸວົງ"},  // 1
			{ProvinceId: 18, Name: "ລ້ອງແຈ້ງ"}, // 2
			{ProvinceId: 18, Name: "ລ້ອງຊານ"},  // 3
			{ProvinceId: 18, Name: "ຮົ່ມ"},     // 4
			{ProvinceId: 18, Name: "ທ່າໂທມ"},   // 5
		}
		db.Create(&districts)
	}
}
