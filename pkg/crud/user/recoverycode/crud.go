package user

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entrecoverycode "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/recoverycode"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Conds struct {
	ID     *cruder.Cond
	EntID  *cruder.Cond
	UserID *cruder.Cond
	AppID  *cruder.Cond
	Code   *cruder.Cond
	Used   *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.RecoveryCodeQuery, conds *Conds) (*ent.RecoveryCodeQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entrecoverycode.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entrecoverycode.EntID(id))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.UserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("userid appid")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entrecoverycode.UserID(id))
		default:
			return nil, fmt.Errorf("invalid userid field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entrecoverycode.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appid field")
		}
	}
	if conds.Code != nil {
		code, ok := conds.Code.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid code")
		}
		switch conds.Code.Op {
		case cruder.EQ:
			q.Where(entrecoverycode.Code(code))
		default:
			return nil, fmt.Errorf("invalid code field")
		}
	}
	if conds.Used != nil {
		used, ok := conds.Used.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid used")
		}
		switch conds.Used.Op {
		case cruder.EQ:
			q.Where(entrecoverycode.Used(used))
		default:
			return nil, fmt.Errorf("invalid used field")
		}
	}
	q.Where(entrecoverycode.DeletedAt(0))
	return q, nil
}
