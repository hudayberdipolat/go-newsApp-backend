package dto

import (
	"github.com/hudayberdipolat/go-newsApp-backend/internal/models"
)

type AllPostResponse struct {
	ID           int                  `json:"id"`
	PostTitle    string               `json:"post_title"`
	PostSlug     string               `json:"post_slug"`
	ClickCount   int                  `json:"click_count"`
	PostStatus   string               `json:"post_status"`
	CategoryID   int                  `json:"category_id"`
	ImageUrl     *string              `json:"image_url"`
	CreatedAt    string               `json:"created_at"`
	UpdatedAt    string               `json:"updated_at"`
	PostCategory postCategoryResponse `json:"post_category"`
	TagCount     int                  `json:"tag_count"`
}

func NewAllPostResponse(posts []models.Post) []AllPostResponse {
	var allPostResponses []AllPostResponse

	tagCount := 0
	for _, post := range posts {
		for i := 0; i < len(post.PostTags); i++ {
			tagCount = tagCount + 1
		}

		postResponse := AllPostResponse{
			ID:         post.ID,
			PostTitle:  post.PostTitle,
			PostSlug:   post.PostSlug,
			ClickCount: post.ClickCount,
			PostStatus: post.PostStatus,
			CategoryID: post.CategoryID,
			ImageUrl:   post.ImageUrl,
			CreatedAt:  post.CreatedAt.Format("01-02-2006"),
			UpdatedAt:  post.UpdatedAt.Format("01-02-2006"),
			TagCount:   tagCount,
			PostCategory: postCategoryResponse{
				Id:           post.Category.ID,
				CategoryName: post.Category.CategoryName,
				CategorySlug: post.Category.CategorySlug,
			},
		}
		allPostResponses = append(allPostResponses, postResponse)
	}
	return allPostResponses
}

type OnePostResponse struct {
	ID           int                  `json:"id"`
	PostTitle    string               `json:"post_title"`
	PostSlug     string               `json:"post_slug"`
	PostDesc     string               `json:"post_desc"`
	ClickCount   int                  `json:"click_count"`
	PostStatus   string               `json:"post_status"`
	CategoryID   int                  `json:"category_id"`
	ImageUrl     *string              `json:"image_url"`
	CreatedAt    string               `json:"created_at"`
	UpdatedAt    string               `json:"updated_at"`
	PostCategory postCategoryResponse `json:"post_category"`
	PostTags     []postTagResponse    `json:"post_tags"`
}

func NewOnePostResponse(post *models.Post) OnePostResponse {
	var postTagResponses []postTagResponse
	for _, postTag := range post.PostTags {
		onePostTagResponse := postTagResponse{
			Id:      postTag.ID,
			TagName: postTag.TagName,
		}
		postTagResponses = append(postTagResponses, onePostTagResponse)
	}
	return OnePostResponse{
		ID:         post.ID,
		PostTitle:  post.PostTitle,
		PostSlug:   post.PostSlug,
		PostDesc:   post.PostDesc,
		ClickCount: post.ClickCount,
		PostStatus: post.PostStatus,
		CategoryID: post.CategoryID,
		ImageUrl:   post.ImageUrl,
		CreatedAt:  post.CreatedAt.Format("01-02-2006"),
		UpdatedAt:  post.UpdatedAt.Format("01-02-2006"),
		PostCategory: postCategoryResponse{
			Id:           post.Category.ID,
			CategoryName: post.Category.CategoryName,
			CategorySlug: post.Category.CategorySlug,
		},
		PostTags: postTagResponses,
	}
}

type postCategoryResponse struct {
	Id           int    `json:"id"`
	CategoryName string `json:"category_name"`
	CategorySlug string `josn:"category_slug"`
}

type postTagResponse struct {
	Id      int    `json:"id"`
	TagName string `json:"tag_name"`
}

////////////////////////////
//------------------------
// frontend for responses

