package backend

import (
	"errors"
	"testing"

	"github.com/sensu/sensu-go/testing/mockstore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSeedDefaultRole(t *testing.T) {
	store := &mockstore.MockStore{}
	store.On("UpdateRole", mock.AnythingOfType("*types.Role")).Return(nil)
	store.On("UpdateOrganization", mock.AnythingOfType("*types.Organization")).Return(nil)

	seedInitialData(store)
	store.AssertCalled(t, "UpdateRole", mock.AnythingOfType("*types.Role"))
}

func TestSeedDefaultRoleWithError(t *testing.T) {
	assert := assert.New(t)
	store := &mockstore.MockStore{}

	store.On("UpdateRole", mock.AnythingOfType("*types.Role")).Return(errors.New(""))
	assert.Error(seedInitialData(store))
}
