package account

import (
	"database/sql"

	"github.com/achwanyusuf/carrent-accountsvc/src/model"
	"github.com/achwanyusuf/carrent-accountsvc/src/model/psqlmodel"
	"github.com/achwanyusuf/carrent-accountsvc/src/model/svcerr"
	"github.com/achwanyusuf/carrent-lib/pkg/errormsg"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func (a *AccountDep) insertPSQL(ctx *gin.Context, data *psqlmodel.Account) error {
	tx, err := a.DB.BeginTx(ctx, nil)
	if err != nil {
		return errormsg.WrapErr(svcerr.AccountSVCPSQLErrorTransaction, err, "error begin transaction")
	}

	err = data.Insert(ctx, tx, boil.Infer())
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			a.Log.Warn(ctx, errormsg.WrapErr(svcerr.AccountSVCPSQLErrorRollback, err, "error rollback"))
		}
		return errormsg.WrapErr(svcerr.AccountSVCPSQLErrorInsert, err, "error insert")
	}
	err = tx.Commit()
	if err != nil {
		return errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error commit")
	}
	return nil
}

func (a *AccountDep) getSingleByParamPSQL(ctx *gin.Context, param *model.GetAccountByParam) (psqlmodel.Account, error) {
	var res psqlmodel.Account
	qr := param.GetQuery()
	account, err := psqlmodel.Accounts(qr...).One(ctx, a.DB)
	if err == sql.ErrNoRows {
		return res, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error get accounts")
	}

	if err != nil {
		return res, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error get accounts")
	}

	return *account, nil
}

func (a *AccountDep) updatePSQL(ctx *gin.Context, account *psqlmodel.Account) error {
	tx, err := a.DB.BeginTx(ctx, nil)
	if err != nil {
		return errormsg.WrapErr(svcerr.AccountSVCPSQLErrorTransaction, err, "error begin transaction")
	}

	_, err = account.Update(ctx, tx, boil.Infer())
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			a.Log.Warn(ctx, errormsg.WrapErr(svcerr.AccountSVCPSQLErrorRollback, err, "error rollback"))
		}
		return errormsg.WrapErr(svcerr.AccountSVCPSQLErrorUpdate, err, "error update")
	}
	err = tx.Commit()
	if err != nil {
		return errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error commit")
	}
	return nil
}

func (a *AccountDep) deletePSQL(ctx *gin.Context, account *psqlmodel.Account, id int64, isHardDelete bool) error {
	tx, err := a.DB.BeginTx(ctx, nil)
	if err != nil {
		return errormsg.WrapErr(svcerr.AccountSVCPSQLErrorTransaction, err, "error begin transaction")
	}

	_, err = account.Delete(ctx, tx, isHardDelete)
	if err != nil {
		if errRollback := tx.Rollback(); errRollback != nil {
			a.Log.Warn(ctx, errormsg.WrapErr(svcerr.AccountSVCPSQLErrorRollback, err, "error rollback"))
		}
		return errormsg.WrapErr(svcerr.AccountSVCPSQLErrorUpdate, err, "error delete")
	}

	if !isHardDelete {
		account.DeletedBy = null.NewInt(int(id), true)
		_, err = account.Update(ctx, tx, boil.Infer())
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				a.Log.Warn(ctx, errormsg.WrapErr(svcerr.AccountSVCPSQLErrorRollback, err, "error rollback"))
			}
			return errormsg.WrapErr(svcerr.AccountSVCPSQLErrorUpdate, err, "error update")
		}
	}
	err = tx.Commit()
	if err != nil {
		return errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error commit")
	}
	return nil
}

func (a *AccountDep) getByParamPSQL(ctx *gin.Context, param *model.GetAccountsByParam) (psqlmodel.AccountSlice, model.Pagination, error) {
	var totalPages int64 = 1
	if param.Limit == 0 {
		param.Limit = int64(a.Conf.DefaultPageLimit)
	}

	if param.Page == 0 {
		param.Page = 1
	}

	qr := param.GetQuery()
	count, err := psqlmodel.Accounts(qr...).Count(ctx, a.DB)
	if err != nil {
		return psqlmodel.AccountSlice{}, model.Pagination{}, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error count data")
	}
	qr = append(qr, qm.Offset(int((param.Page-1)*param.Limit)))
	qr = append(qr, qm.Limit(int(param.Limit)))
	accounts, err := psqlmodel.Accounts(qr...).All(ctx, a.DB)
	if err == sql.ErrNoRows {
		return accounts, model.Pagination{}, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error get accounts")
	}
	if err != nil {
		return accounts, model.Pagination{}, errormsg.WrapErr(svcerr.AccountSVCBadRequest, err, "error get accounts")
	}
	if count > 0 {
		totalPages = (count / param.Limit) + 1
	}
	return accounts, model.Pagination{
		CurrentPage:     param.Page,
		CurrentElements: int64(len(accounts)),
		TotalElements:   count,
		TotalPages:      totalPages,
		SortBy:          param.OrderBy.String,
	}, nil
}