type GetAllPostsResponse struct {
	ID           int                  `json:"id"`
	PostTitle    string               `json:"post_title"`
	PostSlug     string               `json:"post_slug"`
	ClickCount   int                  `json:"click_count"`
	ImageUrl     *string              `json:"image_url"`
	CreatedAt    string               `json:"created_at"`
	PostCategory postCategoryResponse `json:"post_category"`
	LikeCount    int                  `json:"like_count"`
	DislikeCount int                  `json:"dislikeCount"`
}

func NewGetAllPostsResponse(posts []models.Post) []GetAllPostsResponse {
	var allPostResponses []GetAllPostsResponse
	for _, post := range posts {
		likeCount := 0
		dislikeCount := 0
		for _, like := range post.Liked {
			if like.LikeType == "like" {
				likeCount = likeCount + 1
			} else if like.LikeType == "dislike" {
				dislikeCount = dislikeCount + 1
			}
		}
		getPostResponse := GetAllPostsResponse{
			ID:         post.ID,
			PostTitle:  post.PostTitle,
			PostSlug:   post.PostSlug,
			ClickCount: post.ClickCount,
			ImageUrl:   post.ImageUrl,
			CreatedAt:  post.CreatedAt.Format("01-02-2006"),
			PostCategory: postCategoryResponse{
				Id:           post.Category.ID,
				CategoryName: post.Category.CategoryName,
				CategorySlug: post.Category.CategorySlug,
			},
			LikeCount:    likeCount,
			DislikeCount: dislikeCount,
		}
		allPostResponses = append(allPostResponses, getPostResponse)
	}
	return allPostResponses
}

// get one post response

type GetOnePostResponse struct {
	ID           int                  `json:"id"`
	PostTitle    string               `json:"post_title"`
	PostSlug     string               `json:"post_slug"`
	PostDesc     string               `json:"post_desc"`
	ClickCount   int                  `json:"click_count"`
	ImageUrl     *string              `json:"image_url"`
	CreatedAt    string               `json:"created_at"`
	PostCategory postCategoryResponse `json:"post_category"`
	PostTags     []postTagResponse    `json:"post_tags"`
	PostComments []postComment        `json:"post_comments"`
	LikeCount    int                  `json:"like_count"`
	DislikeCount int                  `json:"dislike_count"`
}

type postComment struct {
	ID          int    `json:"id"`
	WrittenUser string `json:"written_user"`
	UserComment string `json:"user_comment"`
	WriteTime   string `json:"write_time"`
}

// return get one post response

func NewGetOnePostResponse(post *models.Post) GetOnePostResponse {
	var postTagResponses []postTagResponse
	var postComments []postComment
	for _, postTag := range post.PostTags {
		onePostTagResponse := postTagResponse{
			Id:      postTag.ID,
			TagName: postTag.TagName,
		}
		postTagResponses = append(postTagResponses, onePostTagResponse)
	}

	for _, comment := range post.Comments {
		onePostComment := postComment{
			ID:          comment.ID,
			WrittenUser: "",
			UserComment: comment.PostComment,
			WriteTime:   comment.CreatedAt.Format("02.01.2006 "),
		}
		postComments = append(postComments, onePostComment)
	}

	likeCount := 0
	dislikeCount := 0
	for _, postLike := range post.Liked {
		if postLike.LikeType == "like" {
			likeCount = likeCount + 1
		} else if postLike.LikeType == "dislike" {
			dislikeCount = dislikeCount + 1
		}
	}

	return GetOnePostResponse{
		ID:         post.ID,
		PostTitle:  post.PostTitle,
		PostSlug:   post.PostSlug,
		PostDesc:   post.PostDesc,
		ClickCount: post.ClickCount,
		ImageUrl:   post.ImageUrl,
		CreatedAt:  post.CreatedAt.Format("01-02-2006"),
		PostCategory: postCategoryResponse{
			Id:           post.Category.ID,
			CategoryName: post.Category.CategoryName,
			CategorySlug: post.Category.CategorySlug,
		},
		PostTags:     postTagResponses,
		PostComments: postComments,
		LikeCount:    likeCount,
		DislikeCount: dislikeCount,
	}
}
