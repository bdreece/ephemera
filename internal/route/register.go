package route

// import (
// 	"net/http"
//
// 	"github.com/ajg/form"
// 	"github.com/bdreece/ephemera/pkg/database"
// 	"github.com/bdreece/ephemera/pkg/identity"
// 	"github.com/bdreece/ephemera/pkg/security"
// 	"github.com/gofrs/uuid"
// )
//
// func (route *IdentityRoute) Register(w http.ResponseWriter, r *http.Request) {
// 	var input struct {
// 		FirstName   string            `form:"firstName"`
// 		LastName    string            `form:"lastName"`
// 		DisplayName string            `form:"username"`
// 		Password    identity.Password `form:"password"`
// 	}
//
// 	if err := form.NewDecoder(r.Body).Decode(&input); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
//
// 	uuid, _ := uuid.NewV4()
// 	hash, salt := input.Password.Hash()
//
//
// }
