package entity

type PostRepository interface {
	CreatePost(post Post) error
	GetPostList() ([]Post, error)
	GetPost(id string) (Post, error)
	UpdatePost(id string, post Post) error
	DeletePost(id string) error
	CreateComment(postId string, comment Comment) error
	GetCommentList(postId string) ([]Comment, error)
	DeletePostComment(postId string, commentID string) error
}

type Blog struct {
	PostRepository PostRepository
}

type PostListDto struct {
	Posts []Post `json:"posts"`
}

func NewBlog(postRepository PostRepository) Blog {
	return Blog{
		PostRepository: postRepository,
	}
}

func (b *Blog) CreatePost(post Post) error {
	return b.PostRepository.CreatePost(post)
}

func (b *Blog) GetPostList() (PostListDto, error) {
	postList := PostListDto{}
	response, err := b.PostRepository.GetPostList()
	if err != nil {
		return PostListDto{}, err
	}
	postList.Posts = response
	return postList, nil
}

func (b *Blog) GetPost(id string) (Post, error) {
	return b.PostRepository.GetPost(id)
}

func (b *Blog) UpdatePost(id string, post Post) error {
	return b.PostRepository.UpdatePost(id, post)
}

func (b *Blog) DeletePost(id string) error {
	return b.PostRepository.DeletePost(id)
}

func (b *Blog) CreateComment(postId string, comment Comment) error {
	return b.PostRepository.CreateComment(postId, comment)
}

func (b *Blog) GetCommentList(postID string) ([]Comment, error) {
	return b.PostRepository.GetCommentList(postID)
}

func (b *Blog) DeletePostComment(postID string, commentID string) error {
	return b.PostRepository.DeletePostComment(postID, commentID)
}
