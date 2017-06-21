package bin_api_reflect

import "reflect"

const SwIfIndex = "SwIfIndex"
const Retval = "Retval"
const Reply = "Reply"

func IsReplySwIfIdx(reply reflect.Type) bool {
	_, found := reply.FieldByName(SwIfIndex)
	return found
}

func SetSwIfIdx(reply reflect.Value, swIfIndex uint32) {
	if field := reply.FieldByName(SwIfIndex); field.IsValid() {
		field.Set(reflect.ValueOf(swIfIndex))
	}
}

func SetRetVal(reply reflect.Value, retVal int32) {
	if field := reply.FieldByName(Retval); field.IsValid() {
		field.Set(reflect.ValueOf(retVal))
	}
}

func ReplyNameFor(request string) (string, bool) {
	return request + Reply, true
}