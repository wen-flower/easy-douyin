// Code generated by Validator v0.1.4. DO NOT EDIT.

package video

import (
	"bytes"
	"fmt"
	"reflect"
	"regexp"
	"strings"
	"time"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = (*regexp.Regexp)(nil)
	_ = time.Nanosecond
)

func (p *CreateVideoParam) IsValid() error {
	return nil
}
func (p *CreateVideoResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *PublishListParam) IsValid() error {
	return nil
}
func (p *PublishListResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *VideoFeedParam) IsValid() error {
	return nil
}
func (p *VideoFeedResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *FavoriteVideoParam) IsValid() error {
	return nil
}
func (p *FavoriteVideoResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *FavoriteListParam) IsValid() error {
	return nil
}
func (p *FavoriteListResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *CommentVideoParam) IsValid() error {
	return nil
}
func (p *CommentVideoResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	if p.Comment != nil {
		if err := p.Comment.IsValid(); err != nil {
			return fmt.Errorf("filed Comment not valid, %w", err)
		}
	}
	return nil
}
func (p *DeleteCommentParam) IsValid() error {
	return nil
}
func (p *DeleteCommentResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
func (p *CommentListParam) IsValid() error {
	return nil
}
func (p *CommentListResp) IsValid() error {
	if p.BaseResp != nil {
		if err := p.BaseResp.IsValid(); err != nil {
			return fmt.Errorf("filed BaseResp not valid, %w", err)
		}
	}
	return nil
}
