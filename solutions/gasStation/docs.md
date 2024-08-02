#Gas Station
##Question
There are n gas stations along a circular route, 
where the amount of gas at the ith station is gas[i].

You have a car with an unlimited gas tank and it costs cost[i]
of gas to travel from the ith station to its next (i + 1)th station.
You begin the journey with an empty tank at one of the gas stations.

Given two integer arrays gas and cost, return the starting gas station's
index if you can travel around the circuit once in the clockwise direction, 
otherwise return -1. If there exists a solution, it is guaranteed to be unique

#Observations:
If the total amount of gas is less than the total cost required to complete the
circuit, it's impossible to complete the circuit. This is because you need at least as much gas as the total cost to make a full loop.
If you can complete the circuit starting from any station, it will be unique.

#SOLUTION--Algorithm
1. Check Feasibility:
First, calculate the total gas and total cost.
If total_gas < total_cost, return -1 because it’s impossible to complete the circuit.

2. Find the Starting Point:
Traverse the gas stations while keeping track of the current balance of gas.
Use two variables: total_tank to track the overall gas balance and current_tank
to track the balance from the current starting index.

If current_tank drops below zero during traversal, it indicates that the journey 
from the current start index to the current station isn’t feasible. Therefore, reset the starting index to the next station and reset current_tank to zero.

##Example
```
Input: gas = [1,2,3,4,5], cost = [3,4,5,1,2]
Output: 3
Explanation:
Start at station 3 (index 3) and fill up with 4 unit of gas. Your tank = 0 + 4 = 4
Travel to station 4. Your tank = 4 - 1 + 5 = 8
Travel to station 0. Your tank = 8 - 2 + 1 = 7
Travel to station 1. Your tank = 7 - 3 + 2 = 6
Travel to station 2. Your tank = 6 - 4 + 3 = 5
Travel to station 3. The cost is 5. Your gas is just enough to travel back to station 3.
Therefore, return 3 as the starting index.
```