package handlers

import (
	"member-service/domain"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ListMembersHandler struct {
	memberService domain.MemberService
}

func NewListMembersHandler(memberService domain.MemberService) *ListMembersHandler {
	return &ListMembersHandler{
		memberService: memberService,
	}
}

func (h *ListMembersHandler) Handle(c *gin.Context) {
	members, err := h.memberService.GetAllMembers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}
