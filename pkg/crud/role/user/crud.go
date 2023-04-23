//nolint:nolintlint,dupl
package approleuser

import (
	"fmt"

	"github.com/NpoolPlatform/appuser-middleware/pkg/db/ent"
	entapproleuser "github.com/NpoolPlatform/appuser-middleware/pkg/db/ent/approleuser"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/google/uuid"
)

type Req struct {
	ID        *uuid.UUID
	AppID     *uuid.UUID
	RoleID    *uuid.UUID
	UserID    *uuid.UUID
	DeletedAt *uint32
}

func CreateSet(c *ent.AppRoleUserCreate, req *Req) *ent.AppRoleUserCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.RoleID != nil {
		c.SetRoleID(*req.RoleID)
	}
	if req.UserID != nil {
		c.SetUserID(*req.UserID)
	}
	return c
}

func UpdateSet(u *ent.AppRoleUserUpdateOne, req *Req) *ent.AppRoleUserUpdateOne {
	if req.RoleID != nil {
		u.SetRoleID(*req.RoleID)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	AppID   *cruder.Cond
	RoleID  *cruder.Cond
	UserID  *cruder.Cond
	AppIDs  *cruder.Cond
	RoleIDs *cruder.Cond
	Genesis *cruder.Cond
}

//nolint
func SetQueryConds(q *ent.AppRoleUserQuery, conds *Conds) (*ent.AppRoleUserQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entapproleuser.ID(id))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entapproleuser.AppID(id))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entapproleuser.AppIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	if conds.UserID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.UserID.Op {
		case cruder.EQ:
			q.Where(entapproleuser.UserID(id))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	if conds.RoleID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.RoleID.Op {
		case cruder.EQ:
			q.Where(entapproleuser.RoleID(id))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	if conds.RoleIDs != nil {
		ids, ok := conds.RoleIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.RoleIDs.Op {
		case cruder.IN:
			q.Where(entapproleuser.RoleIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid approleuser field")
		}
	}
	return q, nil
}
