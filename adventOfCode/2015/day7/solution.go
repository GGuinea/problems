package solution

import (
	"os"
	"strconv"
	"strings"
)

const (
	SIZE_X = 1000
	SIZE_Y = 1000
)

func readFile(filename string) string {
	res, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(res)
}

func part1() uint16 {
	input := readFile("input")
	return handlePart1(input, "a")
}

func part2() uint16 {
	input := readFile("input")
	return handlePart2(input, "a")
}

type wireCondition struct {
	left  string
	right string
	op    string
	value uint16
	done  bool
}

func getInt(val string) uint16 {
	res, _ := strconv.Atoi(val)
	return uint16(res)
}

func handlePart1(input, cable string) uint16 {
	conditionMap := parseInputWires(input)
	return solve(conditionMap, cable)
}

func handlePart2(input, cable string) uint16 {
	cableA := handlePart1(input, "a")
	conditionMap := parseInputWires(input)
	bObj := conditionMap["b"]

	bObj.done = true
	bObj.value = cableA

	return solve(conditionMap, cable)
}

func getObjValue(conditionMap map[string]*wireCondition, cable string) (res uint16) {
	if cable != "-" {
		lObj, found := conditionMap[cable]
		if !found {
			res = getInt(cable)
		} else if !lObj.done {
			res = solve(conditionMap, cable)
			lObj.done = true
			lObj.value = res
		} else if lObj.done {
			res = lObj.value
		}
	}
	return
}

func solve(conditionMap map[string]*wireCondition, requestedCable string) uint16 {
	obj := conditionMap[requestedCable]

	if obj.done {
		return obj.value
	}

	var lObjVal, rObjVal uint16
	lObjVal = getObjValue(conditionMap, obj.left)
	rObjVal = getObjValue(conditionMap, obj.right)

	switch obj.op {
	case "AND":
		return lObjVal & rObjVal
	case "OR":
		return lObjVal | rObjVal
	case "LSHIFT":
		return lObjVal << rObjVal
	case "RSHIFT":
		return lObjVal >> rObjVal
	case "NOT":
		return ^rObjVal
	case "-":
		return lObjVal
	default:
		panic(obj.op)
	}
}

func parseInputWires(input string) map[string]*wireCondition {
	res := make(map[string]*wireCondition, 0)
	conditions := strings.Split(input, "\n")
	for _, line := range conditions {
		if len(line) == 0 {
			continue
		}
		split := strings.Split(line, " -> ")
		cable := split[1]
		splitCondition := strings.Split(split[0], " ")
		splitConditionLen := len(splitCondition)
		switch splitConditionLen {
		case 1:
			val, err := strconv.Atoi(splitCondition[0])
			if err != nil {
				res[cable] = &wireCondition{
					left:  splitCondition[0],
					right: "-",
					op:    "-",
					value: 0,
					done:  false,
				}
			} else {
				res[cable] = &wireCondition{
					left:  "-",
					right: "-",
					op:    "-",
					value: uint16(val),
					done:  true,
				}
			}
		case 2:
			if splitCondition[0] == "NOT" {
				res[cable] = &wireCondition{
					left:  "-",
					right: splitCondition[1],
					op:    "NOT",
					value: 0,
					done:  false,
				}
			}
		case 3:
			res[cable] = &wireCondition{
				left:  splitCondition[0],
				op:    splitCondition[1],
				right: splitCondition[2],
				value: 0,
				done:  false,
			}
		}
	}
	return res
}
