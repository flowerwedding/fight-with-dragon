// Time : 2019/10/19 14:38
// Author : MashiroC

// dragongame
package game

// Card.go something

// 回合开始 回合结束 攻击

func NewCard1(name string, cost, hurt int, eff ...Effect) *BaseCard {
	if cost < 0 || hurt < 0 {
		panic("some failed game status")
	}
	return &BaseCard{
		cardName: name,
		cost:     cost,
		hurtNum:  hurt,
		effect:   eff,
	}
}
