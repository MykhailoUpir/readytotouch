package cac

import (
	"net/http"
	"strings"
	"time"

	"github.com/readytotouch/readytotouch/internal/db/postgres"
	"github.com/readytotouch/readytotouch/internal/domain"
	template "github.com/readytotouch/readytotouch/internal/templates/v1"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	userToLinkedInCompanyRepository *postgres.UserToLinkedInCompanyRepository
	service                         *Service
}

func NewController(userToLinkedInCompanyRepository *postgres.UserToLinkedInCompanyRepository, service *Service) *Controller {
	return &Controller{userToLinkedInCompanyRepository: userToLinkedInCompanyRepository, service: service}
}

func (c *Controller) Index(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(template.WipCompaniesAndConnections()))
}

func (c *Controller) Ukraine(ctx *gin.Context) {
	c.Index(ctx)
}

func (c *Controller) Czechia(ctx *gin.Context) {
	c.Index(ctx)
}

func (c *Controller) Poland(ctx *gin.Context) {
	c.Index(ctx)
}

func (c *Controller) Companies(ctx *gin.Context) {
	var (
		authUserID = domain.ContextGetUserID(ctx)
	)
	if authUserID == 0 {
		// For the sake of the demo, we're going to return a list of companies
		ctx.JSON(http.StatusOK, []domain.LinkedInProfileResponse{
			{
				ID:    14136494,
				Alias: "dochq",
				Name:  "DocHQ",
			},
			{
				ID:    1441,
				Alias: "google",
				Name:  "Google",
			},
			{
				ID:    1035,
				Alias: "microsoft",
				Name:  "Microsoft",
			},
		})

		return
	}

	c.companies(ctx, authUserID)
}

func (c *Controller) AddCompany(ctx *gin.Context) {
	type addRequestBody struct {
		Alias string `json:"alias"`
	}

	var (
		authUserID = domain.ContextGetUserID(ctx)
	)
	if authUserID == 0 {
		ctx.JSON(http.StatusUnauthorized, &domain.ErrorResponse{
			ErrorMessage: "Unauthorized",
		})
		return
	}

	var body addRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	var companyVanityName = strings.ToLower(strings.TrimSpace(body.Alias))
	if companyVanityName == "" {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: "Alias is required",
		})
		return
	}

	var now = time.Now().UTC()

	linkedinCompanyID, err := c.service.Load(ctx, companyVanityName, authUserID, now)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	err = c.userToLinkedInCompanyRepository.Add(ctx, authUserID, linkedinCompanyID, now)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	c.companies(ctx, authUserID)
}

func (c *Controller) DeleteCompany(ctx *gin.Context) {
	type deleteRequestBody struct {
		ID int64 `json:"id"`
	}

	var (
		authUserID = domain.ContextGetUserID(ctx)
	)
	if authUserID == 0 {
		ctx.JSON(http.StatusUnauthorized, &domain.ErrorResponse{
			ErrorMessage: "Unauthorized",
		})
		return
	}

	var body deleteRequestBody
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}
	if body.ID == 0 {
		ctx.JSON(http.StatusBadRequest, &domain.ErrorResponse{
			ErrorMessage: "ID is required",
		})
		return
	}

	var now = time.Now().UTC()

	err := c.userToLinkedInCompanyRepository.Delete(ctx, authUserID, body.ID, now)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (c *Controller) companies(ctx *gin.Context, authUserID int64) {
	companies, err := c.userToLinkedInCompanyRepository.List(ctx, authUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &domain.ErrorResponse{
			ErrorMessage: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, companies)
}