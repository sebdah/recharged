package types_test

import (
	"testing"

	"github.com/sebdah/recharged/central-system/types"
	"github.com/stretchr/testify/assert"
)

func TestNewIdTagInfo(t *testing.T) {
	idTagInfo := types.NewIdTagInfo()
	assert.Equal(t, "", idTagInfo.ExpiryDate)
	assert.Equal(t, "en", idTagInfo.Language)
	assert.Equal(t, types.AuthorizationStatusInvalid, idTagInfo.Status)
	assert.Nil(t, idTagInfo.GroupTagId)

	return
}
