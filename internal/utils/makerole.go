package utils

func MakeRole(phone string) string {
    if phone == "09911732328" {
        return "admin"
    }
    return "user"
}
