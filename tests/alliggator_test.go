package tests_test

import (
	"Alliggator"
	// "log"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("Alliggator", func() {

	all := alliggator.New()
	var jsonString string

	BeforeEach(func() {
		all = alliggator.New()
		jsonString = `[{"$match":{"domain": "carrierexpress.com.br"}},{"$project": {"_id": 0,"domain": 1,"ipPort": 1}},{"$sort": {"ipPort": 1}},{"$limit": 17},{"$skip": 7}]`
	})

	Describe("Getting data type from alliggator obj (all)", func() {
		Context("Instance from alliggator", func() {
			It("should be *alliggator.alliggator", func() {
				Expect(reflect.TypeOf(all).String()).To(Equal("*alliggator.alliggator"))
			})
		})
	})

	Describe("Getting return from all.TreatDollarSign", func() {
		Context("Sending a json with mongodb aggregation format to get a json without dollar signs", func() {
			It("should be the same json string, but without dollar signs", func() {
				Expect(all.TreatDollarSign(jsonString)).To(Equal(`[{"match":{"domain": "carrierexpress.com.br"}},{"project": {"_id": 0,"domain": 1,"ipPort": 1}},{"sort": {"ipPort": 1}},{"limit": 17},{"skip": 7}]`))
			})
		})
	})

	Describe("Getting return from all.TreatDollarSign", func() {
		Context("Sending a json with mongodb aggregation format to get a json without dollar signs", func() {
			It("should be the same json string, but without dollar signs", func() {
				all.FromString(jsonString)
			})
		})
	})
})
