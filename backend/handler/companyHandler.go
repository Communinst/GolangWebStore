package handler

import (
	"net/http"
	"strconv"

	entities "github.com/Communinst/GolangWebStore/backend/entity"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createCompany(c *gin.Context) {
	var company entities.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.service.CompanyServiceInterface.PostCompany(c.Request.Context(), &company); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company created successfully"})
}

func (h *Handler) getCompany(c *gin.Context) {
	companyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	company, err := h.service.CompanyServiceInterface.GetCompany(c.Request.Context(), companyId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *Handler) getAllCompanies(c *gin.Context) {
	companies, err := h.service.CompanyServiceInterface.GetAllCompanies(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve companies"})
		return
	}

	c.JSON(http.StatusOK, companies)
}

func (h *Handler) deleteCompany(c *gin.Context) {
	companyId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid company ID"})
		return
	}

	if err := h.service.CompanyServiceInterface.DeleteCompany(c.Request.Context(), companyId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}

func (h *Handler) getCompanyByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company name is required"})
		return
	}

	company, err := h.service.CompanyServiceInterface.GetCompanyByName(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	c.JSON(http.StatusOK, company)
}

func (h *Handler) deleteCompanyByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Company name is required"})
		return
	}

	if err := h.service.CompanyServiceInterface.DeleteCompanyByName(c.Request.Context(), name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete company"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Company deleted successfully"})
}
