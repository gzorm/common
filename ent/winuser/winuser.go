// Code generated by ent, DO NOT EDIT.

package winuser

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the winuser type in the database.
	Label = "win_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldFcoin holds the string denoting the fcoin field in the database.
	FieldFcoin = "fcoin"
	// FieldCoinCommission holds the string denoting the coin_commission field in the database.
	FieldCoinCommission = "coin_commission"
	// FieldLevelID holds the string denoting the level_id field in the database.
	FieldLevelID = "level_id"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldIsPromoter holds the string denoting the is_promoter field in the database.
	FieldIsPromoter = "is_promoter"
	// FieldFlag holds the string denoting the flag field in the database.
	FieldFlag = "flag"
	// FieldRealName holds the string denoting the real_name field in the database.
	FieldRealName = "real_name"
	// FieldSignature holds the string denoting the signature field in the database.
	FieldSignature = "signature"
	// FieldBirthday holds the string denoting the birthday field in the database.
	FieldBirthday = "birthday"
	// FieldAreaCode holds the string denoting the area_code field in the database.
	FieldAreaCode = "area_code"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldSex holds the string denoting the sex field in the database.
	FieldSex = "sex"
	// FieldBindBank holds the string denoting the bind_bank field in the database.
	FieldBindBank = "bind_bank"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldPromoCode holds the string denoting the promo_code field in the database.
	FieldPromoCode = "promo_code"
	// FieldSupUID1 holds the string denoting the sup_uid_1 field in the database.
	FieldSupUID1 = "sup_uid_1"
	// FieldSupUsername1 holds the string denoting the sup_username_1 field in the database.
	FieldSupUsername1 = "sup_username_1"
	// FieldSupUID2 holds the string denoting the sup_uid_2 field in the database.
	FieldSupUID2 = "sup_uid_2"
	// FieldSupUID3 holds the string denoting the sup_uid_3 field in the database.
	FieldSupUID3 = "sup_uid_3"
	// FieldSupUID4 holds the string denoting the sup_uid_4 field in the database.
	FieldSupUID4 = "sup_uid_4"
	// FieldSupUID5 holds the string denoting the sup_uid_5 field in the database.
	FieldSupUID5 = "sup_uid_5"
	// FieldSupUID6 holds the string denoting the sup_uid_6 field in the database.
	FieldSupUID6 = "sup_uid_6"
	// FieldSupUIDTop holds the string denoting the sup_uid_top field in the database.
	FieldSupUIDTop = "sup_uid_top"
	// FieldSupUsernameTop holds the string denoting the sup_username_top field in the database.
	FieldSupUsernameTop = "sup_username_top"
	// FieldSupLevelTop holds the string denoting the sup_level_top field in the database.
	FieldSupLevelTop = "sup_level_top"
	// FieldPasswordHash holds the string denoting the password_hash field in the database.
	FieldPasswordHash = "password_hash"
	// FieldPasswordCoin holds the string denoting the password_coin field in the database.
	FieldPasswordCoin = "password_coin"
	// FieldIP holds the string denoting the ip field in the database.
	FieldIP = "ip"
	// FieldIPRegion holds the string denoting the ip_region field in the database.
	FieldIPRegion = "ip_region"
	// FieldThirdLoginType holds the string denoting the third_login_type field in the database.
	FieldThirdLoginType = "third_login_type"
	// FieldFreezeCause holds the string denoting the freeze_cause field in the database.
	FieldFreezeCause = "freeze_cause"
	// FieldFreezeAt holds the string denoting the freeze_at field in the database.
	FieldFreezeAt = "freeze_at"
	// FieldOperatorName holds the string denoting the operator_name field in the database.
	FieldOperatorName = "operator_name"
	// FieldCreatedName holds the string denoting the created_name field in the database.
	FieldCreatedName = "created_name"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldLastLoginIP holds the string denoting the last_login_ip field in the database.
	FieldLastLoginIP = "last_login_ip"
	// FieldLastLoginIPRegion holds the string denoting the last_login_ip_region field in the database.
	FieldLastLoginIPRegion = "last_login_ip_region"
	// FieldLastLoginTime holds the string denoting the last_login_time field in the database.
	FieldLastLoginTime = "last_login_time"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the winuser in the database.
	Table = "win_user"
)

// Columns holds all SQL columns for winuser fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldAvatar,
	FieldFcoin,
	FieldCoinCommission,
	FieldLevelID,
	FieldRole,
	FieldIsPromoter,
	FieldFlag,
	FieldRealName,
	FieldSignature,
	FieldBirthday,
	FieldAreaCode,
	FieldMobile,
	FieldEmail,
	FieldSex,
	FieldBindBank,
	FieldAddress,
	FieldScore,
	FieldPromoCode,
	FieldSupUID1,
	FieldSupUsername1,
	FieldSupUID2,
	FieldSupUID3,
	FieldSupUID4,
	FieldSupUID5,
	FieldSupUID6,
	FieldSupUIDTop,
	FieldSupUsernameTop,
	FieldSupLevelTop,
	FieldPasswordHash,
	FieldPasswordCoin,
	FieldIP,
	FieldIPRegion,
	FieldThirdLoginType,
	FieldFreezeCause,
	FieldFreezeAt,
	FieldOperatorName,
	FieldCreatedName,
	FieldStatus,
	FieldLastLoginIP,
	FieldLastLoginIPRegion,
	FieldLastLoginTime,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the WinUser queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByAvatar orders the results by the avatar field.
