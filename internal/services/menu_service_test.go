package services_test

import (
	"pizza-hub/internal/models"
	"pizza-hub/internal/services"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("MenuService", func() {
	var service *services.MenuService

	BeforeEach(func() {
		service = services.NewMenuService()
	})

	Describe("FindMenuByID", func() {
		var menu1, menu2 *models.Menu

		BeforeEach(func() {
			menu1 = service.AddMenu("Menu A", 30)
			menu2 = service.AddMenu("Menu B", 45)
		})

		It("should find a menu by ID", func() {
			menu, err := service.FindMenuByID(menu1.ID)
			Expect(err).To(BeNil())
			Expect(menu).To(Equal(menu1))

			menu, err = service.FindMenuByID(menu2.ID)
			Expect(err).To(BeNil())
			Expect(menu).To(Equal(menu2))
		})

		It("should return an error if menu not found", func() {
			menu, err := service.FindMenuByID(999)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("menu not found"))
			Expect(menu).To(BeNil())
		})
	})
})
