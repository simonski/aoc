#!/usr/bin/env python3
from typing import List
import sys
import os

DEBUG = os.environ.get("DEBUG", "0") == "1"

INPUT_1 = [ "7,13,x,x,59,x,31,19", 1068781 ]
INPUT_2 = [ "17,x,13,19", 3417 ]
INPUT_3 = [ "67,7,59,61", 754018 ]
INPUT_4 = [ "67,x,7,59,61", 779210 ]
INPUT_5 = [ "67,7,x,59,61", 1261476 ]
INPUT_6 = [ "1789,37,47,1889", 1202161486 ]
INPUT_7 = [ "29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,433,x,x,x,x,x,x,x,x,x,x,x,x,13,17,x,x,x,x,19,x,x,x,23,x,x,x,x,x,x,x,977,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,41", 1 ]

INPUTS = [ INPUT_1, INPUT_2, INPUT_3, INPUT_4, INPUT_5, INPUT_6, INPUT_7 ]

class InputData:
    def __init__(self, values, target=0):
        # values = "7,13,x,x,59,x,31,19"
        # target = 1068781 
        ints_and_index = []
        self.splits = values.split(",")
        for index in range(len(self.splits)):
            entry = self.splits[index]
            if entry != 'x':
                ints_and_index.append([int(entry), index ])

        self.values = ints_and_index
        self.original_values = values
        self.target = target

    def size(self):
        return len(self.values)

    def get(self, index):
        return self.values[index]

    def get_ints_original(self):
        ints = []
        for index in range(len(self.values)):
            entry = self.values[index]
            value = entry[0]
            ints.append(value)
        return ints

    def get_ints_modified(self):
        ints = []
        for index in range(len(self.values)):
            entry = self.values[index]
            value = entry[0] + entry[1]
            ints.append(value)
        return ints

        
class Pair:
    def __init__(self, first, last, index=0, length=0):
        self.first = first
        self.last = last
        self.index = index
        self.length = length
        self.offset = self.calculate_offset(first, last, length)
        self.lcm = calculate_lcm(first, last)
        self.gcd = calculate_gcd(first, last)
        self.coprime = self.gcd == 1

    def __str__(self):
        return "Pair({},{}) lcm={}, gcd={}, index={}, length={}, offset={}".format(self.first, self.last, self.lcm, self.gcd, self.index, self.length, self.offset)

    def calculate_offset(self, first, last, length):
        """ 
        works out the first index that t (from 0....n) 
        t % first = 0
        t + length % last = 0
        """
        t = 0
        while True:
            t1 = t % first
            t2 = (t+length) % last
            if (t1 + t2) == 0:
                return t
            t += 1

def calculate_lcm( value1, value2):
    """
    calculates largest common multiplier
    """
    for t in range(1, value1*value2):
        if (t % value1 ==0) and (t % value2) == 0:
            return t
    return value1 * value2

def calculate_gcd(value1, value2):
    """ 
    calcualtes greatest common divisor 
    """
    return (value1 * value2) / calculate_lcm(value1, value2)

def calculate_gcd_euclid_recursive_values(values, depth=0):
    while len(values) != 1:
        # print(values)
        value1 = values[0]
        value2 = values[1]
        gcd = calculate_gcd_euclid_recursive(value1, value2, depth)
        new_values = [ gcd ]
        new_values += values[2:]
        values = new_values
    return values[0]

def calculate_gcd_euclid_recursive(value1, value2, depth=0):
    """
    1. Given two whole numbers, subtract the smaller number from the larger number and note the result.
    2. Repeat the process subtracting the smaller number from the result until the result is smaller than the original small number.
    3. Use the original small number as the new larger number. Subtract the result from Step 2 from the new larger number.
    4. Repeat the process for every new larger number and smaller number until you reach zero.
    5. When you reach zero, go back one calculation: the GCF is the number you found just before the zero result.        
    """

    # if DEBUG:
    #     print("> euclid[{}] {},{}".format(depth, value1, value2))
    min_value = min(value1, value2)
    max_value = max(value1, value2)

    new_max = max_value
    while new_max >= min_value:
        new_max -= min_value
    
    if new_max == 0:
        # if DEBUG:
        #     print("< euclid[{}] {},{} = {}".format(depth, value1, value2, min_value))
        return min_value
    else:
        new_min = min_value - new_max 
        depth += 1
        return calculate_gcd_euclid_recursive(new_max, new_min, depth)

    
def build_pairs_first_variant_only(data_str:str) -> List[Pair]:
    """
    returns the pairs for index 0,1... 0,2... 0,n
    """
    data = data_str.split(",")
    results = []
    value1 = data[0]
    for index in range(1, len(data)):
        value2 = data[index]
        if value2 == 'x':
            continue
        else:
            first = int(value1)
            last =int(value2)
            length = index
            p = Pair(first, last, 0, length)
            results.append(p)

    return results
    
def build_pairs_all_variants(data_str:str) -> (Pair, List[Pair]):
    """
    returns the largest pair and all other pairs in a list
    """
    data = data_str.split(",")
    results = []
    for index1 in range(0, len(data)):
        value1 = data[index1]
        if value1 == 'x':
            continue
        for index2 in range(index1+1, len(data)):
            value2 = data[index2]
            if value2 == 'x':
                continue
            else:
                first = int(value1)
                last =int(value2)
                index = index1
                length = index2 - index1
                p = Pair(first, last, index, length)
                results.append(p)

    largest = results[0]
    for candidate in results:
        if candidate.lcm > largest.lcm:
            largest = candidate
    return largest, results