func ByAvatar(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAvatar, opts...).ToFunc()
}

// ByFcoin orders the results by the fcoin field.
func ByFcoin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFcoin, opts...).ToFunc()
}

// ByCoinCommission orders the results by the coin_commission field.
func ByCoinCommission(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCoinCommission, opts...).ToFunc()
}

// ByLevelID orders the results by the level_id field.
func ByLevelID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevelID, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByIsPromoter orders the results by the is_promoter field.
func ByIsPromoter(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsPromoter, opts...).ToFunc()
}

// ByFlag orders the results by the flag field.
func ByFlag(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFlag, opts...).ToFunc()
}

// ByRealName orders the results by the real_name field.
func ByRealName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRealName, opts...).ToFunc()
}

// BySignature orders the results by the signature field.
func BySignature(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignature, opts...).ToFunc()
}

// ByBirthday orders the results by the birthday field.
func ByBirthday(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthday, opts...).ToFunc()
}

// ByAreaCode orders the results by the area_code field.
func ByAreaCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAreaCode, opts...).ToFunc()
}

// ByMobile orders the results by the mobile field.
func ByMobile(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMobile, opts...).ToFunc()
}

// ByEmail orders the results by the email field.
func ByEmail(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEmail, opts...).ToFunc()
}

// BySex orders the results by the sex field.
func BySex(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSex, opts...).ToFunc()
}

// ByBindBank orders the results by the bind_bank field.
func ByBindBank(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBindBank, opts...).ToFunc()
}

// ByAddress orders the results by the address field.
func ByAddress(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAddress, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByPromoCode orders the results by the promo_code field.
func ByPromoCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPromoCode, opts...).ToFunc()
}

// BySupUID1 orders the results by the sup_uid_1 field.
func BySupUID1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupUID1, opts...).ToFunc()
}

// BySupUsername1 orders the results by the sup_username_1 field.
func BySupUsername1(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupUsername1, opts...).ToFunc()
}

// BySupUID2 orders the results by the sup_uid_2 field.
func BySupUID2(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupUID2, opts...).ToFunc()
}

// BySupUID3 orders the results by the sup_uid_3 field.
func BySupUID3(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupUID3, opts...).ToFunc()
}

// BySupUID4 orders the results by the sup_uid_4 field.
func BySupUID4(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupUID4, opts...).ToFunc()
}

// BySupUID5 orders the results by the sup_uid_5 field.
func BySupUID5(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupUID5, opts...).ToFunc()
}

// BySupUID6 orders the results by the sup_uid_6 field.
func BySupUID6(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupUID6, opts...).ToFunc()
}

// BySupUIDTop orders the results by the sup_uid_top field.
func BySupUIDTop(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupUIDTop, opts...).ToFunc()
}

// BySupUsernameTop orders the results by the sup_username_top field.
func BySupUsernameTop(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupUsernameTop, opts...).ToFunc()
}

// BySupLevelTop orders the results by the sup_level_top field.
func BySupLevelTop(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSupLevelTop, opts...).ToFunc()
}

// ByPasswordHash orders the results by the password_hash field.
func ByPasswordHash(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordHash, opts...).ToFunc()
}

// ByPasswordCoin orders the results by the password_coin field.
func ByPasswordCoin(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPasswordCoin, opts...).ToFunc()
}

// ByIP orders the results by the ip field.
func ByIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIP, opts...).ToFunc()
}

// ByIPRegion orders the results by the ip_region field.
func ByIPRegion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIPRegion, opts...).ToFunc()
}

// ByThirdLoginType orders the results by the third_login_type field.
func ByThirdLoginType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThirdLoginType, opts...).ToFunc()
}

// ByFreezeCause orders the results by the freeze_cause field.
func ByFreezeCause(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFreezeCause, opts...).ToFunc()
}

// ByFreezeAt orders the results by the freeze_at field.
func ByFreezeAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFreezeAt, opts...).ToFunc()
}

// ByOperatorName orders the results by the operator_name field.
func ByOperatorName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldOperatorName, opts...).ToFunc()
}

// ByCreatedName orders the results by the created_name field.
func ByCreatedName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedName, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByLastLoginIP orders the results by the last_login_ip field.
func ByLastLoginIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginIP, opts...).ToFunc()
}

// ByLastLoginIPRegion orders the results by the last_login_ip_region field.
func ByLastLoginIPRegion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginIPRegion, opts...).ToFunc()
}

// ByLastLoginTime orders the results by the last_login_time field.
func ByLastLoginTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastLoginTime, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}
