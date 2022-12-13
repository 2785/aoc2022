package main

import (
	"sort"
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

const d11Sample = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`

func TestD11P1(t *testing.T) {
	monkeysWithN := d11ParseMonkeys(d11Sample)

	monkeys := lo.Map(monkeysWithN, func(m lo.Tuple2[day11Monkey, int], _ int) day11Monkey {
		return m.A
	})

	d11Round(monkeys)

	require.Equal(t, []int{20, 23, 27, 26}, monkeys[0].items)
	require.Equal(t, []int{2080, 25, 167, 207, 401, 1046}, monkeys[1].items)

	require.Equal(t, []int{}, monkeys[2].items)
	require.Equal(t, []int{}, monkeys[3].items)

	for i := 0; i < 19; i++ {
		d11Round(monkeys)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectCount > monkeys[j].inspectCount
	})

	business := monkeys[0].inspectCount * monkeys[1].inspectCount

	require.Equal(t, 10605, business)
}

func TestD11P2(t *testing.T) {
	monkeysWithN := d11ParseMonkeys(d11Sample)

	monkeys := lo.Map(monkeysWithN, func(m lo.Tuple2[day11Monkey, int], _ int) day11Monkey {
		return m.A
	})

	divisors := lo.Map(monkeysWithN, func(m lo.Tuple2[day11Monkey, int], _ int) int {
		return m.B
	})

	base := lcm(divisors...)

	for i := 0; i < 10000; i++ {
		d11RoundWithBase(monkeys, base)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectCount > monkeys[j].inspectCount
	})

	business := monkeys[0].inspectCount * monkeys[1].inspectCount

	require.Equal(t, 2713310158, business)
}
