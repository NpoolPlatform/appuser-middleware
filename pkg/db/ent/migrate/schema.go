// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppsColumns holds the columns for the "apps" table.
	AppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "logo", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// AppsTable holds the schema information for the "apps" table.
	AppsTable = &schema.Table{
		Name:       "apps",
		Columns:    AppsColumns,
		PrimaryKey: []*schema.Column{AppsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "app_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppsColumns[4]},
			},
		},
	}
	// AppControlsColumns holds the columns for the "app_controls" table.
	AppControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "signup_methods", Type: field.TypeJSON, Nullable: true},
		{Name: "extern_signin_methods", Type: field.TypeJSON, Nullable: true},
		{Name: "recaptcha_method", Type: field.TypeString, Nullable: true, Default: "GoogleRecaptchaV3"},
		{Name: "kyc_enable", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "signin_verify_enable", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "invitation_code_must", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "create_invitation_code_when", Type: field.TypeString, Nullable: true, Default: "Registration"},
		{Name: "max_typed_coupons_per_order", Type: field.TypeUint32, Nullable: true, Default: 1},
		{Name: "maintaining", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "coupon_withdraw_enable", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "commit_button_targets", Type: field.TypeJSON, Nullable: true},
	}
	// AppControlsTable holds the schema information for the "app_controls" table.
	AppControlsTable = &schema.Table{
		Name:       "app_controls",
		Columns:    AppControlsColumns,
		PrimaryKey: []*schema.Column{AppControlsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appcontrol_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppControlsColumns[4]},
			},
		},
	}
	// AppOauthThirdPartiesColumns holds the columns for the "app_oauth_third_parties" table.
	AppOauthThirdPartiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "third_party_id", Type: field.TypeUUID, Nullable: true},
		{Name: "client_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "client_secret", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "callback_url", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "salt", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// AppOauthThirdPartiesTable holds the schema information for the "app_oauth_third_parties" table.
	AppOauthThirdPartiesTable = &schema.Table{
		Name:       "app_oauth_third_parties",
		Columns:    AppOauthThirdPartiesColumns,
		PrimaryKey: []*schema.Column{AppOauthThirdPartiesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appoauththirdparty_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppOauthThirdPartiesColumns[4]},
			},
		},
	}
	// AppRolesColumns holds the columns for the "app_roles" table.
	AppRolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "created_by", Type: field.TypeUUID, Nullable: true},
		{Name: "role", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "description", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "default", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "genesis", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// AppRolesTable holds the schema information for the "app_roles" table.
	AppRolesTable = &schema.Table{
		Name:       "app_roles",
		Columns:    AppRolesColumns,
		PrimaryKey: []*schema.Column{AppRolesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "approle_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppRolesColumns[4]},
			},
		},
	}
	// AppRoleUsersColumns holds the columns for the "app_role_users" table.
	AppRoleUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "role_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
	}
	// AppRoleUsersTable holds the schema information for the "app_role_users" table.
	AppRoleUsersTable = &schema.Table{
		Name:       "app_role_users",
		Columns:    AppRoleUsersColumns,
		PrimaryKey: []*schema.Column{AppRoleUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "approleuser_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppRoleUsersColumns[4]},
			},
		},
	}
	// AppSubscribesColumns holds the columns for the "app_subscribes" table.
	AppSubscribesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "subscribe_app_id", Type: field.TypeUUID, Nullable: true},
	}
	// AppSubscribesTable holds the schema information for the "app_subscribes" table.
	AppSubscribesTable = &schema.Table{
		Name:       "app_subscribes",
		Columns:    AppSubscribesColumns,
		PrimaryKey: []*schema.Column{AppSubscribesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appsubscribe_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppSubscribesColumns[4]},
			},
		},
	}
	// AppUsersColumns holds the columns for the "app_users" table.
	AppUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "email_address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "phone_no", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "import_from_app", Type: field.TypeUUID, Nullable: true},
	}
	// AppUsersTable holds the schema information for the "app_users" table.
	AppUsersTable = &schema.Table{
		Name:       "app_users",
		Columns:    AppUsersColumns,
		PrimaryKey: []*schema.Column{AppUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appuser_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppUsersColumns[4]},
			},
		},
	}
	// AppUserControlsColumns holds the columns for the "app_user_controls" table.
	AppUserControlsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "signin_verify_by_google_authentication", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "google_authentication_verified", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "signin_verify_type", Type: field.TypeString, Nullable: true, Default: "Email"},
		{Name: "kol", Type: field.TypeBool, Default: false},
		{Name: "kol_confirmed", Type: field.TypeBool, Default: false},
		{Name: "selected_lang_id", Type: field.TypeUUID, Nullable: true},
	}
	// AppUserControlsTable holds the schema information for the "app_user_controls" table.
	AppUserControlsTable = &schema.Table{
		Name:       "app_user_controls",
		Columns:    AppUserControlsColumns,
		PrimaryKey: []*schema.Column{AppUserControlsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appusercontrol_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppUserControlsColumns[4]},
			},
			{
				Name:    "appusercontrol_app_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{AppUserControlsColumns[5], AppUserControlsColumns[6]},
			},
		},
	}
	// AppUserExtrasColumns holds the columns for the "app_user_extras" table.
	AppUserExtrasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "username", Type: field.TypeString, Default: ""},
		{Name: "first_name", Type: field.TypeString, Default: ""},
		{Name: "last_name", Type: field.TypeString, Default: ""},
		{Name: "address_fields", Type: field.TypeJSON},
		{Name: "gender", Type: field.TypeString, Default: ""},
		{Name: "postal_code", Type: field.TypeString, Default: ""},
		{Name: "age", Type: field.TypeUint32, Default: 0},
		{Name: "birthday", Type: field.TypeUint32, Default: 0},
		{Name: "avatar", Type: field.TypeString, Default: ""},
		{Name: "organization", Type: field.TypeString, Default: ""},
		{Name: "id_number", Type: field.TypeString, Default: ""},
		{Name: "action_credits", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// AppUserExtrasTable holds the schema information for the "app_user_extras" table.
	AppUserExtrasTable = &schema.Table{
		Name:       "app_user_extras",
		Columns:    AppUserExtrasColumns,
		PrimaryKey: []*schema.Column{AppUserExtrasColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appuserextra_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppUserExtrasColumns[4]},
			},
			{
				Name:    "appuserextra_app_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{AppUserExtrasColumns[5], AppUserExtrasColumns[6]},
			},
		},
	}
	// AppUserSecretsColumns holds the columns for the "app_user_secrets" table.
	AppUserSecretsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "password_hash", Type: field.TypeString},
		{Name: "salt", Type: field.TypeString},
		{Name: "google_secret", Type: field.TypeString, Default: ""},
	}
	// AppUserSecretsTable holds the schema information for the "app_user_secrets" table.
	AppUserSecretsTable = &schema.Table{
		Name:       "app_user_secrets",
		Columns:    AppUserSecretsColumns,
		PrimaryKey: []*schema.Column{AppUserSecretsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appusersecret_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppUserSecretsColumns[4]},
			},
			{
				Name:    "appusersecret_app_id_user_id",
				Unique:  true,
				Columns: []*schema.Column{AppUserSecretsColumns[5], AppUserSecretsColumns[6]},
			},
		},
	}
	// AppUserThirdPartiesColumns holds the columns for the "app_user_third_parties" table.
	AppUserThirdPartiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "third_party_user_id", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "third_party_id", Type: field.TypeUUID, Nullable: true},
		{Name: "third_party_username", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "third_party_avatar", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
	}
	// AppUserThirdPartiesTable holds the schema information for the "app_user_third_parties" table.
	AppUserThirdPartiesTable = &schema.Table{
		Name:       "app_user_third_parties",
		Columns:    AppUserThirdPartiesColumns,
		PrimaryKey: []*schema.Column{AppUserThirdPartiesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appuserthirdparty_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppUserThirdPartiesColumns[4]},
			},
		},
	}
	// AuthsColumns holds the columns for the "auths" table.
	AuthsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "role_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "resource", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "method", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// AuthsTable holds the schema information for the "auths" table.
	AuthsTable = &schema.Table{
		Name:       "auths",
		Columns:    AuthsColumns,
		PrimaryKey: []*schema.Column{AuthsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "auth_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AuthsColumns[4]},
			},
		},
	}
	// AuthHistoriesColumns holds the columns for the "auth_histories" table.
	AuthHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "resource", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "method", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "allowed", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// AuthHistoriesTable holds the schema information for the "auth_histories" table.
	AuthHistoriesTable = &schema.Table{
		Name:       "auth_histories",
		Columns:    AuthHistoriesColumns,
		PrimaryKey: []*schema.Column{AuthHistoriesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "authhistory_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AuthHistoriesColumns[4]},
			},
		},
	}
	// BanAppsColumns holds the columns for the "ban_apps" table.
	BanAppsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "message", Type: field.TypeString, Default: ""},
	}
	// BanAppsTable holds the schema information for the "ban_apps" table.
	BanAppsTable = &schema.Table{
		Name:       "ban_apps",
		Columns:    BanAppsColumns,
		PrimaryKey: []*schema.Column{BanAppsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "banapp_ent_id",
				Unique:  true,
				Columns: []*schema.Column{BanAppsColumns[4]},
			},
		},
	}
	// BanAppUsersColumns holds the columns for the "ban_app_users" table.
	BanAppUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "message", Type: field.TypeString, Default: ""},
	}
	// BanAppUsersTable holds the schema information for the "ban_app_users" table.
	BanAppUsersTable = &schema.Table{
		Name:       "ban_app_users",
		Columns:    BanAppUsersColumns,
		PrimaryKey: []*schema.Column{BanAppUsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "banappuser_ent_id",
				Unique:  true,
				Columns: []*schema.Column{BanAppUsersColumns[4]},
			},
		},
	}
	// KycsColumns holds the columns for the "kycs" table.
	KycsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "document_type", Type: field.TypeString, Nullable: true, Default: "DefaultKycDocumentType"},
		{Name: "id_number", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "front_img", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "back_img", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "selfie_img", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "entity_type", Type: field.TypeString, Nullable: true, Default: "Individual"},
		{Name: "review_id", Type: field.TypeUUID, Nullable: true},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultState"},
	}
	// KycsTable holds the schema information for the "kycs" table.
	KycsTable = &schema.Table{
		Name:       "kycs",
		Columns:    KycsColumns,
		PrimaryKey: []*schema.Column{KycsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "kyc_ent_id",
				Unique:  true,
				Columns: []*schema.Column{KycsColumns[4]},
			},
		},
	}
	// LoginHistoriesColumns holds the columns for the "login_histories" table.
	LoginHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "client_ip", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "user_agent", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "location", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "login_type", Type: field.TypeString, Nullable: true, Default: "FreshLogin"},
	}
	// LoginHistoriesTable holds the schema information for the "login_histories" table.
	LoginHistoriesTable = &schema.Table{
		Name:       "login_histories",
		Columns:    LoginHistoriesColumns,
		PrimaryKey: []*schema.Column{LoginHistoriesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "loginhistory_ent_id",
				Unique:  true,
				Columns: []*schema.Column{LoginHistoriesColumns[4]},
			},
		},
	}
	// OauthThirdPartiesColumns holds the columns for the "oauth_third_parties" table.
	OauthThirdPartiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "client_name", Type: field.TypeString, Nullable: true, Default: "DefaultSignMethod"},
		{Name: "client_tag", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "client_logo_url", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "client_oauth_url", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "response_type", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "scope", Type: field.TypeString, Nullable: true, Default: ""},
	}
	// OauthThirdPartiesTable holds the schema information for the "oauth_third_parties" table.
	OauthThirdPartiesTable = &schema.Table{
		Name:       "oauth_third_parties",
		Columns:    OauthThirdPartiesColumns,
		PrimaryKey: []*schema.Column{OauthThirdPartiesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "oauththirdparty_ent_id",
				Unique:  true,
				Columns: []*schema.Column{OauthThirdPartiesColumns[4]},
			},
		},
	}
	// PubsubMessagesColumns holds the columns for the "pubsub_messages" table.
	PubsubMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "message_id", Type: field.TypeString, Nullable: true, Default: "DefaultMsgID"},
		{Name: "state", Type: field.TypeString, Nullable: true, Default: "DefaultMsgState"},
		{Name: "resp_to_id", Type: field.TypeUUID, Nullable: true},
		{Name: "undo_id", Type: field.TypeUUID, Nullable: true},
		{Name: "arguments", Type: field.TypeString, Nullable: true, Size: 2147483647, Default: ""},
	}
	// PubsubMessagesTable holds the schema information for the "pubsub_messages" table.
	PubsubMessagesTable = &schema.Table{
		Name:       "pubsub_messages",
		Columns:    PubsubMessagesColumns,
		PrimaryKey: []*schema.Column{PubsubMessagesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "pubsubmessage_ent_id",
				Unique:  true,
				Columns: []*schema.Column{PubsubMessagesColumns[4]},
			},
			{
				Name:    "pubsubmessage_state_resp_to_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[6], PubsubMessagesColumns[7]},
			},
			{
				Name:    "pubsubmessage_state_undo_id",
				Unique:  false,
				Columns: []*schema.Column{PubsubMessagesColumns[6], PubsubMessagesColumns[8]},
			},
		},
	}
	// RecoveryCodesColumns holds the columns for the "recovery_codes" table.
	RecoveryCodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "code", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "used", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// RecoveryCodesTable holds the schema information for the "recovery_codes" table.
	RecoveryCodesTable = &schema.Table{
		Name:       "recovery_codes",
		Columns:    RecoveryCodesColumns,
		PrimaryKey: []*schema.Column{RecoveryCodesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "recoverycode_ent_id",
				Unique:  true,
				Columns: []*schema.Column{RecoveryCodesColumns[4]},
			},
		},
	}
	// SubscribersColumns holds the columns for the "subscribers" table.
	SubscribersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "email_address", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "registered", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// SubscribersTable holds the schema information for the "subscribers" table.
	SubscribersTable = &schema.Table{
		Name:       "subscribers",
		Columns:    SubscribersColumns,
		PrimaryKey: []*schema.Column{SubscribersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "subscriber_ent_id",
				Unique:  true,
				Columns: []*schema.Column{SubscribersColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppsTable,
		AppControlsTable,
		AppOauthThirdPartiesTable,
		AppRolesTable,
		AppRoleUsersTable,
		AppSubscribesTable,
		AppUsersTable,
		AppUserControlsTable,
		AppUserExtrasTable,
		AppUserSecretsTable,
		AppUserThirdPartiesTable,
		AuthsTable,
		AuthHistoriesTable,
		BanAppsTable,
		BanAppUsersTable,
		KycsTable,
		LoginHistoriesTable,
		OauthThirdPartiesTable,
		PubsubMessagesTable,
		RecoveryCodesTable,
		SubscribersTable,
	}
)

func init() {
}