def is_match(time, inputs):
    """
    checks all input values match their modulus
    """
    success = True
    output= ""
    for entry in inputs:
        if entry != "x":
            mod = time % entry
            if mod != 0:
                success = False
            output += " " + str(mod)
        else:
            output += " x"
        time += 1
    if DEBUG:
        print(output)
    return success    

def test_inputs(data_to_test_with):
    #data = sys.argv[1]
    #largest, pairs = build_pairs(data)
    #for pair in pairs:
    #    print(pair) #"({},{}) length {} offset {} lcm {}".format(pair.first, pair.last, pair.length, pair.offset, pair.lcm))

    for data in data_to_test_with:
        print(">>>>>>>>>>>>>>>>>>>>>>")
        print("Testing {}".format(data))
        found, attempts, search_space, increment = attempt_5(data[0])
        pct = (100.0 / search_space) * attempts
        print("\n Attempt 5 (First/Last Pair): found {}, tried {}, increment size {} , search space {} ({}%).".format(found, attempts, increment, search_space, pct))

        # found, attempts, search_space, increment = attempt_4(data[0])
        # pct = (100.0 / search_space) * attempts
        # print(data)
        # print("\nAttempt 4 : found {}, tried {}, increment size {} , search space {} ({}%).".format(found, attempts, increment, search_space, pct))

        # found, attempts, search_space, increment = attempt_3(data[0])
        # pct = (100.0 / search_space) * attempts
        # print("\n Attempt 3 (First/Last Pair): found {}, tried {}, increment size {} , search space {} ({}%).".format(found, attempts, increment, search_space, pct))
        print(">>>>>>>>>>>>>>>>>>>>>>\n\n")
    
def test_logic(test_input):

    print("\n>>>>>>>>>>>>>>>>>>>>>>")
    print("Test Logic: {}".format(test_input))
    pairs = build_pairs_first_variant_only(test_input[0])
    for p in pairs:
         print(p)

    p1 = pairs[0]
    pairs = pairs[1:]
    round = 1
    t = 0
    while len(pairs) > 0:
        print("")
        print("Round {}".format(round))
        p2 = pairs[0]
        print("P1: {}", p1)
        print("P2: {}", p2)

        # print(p1)
        # for p in pairs:
        #     print(p)


        p1_iterations, p2_iterations, iterations, t = find_number_of_iterations_to_match_pairs_from_respective_offsets(p1, p2)

        # gives me the first t value they meet at
        
        print("Step1/2 [p1,p2]: p1_Offset+p0_iterations={}, p2_Offset+p1_iterations={}, iterations={}, t={}".format(p1_iterations, p2_iterations, iterations, t))
        p1_iterations, p2_iterations, iterations, t1 = find_number_of_iterations_to_match_pairs_from_index(p1, p2, 0)
        print("Step2/2 [p1,p2]: p0_iterations={}, p3_iterations={}, iterations={}, t={}".format(p1_iterations, p2_iterations, iterations, t))

        p1 = Pair(1, 1)

        p1.offset = t
        p1.lcm = t1

        print("New pair {}".format(p1))

        pairs = pairs[1:]
        round += 1
        print("")


    print("RESULT: >>>>>>>>>>>>>>>>>>>>>>")
    print("RESULT: Test Input was {}".format(test_input))
    print("RESULT: Final pair was {}".format(p1))
    print("RESULT: Answer was {}".format(p1.offset))
    print("RESULT: >>>>>>>>>>>>>>>>>>>>>>")


def find_number_of_iterations_to_match_pairs_from_respective_offsets(p1, p2):
    iterations = 0
    p1_t = p1.offset
    p2_t = p2.offset
    p1_count = 0
    p2_count = 0
    #print("find_number_of_iterations_to_match_pairs_from_respective_offset, p1t={}, p2t={}".format(p1_t, p2_t))
    while True:
        if p1_t < p2_t:
            difference = p2_t - p1_t
            multiple = int(difference/p1.lcm)
            multiple = max(1, multiple)
            p1_t += (p1.lcm*multiple)
            p1_count += 1
        elif p2_t < p1_t:
            difference = p1_t - p2_t
            multiple = int(difference/p2.lcm)
            multiple = max(1, multiple)
            p2_t += (p2.lcm*multiple)
            p2_count += 1
        elif p1_t == p2_t:
            break

        # print("p1={}".format(p1))
        # print("p2={}".format(p2))
        iterations +=1 
    return p1_count, p2_count, iterations, p1_t

def find_number_of_iterations_to_match_pairs_from_index(p1, p2, index):
    iterations = 0
    p1_t = index
    p2_t = index + p2.lcm
    p1_count = 0
    p2_count = 0
    while True:
        if p1_t < p2_t:
            difference = p2_t - p1_t
            multiple = int(difference/p1.lcm)
            multiple = max(1, multiple)
            p1_t += (p1.lcm*multiple)
            p1_count += 1
        elif p2_t < p1_t:
            difference = p1_t - p2_t
            multiple = int(difference/p2.lcm)
            multiple = max(1, multiple)
            p2_t += (p2.lcm*multiple)
            p2_count += 1
        elif p1_t == p2_t:
            break
        iterations +=1 
    return p1_count, p2_count, iterations, p1_t


def main():
    for test_input in INPUTS:
        test_logic(test_input)

if __name__ == "__main__":
    main()

