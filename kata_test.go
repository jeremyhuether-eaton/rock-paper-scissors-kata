/*
	Player one paper, player two rock; player one wins
	Player one paper, player two scissors; player one loses
	Player one paper, player two paper; tie
	Player one rock, player two scissors; player one wins
	Player one rock, player two paper; player one loses
	Player one rock, player two rock; tie
	Player one scissors, player two paper; player one wins
	Player one scissors, player two rock; player one loses
	Player one scissors, player two scissors; tie
*/

package main

import (
	"testing"
)

func TestRpcKata(t *testing.T) {
	t.Run("Player one paper, player two rock; player one wins", func(t *testing.T) {

		got := Kata(PAPER, ROCK)
		want := PLAYER_ONE_WINS

		if got != want {
			t.Fatalf("Fail")
		}
	})

	t.Run("Player one paper, player two scissors; player one loses", func(t *testing.T) {

		got := Kata(PAPER, SCISSORS)
		want := PLAYER_ONE_LOSES

		if got != want {
			t.Fatalf("Fail")
		}
	})
	t.Run("Player one paper, player two paper; tie", func(t *testing.T) {

		got := Kata(PAPER, PAPER)
		want := TIE

		if got != want {
			t.Fatalf("Fail")
		}
	})
	t.Run("Player one rock, player two scissors; player one wins",
		func(t *testing.T) {

			got := Kata(ROCK, SCISSORS)
			want := PLAYER_ONE_WINS

			if got != want {
				t.Fatalf("Fail")
			}
		})
	t.Run("	Player one rock, player two paper; player one loses", func(t *testing.T) {
		// line 6

		got := Kata(ROCK, PAPER)
		want := PLAYER_ONE_LOSES

		if got != want {
			t.Fatalf("Fail")
		}
	})
	t.Run("	Player one rock, player two rock; tie", func(t *testing.T) {
		// line 7

		got := Kata(ROCK, ROCK)
		want := TIE

		if got != want {
			t.Fatalf("Fail")
		}
	})
	t.Run("	Player one scissors, player two paper; player one wins", func(t *testing.T) {
		// line 8

		got := Kata(SCISSORS, PAPER)
		want := PLAYER_ONE_WINS

		if got != want {
			t.Fatalf("Fail")
		}
	})
	t.Run(" Player one scissors, player two rock; player one loses", func(t *testing.T) {
		// line 9

		got := Kata(SCISSORS, ROCK)
		want := PLAYER_ONE_LOSES

		if got != want {
			t.Fatalf("Fail")
		}
	})
	t.Run(" Player one scissors, player two scissors; tie", func(t *testing.T) {
		// line 10

		got := Kata(SCISSORS, SCISSORS)
		want := TIE

		if got != want {
			t.Fatalf("Fail")
		}
	})

}

/*
Got: player one wins
            Want: tie

*/
