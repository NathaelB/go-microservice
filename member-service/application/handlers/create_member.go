package handlers

import (
	"member-service/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateMemberRequest struct {
	Name    string `json:"name" binding:"required"`
	GuildID string `json:"guild_id" binding:"required"`
}

type CreateMemberHandler struct {
	memberService domain.MemberService
}

func NewCreateMemberHandler(memberService domain.MemberService) *CreateMemberHandler {
	return &CreateMemberHandler{
		memberService: memberService,
	}
}

func (h *CreateMemberHandler) Handle(c *gin.Context) {
	var req CreateMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member, err := h.memberService.CreateMember(req.Name, req.GuildID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, member)
}
