package d11

const TEST_DATA = `Monkey 0:
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

const REAL_DATA = `Monkey 0:
Starting items: 54, 98, 50, 94, 69, 62, 53, 85
Operation: new = old * 13
Test: divisible by 3
  If true: throw to monkey 2
  If false: throw to monkey 1

Monkey 1:
Starting items: 71, 55, 82
Operation: new = old + 2
Test: divisible by 13
  If true: throw to monkey 7
  If false: throw to monkey 2

Monkey 2:
Starting items: 77, 73, 86, 72, 87
Operation: new = old + 8
Test: divisible by 19
  If true: throw to monkey 4
  If false: throw to monkey 7

Monkey 3:
Starting items: 97, 91
Operation: new = old + 1
Test: divisible by 17
  If true: throw to monkey 6
  If false: throw to monkey 5

Monkey 4:
Starting items: 78, 97, 51, 85, 66, 63, 62
Operation: new = old * 17
Test: divisible by 5
  If true: throw to monkey 6
  If false: throw to monkey 3

Monkey 5:
Starting items: 88
Operation: new = old + 3
Test: divisible by 7
  If true: throw to monkey 1
  If false: throw to monkey 0

Monkey 6:
Starting items: 87, 57, 63, 86, 87, 53
Operation: new = old * old
Test: divisible by 11
  If true: throw to monkey 5
  If false: throw to monkey 0

Monkey 7:
Starting items: 73, 59, 82, 65
Operation: new = old + 6
Test: divisible by 2
  If true: throw to monkey 4
  If false: throw to monkey 3`
