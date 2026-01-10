func canCompleteCircuit(gas []int, cost []int) int {
    length := len(gas)
    diff, total, start, tank := make([]int, length), 0, 0, 0
    for i := 0; i < length; i++ {
        diff[i] = gas[i] - cost[i]
        total += diff[i]
        tank += diff[i]
        if tank < 0 {
            tank = 0
            start = i + 1
        }
    }
    if total < 0 {
        return -1
    }
    return start
}