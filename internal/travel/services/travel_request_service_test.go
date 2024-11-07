package services_test

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/dtos"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/models"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/infra/repositories/mocks"
	"github.com/rodrigoPQF/travel-request-backend/internal/travel/services"
	"github.com/stretchr/testify/assert"
)

func TestNewTravelRequestService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTravelRequestRepository(ctrl)
	service := services.NewTravelRequestService(mockRepo)

	assert.NotNil(t, service)
	assert.Implements(t, (*services.TravelRequestService)(nil), service)
}

func TestCreateTravelRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTravelRequestRepository(ctrl)
	service := services.NewTravelRequestService(mockRepo)
	ctx := context.Background()

	t.Run("should return error if findById fails", func(t *testing.T) {
		input := &dtos.CreateTravelRequestInputDto{Id: "unique-id"}
		mockRepo.EXPECT().FindById(ctx, input.Id).Return(nil, errors.New("database error"))

		output, err := service.Create(ctx, input)

		assert.Nil(t, output)
		assert.Error(t, err)
		assert.Equal(t, "database error", err.Error())
	})

	t.Run("should return error if travel request already exists", func(t *testing.T) {
		genUuid := uuid.New()
		existingRequest := &models.TravelRequest{Id: genUuid}
		input := &dtos.CreateTravelRequestInputDto{Id: genUuid.String()}

		mockRepo.EXPECT().FindById(ctx, input.Id).Return(existingRequest, nil)

		output, err := service.Create(ctx, input)

		assert.Nil(t, output)
		assert.Error(t, err)
		assert.Equal(t, "error travel request already exist", err.Error())
	})

	t.Run("should create travel request successfully", func(t *testing.T) {
		input := &dtos.CreateTravelRequestInputDto{
			Id:            "non_id",
			ApplicantName: "Rodrigo",
			Destination:   "Paris",
			DepartureDate: "2024-12-01",
			ReturnDate:    "2024-12-15",
		}

		mockRepo.EXPECT().FindById(ctx, input.Id).Return(nil, nil)
		mockRepo.EXPECT().Create(ctx, gomock.Any()).Return(nil)

		output, err := service.Create(ctx, input)

		assert.NoError(t, err)
		assert.NotNil(t, output)
	})
}

func TestGetAllTravelRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTravelRequestRepository(ctrl)
	service := services.NewTravelRequestService(mockRepo)
	ctx := context.Background()

	t.Run("should find all travel requests successfully", func(t *testing.T) {
		input := &dtos.GetAllTravelRequestInputDto{Status: "REQUESTED", Page: 1, Limit: 10}
		genUuid := uuid.New()
		expectedResult := []*models.TravelRequest{
			{
				Id:            genUuid,
				ApplicantName: "Rodrigo",
				Destination:   "Paris",
				Status:        "REQUESTED",
			},
			{
				Id:            genUuid,
				ApplicantName: "Pereira",
				Destination:   "New York",
				Status:        "REQUESTED",
			},
		}

		mockRepo.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(expectedResult, nil).Times(1)

		output, err := service.GetAll(ctx, input)

		assert.NoError(t, err)
		assert.Len(t, output, len(expectedResult))
		assert.Equal(t, output[0].Id, expectedResult[0].Id.String())
		assert.Equal(t, output[1].ApplicantName, expectedResult[1].ApplicantName)
	})

	t.Run("should return error if no travel requests found", func(t *testing.T) {
		input := &dtos.GetAllTravelRequestInputDto{Status: "REQUESTED", Page: 1, Limit: 10}
		mockRepo.EXPECT().Find(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, fmt.Errorf("database error"))

		output, err := service.GetAll(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
	})
}

func TestGetTravelRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTravelRequestRepository(ctrl)
	service := services.NewTravelRequestService(mockRepo)
	ctx := context.Background()

	t.Run("should find travel request successfully", func(t *testing.T) {
		genUuid := uuid.New()
		input := &dtos.GetTravelRequestInputDto{Id: genUuid.String()}
		expectedResult := &models.TravelRequest{
			Id:            genUuid,
			ApplicantName: "Rodrigo",
			Destination:   "Paris",
			Status:        models.Requested,
		}

		mockRepo.EXPECT().FindById(gomock.Any(), input.Id).Return(expectedResult, nil)

		output, err := service.Get(ctx, input)

		assert.NoError(t, err)
		assert.Equal(t, output.Id, expectedResult.Id.String())
		assert.Equal(t, output.ApplicantName, expectedResult.ApplicantName)
		assert.Equal(t, output.Destination, expectedResult.Destination)
		assert.Equal(t, output.Status, string(expectedResult.Status))
	})

	t.Run("should return error if findById fails", func(t *testing.T) {
		input := &dtos.GetTravelRequestInputDto{Id: "non_id"}
		mockRepo.EXPECT().FindById(gomock.Any(), input.Id).Return(nil, fmt.Errorf("travel request not found"))

		output, err := service.Get(ctx, input)

		assert.Error(t, err)
		assert.Nil(t, output)
	})
}

func TestUpdateTravelRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTravelRequestRepository(ctrl)
	service := services.NewTravelRequestService(mockRepo)
	ctx := context.Background()

	t.Run("should update travel request successfully", func(t *testing.T) {
		genUuid := uuid.New()
		input := &dtos.UpdateTravelRequestInputDto{
			UpdateTravelRequestBodyInputDto: dtos.UpdateTravelRequestBodyInputDto{
				Status: "APPROVED",
			},
			UpdateTravelRequestParamsInputDto: dtos.UpdateTravelRequestParamsInputDto{
				Id: genUuid.String(),
			},
		}

		expectedResult := &models.TravelRequest{
			Id:            genUuid,
			ApplicantName: "Rodrigo",
			Destination:   "Paris",
			Status:        models.Requested,
		}

		mockRepo.EXPECT().FindById(gomock.Any(), input.Id).Return(expectedResult, nil)
		mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil)

		err := service.Update(ctx, input)

		assert.NoError(t, err)
	})

	t.Run("should return error if travel request not found", func(t *testing.T) {
		input := &dtos.UpdateTravelRequestInputDto{
			UpdateTravelRequestBodyInputDto: dtos.UpdateTravelRequestBodyInputDto{
				Status: "APPROVED",
			},
			UpdateTravelRequestParamsInputDto: dtos.UpdateTravelRequestParamsInputDto{
				Id: "non_id",
			},
		}

		mockRepo.EXPECT().FindById(gomock.Any(), input.Id).Return(nil, nil)

		err := service.Update(ctx, input)

		assert.Error(t, err)
		assert.Equal(t, "error travel request not exist", err.Error())
	})

	t.Run("should return error if repository update fails", func(t *testing.T) {
		genUuid := uuid.New()
		input := &dtos.UpdateTravelRequestInputDto{
			UpdateTravelRequestBodyInputDto: dtos.UpdateTravelRequestBodyInputDto{
				Status: "APPROVED",
			},
			UpdateTravelRequestParamsInputDto: dtos.UpdateTravelRequestParamsInputDto{
				Id: genUuid.String(),
			},
		}

		expectedResult := &models.TravelRequest{
			Id:            genUuid,
			Status:        models.Requested,
		}

		mockRepo.EXPECT().FindById(gomock.Any(), input.Id).Return(expectedResult, nil)
		mockRepo.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors.New("repository error"))

		err := service.Update(ctx, input)

		assert.Error(t, err)
		assert.Equal(t, "repository error", err.Error())
	})
}
