package types

/*
import "encoding/json"

func Header(c string){
	return alg(type+alg+c)
}

type claims struct {
	Sub    string `json:"sub"`
	Iat    int    `json:"iat"`
	Exp    int    `json:"exp"`			 // header.alg(json.Marshal(claims))	string
	Role   string `json:"role"`
	UserId string `json:"user_id"`
}

type signature struct {
	hc string //hash(header + claims, key)
}
//															header.claims.sing -> server -> h + c -> newsing ?? sing
type JWT struct {
	hc string
	s signature
}

func testJWT(){
	j := &JWT{
		c: claims{},
		s: signature{},
	}
}
*/
