package main

import "log"

// --- command ---

type Command interface {
	Execute() error
	Rollback() error
}

type CartCommand struct{}

func (c CartCommand) Execute() error {
	log.Println("add product")
	return nil
}

func (c CartCommand) Rollback() error {
	log.Println("drop product")
	return nil
}

type StockCommand struct{}

func (c StockCommand) Execute() error {
	log.Println("pick up")
	return nil
}

func (c StockCommand) Rollback() error {
	log.Println("pick off")
	return nil
}

type ShippingCommand struct{}

func (c ShippingCommand) Execute() error {
	log.Println("pick up")
	return nil
}

func (c ShippingCommand) Rollback() error {
	log.Println("pick off")
	return nil
}

type SettlementCommand struct{}

func (c SettlementCommand) Execute() error {
	log.Println("pay")
	return nil
}

func (c SettlementCommand) Rollback() error {
	log.Println("refund")
	return nil
}

// --- order ---

type Order struct {
	commands []Command
}

func NewOrder(commands []Command) Order {
	return Order{commands: commands}
}

func (o Order) add_command(command Command) {
	o.commands = append(o.commands, command)
}

func (o Order) execute() error {
	for _, command := range o.commands {
		if err := command.Execute(); err != nil {
			return err
		}
	}
	return nil
}

func (o Order) rollback() error {
	// 後入れ先出し
	for i := len(o.commands); i >= 0; i-- {
		if err := o.commands[i].Rollback(); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	/*
		1. カート情報取得
		2. 在庫検索
		3. 決済
		4. 出荷
	*/
	cart := CartCommand{}
	stock := StockCommand{}
	shipping := ShippingCommand{}
	pay := SettlementCommand{}

	order := NewOrder([]Command{
		cart, stock, pay, shipping,
	})
	if err := order.execute(); err != nil {
		order.rollback()
	}
}
