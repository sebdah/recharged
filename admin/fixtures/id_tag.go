package fixtures
import "github.com/sebdah/recharged/admin/models"

// Setup all fixtures
func SetupIdTag() {
	log.Info("Creating fixtures for IdTag")

	one := models.NewIdTag()
	one.IdTag = "1"
	models.Save(one)

	two := models.NewIdTag()
	two.IdTag = "2"
	models.Save(two)
}
