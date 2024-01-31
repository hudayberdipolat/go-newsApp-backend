package constructor

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/handler"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/repository"
	"github.com/hudayberdipolat/go-newsApp-backend/internal/domain/adminPanel/posts/service"
	"github.com/hudayberdipolat/go-newsApp-backend/pkg/config"
	"gorm.io/gorm"
)

var postRepo repository.PostRepository
var postService service.PostService
var PostHandler handler.PostHandler

func PostRequirementCreator(db *gorm.DB, config *config.Config) {
	postRepo = repository.NewPostRepository(db)
	postService = service.NewPostService(postRepo)
	PostHandler = handler.NewPostHandler(postService, *config)
}
