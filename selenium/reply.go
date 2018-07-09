package selenium

import (
	"errors"
	"strings"
)

/*
type Reply struct {
	SessionId string      `json:"sessionId,omitempty"`
	Status    int         `json:"status,omitempty"`
	Value     interface{} `json:"value,omitempty"`

} */

type Reply struct {
	StatusCode int
	Data       map[string]interface{}
}

func (reply *Reply) Get(name string, useDotNotation bool) (interface{}, error) {

	if useDotNotation {

		reply, err := parseReply(name, reply.Data)
		if err != nil {
			return nil, err
		}

		return reply, nil

	}

	if value, ok := reply.Data[name]; ok {
		return value, nil
	}
	return nil, errors.New("could not parse: " + name)
}

func parseReply(name string, data interface{}) (interface{}, error) {

	keys := strings.Split(name, ".")

	if len(keys) != 0 {

		if object, ok := data.(map[string]interface{}); ok {

			if value, ok := object[keys[0]]; ok {
				if len(keys) == 1 {
					return value, nil
				}
				return parseReply(strings.Join(keys[1:], "."), value)
			}

			return nil, errors.New("could not parse: " + keys[0])

		}

	}

	return nil, errors.New("could not parse: " + keys[0])
}

func (reply *Reply) GetInt(name string, useDotNotation bool) (int, error) {

	value, err := reply.Get(name, useDotNotation)
	if err != nil {
		return 0, err
	}

	if integer, ok := value.(int); ok {
		return integer, nil
	}
	return 0, errors.New("could not parse int: " + name)

}

func (reply *Reply) GetFloat(name string, useDotNotation bool) (float64, error) {

	value, err := reply.Get(name, useDotNotation)
	if err != nil {
		return 0, err
	}

	if integer, ok := value.(float64); ok {
		return integer, nil
	}
	return 0, errors.New("could not parse int: " + name)

}

func (reply *Reply) GetBool(name string, useDotNotation bool) (bool, error) {

	value, err := reply.Get(name, useDotNotation)
	if err != nil {
		return false, err
	}

	if boolean, ok := value.(bool); ok {
		return boolean, nil
	}
	return false, errors.New("could not parse bool: " + name)

}

func (reply *Reply) GetString(name string, useDotNotation bool) (string, error) {

	value, err := reply.Get(name, useDotNotation)
	if err != nil {
		return "", err
	}

	if str, ok := value.(string); ok {
		return str, nil
	}
	return "", errors.New("could not parse string: " + name)

}

func (reply *Reply) GetStringSlice(name string, useDotNotation bool) ([]string, error) {

	value, err := reply.Get(name, useDotNotation)
	if err != nil {
		return nil, err
	}

	if str, ok := value.([]string); ok {
		return str, nil
	}
	return nil, errors.New("could not parse string slice: " + name)

}

func (reply *Reply) GetMap(name string, useDotNotation bool) (map[string]interface{}, error) {

	value, err := reply.Get(name, useDotNotation)
	if err != nil {
		return nil, err
	}

	if dict, ok := value.(map[string]interface{}); ok {
		return dict, nil
	}
	return nil, errors.New("could not parse map: " + name)

}

func (reply *Reply) GetStringMap(name string, useDotNotation bool) (map[string]string, error) {

	value, err := reply.Get(name, useDotNotation)
	if err != nil {
		return nil, err
	}

	if dict, ok := value.(map[string]interface{}); ok {

		stringMap := make(map[string]string, 0)

		for key, value := range dict {
			if str, ok := value.(string); ok {
				stringMap[key] = str
			} else {
				return nil, errors.New("could not parse string map(1): " + name)
			}

		}

		return stringMap, nil
	}

	return nil, errors.New("could not parse string map(2): " + name)

}
