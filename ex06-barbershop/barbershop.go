package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	procs = new(sync.WaitGroup)
)

type Shop struct {
	barber   *Barber
	WaitRoom chan *Client
}

type Client struct {
	name string
}

type Barber struct {
	name      string
	SleepChan chan *Client
}

func CreateShop(barber *Barber, seats int) *Shop {
	shop := new(Shop)
	shop.barber = barber
	shop.WaitRoom = make(chan *Client, seats)
	return shop
}

func CreateBarber(name string) *Barber {
	barber := new(Barber)
	barber.name = name
	barber.SleepChan = make(chan *Client)
	return barber
}

func NewCustomer(name string) *Client {
	customer := new(Client)
	customer.name = name
	return customer
}

func (c *Client) LetsHaveHaircut(shop *Shop) {
	fmt.Printf("Client %s entered the barbershop\n", c.name)
	select {
	case shop.WaitRoom <- c:
		fmt.Printf("Client %s found the seat in Waitting Room!\n", c.name)
		select {
		case shop.barber.SleepChan <- c:
			fmt.Printf("Customer %s waked up barber %s", c.name, shop.barber.name)
		default:
		}
	default:
		fmt.Printf("No seats available!")
		time.Sleep(5 * time.Second)
		c.LetsHaveHaircut(shop)
	}
}

func (barber *Barber) StartWork(shop *Shop) {
	for {
		select {
		case c := <-shop.WaitRoom:
			fmt.Printf("Barber %s doing haircut to %s...\n", barber.name, c.name)
			time.Sleep(3 * time.Second)
			fmt.Printf("%s got new haircut!\n", c.name)
			procs.Done()
		default:
			fmt.Printf("Barber %s is sleeping now\n", barber.name)
			c := <-barber.SleepChan
			fmt.Printf("Barber %s was waked by %s\n", barber.name, c.name)
		}
	}
}

func main() {
	procs.Add(8)
	barber := CreateBarber("Mr. Routine")
	barbershop := CreateShop(barber, 1)
	go barber.StartWork(barbershop)
	time.Sleep(3 * time.Second)
	clients := []string{"A", "B", "C", "D", "E", "F", "G", "H"}
	for _, i := range clients {
		customer := NewCustomer(i)
		time.Sleep(1 * time.Second)
		go customer.LetsHaveHaircut(barbershop)
	}
	procs.Wait()
}
