package delivery

import "three/v1/user/domain"

func GetAllUserMapper(users []domain.User) []domain.GetAllUserRespose {
	if users == nil {
		return []domain.GetAllUserRespose{}
	}
	var res []domain.GetAllUserRespose
	for _, user := range users {
		var r domain.GetAllUserRespose
		r.ID = user.ID
		r.Username = user.Username
		res = append(res, r)
	}
	return res
}
