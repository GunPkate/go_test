package main

import (
	"fmt"
	"math"
)

type Plays map[string]Play

type Play struct {
	Name string
	Type string
}

type Performance struct {
	PlayID   string `json:"playID"`
	Audience int    `json:"audience"`
}

type Invoice struct {
	Customer     string        `json:"customer"`
	Performances []Performance `json:"performances"`
}

func playType(play Play) string {
	return play.Type
}

func playName(play Play) string {
	return play.Name
}

func playFor(plays Plays, perf Performance) Play {
	return plays[perf.PlayID]
}

func amountFor(perf Performance, play Play) float64 {
	result := 0.0

	switch playType(play) {
	case "tragedy":
		result = 40000
		if perf.Audience > 30 {
			result += 1000 * (float64(perf.Audience - 30))
		}
	case "comedy":
		result = 30000
		if perf.Audience > 20 {
			result += 10000 + 500*(float64(perf.Audience-20))
		}
		result += 300 * float64(perf.Audience)
	default:
		panic(fmt.Sprintf("unknow type: %s", playType(play)))
	}
	return result
}

func totalAmountFor(invoice Invoice, plays Plays) float64 {
	result := 0.0
	for _, perf := range invoice.Performances {
		result += amountFor(perf, playFor(plays, perf))
	}
	return result
}

// invoice.Performances (push up dependency)
func totalVolumeCreditsFor(performances []Performance, plays Plays) float64 {
	result := 0.0
	for _, perf := range performances {
		result += math.Max(float64(perf.Audience-30), 0)
		// add extra credit for every ten comedy attendees
		if "comedy" == playType(playFor(plays, perf)) {
			result += math.Floor(float64(perf.Audience / 5))
		}
	}
	return result
}

type Bill struct {
	Customer           string
	Rates              []Rate
	TotalAmount        float64
	TotalVolumeCredits float64
}

type Rate struct {
	Play          Play
	Amount        float64
	VolumeCredits float64
	Audience      int
}

func statement(invoice Invoice, plays Plays) string {
	rates := []Rate{}

	for _, perf := range invoice.Performances {
		r := Rate{
			Play:   playFor(plays, perf),
			Amount: amountFor(perf, playFor(plays, perf)),
			// VolumeCredits: totalVolumeCreditsFor(perf, plays),
			Audience: perf.Audience,
		}
		rates = append(rates, r)
	}

	bill := Bill{
		Customer:           invoice.Customer,
		Rates:              rates,
		TotalAmount:        totalAmountFor(invoice, plays),
		TotalVolumeCredits: totalVolumeCreditsFor(invoice.Performances, plays),
	}

	return renderPlainText(invoice, plays, bill)
}

func renderPlainText(invoice Invoice, plays Plays, bill Bill) string {

	result := fmt.Sprintf("Statement for %s\n", bill.Customer)
	for _, r := range bill.Rates {
		result += fmt.Sprintf("  %s: $%.2f (%d seats)\n", r.Play.Name, r.Amount/100, r.Audience)
	}
	result += fmt.Sprintf("Amount owed is $%.2f\n", bill.TotalAmount/100)
	result += fmt.Sprintf("you earned %.0f credits\n", bill.TotalVolumeCredits)
	return result
}

func main() {
	inv := Invoice{
		Customer: "Bigco",
		Performances: []Performance{
			{PlayID: "hamlet", Audience: 55},
			{PlayID: "as-like", Audience: 35},
			{PlayID: "othello", Audience: 40},
		}}
	plays := map[string]Play{
		"hamlet":  {Name: "Hamlet", Type: "tragedy"},
		"as-like": {Name: "As You Like It", Type: "comedy"},
		"othello": {Name: "Othello", Type: "tragedy"},
	}

	bill := statement(inv, plays)
	fmt.Println(bill)
}
