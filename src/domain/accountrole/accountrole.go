package accountrole

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/achwanyusuf/carrent-accountsvc/src/model"
	"github.com/achwanyusuf/carrent-accountsvc/src/model/psqlmodel"
	"github.com/achwanyusuf/carrent-accountsvc/src/model/svcerr"
	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/achwanyusuf/carrent-lib/pkg/logger"

	"github.com/gin-gonic/gin"
	goredislib "github.com/redis/go-redis/v9"
)

type AccountRoleDep struct {
	Log   logger.Logger
	DB    *sql.DB
	Redis *goredislib.Client
	Conf  Conf
}

type Conf struct {
	DefaultPageLimit    int           `mapstructure:"page_limit"`
	RedisExpirationTime time.Duration `mapstructure:"expiration_time"`
}

type AccountRoleInterface interface {
	Insert(ctx *gin.Context, data *psqlmodel.AccountRole) error
	GetSingleByParam(ctx *gin.Context, cacheControl string, param *model.GetAccountRoleByParam) (psqlmodel.AccountRole, error)
	Update(ctx *gin.Context, AccountRole *psqlmodel.AccountRole) error
	Delete(ctx *gin.Context, AccountRole *psqlmodel.AccountRole, id int64, isHardDelete bool) error
	GetByParam(ctx *gin.Context, cacheControl string, param *model.GetAccountRolesByParam) (psqlmodel.AccountRoleSlice, model.Pagination, error)
}

func New(conf Conf, log *logger.Logger, db *sql.DB, rds *goredislib.Client) AccountRoleInterface {
	return &AccountRoleDep{
		Log:   *log,
		DB:    db,
		Redis: rds,
		Conf:  conf,
	}
}

func (a *AccountRoleDep) Insert(ctx *gin.Context, data *psqlmodel.AccountRole) error {
	return a.insertPSQL(ctx, data)
}

func (a *AccountRoleDep) GetSingleByParam(ctx *gin.Context, cacheControl string, param *model.GetAccountRoleByParam) (psqlmodel.AccountRole, error) {
	str, err := json.Marshal(param)
	if err != nil {
		return psqlmodel.AccountRole{}, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error marshal param")
	}

	key := fmt.Sprintf(model.GetSingleByParamAccountRoleKey, str)
	if cacheControl != model.MustRevalidate {
		res, err := a.getSingleByParamRedis(ctx, key)
		if err != nil {
			if err == goredislib.Nil {
				res, err := a.getSingleByParamPSQL(ctx, param)
				if err == nil {
					dataStr, err := json.Marshal(&res)
					if err != nil {
						return res, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error get psql")
					}
					err = a.setRedis(ctx, key, string(dataStr))
					if err != nil {
						return res, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error set redis")
					}
				}
				return res, err
			}
			return res, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error marshal param")
		}
		return res, nil
	}

	res, err := a.getSingleByParamPSQL(ctx, param)
	if err == nil {
		dataStr, err := json.Marshal(&res)
		if err != nil {
			return res, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error get psql")
		}
		err = a.setRedis(ctx, key, string(dataStr))
		if err != nil {
			return res, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error set redis")
		}
	}
	return res, nil
}

func (a *AccountRoleDep) Update(ctx *gin.Context, AccountRole *psqlmodel.AccountRole) error {
	return a.updatePSQL(ctx, AccountRole)
}

func (a *AccountRoleDep) Delete(ctx *gin.Context, AccountRole *psqlmodel.AccountRole, id int64, isHardDelete bool) error {
	return a.deletePSQL(ctx, AccountRole, id, isHardDelete)
}
func (a *AccountRoleDep) GetByParam(ctx *gin.Context, cacheControl string, param *model.GetAccountRolesByParam) (psqlmodel.AccountRoleSlice, model.Pagination, error) {
	var pg model.Pagination
	var res psqlmodel.AccountRoleSlice

	str, err := json.Marshal(param)
	if err != nil {
		return psqlmodel.AccountRoleSlice{}, pg, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error marshal param")
	}

	key := fmt.Sprintf(model.GetByParamAccountRoleKey, str)
	keyPg := fmt.Sprintf(model.GetByParamAccountRolePgKey, str)
	if cacheControl != model.MustRevalidate {
		res, err1 := a.getByParamRedis(ctx, key)
		pg, err2 := a.getByParamPaginationRedis(ctx, keyPg)
		if err1 != nil || err2 != nil {
			if err1 == goredislib.Nil || err2 == goredislib.Nil {
				res, pg, err := a.getByParamPSQL(ctx, param)
				if err == nil {
					dataStr, err := json.Marshal(&res)
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error get psql")
					}
					err = a.setRedis(ctx, key, string(dataStr))
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error set redis")
					}
					dataStr, err = json.Marshal(&pg)
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error get psql")
					}
					err = a.setRedis(ctx, key, string(dataStr))
					if err != nil {
						return res, pg, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error set redis")
					}
				}
				return res, pg, err
			}
			return res, pg, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error marshal param")
		}
		return res, pg, nil
	}

	res, pg, err = a.getByParamPSQL(ctx, param)
	if err == nil {
		dataStr, err := json.Marshal(&res)
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error get psql")
		}
		err = a.setRedis(ctx, key, string(dataStr))
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error set redis")
		}
		dataStr, err = json.Marshal(&pg)
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error get psql")
		}
		err = a.setRedis(ctx, key, string(dataStr))
		if err != nil {
			return res, pg, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error set redis")
		}
	}
	return res, pg, err
}
