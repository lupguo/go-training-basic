package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"text/tabwriter"
	"time"
)

func init() {
	//helper.StatusLog(helper.ToFile, 800*time.Millisecond)
	rand.Seed(time.Now().UnixNano())
}

const (
	TimeLayout = "15:04:05"
)

type Dice rune
type Reward struct {
	score int
	prize rune
}

var (
	dcList = []Dice{
		'\u2680',
		'\u2681',
		'\u2682',
		'\u2683',
		'\u2684',
		'\u2685',
	}
	dcScore = map[Dice]int{
		'\u2680': 1,
		'\u2681': 2,
		'\u2682': 3,
		'\u2683': 4,
		'\u2684': 5,
		'\u2685': 6,
	}
	rewards = map[int]*Reward{
		18: {18, 0x1F947},
		//15:  {15, 0x1F948},
		//12:  {12, 0x1F948},
		//9:  {9, 0x1F949},
		//6:  {6, 0x1F949},
		3:  {3, 0x1F949},
	}
)

type Player struct {
	coins int
}

// Game statistics
type GameStat struct {
	turn  int
	score int
}

// Game information
type Game struct {
	begin   time.Time
	end     time.Time
	elems   []Dice
	rewards map[int]*Reward
	dch     chan Dice
	win     chan *Reward
	player  *Player
	stat    *GameStat
}

// NewGame new an game
func NewGame() *Game {
	g := &Game{
		begin:   time.Now(),
		elems:   dcList,
		rewards: rewards,
		dch:     make(chan Dice),
		win:     make(chan *Reward),
		player:  new(Player),
		stat:    new(GameStat),
	}
	return g
}

// Prepare ticker send rand dice to dice chan
func (g *Game) Prepare(d time.Duration) {
	for i := 0; ; i++ {
		// check chan whether close
		//if _, ok := <-g.dch; !ok {
		//	return
		//}
		select {
		case <-time.Tick(d):
			g.dch <- dcList[rand.Intn(6)]
		}
	}
}

// Wait for the game stop
func (g *Game) Wait() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	signCh := make(chan os.Signal)
	signal.Notify(signCh, os.Interrupt)
	go func() {
		select {
		case <-signCh:
			fmt.Println("\nPlayer shutdown the game!")
			wg.Done()
		case w := <-g.win:
			fmt.Printf("\nCongratulate!! You got leopard point(%d) => %c!!\n", w.score, w.prize)
			wg.Done()
		}
	}()
	wg.Wait()
	fmt.Println("Game Over!!")
}

// Halt the game
func (g *Game) Halt() {
	close(g.dch)
}

// Start run the game and show game information
func (g *Game) Start() {
	// format print
	tw := tabwriter.NewWriter(os.Stdout, 8, 4, 2, ' ', 0)
	_, _ = fmt.Fprintf(tw, "Time\tTurns\tResult\tScore\n---\t---\t---\t---\n")

	// print dice value
	dicNum, num, turn := 3, 0, 0
	score := 0
	for dice := range g.dch {
		// write turn
		if turn++; (turn-1)%dicNum == 0 {
			_, _ = fmt.Fprintf(tw, "%s\t%d\t", time.Now().Format(TimeLayout), turn/dicNum+1)
		}
		// dice print
		_, _ = fmt.Fprintf(tw, "%c ", dice)
		// score calc
		score += dcScore[dice]
		// new line
		if num++; num%dicNum == 0 {
			_, _ = fmt.Fprintf(tw, "\tpoint(%d)\n", score)
			_ = tw.Flush()
			// hand score
			g.ScoreHandle(&score)
		}
	}
}

// ScoreHandle handle score
func (g *Game) ScoreHandle(score *int) {
	if reward, ok := g.rewards[*score]; ok {
		g.win <- reward
	}
	*score = 0
}

func main() {
	g := NewGame()
	// background prepare running
	go g.Prepare(200 * time.Millisecond)
	// game start
	go g.Start()
	// wait
	g.Wait()
}
