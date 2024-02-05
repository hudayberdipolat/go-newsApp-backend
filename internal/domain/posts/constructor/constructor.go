package constructor

import (
	handler2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/handler"
	repository2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/repository"
	service2 "github.com/hudayberdipolat/go-newsApp-backend/internal/domain/posts/service"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
	"gorm.io/gorm"
)

var postRepo repository2.PostRepository
var postService service2.PostService
var PostHandler handler2.PostHandler

func PostRequirementCreator(db *gorm.DB, config *config.Config) {
	postRepo = repository2.NewPostRepository(db)
	postService = service2.NewPostService(postRepo)
	PostHandler = handler2.NewPostHandler(postService, *config)
}
