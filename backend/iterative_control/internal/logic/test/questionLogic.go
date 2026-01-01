// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package test

import (
	"context"
	"iterative_control/internal/svc"
	"iterative_control/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionLogic {
	return &QuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuestionLogic) Question(req *types.QuestionReq) (resp *types.AnswerResp, err error) {

	ans := l.svcCtx.BLClient.Chat(req.Question)

	return &types.AnswerResp{
		Answer: ans,
	}, nil
}
