package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"money-api-transfer/api/commons"
	"money-api-transfer/api/entity"
	"money-api-transfer/api/repository"
	"net/http"
)

type ValidateAccountUsecase struct {
	repo repository.Repository
}

func NewValidateAccountUsecase(rp repository.Repository) ValidateAccountUsecase {
	return ValidateAccountUsecase{
		repo: rp,
	}
}

func (vau ValidateAccountUsecase) ValidateAccount(ctx context.Context, req entity.ValidateBankAccountReq) (res entity.ValidateBankAccountRes, code int, err error) {
	// set as array because response from Mock API is in array json form
	arrRes := []entity.ValidateBankAccountRes{}

	id, err := vau.repo.AccountRepo.GetAccountId(req.Account, req.Name)
	if err != nil {
		return res, http.StatusBadRequest, err
	}

	resp, err := http.Get(fmt.Sprintf(commons.AccountValidationUrl, id))
	if err != nil {

		return res, resp.StatusCode, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&arrRes); err != nil {
		err = errors.New(resp.Status)
		return res, http.StatusBadRequest, err
	}

	// get first element because mock API return in array format
	res = arrRes[0]

	return res, http.StatusOK, nil
}
