package d2

const TEST_DATA = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

const REAL_DATA = `Game 1: 1 green, 6 red, 4 blue; 2 blue, 6 green, 7 red; 3 red, 4 blue, 6 green; 3 green; 3 blue, 2 green, 1 red
Game 2: 2 blue, 4 red, 7 green; 17 red, 3 blue, 2 green; 3 green, 14 red, 1 blue
Game 3: 12 blue, 3 red, 1 green; 8 blue, 9 red; 1 blue, 1 green, 9 red; 4 blue, 1 green, 9 red
Game 4: 2 red, 10 green, 5 blue; 11 blue, 4 green; 6 green, 7 blue, 2 red; 4 blue, 9 green; 6 green, 1 red, 5 blue
Game 5: 10 green, 5 blue, 5 red; 10 blue, 13 green; 2 red, 12 blue; 9 green, 9 red
Game 6: 2 red, 3 green; 1 blue, 15 red, 2 green; 1 green, 7 red
Game 7: 16 blue, 4 green, 9 red; 6 red, 2 blue, 12 green; 2 red, 5 green, 14 blue; 11 blue, 13 red; 10 blue, 3 red, 17 green; 1 green, 12 blue
Game 8: 14 red, 12 green, 1 blue; 5 blue, 7 green, 12 red; 8 green, 1 red, 8 blue; 8 blue, 2 green, 15 red; 9 blue, 12 red, 10 green; 4 blue, 15 red, 1 green
Game 9: 2 red, 7 green, 5 blue; 1 red, 5 blue, 13 green; 5 blue
Game 10: 4 red, 1 green, 4 blue; 7 green, 8 blue, 4 red; 9 green, 3 red, 8 blue; 5 red, 2 green, 7 blue
Game 11: 4 green, 1 blue, 1 red; 3 green, 3 red, 1 blue; 3 green, 1 red, 1 blue
Game 12: 7 red, 6 green, 12 blue; 6 blue, 8 green, 3 red; 12 green, 5 blue, 4 red; 3 red, 16 blue, 8 green; 12 red, 11 green, 6 blue
Game 13: 2 green, 5 red, 12 blue; 8 green, 12 red, 4 blue; 6 blue, 7 green, 13 red
Game 14: 1 blue, 7 green, 5 red; 1 blue, 8 green, 6 red; 3 green, 1 blue, 4 red
Game 15: 11 red, 8 blue, 1 green; 11 red, 1 green; 3 green, 8 red, 2 blue; 4 blue, 11 red, 1 green; 5 blue, 5 red, 2 green
Game 16: 18 green, 4 blue, 2 red; 5 blue, 11 green, 10 red; 8 red, 2 blue, 14 green; 8 red, 7 blue, 1 green; 3 red, 5 blue, 17 green; 6 blue, 5 green, 11 red
Game 17: 3 blue, 3 red, 7 green; 4 blue, 1 red, 2 green; 5 blue, 3 green, 3 red
Game 18: 2 blue, 2 red, 1 green; 4 blue, 2 red, 7 green; 10 blue, 4 red, 3 green; 5 blue, 3 red, 2 green; 4 green, 3 red, 4 blue; 3 green, 5 red, 5 blue
Game 19: 2 red, 1 green, 1 blue; 8 red, 8 blue, 10 green; 16 green, 5 blue, 2 red; 4 red, 9 green
Game 20: 12 red, 1 blue, 1 green; 4 blue, 2 green, 2 red; 3 blue; 5 red, 8 green; 14 red, 4 blue, 6 green
Game 21: 9 red, 7 green, 1 blue; 5 green, 17 red, 11 blue; 14 red, 7 blue, 10 green; 7 green, 7 red, 10 blue; 6 blue, 6 green, 17 red; 16 red, 13 green, 7 blue
Game 22: 4 blue, 1 red; 1 green, 8 blue; 1 green; 6 blue, 1 red
Game 23: 13 red, 7 blue, 1 green; 4 green, 2 blue, 7 red; 4 green, 10 blue, 12 red
Game 24: 9 green, 10 blue; 2 blue, 4 green, 4 red; 9 green, 1 red, 9 blue; 4 green, 5 red, 12 blue
Game 25: 4 red, 1 green; 10 green, 6 red, 4 blue; 4 red, 1 blue, 7 green; 10 green, 3 red, 7 blue
Game 26: 8 red, 1 green, 2 blue; 5 green, 5 red; 6 green, 19 red; 11 red, 2 blue, 8 green; 13 red, 2 blue, 5 green; 15 red, 2 blue, 10 green
Game 27: 17 blue, 1 green; 2 red, 12 blue, 11 green; 16 green, 16 blue; 18 green, 4 blue; 10 blue, 1 red, 8 green
Game 28: 5 red, 1 green, 1 blue; 3 blue, 8 green, 4 red; 6 green, 2 red, 2 blue
Game 29: 3 green, 12 red, 11 blue; 2 green, 15 red, 8 blue; 13 red, 4 green; 17 red, 9 blue, 5 green
Game 30: 10 green; 4 blue, 1 green; 2 blue, 2 red, 7 green; 5 green, 4 blue, 1 red; 4 red, 10 green, 1 blue
Game 31: 15 blue, 2 red; 17 blue, 2 green; 19 blue, 6 red
Game 32: 1 green, 7 red; 8 red, 1 blue; 5 red, 1 blue, 11 green; 3 blue, 17 red
Game 33: 11 red, 9 green, 1 blue; 3 green, 8 blue; 10 red, 4 green, 8 blue; 6 red, 9 blue, 17 green; 15 green, 10 red, 4 blue; 1 red, 2 blue, 7 green
Game 34: 13 red, 6 green; 6 red, 14 green, 2 blue; 3 red, 19 green; 9 green, 9 red
Game 35: 7 green, 3 red; 12 green, 7 blue; 13 green, 7 red, 6 blue; 3 blue, 12 red
Game 36: 6 blue, 11 green, 14 red; 3 blue, 12 green, 4 red; 18 red, 1 blue; 7 red, 9 green, 6 blue
Game 37: 3 red, 16 blue, 6 green; 2 green, 7 blue; 8 blue, 3 red
Game 38: 16 blue, 3 green, 14 red; 8 red, 15 blue; 17 red, 15 blue, 4 green; 1 green, 11 blue, 17 red; 3 green, 10 blue, 17 red
Game 39: 1 green, 2 red, 5 blue; 12 blue, 12 green; 3 blue, 1 red
Game 40: 1 red, 2 blue, 1 green; 7 green, 1 red, 6 blue; 8 blue, 1 red, 6 green; 12 blue, 1 red, 3 green; 4 green, 8 blue
Game 41: 2 red, 2 blue, 5 green; 5 red, 8 blue; 4 green, 4 blue; 1 red, 11 blue
Game 42: 1 red, 3 green, 13 blue; 13 blue, 7 green; 13 green; 1 red, 3 blue, 4 green; 13 blue, 7 green
Game 43: 3 red, 4 green; 7 red, 11 blue, 3 green; 3 green, 12 red, 7 blue; 9 blue, 5 green
Game 44: 4 blue, 9 red, 2 green; 10 blue, 5 red, 2 green; 9 red, 9 blue, 1 green; 8 blue, 2 green, 14 red; 3 blue, 3 green, 6 red; 4 blue, 3 green, 14 red
Game 45: 1 red, 2 green, 2 blue; 2 green, 1 red; 1 green, 2 blue; 1 green, 1 red, 2 blue; 2 red, 2 blue, 1 green
Game 46: 1 green, 3 red, 3 blue; 6 green, 2 blue, 4 red; 1 green, 3 blue, 1 red; 3 green, 1 blue, 5 red; 6 green; 1 red, 1 green, 2 blue
Game 47: 18 green, 1 red, 7 blue; 6 blue, 19 green, 1 red; 5 blue, 7 green, 1 red; 1 red, 5 blue, 16 green; 15 green, 3 blue
Game 48: 4 green, 8 blue, 8 red; 13 green, 5 red, 12 blue; 9 red, 6 blue, 10 green; 18 green, 3 blue, 4 red; 2 blue, 9 red, 8 green
Game 49: 9 blue, 5 red, 9 green; 5 blue, 11 green, 5 red; 12 green, 6 blue
Game 50: 13 red, 8 green, 3 blue; 2 red, 11 green, 3 blue; 16 red, 7 green; 3 blue, 11 green, 15 red; 10 red, 2 blue, 5 green; 7 green, 2 blue, 4 red
Game 51: 2 red, 1 green, 3 blue; 2 green, 11 red, 17 blue; 2 red, 3 green, 6 blue; 4 red, 3 green, 6 blue; 13 red, 12 blue
Game 52: 1 blue, 5 green; 20 green, 6 blue; 9 blue, 6 green; 11 green, 1 red; 1 green, 1 red, 1 blue
Game 53: 8 red, 6 blue; 6 blue, 6 red, 2 green; 5 blue, 2 green, 3 red; 3 green, 3 blue; 4 green, 5 red, 1 blue
Game 54: 4 blue, 1 red, 3 green; 4 green, 10 blue, 9 red; 7 red, 3 blue, 3 green; 9 green, 9 red, 1 blue; 9 blue, 6 red, 7 green; 6 blue, 7 green, 9 red
Game 55: 15 red, 1 blue, 6 green; 11 blue, 3 red; 9 blue, 3 red, 1 green
Game 56: 8 green, 8 red, 9 blue; 8 red, 8 green, 1 blue; 7 red, 10 green, 4 blue; 10 blue, 2 green, 9 red
Game 57: 10 red, 3 green, 2 blue; 1 red, 4 green; 7 red, 1 green, 3 blue; 12 red, 4 blue; 14 red, 5 green, 4 blue
Game 58: 8 green, 3 blue, 7 red; 7 red, 14 blue, 5 green; 3 green, 7 red; 16 blue, 15 green; 1 red, 10 blue
Game 59: 3 red, 13 green, 2 blue; 10 blue, 3 green, 6 red; 3 green, 2 blue; 7 green, 2 blue, 7 red; 17 green, 6 blue, 15 red
Game 60: 2 blue, 2 red, 6 green; 11 green, 1 blue, 2 red; 1 blue, 9 green; 1 red, 4 green, 2 blue; 1 red, 2 blue, 10 green
Game 61: 3 red, 12 blue, 1 green; 3 red, 1 green, 18 blue; 5 blue, 2 red
Game 62: 4 red, 3 blue, 8 green; 2 blue, 8 red, 9 green; 8 blue, 15 green, 1 red
Game 63: 14 green, 2 red, 1 blue; 7 green, 11 blue, 1 red; 7 blue, 3 red; 4 green, 10 blue, 3 red
Game 64: 8 blue, 18 green, 2 red; 3 red, 17 green; 7 green, 1 red, 12 blue; 15 green, 2 red, 4 blue; 7 green, 8 red, 13 blue
Game 65: 6 blue, 5 green, 2 red; 1 red, 4 green; 5 green, 1 blue; 6 blue, 3 red, 2 green; 4 blue, 5 green
Game 66: 11 red, 9 blue, 4 green; 8 red, 8 blue; 9 red, 7 blue; 1 blue, 12 green, 4 red; 2 red, 11 blue, 10 green
Game 67: 1 red, 4 blue, 1 green; 7 red, 1 blue; 3 green, 4 blue, 6 red; 6 green, 3 blue, 14 red; 11 red, 1 blue, 1 green; 4 green, 8 red
Game 68: 3 red, 1 green, 2 blue; 1 red, 9 blue; 2 red, 1 green
Game 69: 3 green, 2 blue, 2 red; 1 red, 6 green; 13 red, 2 blue, 4 green; 4 blue, 13 red, 6 green; 12 red, 2 blue
Game 70: 15 blue, 2 green, 7 red; 3 red, 14 blue; 6 blue, 1 green; 1 red, 2 green, 4 blue; 2 green, 13 red; 12 blue, 3 red
Game 71: 7 red, 3 blue; 1 red, 4 blue; 2 red, 5 green, 1 blue; 6 blue, 8 red, 1 green; 3 green, 7 blue, 8 red
Game 72: 7 green; 4 green, 2 red, 8 blue; 1 blue, 5 green
Game 73: 5 red, 5 green, 2 blue; 8 red, 1 blue, 8 green; 1 red, 3 blue, 7 green
Game 74: 17 green, 9 blue, 4 red; 20 green, 2 red, 7 blue; 7 blue, 2 green, 4 red; 2 blue, 5 red, 20 green; 1 blue, 1 red, 12 green; 19 green, 9 blue, 3 red
Game 75: 1 red, 8 green, 9 blue; 7 blue, 3 green, 1 red; 2 green, 1 red, 9 blue; 5 blue, 1 red, 8 green; 2 green, 1 red, 11 blue; 5 green, 1 red
Game 76: 3 blue, 16 green, 2 red; 10 green, 3 blue, 1 red; 6 blue, 14 red, 13 green; 7 red, 2 green, 13 blue
Game 77: 7 red, 14 green; 1 blue, 1 red; 4 red, 1 green; 7 green, 11 red
Game 78: 1 red, 19 green; 10 green, 14 red, 1 blue; 3 green, 3 blue, 11 red; 7 blue, 1 green; 15 red, 3 green, 4 blue
Game 79: 7 red, 7 green, 6 blue; 3 red, 7 green, 5 blue; 7 red, 8 green, 12 blue
Game 80: 15 red, 6 blue; 1 red, 5 green, 2 blue; 1 green, 3 blue
Game 81: 3 red, 7 blue, 7 green; 7 green, 2 blue, 4 red; 3 green, 5 blue; 9 blue, 3 red, 6 green; 6 green, 1 red, 3 blue; 8 blue, 2 green, 1 red
Game 82: 5 red, 13 green; 3 blue, 13 green; 6 blue, 4 red, 10 green; 5 red, 1 green, 4 blue; 1 blue, 8 red; 4 red, 5 green
Game 83: 17 red, 1 blue, 2 green; 3 green, 3 red, 2 blue; 1 red, 5 blue, 10 green; 4 blue, 9 red, 11 green
Game 84: 13 green, 14 red, 12 blue; 14 blue, 2 red, 1 green; 4 blue, 8 red
Game 85: 3 red, 1 blue; 6 red, 3 blue, 2 green; 5 green, 3 blue, 3 red; 3 green, 5 blue, 1 red; 1 blue, 12 red, 2 green
Game 86: 16 blue, 17 green, 7 red; 14 blue, 13 green; 18 blue, 8 green
Game 87: 1 blue, 1 red; 4 blue, 1 green, 4 red; 1 green, 16 red; 1 green, 12 red, 1 blue
Game 88: 1 red, 6 green; 3 red, 2 blue, 19 green; 11 green, 2 red; 5 blue, 5 green; 5 blue, 9 green, 1 red; 2 blue, 2 red, 4 green
Game 89: 4 green, 11 red; 8 blue, 14 red; 14 blue, 8 green, 9 red; 14 green, 15 red, 10 blue
Game 90: 8 green, 2 red, 1 blue; 11 green, 4 blue, 2 red; 7 green, 2 blue; 13 green, 1 red
Game 91: 1 blue, 3 green; 1 blue; 4 green, 1 blue, 1 red; 1 blue, 2 red; 1 green, 2 red; 2 red, 5 green, 2 blue
Game 92: 16 red, 4 green, 5 blue; 9 blue, 13 green, 5 red; 13 red, 11 green, 7 blue; 11 red, 8 green, 2 blue
Game 93: 4 blue, 3 red, 3 green; 4 blue, 2 red, 1 green; 1 green, 2 red, 2 blue; 1 green, 2 red, 2 blue; 4 green, 1 blue
Game 94: 8 blue, 11 red, 7 green; 8 red, 6 green; 15 blue, 11 green, 2 red; 9 green, 6 red; 16 blue, 5 red, 7 green
Game 95: 13 blue, 1 red, 10 green; 11 green, 9 blue; 6 blue
Game 96: 1 green, 6 red; 1 red; 12 red, 1 green; 6 red, 1 blue
Game 97: 1 red, 9 blue, 8 green; 2 green, 6 blue, 1 red; 6 green, 1 blue
Game 98: 9 blue, 7 green, 8 red; 6 red, 11 blue, 4 green; 11 green, 9 blue, 15 red; 11 red, 6 blue, 16 green
Game 99: 2 blue, 1 red, 9 green; 8 red, 1 blue, 1 green; 2 red, 7 green, 8 blue; 1 red, 5 green, 7 blue; 7 blue, 10 green, 9 red; 1 green, 1 blue, 1 red
Game 100: 3 blue, 6 red, 9 green; 4 red, 3 green; 4 green, 16 red, 1 blue; 14 blue, 1 green`
