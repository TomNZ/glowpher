package dsl

import (
	"fmt"
	"sort"

	"github.com/tomnz/glowpher/config"
	"github.com/tomnz/glowpher/internal/color"
	"github.com/tomnz/glowpher/internal/effect"
	"github.com/tomnz/glowpher/internal/variable"
)

// Decompile builds the API response from the current Glowpher configuration.
func Decompile(cfg *Config) *config.API {
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
		idx++
	}

	// Build Variables
	vars := make([]config.Variable, len(cfg.Variables))
	idx = 0
	for name, vari := range cfg.Variables {
		vars[idx] = config.Variable{
			Name:   name,
			Type:   vari.Type(),
			Params: decompileParams(vari.Params()),
		}
		idx++
	}
	api.Config.Variables = vars

	// Build Scenes
	scenes := make([]config.Scene, len(cfg.Scenes))
	idx = 0
	for _, scene := range cfg.Scenes {
		scenes[idx] = decompileScene(scene)
		idx++
	}
	sort.Slice(scenes, func(i, j int) bool {
		return scenes[i].Name < scenes[j].Name
	})
	api.Config.Scenes = scenes

	// Build Playlists
	playlists := make([]config.Playlist, len(cfg.Playlists))
	idx = 0
	for _, playlist := range cfg.Playlists {
		playlists[idx] = decompilePlaylist(playlist)
		idx++
	}
	sort.Slice(playlists, func(i, j int) bool {
		return playlists[i].Name < playlists[j].Name
	})
	api.Config.Playlists = playlists

	return api
}

func decompileParams(params map[string]variable.Param) map[string]config.Param {
	cfg := map[string]config.Param{}
	for name, param := range params {
		if pvar, ok := param.(variable.ParamVariable); ok {
			varName := pvar.Variable().Name()
			switch pvarTyped := param.(type) {
			case variable.ParamFloatVariable:
				cfg[name] = config.Param{
					Type:     config.ParamFloatType,
					Variable: &varName,
					VariableParams: &config.VariableParams{
						Multiply: pvarTyped.Multiply(),
						Add:      pvarTyped.Add(),
					},
				}

			case variable.ParamIntVariable:
				cfg[name] = config.Param{
					Type:     config.ParamIntType,
					Variable: &varName,
					VariableParams: &config.VariableParams{
						Multiply: pvarTyped.Multiply(),
						Add:      pvarTyped.Add(),
					},
				}

			default:
				panic(fmt.Sprintf("unhandled variable param type %T", pvar))
			}
		} else {
			switch p := param.(type) {
			case variable.ParamColor:
				cfg[name] = config.Param{
					Type: config.ParamColorType,
					Value: &config.ParamValue{
						Color: decompileColor(p.Value()),
					},
				}

			case variable.ParamFloat:
				value := p.Value()
				cfg[name] = config.Param{
					Type: config.ParamFloatType,
					Value: &config.ParamValue{
						Float: &value,
					},
				}

			case variable.ParamInt:
				value := p.Value()
				cfg[name] = config.Param{
					Type: config.ParamFloatType,
					Value: &config.ParamValue{
						Int: &value,
					},
				}

			case variable.ParamString:
				value := p.Value()
				cfg[name] = config.Param{
					Type: config.ParamFloatType,
					Value: &config.ParamValue{
						String: &value,
					},
				}

			default:
				panic(fmt.Sprintf("unhandled param value type %T", param))
			}
		}
	}
	return cfg
}

func decompileScene(scene *Scene) config.Scene {
	effects := make([]config.SceneEffect, len(scene.Effects))
	for idx := range effects {
		effects[idx] = decompileEffect(scene.Effects[idx])
	}

	return config.Scene{
		Name:    scene.Name,
		Effects: effects,
	}
}

func decompileEffect(eff effect.Effect) config.SceneEffect {
	return config.SceneEffect{
		Type:   eff.Type(),
		Params: decompileParams(eff.Params()),
	}
}

func decompilePlaylist(playlist *Playlist) config.Playlist {
	scenes := make([]config.PlaylistScene, len(playlist.Scenes))
	for idx, scene := range playlist.Scenes {
		var duration *string
		if scene.Duration != 0 {
			durationStr := scene.Duration.String()
			duration = &durationStr
		}

		scenes[idx] = config.PlaylistScene{
			Name:     scene.Name,
			Duration: duration,
		}
	}

	return config.Playlist{
		Name:            playlist.Name,
		Scenes:          scenes,
		DefaultDuration: playlist.DefaultDuration.String(),
	}
}

func decompileColor(col color.Color) *config.Color {
	return &config.Color{
		R: col.R,
		G: col.G,
		B: col.B,
	}
}
