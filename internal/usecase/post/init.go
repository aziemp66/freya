package post

import (
	"context"

	errorCommon "github.com/aziemp66/freya-be/common/error"
	httpCommon "github.com/aziemp66/freya-be/common/http"
	PostDomain "github.com/aziemp66/freya-be/internal/domain/post"
	PostRepository "github.com/aziemp66/freya-be/internal/repository/post"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostUsecaseImplementation struct {
	postRepository PostRepository.Repository
}

func NewPostUsecaseImplementation(postRepository PostRepository.Repository) *PostUsecaseImplementation {
	return &PostUsecaseImplementation{postRepository}
}

func (p *PostUsecaseImplementation) InsertPost(ctx context.Context, authorId, title, content string) (err error) {
	objId, err := primitive.ObjectIDFromHex(authorId)

	if err != nil {
		return errorCommon.NewInvariantError("Invalid author id")
	}

	err = p.postRepository.InsertPost(ctx, PostDomain.Post{
		AuthorId: objId,
		Title:    title,
		Content:  content,
	})

	if err != nil {
		return err
	}

	return nil
}

func (p *PostUsecaseImplementation) GetPostById(ctx context.Context, id string) (post httpCommon.Post, err error) {
	postObj, err := p.postRepository.FindPostByID(ctx, id)

	if err != nil {
		return post, err
	}

	post = httpCommon.Post{
		ID:       postObj.ID.Hex(),
		AuthorID: postObj.AuthorId.Hex(),
		Title:    postObj.Title,
		Content:  postObj.Content,
	}

	return
}

func (p *PostUsecaseImplementation) GetAllPost(ctx context.Context) (posts []httpCommon.Post, err error) {
	postsObj, err := p.postRepository.FindAllPosts(ctx)

	if err != nil {
		return posts, err
	}

	for _, postObj := range postsObj {
		posts = append(posts, httpCommon.Post{
			ID:       postObj.ID.Hex(),
			AuthorID: postObj.AuthorId.Hex(),
			Title:    postObj.Title,
			Content:  postObj.Content,
		})
	}

	return
}

func (p *PostUsecaseImplementation) DeletePost(ctx context.Context, id string) (err error) {
	err = p.postRepository.DeletePost(ctx, id)

	if err != nil {
		return err
	}

	return nil
}

func (p *PostUsecaseImplementation) InsertComment(ctx context.Context, authorId, postId, content string) (err error) {
	authObjId, err := primitive.ObjectIDFromHex(authorId)

	if err != nil {
		return errorCommon.NewInvariantError("Invalid author id")
	}

	postObjId, err := primitive.ObjectIDFromHex(postId)

	if err != nil {
		return errorCommon.NewInvariantError("Invalid post id")
	}

	err = p.postRepository.InsertComment(ctx, PostDomain.Comment{
		AuthorId: authObjId,
		PostId:   postObjId,
		Content:  content,
	})

	if err != nil {
		return err
	}

	return nil
}

func (p *PostUsecaseImplementation) GetAllCommentByPostId(ctx context.Context, postId string) (comments []httpCommon.Comment, err error) {
	commentsObj, err := p.postRepository.FindAllCommentsByPostID(ctx, postId)

	if err != nil {
		return comments, err
	}

	for _, commentObj := range commentsObj {
		comments = append(comments, httpCommon.Comment{
			ID:       commentObj.ID.Hex(),
			AuthorID: commentObj.AuthorId.Hex(),
			PostID:   commentObj.PostId.Hex(),
			Content:  commentObj.Content,
		})
	}

	return
}

func (p *PostUsecaseImplementation) GetCommentById(ctx context.Context, id string) (comment httpCommon.Comment, err error) {
	commentObj, err := p.postRepository.FindCommentByID(ctx, id)

	if err != nil {
		return comment, err
	}

	comment = httpCommon.Comment{
		ID:       commentObj.ID.Hex(),
		AuthorID: commentObj.AuthorId.Hex(),
		PostID:   commentObj.PostId.Hex(),
		Content:  commentObj.Content,
	}

	return
}

func (p *PostUsecaseImplementation) DeleteComment(ctx context.Context, id string) (err error) {
	err = p.postRepository.DeleteComment(ctx, id)

	if err != nil {
		return err
	}

	return nil
}
