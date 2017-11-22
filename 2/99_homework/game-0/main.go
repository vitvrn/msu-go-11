package main

//TODO !!! СОСТОЯНИЯ - КАК РЕАЛИЗОВАТЬ ???
//TODO !!!
//TODO !!! Player <- доступные_действия -> Room
//TODO !!!        ?                    ????
//TODO !!!
//TODO !!! Пока не реализованные интересные идеи из readme.md:
//TODO !!! - функций которые описывают какой-то интерактив в комнате
//TODO !!! - триггер (действие, выполняемое при каком-то событии)
//TODO !!! - у структуры поле может иметь тип "функция"

import (
	"fmt"
	"strings"
)

var kitchen, hall, room, street Location
var player Player
var locationNames map[string]*Location
var actions map[string]PlayerAction
var LookAround, Move, PutOn, Take, Use PlayerAction

type Item struct { //чай(?), рюкзак*, ключи, конспекты, дверь**(?),
	name     string
	wearable bool //*рюкзак: true //??? перенести в движок
	/* также нужно знать, к какому предмету можно применить,
	какой будет результат (ИЛИ отдать эту ф-цию движку игры)
	*/
}

type ActionResult string //TODO перечислимый тип???

//type WearableItem string //TODO перечислимый тип?
//type Wearable interface { //TODO ???
//	Wear()
//}

type Location struct { //дом{кухня, коридор, комната}, улица
	name  string
	items []Item
	//	place string //стол, стул //???
	paths []Path //куда можно пройти отсюда
}
type Path struct {
	where      string //*Location
	doorLocked bool
}

func (r Location) LookAround() ActionResult {
	//1) ты находишься... / ты в... / ты на...
	//2)
	return ""
} //???

type Player struct {
	curRoom   *Location
	puttedOn  []Item //WearableItem?
	inventory []Item
}
type PlayerAction func([]string) ActionResult

func (p Player) handleCommand(commandLine string) (result string) { //???NOT Player's func
	//TODO
	//parse commandline == command param1 param2 param3 == []command
	splittedCommandLine := strings.Split(commandLine, " ")
	fmt.Println(splittedCommandLine)

	if len(splittedCommandLine) < 1 {
		result = "Введите команду (added by vit)"
		return
	}
	action, ok := actions[splittedCommandLine[0]]
	if ok {
		result = string(action(splittedCommandLine[1:]))
	}

	return //???
}

func initGame() {
	//TODO "ты находишься на кухне, на столе чай, надо собрать рюкзак и идти в универ. можно пройти - коридор"
	//init(player)
	locationNames = map[string]*Location{
		"кухня":   &kitchen,
		"коридор": &hall,
		"комната": &room,
		"улица":   &street,
	}
	for i, _ := range locationNames {
		*locationNames[i] = Location{name: i} //hall.name=коридор, !=домой
	}
	//TODO refactor HACK: "домой" - алиас для "коридор" (но в предыдущем цикле не участвует, чтобы не заменить hall.name)
	//Location aliases (фактически, locationNames теперь мог бы называться locationAliases):
	locationNames["домой"] = &hall

	actions = map[string]PlayerAction{
		"осмотреться": LookAround,
		"идти":        Move,
		"одеть":       PutOn,
		"взять":       Take,
		"применить":   Use,
	}

	player = Player{curRoom: &kitchen}
	//init(rooms: things_avail)
}

func main() {
	initGame()
	fmt.Println(player)
	fmt.Println(kitchen)
}
