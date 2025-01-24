package services_test

import (
	"errors"
	"pizza-hub/internal/services"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ChefService", func() {
	var service *services.ChefService

	BeforeEach(func() {
		service = services.NewChefService()
	})

	Describe("AddChef", func() {
		It("should add a new chef", func() {
			chef := service.AddChef("Chef A")
			Expect(chef).NotTo(BeNil())
			Expect(chef.Name).To(Equal("Chef A"))
			Expect(chef.Status).To(Equal("available"))
			Expect(service.GetChefs()).To(HaveLen(1))
		})
	})

	Describe("FindChefByID", func() {
		It("should find a chef by ID", func() {
			chef := service.AddChef("Chef A")
			foundChef, err := service.FindChefByID(chef.ID)
			Expect(err).To(BeNil())
			Expect(foundChef).To(Equal(chef))
		})

		It("should return an error if chef is not found", func() {
			_, err := service.FindChefByID(999)
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(errors.New("chef not found")))
		})
	})

	Describe("FindChefsByStatus", func() {
		BeforeEach(func() {
			service.AddChef("Chef A")
			service.AddChef("Chef B")
			service.AddChef("Chef C")
			service.GetChefs()[1].Status = "busy"
			service.GetChefs()[2].Status = "busy"
		})

		It("should find chefs by status", func() {
			chefs, err := service.FindChefsByStatus("available")
			Expect(err).To(BeNil())
			Expect(chefs).To(HaveLen(1))

			chefs, err = service.FindChefsByStatus("busy")
			Expect(err).To(BeNil())
			Expect(chefs).To(HaveLen(2))
		})

		It("should return an error if no chefs are available", func() {
			_, err := service.FindChefsByStatus("on break")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(errors.New("no chefs available")))
		})
	})
})
