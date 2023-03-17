package post

import (
	errorCommon "github.com/aziemp66/freya-be/common/error"
	httpCommon "github.com/aziemp66/freya-be/common/http"
	"github.com/aziemp66/freya-be/common/http/middleware"
	"github.com/aziemp66/freya-be/common/jwt"
	PostUsecase "github.com/aziemp66/freya-be/internal/usecase/post"
	UserUseCase "github.com/aziemp66/freya-be/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type PostDeliveryImplementation struct {
	postUsecase PostUsecase.Usecase
	userUsecase UserUseCase.Usecase
	jwtManager  *jwt.JWTManager
}

func NewPostDeliveryImplementation(router *gin.RouterGroup, postUsecase PostUsecase.Usecase, userUsecase UserUseCase.Usecase, jwtManager *jwt.JWTManager) *PostDeliveryImplementation {
	postDelivery := &PostDeliveryImplementation{
		postUsecase: postUsecase,
		userUsecase: userUsecase,
		jwtManager:  jwtManager,
	}

	router.GET("/", postDelivery.GetAllPost)
	router.GET("/:id", postDelivery.GetPostById)
	router.GET("/:id/comment", postDelivery.GetAllCommentByPostId)

	authGroup := router.Group("/", middleware.JWTAuth(jwtManager), middleware.RoleAuth("base"))
	authGroup.POST("/", postDelivery.CreatePost)
	authGroup.DELETE("/:id", postDelivery.DeletePost)
	authGroup.POST("/:id/comment", postDelivery.CreateComment)
	authGroup.DELETE("/comment/:commentId", postDelivery.DeleteComment)

	return postDelivery
}

func (p *PostDeliveryImplementation) CreatePost(c *gin.Context) {
	var postRequest httpCommon.AddPost

	if err := c.ShouldBindJSON(&postRequest); err != nil {
		return
	}

	authorId := c.GetString("user_id")

	err := p.postUsecase.InsertPost(c, authorId, postRequest.Title, postRequest.Content)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Create post success",
	})
}

func (p *PostDeliveryImplementation) GetAllPost(c *gin.Context) {
	posts, err := p.postUsecase.GetAllPost(c)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Get all post success",
		Value:   gin.H{"posts": posts},
	})
}

func (p *PostDeliveryImplementation) GetPostById(c *gin.Context) {
	postId := c.Param("id")

	post, err := p.postUsecase.GetPostById(c, postId)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Get post by id success",
		Value:   gin.H{"post": post},
	})
}

func (p *PostDeliveryImplementation) DeletePost(c *gin.Context) {
	postId := c.Param("id")

	postData, err := p.postUsecase.GetPostById(c, postId)

	if err != nil {
		c.Error(err)
		return
	}

	userID := c.GetString("user_id")

	if postData.AuthorID != userID {
		c.Error(errorCommon.NewUnauthorizedError("You are not authorized to delete this post"))
		return
	}

	err = p.postUsecase.DeletePost(c, postId)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Delete post success",
	})
}

func (p *PostDeliveryImplementation) CreateComment(c *gin.Context) {
	var commentRequest httpCommon.AddComment

	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		return
	}

	authorId := c.GetString("user_id")
	postId := c.Param("id")

	err := p.postUsecase.InsertComment(c, authorId, postId, commentRequest.Content)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Create comment success",
	})
}

func (p *PostDeliveryImplementation) GetAllCommentByPostId(c *gin.Context) {
	postId := c.Param("id")

	comments, err := p.postUsecase.GetAllCommentByPostId(c, postId)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Get all comment by post id success",
		Value:   gin.H{"comments": comments},
	})
}

func (p *PostDeliveryImplementation) DeleteComment(c *gin.Context) {
	commentId := c.Param("commentId")

	commentData, err := p.postUsecase.GetCommentById(c, commentId)

	if err != nil {
		c.Error(err)
		return
	}

	userID := c.GetString("user_id")

	if commentData.AuthorID != userID {
		c.Error(errorCommon.NewUnauthorizedError("You are not authorized to delete this comment"))
		return
	}

	err = p.postUsecase.DeleteComment(c, commentId)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, httpCommon.Response{
		Code:    200,
		Message: "Delete comment success",
	})
}
