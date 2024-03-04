package account

import (
	"encoding/json"

	"github.com/achwanyusuf/carrent-accountsvc/src/model"
	"github.com/achwanyusuf/carrent-accountsvc/src/model/psqlmodel"
	"github.com/gin-gonic/gin"
)

func (a *AccountDep) getSingleByParamRedis(ctx *gin.Context, key string) (psqlmodel.Account, error) {
	var res psqlmodel.Account
	data, err := a.Redis.Get(ctx, key).Result()
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (a *AccountDep) setRedis(ctx *gin.Context, key string, data string) error {
	expTime := a.Conf.RedisExpirationTime
	if a.Conf.RedisExpirationTime == 0 {
		expTime = model.DefaultRedisExpiration
	}
	_, err := a.Redis.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	_, err = a.Redis.Set(ctx, key, data, expTime).Result()
	return err
}

func (a *AccountDep) getByParamRedis(ctx *gin.Context, key string) (psqlmodel.AccountSlice, error) {
	var res psqlmodel.AccountSlice
	data, err := a.Redis.Get(ctx, key).Result()
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (a *AccountDep) getByParamPaginationRedis(ctx *gin.Context, key string) (model.Pagination, error) {
	var res model.Pagination
	data, err := a.Redis.Get(ctx, key).Result()
	if err != nil {
		return res, err
	}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}
