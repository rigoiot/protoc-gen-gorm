// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/feature_demo/demo_multi_file_service.proto

package example

import context "context"

import gorm1 "github.com/jinzhu/gorm"

import fmt "fmt"
import math "math"
import _ "github.com/rigoiot/atlas-app-toolkit/query"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type BlogPostServiceDefaultServer struct {
	DB *gorm1.DB
}

// Read ...
func (m *BlogPostServiceDefaultServer) Read(ctx context.Context, in *ReadAccountRequest) (*ReadBlogPostsResponse, error) {
	return &ReadBlogPostsResponse{}, nil
}
