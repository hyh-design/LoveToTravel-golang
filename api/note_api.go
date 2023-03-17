package api

import (
	"github.com/gin-gonic/gin"
	"ltt-gc/model/vo"
	"ltt-gc/service"
)

func GetNoteById(c *gin.Context) {
	noteService := service.NoteService{}
	res := noteService.GetNoteById(c.Param("id"))
	c.JSON(200, res)
}

func GetNoteList(c *gin.Context) {
	noteService := service.NoteService{}
	res := noteService.GetNoteList()
	c.JSON(200, res)
}

func GetNotePage(c *gin.Context) {
	noteService := service.NoteService{}
	p := vo.Page{}
	if err := c.ShouldBind(&p); err == nil {
		res := noteService.GetNotePage(p)
		c.JSON(200, res)
	}
}

func GetNotePageFuzzy(c *gin.Context) {
	noteService := service.NoteService{}
	p := vo.Page{}
	if err := c.ShouldBind(&p); err == nil {
		res := noteService.GetNotePageFuzzy(p)
		c.JSON(200, res)
	}
}

func CreateNote(c *gin.Context) {
	noteService := service.NoteService{}
	if err := c.ShouldBind(&noteService); err == nil {
		res := noteService.CreateNote()
		c.JSON(200, res)
	}
}

func UpdateNote(c *gin.Context) {
	noteService := service.NoteService{}
	if err := c.ShouldBind(&noteService); err == nil {
		res := noteService.UpdateNote()
		c.JSON(200, res)
	}
}

func DeleteNoteById(c *gin.Context) {
	noteService := service.NoteService{}
	res := noteService.DeleteNoteById(c.Param("id"))
	c.JSON(200, res)
}
