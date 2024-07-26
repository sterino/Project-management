package user

//type Request struct {
//	Title       string `json:"title"`
//	Description string `json:"description"`
//	ManagerID   string `json:"manager_id"`
//}
//
//type Response struct {
//	Title       string `json:"title"`
//	Description string `json:"description"`
//	ManagerID   string `json:"manager_id"`
//}
//
//func ParseFromEntity(entity Entity) Response {
//	return Response{
//		Title:       entity.Title,
//		Description: entity.Description,
//		ManagerID:   entity.ManagerID,
//	}
//}
//
//func ParseFromEntities(data []Entity) (res []Response) {
//	res = make([]Response, 0)
//	for _, entity := range data {
//		res = append(res, ParseFromEntity(entity))
//	}
//	return
//}
