package model

func ReturnResult(data interface{},code int,err string)map[string]interface{}  {
	result := make(map[string]interface{})
	status := true
	if code < 9000 && code != 200{
		status = false
	}
	result = map[string]interface{}{
		"status": status,
		"code"  : code,
		"data"  : data,
		"err"   : err,
	}
	return result
}
