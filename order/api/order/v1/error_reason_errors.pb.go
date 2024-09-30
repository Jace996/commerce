// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	context "context"
	errors "github.com/go-kratos/kratos/v2/errors"
	i18n "github.com/go-saas/go-i18n/v2/i18n"
	localize "github.com/go-saas/kit/pkg/localize"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsContentMissing(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_CONTENT_MISSING.String() && e.Code == 400
}

func ErrorContentMissingLocalized(ctx context.Context, data map[string]interface{}, pluralCount interface{}) *errors.Error {
	localizer := localize.FromContext(ctx)
	if localizer == nil {
		return errors.New(400, ErrorReason_CONTENT_MISSING.String(), "")
	}
	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID: "ContentMissing",
		},
		TemplateData: data,
		PluralCount:  pluralCount,
	})
	if err == nil {
		return errors.New(400, ErrorReason_CONTENT_MISSING.String(), msg)
	} else {
		return errors.New(400, ErrorReason_CONTENT_MISSING.String(), "")
	}
}
