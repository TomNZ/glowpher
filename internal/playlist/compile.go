package playlist

import (
	"fmt"
	"time"

	"github.com/tomnz/glowpher/config"
	"github.com/tomnz/glowpher/internal/effect"
	"github.com/tomnz/glowpher/internal/variable"
)

// Compile compiles the given configuration into a playlist.
func Compile(cfg config.Config) (*Playlist, error) {
	variables := map[string]variable.Variable{}
	for _, cfgVariable := range cfg.Variables {
		vari, err := compileVariable(cfgVariable)
		if err != nil {
			return nil, err
		}
		variables[vari.Name()] = vari
	}

	for name, vari := range variables {
		if err := wireParamVariables(vari.Params(), variables); err != nil {
			return nil, fmt.Errorf("variable %q: %s", name, err)
		}
	}

	scenes := make([]*Scene, len(cfg.Playlist.Scenes))
	for idx, cfgScene := range cfg.Playlist.Scenes {
		scene, err := compileScene(cfgScene, variables)
		if err != nil {
			return nil, err
		}
		scenes[idx] = scene
	}

	defaultDuration, err := compileDuration(&cfg.Playlist.DefaultDuration)
	if err != nil {
		return nil, fmt.Errorf("playlist default duration %q: %s", cfg.Playlist.DefaultDuration, err)
	}

	return &Playlist{
		scenes:          scenes,
		defaultDuration: defaultDuration,
	}, nil
}

func compileVariable(cfg config.Variable) (variable.Variable, error) {
	vari, ok := variable.Registry[cfg.Type]
	if !ok {
		return nil, fmt.Errorf("unknown variable type %q", cfg.Type)
	}

	params, err := compileParams(cfg.Params)
	if err != nil {
		return nil, fmt.Errorf("variable %q: %s", cfg.Name, err)
	}

	newVari, err := vari.New(cfg.Name, params)
	if err != nil {
		return nil, fmt.Errorf("variable %q: %s", cfg.Name, err)
	}

	return newVari, nil
}

func compileParams(cfg map[string]config.Param) (map[string]variable.Param, error) {
	if len(cfg) == 0 {
		return nil, nil
	}

	params := map[string]variable.Param{}
	for name, cfgParam := range cfg {
		if cfgParam.Variable != nil && cfgParam.Value != nil {
			return nil, fmt.Errorf("param %q must define a value or variable, not both", name)
		} else if cfgParam.Value != nil {
			switch cfgParam.Type {
			case config.ParamStringType:
				if cfgParam.Value.String == nil {
					return nil, fmt.Errorf("param %q must define a string value", name)
				}
				params[name] = variable.ParamStringLiteral(*cfgParam.Value.String)

			case config.ParamIntType:
				if cfgParam.Value.Int == nil {
					return nil, fmt.Errorf("param %q must define an integer value", name)
				}
				params[name] = variable.ParamIntLiteral(*cfgParam.Value.Int)

			case config.ParamFloatType:
				if cfgParam.Value.Float == nil {
					return nil, fmt.Errorf("param %q must define a float value", name)
				}
				params[name] = variable.ParamFloatLiteral(*cfgParam.Value.Float)

			case config.ParamColorType:
				if cfgParam.Value.Color == nil {
					return nil, fmt.Errorf("param %q must define a color value", name)
				}
				// This compiles/works, because the two color types have the same fields
				params[name] = variable.ParamColorLiteral(*cfgParam.Value.Color)

			default:
				panic(fmt.Sprintf("unhandled param value type %q", cfgParam.Type))
			}
		} else {
			switch cfgParam.Type {
			case config.ParamIntType:
				var (
					multiply float32 = 1.0
					add      float32
				)
				if cfgParam.VariableParams != nil {
					multiply = cfgParam.VariableParams.Multiply
					add = cfgParam.VariableParams.Add
				}
				params[name] = variable.NewParamIntVariable(*cfgParam.Variable, multiply, add)

			case config.ParamFloatType:
				var (
					multiply float32 = 1.0
					add      float32
				)
				if cfgParam.VariableParams != nil {
					multiply = cfgParam.VariableParams.Multiply
					add = cfgParam.VariableParams.Add
				}
				params[name] = variable.NewParamFloatVariable(*cfgParam.Variable, multiply, add)

			default:
				panic(fmt.Sprintf("unhandled param variable type %q", cfgParam.Type))
			}
		}
	}

	return params, nil
}

func wireParamVariables(params map[string]variable.Param, variables map[string]variable.Variable) error {
	for paramName, param := range params {
		paramVari, ok := param.(variable.ParamVariable)
		if !ok {
			continue
		}
		if err := paramVari.WireVariable(variables); err != nil {
			return fmt.Errorf("param %q: %s", paramName, err)
		}
	}
	return nil
}

func compileScene(cfg config.Scene, variables map[string]variable.Variable) (*Scene, error) {
	effects := make([]effect.Effect, len(cfg.Effects))
	for idx, cfgEffect := range cfg.Effects {
		eff, err := compileEffect(cfgEffect, variables)
		if err != nil {
			return nil, fmt.Errorf("scene %q: %s", cfg.Name, err)
		}
		effects[idx] = eff
	}

	duration, err := compileDuration(cfg.Duration)
	if err != nil {
		return nil, fmt.Errorf("scene %q: %s", cfg.Name, err)
	}

	return &Scene{
		name:     cfg.Name,
		effects:  effects,
		duration: duration,
	}, nil
}

func compileDuration(cfg *string) (time.Duration, error) {
	if cfg != nil {
		return time.ParseDuration(*cfg)
	}
	var duration time.Duration
	return duration, nil
}
func compileEffect(cfg config.SceneEffect, variables map[string]variable.Variable) (effect.Effect, error) {
	eff, ok := effect.Registry[cfg.Type]
	if !ok {
		return nil, fmt.Errorf("unknown effect type %q", cfg.Type)
	}

	params, err := compileParams(cfg.Params)
	if err != nil {
		return nil, fmt.Errorf("effect type %q: %s", cfg.Type, err)
	}

	if err := wireParamVariables(params, variables); err != nil {
		return nil, fmt.Errorf("effect type %q: %s", cfg.Type, err)
	}

	newEff, err := eff.New(params)
	if err != nil {
		return nil, fmt.Errorf("effect type %q: %s", cfg.Type, err)
	}
	return newEff, nil
}
