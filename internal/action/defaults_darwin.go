package action

// DefaultBindings returns a map containing micro's default keybindings
func DefaultBindings() map[string]string {
	return map[string]string{
		"Up":             "CursorUp",
		"Down":           "CursorDown",
		"Right":          "CursorRight",
		"Left":           "CursorLeft",
		"ShiftUp":        "SelectUp",
		"ShiftDown":      "SelectDown",
		"ShiftLeft":      "SelectLeft",
		"ShiftRight":     "SelectRight",
		"AltLeft":        "WordLeft",
		"AltRight":       "WordRight",
		"AltUp":          "MoveLinesUp",
		"AltDown":        "MoveLinesDown",
		"AltShiftRight":  "SelectWordRight",
		"AltShiftLeft":   "SelectWordLeft",
		"CtrlLeft":       "StartOfLine",
		"CtrlRight":      "EndOfLine",
		"CtrlShiftLeft":  "SelectToStartOfLine",
		"ShiftHome":      "SelectToStartOfLine",
		"CtrlShiftRight": "SelectToEndOfLine",
		"ShiftEnd":       "SelectToEndOfLine",
		"CtrlUp":         "CursorStart",
		"CtrlDown":       "CursorEnd",
		"CtrlShiftUp":    "SelectToStart",
		"CtrlShiftDown":  "SelectToEnd",
		"Alt-{":          "ParagraphPrevious",
		"Alt-}":          "ParagraphNext",
		"Enter":          "InsertNewline",
		"CtrlH":          "Backspace",
		"Backspace":      "Backspace",
		"Alt-CtrlH":      "DeleteWordLeft",
		"Alt-Backspace":  "DeleteWordLeft",
		"Tab":            "Autocomplete|IndentSelection|InsertTab",
		"Backtab":        "CycleAutocompleteBack|OutdentSelection|OutdentLine",
		"CtrlO":          "OpenFile",
		"CtrlS":          "Save",
		"CtrlF":          "Find",
		"CtrlN":          "FindNext",
		"CtrlP":          "FindPrevious",
		"CtrlZ":          "Undo",
		"CtrlY":          "Redo",
		"CtrlC":          "Copy",
		"CtrlX":          "Cut",
		"CtrlK":          "CutLine",
		"CtrlD":          "DuplicateLine",
		"CtrlV":          "Paste",
		"CtrlA":          "SelectAll",
		"CtrlT":          "AddTab",
		"Alt,":           "PreviousTab",
		"Alt.":           "NextTab",
		"Home":           "StartOfLine",
		"End":            "EndOfLine",
		"CtrlHome":       "CursorStart",
		"CtrlEnd":        "CursorEnd",
		"PageUp":         "CursorPageUp",
		"PageDown":       "CursorPageDown",
		"CtrlPageUp":     "PreviousTab",
		"CtrlPageDown":   "NextTab",
		"CtrlG":          "ToggleHelp",
		"Alt-g":          "ToggleKeyMenu",
		"CtrlR":          "ToggleRuler",
		"CtrlL":          "command-edit:goto ",
		"Delete":         "Delete",
		"CtrlB":          "ShellMode",
		"CtrlQ":          "Quit",
		"CtrlE":          "CommandMode",
		"CtrlW":          "NextSplit",
		"CtrlU":          "ToggleMacro",
		"CtrlJ":          "PlayMacro",
		"Insert":         "ToggleOverwriteMode",

		// Emacs-style keybindings
		"Alt-f": "WordRight",
		"Alt-b": "WordLeft",
		"Alt-a": "StartOfLine",
		"Alt-e": "EndOfLine",
		// "Alt-p": "CursorUp",
		// "Alt-n": "CursorDown",

		// Integration with file managers
		"F2":  "Save",
		"F3":  "Find",
		"F4":  "Quit",
		"F7":  "Find",
		"F10": "Quit",
		"Esc": "Escape",

		// Mouse bindings
		"MouseWheelUp":   "ScrollUp",
		"MouseWheelDown": "ScrollDown",
		"MouseLeft":      "MousePress",
		"MouseMiddle":    "PastePrimary",
		"Ctrl-MouseLeft": "MouseMultiCursor",

		"Alt-n": "SpawnMultiCursor",
		"Alt-m": "SpawnMultiCursorSelect",
		"Alt-p": "RemoveMultiCursor",
		"Alt-c": "RemoveAllMultiCursors",
		"Alt-x": "SkipMultiCursor",
	}
}
