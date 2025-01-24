package services_test

import (
	"pizza-hub/internal/models"
	"pizza-hub/internal/services"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrderService", func() {
	var (
		chefService  *services.ChefService
		menuService  *services.MenuService
		orderService *services.OrderService
		chef1        *models.Chef
		menu1        *models.Menu
	)

	BeforeEach(func() {
		chefService = services.NewChefService()
		menuService = services.NewMenuService()
		orderService = services.NewOrderService(chefService, menuService)

		// Add chefs and menus
		chef1 = chefService.AddChef("Chef A")
		menu1 = menuService.AddMenu("Menu A", 1)
	})

	Describe("AddOrder", func() {
		It("should add an order successfully", func() {
			order, err := orderService.AddOrder(menu1.ID, chef1.ID)
			Expect(err).To(BeNil())
			Expect(order).NotTo(BeNil())
			Expect(order.MenuID).To(Equal(menu1.ID))
			Expect(order.ChefID).To(Equal(chef1.ID))
		})

		It("should return an error if menu not found", func() {
			order, err := orderService.AddOrder(999, chef1.ID)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("menu not found"))
			Expect(order).To(BeNil())
		})

		It("should return an error if no chefs or menus available", func() {
			chefService = services.NewChefService()
			menuService = services.NewMenuService()
			orderService = services.NewOrderService(chefService, menuService)

			order, err := orderService.AddOrder(menu1.ID, chef1.ID)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("no chefs or menus available"))
			Expect(order).To(BeNil())
		})

		It("should return an error if chef is busy", func() {
			chef1.Status = "busy"
			order, err := orderService.AddOrder(menu1.ID, chef1.ID)
			Expect(err).To(HaveOccurred())
			Expect(err).To(MatchError("chef is busy"))
			Expect(order).To(BeNil())
		})
	})
})
