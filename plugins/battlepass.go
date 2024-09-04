package plugins

import (
	"log/slog"
	"math"
	"reflect"
	"slices"
	"strings"

	"github.com/Tesohh/femto/editor"
	"github.com/Tesohh/femto/storage"
	"github.com/gdamore/tcell/v2"
)

type BattlePassData struct {
	Xp                   int
	SelectedLanguages    []string       // eg .go, .rs, .py
	CharsTyped           map[string]int // gets reset everyday
	CharsTypedObjectives map[string]int
}

const XpPerLevel = 1000

func (bp BattlePassData) Level() int {
	return int(math.Floor(float64(bp.Xp / XpPerLevel)))
}

type BattlePass struct {
	Storage storage.PersistentStorage[BattlePassData]
}

func (p *BattlePass) GetInfo() editor.PluginInfo {
	return editor.PluginInfo{
		Id:          "femto.battlepass",
		Author:      "femto",
		Name:        "Battle Pass",
		Description: "The optimal way to edit text, is through a battle pass",
	}
}

func (p *BattlePass) Startup(e *editor.Editor) error {
	p.Storage = storage.PersistentStorage[BattlePassData]{Id: "femto.battlepass"}
	err := p.Storage.Load()
	if err != nil {
		return err
	}

	// Default data
	if reflect.DeepEqual(p.Storage.Data, BattlePassData{}) {
		p.Storage.Data = BattlePassData{
			Xp:                   0,
			SelectedLanguages:    []string{"go", "rs", "testfile"},
			CharsTyped:           map[string]int{},
			CharsTypedObjectives: map[string]int{"go": 100, "rs": 100, "testfile": 60},
		}
		err = p.Storage.Save()
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *BattlePass) Update(e *editor.Editor, event tcell.Event) tcell.Event {
	// Check for new insertions
	if _, ok := event.(*editor.CharInsertedEvent); ok {
		sections := strings.Split(e.Win().FilePath, ".")
		ext := sections[len(sections)-1]
		slog.Info(ext)
		if slices.Contains(p.Storage.Data.SelectedLanguages, ext) {
			if _, ok := p.Storage.Data.CharsTyped[ext]; !ok {
				p.Storage.Data.CharsTyped[ext] = 0
			}
			p.Storage.Data.CharsTyped[ext] += 1
			p.Storage.Save()
		}
	}

	return nil
}

func (p *BattlePass) Draw(e *editor.Editor) error {
	return nil
}
