// @Title  用户账户登录业务逻辑
// @Description  用户账户登录业务逻辑
// @Author  noobwu(2021-07-29)
// @Update  用户账户登录业务逻辑(2021-07-29)

package logic

import (
	"context"
	"encoding/json"
	"go-zero-admin/api/internal/common/errorx"
	"go-zero-admin/rpc/sys/sysclient"
	"strings"

	"go-zero-admin/api/internal/svc"
	"go-zero-admin/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

// UserLoginLogic   账户登录业务逻辑对象
type UserLoginLogic struct {
	logx.Logger                     //日志
	ctx         context.Context     //上下文
	svcCtx      *svc.ServiceContext //服务上下文
}

// @title  实例化一个UserLoginLogic对象
// @description  实例化一个UserLoginLogic对象
// @auth      noobwu(2021-07-27)
// @param     ctx              context.Context        "上下文"
// @param     svcCtx           *svc.ServiceContext    "服务上下文"
// @return    UserLoginLogic        UserLoginLogic     "UserLoginLogic实例"
func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserLoginLogic {
	return UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// @title   账户登录
// @description  根据用户名和密码登录
// @auth      noobwu(2021-07-27)
// @param     req              types.LoginReq        "请求参数"
// @param     ip               string                "ip地址"
// @return    LoginResp        types.LoginResp       "用户信息"
// @return    error        error                 "错误信息"

func (l *UserLoginLogic) UserLogin(req types.LoginReq, ip string) (*types.LoginResp, error) {

	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		reqStr, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("用户名或密码不能为空,请求信息失败,参数:%s", reqStr)
		return nil, errorx.NewDefaultError("用户名或密码不能为空")
	}

	resp, err := l.svcCtx.Sys.Login(l.ctx, &sysclient.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据用户名: %s和密码: %s查询用户异常:%s", req.UserName, req.Password, err.Error())
		return nil, errorx.NewDefaultError("登录失败")
	}

	//保存登录日志
	_, _ = l.svcCtx.Sys.LoginLogAdd(l.ctx, &sysclient.LoginLogAddReq{
		UserName: resp.UserName,
		Status:   "login",
		Ip:       ip,
		CreateBy: resp.UserName,
	})

	return &types.LoginResp{
		Code:             "000000",
		Message:          "登录成功",
		Status:           resp.Status,
		CurrentAuthority: resp.CurrentAuthority,
		Id:               resp.Id,
		UserName:         resp.UserName,
		AccessToken:      resp.AccessToken,
		AccessExpire:     resp.AccessExpire,
		RefreshAfter:     resp.RefreshAfter,
	}, nil
}
