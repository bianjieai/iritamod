package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/manifoldco/promptui"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	types "github.com/bianjieai/iritamod/modules/params"
)

const (
	draftUpdateParamsFileName = "draft_update_params.json"
)

// Prompt prompts the user for all values of the given type.
// data is the struct to be filled
// namePrefix is the name to be displayed as "Enter <namePrefix> <field>"
func Prompt[T any](data T, namePrefix string) (T, error) {
	v := reflect.ValueOf(&data).Elem()
	if v.Kind() == reflect.Interface {
		v = reflect.ValueOf(data)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
	}

	for i := 0; i < v.NumField(); i++ {
		// if the field is a struct skip or not slice of string or int then skip
		switch v.Field(i).Kind() {
		case reflect.Struct:
			// TODO(@julienrbrt) in the future we can add a recursive call to Prompt
			continue
		case reflect.Slice:
			if v.Field(i).Type().Elem().Kind() != reflect.String &&
				v.Field(i).Type().Elem().Kind() != reflect.Int {
				continue
			}
		}

		// create prompts
		prompt := promptui.Prompt{
			Label: fmt.Sprintf(
				"Enter %s's %s",
				namePrefix,
				strings.ToLower(client.CamelCaseToString(v.Type().Field(i).Name)),
			),
			Validate: client.ValidatePromptNotEmpty,
		}

		fieldName := strings.ToLower(v.Type().Field(i).Name)

		if strings.EqualFold(fieldName, "authority") {
			// pre-fill with iritamod/params address
			prompt.Default = authtypes.NewModuleAddress(types.ModuleName).String()
			prompt.Validate = client.ValidatePromptAddress
		}

		result, err := prompt.Run()
		if err != nil {
			return data, fmt.Errorf("failed to prompt for %s: %w", fieldName, err)
		}

		switch v.Field(i).Kind() {
		case reflect.String:
			v.Field(i).SetString(result)
		case reflect.Int:
			resultInt, err := strconv.ParseInt(result, 10, 0)
			if err != nil {
				return data, fmt.Errorf("invalid value for int: %w", err)
			}
			v.Field(i).SetInt(resultInt)
		case reflect.Slice:
			switch v.Field(i).Type().Elem().Kind() {
			case reflect.String:
				v.Field(i).Set(reflect.ValueOf([]string{result}))
			case reflect.Int:
				resultInt, err := strconv.ParseInt(result, 10, 0)
				if err != nil {
					return data, fmt.Errorf("invalid value for int: %w", err)
				}
				v.Field(i).Set(reflect.ValueOf([]int{int(resultInt)}))
			}
		default:
			continue
		}
	}

	return data, nil
}

type updateParamsType struct {
	Name    string
	MsgType string
	Msg     sdk.Msg
}

func (p *updateParamsType) Prompt(cdc codec.Codec) (*updateParams, error) {
	updateParams := &updateParams{
		Messages: make([]json.RawMessage, 0),
	}

	if p.Msg == nil {
		return updateParams, nil
	}

	// set messages field
	result, err := Prompt(p.Msg, "msg")
	if err != nil {
		return nil, fmt.Errorf("failed to set update params message: %w", err)
	}

	message, err := cdc.MarshalInterfaceJSON(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal update params message: %w", err)
	}
	updateParams.Messages = append(updateParams.Messages, message)

	return updateParams, nil
}

// writeFile writes the input to the file
func writeFile(fileName string, input any) error {
	raw, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal update params message: %w", err)
	}

	if err := os.WriteFile(fileName, raw, 0o600); err != nil {
		return err
	}

	return nil
}
