package http_handler

import (
	"hacktiv8-msib-final-project-1/dto"
	"hacktiv8-msib-final-project-1/pkg/errs"
	"hacktiv8-msib-final-project-1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *todoHandler {
	return &todoHandler{todoService: todoService}
}

// CreateTodo godoc
//
//	@Summary		Create a todo
//	@Description	Create a todo by json
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			todo	body		dto.NewTodoRequest	true	"Create todo request body"
//	@Success		201		{object}	dto.NewTodoResponse
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/todos [post]
func (t *todoHandler) CreateTodo(ctx *gin.Context) {
	var requestBody dto.NewTodoRequest

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		newError := errs.NewUnprocessableEntity(err.Error())
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdTodo, err := t.todoService.CreateTodo(&requestBody)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	ctx.JSON(http.StatusCreated, createdTodo)
}
