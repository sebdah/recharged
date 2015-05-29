package fixtures
import "github.com/sebdah/recharged/admin/models"

func SetupChargePoint() {
	log.Info("Creating fixtures for ChargePoint")

	one := models.NewChargePoint()
	one.Model = "Model X"
	one.Vendor = "Vendor Y"
	one.SerialNumber = "1234"
	one.Imsi = "12344"
	models.Save(one)

	two := models.NewChargePoint()
	two.Model = "Model 2"
	two.Vendor = "Vendor Y"
	two.SerialNumber = "111222"
	models.Save(two)

	three := models.NewChargePoint()
	three.Model = "Model 1"
	three.Vendor = "Vendor A"
	models.Save(three)
}
