package playlist

import (
	"github.com/tomnz/glowpher/config"
	"github.com/tomnz/glowpher/internal/effect"
	"github.com/tomnz/glowpher/internal/variable"
)

// Decompile builds the API response from the current Glowpher configuration.
func Decompile(playlist *Playlist) (*config.API, error) {
	api := &config.API{}

	// Build options
	api.EffectOptions = make([]config.EffectOption, len(effect.Registry))
	idx := 0
	for typ, eff := range effect.Registry {
		api.EffectOptions[idx] = config.EffectOption{
			Type:          typ,
			DefaultParams: decompileParams(eff.DefaultParams()),
		}
		idx++
	}

	api.VariableOptions = make([]config.VariableOption, len(variable.Registry))
	idx = 0
	for typ, vari := range variable.Registry {
		api.VariableOptions[idx] = config.VariableOption{
			Type:          typ,
			DefaultParams: decompileParams(vari.DefaultParams()),
		}
	}

	return nil, nil
}

func decompileParams(params map[string]variable.Param) map[string]config.Param {
	cfg := map[string]config.Param{}
	for name, param := range params {
		if pvar, ok := param.(variable.ParamVariable); ok {
			switch param.(type) {
			case variable.ParamFloat:
				varName := pvar.Variable().Name()
				cfg[name] = config.Param{
					Type:     config.ParamFloatType,
					Variable: &varName,
					//VariableParams: decompileParams(pvar.Variable().Params()),
				}
			}
		} else {
			//switch p := param.(type) {
			//case variable.ParamColor:
			//	cfg[name] = config.Param{
			//		Type: config.ParamColorType,
			//		//Value: config.ParamValue{
			//		//	Color: p.Value(),
			//		//},
			//	}
			//
			//default:
			//panic(fmt.Sprintf("unhandled param value type %q", param.Type()))
			//			}
		}
	}
	return nil
}
