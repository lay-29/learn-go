package No3

import (
	"encoding/json"
	"fmt"
)

type Bush int

const (
	BushRed Bush = iota
	BushBlue
	BushGreen
)

type Color int

const (
	ColorRed Color = iota
	ColorBlue
	ColorGreen
)

type UserType int

const (
	UserTypeDefault UserType = iota
	UserTypeOrganization
	UserTypeBot
	UserTypeReserved
)

func (u *UserType) String() string {
	switch *u {
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
func ParseUserType(s string) (UserType, error) {
	switch s {
	case "UserTypeDefault":
		return UserTypeDefault, nil
	case "UserTypeOrganization":
		return UserTypeOrganization, nil
	case "UserTypeBot":
		return UserTypeBot, nil
	case "UserTypeReserved":
		return UserTypeReserved, nil
	default:
		return UserTypeDefault, fmt.Errorf("invalid UserType: %s", s)

	}
}
func (u *UserType) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.String())
}
func (u *UserType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if uType, err := ParseUserType(s); err == nil {
		*u = uType
		return nil
	}
	return fmt.Errorf("invalid UserType: %s", string(b))

}

type UserTypeStr string

const (
	UserTypeDefaultStr      UserTypeStr = "UserTypeDefault"
	UserTypeOrganizationStr UserTypeStr = "UserTypeOrganization"
	UserTypeBotStr          UserTypeStr = "UserTypeBot"
	UserTypeReservedStr     UserTypeStr = "UserTypeReserved"
)

func (u UserTypeStr) String() string {
	return string(u)
}

func EnumTest() {
	userType := UserTypeDefault

	fmt.Println("userType: ", userType)
	fmt.Println("userType: ", userType.String())

	userJson, _ := userType.MarshalJSON()
	fmt.Println("userType Json Byte: ", string(userJson))

	data, _ := json.Marshal(userType)
	fmt.Println("JSON:", string(data))

}
