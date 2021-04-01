package helpersmenu

import (
	"errors"
	"fmt"
)

// Menu is a menu
type Menu struct {
	main    map[string]func(args []string) error
	submenu *Menu
}

func (menu *Menu) checkSize(n int, args []string) bool {
	return (n + 1) >= len(args)
}

func (menu *Menu) callRecursive(args []string, menus *Menu, commandCount int) error {
	sizeok := menu.checkSize(commandCount, args)

	function, ok := menus.main[args[commandCount]]

	// se tem funcao
	if ok && function != nil {

		if sizeok {
			function(args[commandCount:])
			return nil
		}

		function(args[commandCount+1:])
		return nil

	}

	if menus.submenu != nil {
		if function == nil {
			if !sizeok {
				commandCount++
			}
		}

		if !sizeok {
			menu.callRecursive(args, menus.submenu, commandCount)
		}

		return errors.New("command not found")

	}

	return nil

}

// CreateSubMenu create a submenu
func (menu *Menu) CreateSubMenu(sub *Menu) {
	menu.submenu = sub

}

// Call call menu
func (menu *Menu) Call(args []string) {
	err := menu.callRecursive(args, menu, 0)
	if err != nil {
		fmt.Println(err)
	}

}

// CreateItem create menu
func (menu *Menu) CreateItem(option string, function func(args []string) error) *Menu {
	menu.main[option] = function
	return menu
}

// NewMenu Ioc
func NewMenu() *Menu {
	return &Menu{
		main: make(map[string]func(args []string) error),
	}
}
