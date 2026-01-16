package No3

type UserType int

const (
	UserTypeDefault UserType = iota
	UserTypeOrganization
	UserTypeBot
	UserTypeReserved
)

func (u UserType) String() string {
	switch u {
	case UserTypeDefault:
		return "Default User"
	case UserTypeOrganization:
		return "Organization"
	case UserTypeBot:
		return "Bot"
	case UserTypeReserved:
		return "Reserved"

	}
	return "Unknown"
}

func EnumTest() {
	userType := UserTypeDefault
	switch userType {
	case UserTypeDefault:
		break
	case UserTypeOrganization:
	default:
		break

	}
}
