package tables

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Access struct {
	Token        string
	TokenRefresh string
}

type User struct {
	bun.BaseModel `bun:"table:hk.users"`

	ID            uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name          string    `bun:"name,notnull,type:varchar"`
	CredentialsID uuid.UUID `bun:"credentials_id,notnull,type:uuid"`
	TypeUserID    uuid.UUID `bun:"type_user_id,notnull,type:uuid"`

	TypeUser    *TypeUser       `bun:"rel:has-one,join:type_user_id=id"`
	Credentials *Credentials    `bun:"rel:has-one,join:credentials_id=id"`
	AccessRols  []UserAccessRol `bun:"-"`
}

type TypeUser struct {
	bun.BaseModel `bun:"table:hk.type_user"`

	ID   uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name string    `bun:"name,notnull,type:varchar"`
}

type UsersStatus struct {
	bun.BaseModel `bun:"table:hk.users_status"`

	ID   uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name string    `bun:"name,notnull,type:varchar"`
}

type Credentials struct {
	bun.BaseModel `bun:"table:hk.credentials"`

	ID       uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Email    string    `bun:"email,notnull,type:varchar"`
	Password string    `bun:"password,notnull,type:varchar"`
	StatusID uuid.UUID `bun:"status_id,notnull,type:uuid"`

	Status *UsersStatus `bun:"rel:has-one,join:status_id=id"`
}

type AccessRol struct {
	bun.BaseModel `bun:"table:hk.access_rol"`

	ID   uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	Name string    `bun:"name,notnull,type:varchar"`
}

type UserAccessRol struct {
	bun.BaseModel `bun:"table:hk.user_access_rol"`

	ID          uuid.UUID `bun:"id,pk,type:uuid,default:gen_random_uuid()"`
	AccessRolID uuid.UUID `bun:"access_rol_id,notnull,type:uuid"`
	UsersID     uuid.UUID `bun:"users_id,notnull,type:uuid"`

	AccessRol *AccessRol `bun:"rel:has-one,join:access_rol_id=id"`
	Users     *User      `bun:"rel:has-one,join:users_id=id"`
}
